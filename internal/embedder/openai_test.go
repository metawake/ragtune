package embedder

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestNewOpenAIEmbedder(t *testing.T) {
	e := NewOpenAIEmbedder()

	if e.model != "text-embedding-3-small" {
		t.Errorf("model = %v, want text-embedding-3-small", e.model)
	}
	if e.dim != 1536 {
		t.Errorf("dim = %v, want 1536", e.dim)
	}
	if e.baseURL != "https://api.openai.com/v1/embeddings" {
		t.Errorf("baseURL = %v, want https://api.openai.com/v1/embeddings", e.baseURL)
	}
}

func TestOpenAIEmbedder_WithURL(t *testing.T) {
	e := NewOpenAIEmbedder(WithOpenAIURL("http://custom:8080"))

	if e.baseURL != "http://custom:8080" {
		t.Errorf("baseURL = %v, want http://custom:8080", e.baseURL)
	}
}

func TestOpenAIEmbedder_MissingAPIKey(t *testing.T) {
	original := os.Getenv("OPENAI_API_KEY")
	defer os.Setenv("OPENAI_API_KEY", original)
	os.Unsetenv("OPENAI_API_KEY")

	e := NewOpenAIEmbedder()
	_, err := e.Embed(context.Background(), "test")

	if err == nil {
		t.Error("expected error for missing API key")
	}
}

func TestOpenAIEmbedder_EmbedSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.Header.Get("Authorization") != "Bearer test-key" {
			t.Errorf("unexpected auth header: %s", r.Header.Get("Authorization"))
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"data": [
				{"index": 0, "embedding": [0.1, 0.2, 0.3]}
			]
		}`))
	}))
	defer server.Close()

	original := os.Getenv("OPENAI_API_KEY")
	defer os.Setenv("OPENAI_API_KEY", original)
	os.Setenv("OPENAI_API_KEY", "test-key")

	e := NewOpenAIEmbedder(WithOpenAIURL(server.URL))
	result, err := e.Embed(context.Background(), "test text")

	if err != nil {
		t.Fatalf("Embed failed: %v", err)
	}
	if len(result) != 3 {
		t.Errorf("expected 3 dimensions, got %d", len(result))
	}
	if result[0] != 0.1 {
		t.Errorf("result[0] = %v, want 0.1", result[0])
	}
}

func TestOpenAIEmbedder_EmbedBatchSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"data": [
				{"index": 0, "embedding": [0.1, 0.2]},
				{"index": 1, "embedding": [0.3, 0.4]}
			]
		}`))
	}))
	defer server.Close()

	original := os.Getenv("OPENAI_API_KEY")
	defer os.Setenv("OPENAI_API_KEY", original)
	os.Setenv("OPENAI_API_KEY", "test-key")

	e := NewOpenAIEmbedder(WithOpenAIURL(server.URL))
	results, err := e.EmbedBatch(context.Background(), []string{"text1", "text2"})

	if err != nil {
		t.Fatalf("EmbedBatch failed: %v", err)
	}
	if len(results) != 2 {
		t.Errorf("expected 2 results, got %d", len(results))
	}
}

func TestOpenAIEmbedder_APIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write([]byte(`{"error": {"message": "rate limit exceeded"}}`))
	}))
	defer server.Close()

	original := os.Getenv("OPENAI_API_KEY")
	defer os.Setenv("OPENAI_API_KEY", original)
	os.Setenv("OPENAI_API_KEY", "test-key")

	e := NewOpenAIEmbedder(WithOpenAIURL(server.URL))
	_, err := e.Embed(context.Background(), "test")

	if err == nil {
		t.Error("expected error for API error response")
	}
}

func TestOpenAIEmbedder_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": {"message": "internal error"}}`))
	}))
	defer server.Close()

	original := os.Getenv("OPENAI_API_KEY")
	defer os.Setenv("OPENAI_API_KEY", original)
	os.Setenv("OPENAI_API_KEY", "test-key")

	e := NewOpenAIEmbedder(WithOpenAIURL(server.URL))
	_, err := e.Embed(context.Background(), "test")

	if err == nil {
		t.Error("expected error for server error")
	}
}

func TestOpenAIEmbedder_InvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{invalid json`))
	}))
	defer server.Close()

	original := os.Getenv("OPENAI_API_KEY")
	defer os.Setenv("OPENAI_API_KEY", original)
	os.Setenv("OPENAI_API_KEY", "test-key")

	e := NewOpenAIEmbedder(WithOpenAIURL(server.URL))
	_, err := e.Embed(context.Background(), "test")

	if err == nil {
		t.Error("expected error for invalid JSON response")
	}
}
