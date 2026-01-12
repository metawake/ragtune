package embedder

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestNewVoyageEmbedder(t *testing.T) {
	// Save and clear env var for testing
	original := os.Getenv("VOYAGE_API_KEY")
	defer os.Setenv("VOYAGE_API_KEY", original)
	os.Setenv("VOYAGE_API_KEY", "test-key")

	t.Run("default configuration", func(t *testing.T) {
		e := NewVoyageEmbedder()

		if e.model != "voyage-2" {
			t.Errorf("model = %v, want voyage-2", e.model)
		}
		if e.dim != 1024 {
			t.Errorf("dim = %v, want 1024", e.dim)
		}
		if e.inputType != "document" {
			t.Errorf("inputType = %v, want document", e.inputType)
		}
	})

	t.Run("with voyage-law-2 model", func(t *testing.T) {
		e := NewVoyageEmbedder(WithVoyageModel("voyage-law-2"))

		if e.model != "voyage-law-2" {
			t.Errorf("model = %v, want voyage-law-2", e.model)
		}
		if e.dim != 1024 {
			t.Errorf("dim = %v, want 1024", e.dim)
		}
	})

	t.Run("with voyage-code-2 model", func(t *testing.T) {
		e := NewVoyageEmbedder(WithVoyageModel("voyage-code-2"))

		if e.model != "voyage-code-2" {
			t.Errorf("model = %v, want voyage-code-2", e.model)
		}
	})

	t.Run("with custom input type", func(t *testing.T) {
		e := NewVoyageEmbedder(WithVoyageInputType("query"))

		if e.inputType != "query" {
			t.Errorf("inputType = %v, want query", e.inputType)
		}
	})
}

func TestVoyageEmbedder_Dim(t *testing.T) {
	tests := []struct {
		model    string
		expected int
	}{
		{"voyage-2", 1024},
		{"voyage-large-2", 1024},
		{"voyage-law-2", 1024},
		{"voyage-code-2", 1024},
		{"voyage-finance-2", 1024},
		{"voyage-lite-02-instruct", 1024},
		{"unknown-model", 1024}, // defaults to 1024
	}

	for _, tt := range tests {
		t.Run(tt.model, func(t *testing.T) {
			e := NewVoyageEmbedder(WithVoyageModel(tt.model))
			if e.Dim() != tt.expected {
				t.Errorf("Dim() = %v, want %v", e.Dim(), tt.expected)
			}
		})
	}
}

func TestVoyageEmbedder_MissingAPIKey(t *testing.T) {
	// Clear env var
	original := os.Getenv("VOYAGE_API_KEY")
	defer os.Setenv("VOYAGE_API_KEY", original)
	os.Unsetenv("VOYAGE_API_KEY")

	e := NewVoyageEmbedder()
	_, err := e.Embed(context.Background(), "test text")

	if err == nil {
		t.Error("expected error for missing API key")
	}
	if err.Error() != "VOYAGE_API_KEY environment variable not set" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestVoyageEmbedder_EmbedBatchMissingAPIKey(t *testing.T) {
	// Clear env var
	original := os.Getenv("VOYAGE_API_KEY")
	defer os.Setenv("VOYAGE_API_KEY", original)
	os.Unsetenv("VOYAGE_API_KEY")

	e := NewVoyageEmbedder()
	_, err := e.EmbedBatch(context.Background(), []string{"text1", "text2"})

	if err == nil {
		t.Error("expected error for missing API key")
	}
}

func TestVoyageEmbedder_EmbedSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.Header.Get("Authorization") != "Bearer test-key" {
			t.Errorf("unexpected auth header: %s", r.Header.Get("Authorization"))
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data": [{"index": 0, "embedding": [0.1, 0.2, 0.3]}]}`))
	}))
	defer server.Close()

	original := os.Getenv("VOYAGE_API_KEY")
	defer os.Setenv("VOYAGE_API_KEY", original)
	os.Setenv("VOYAGE_API_KEY", "test-key")

	e := NewVoyageEmbedder(WithVoyageURL(server.URL))
	result, err := e.Embed(context.Background(), "test text")

	if err != nil {
		t.Fatalf("Embed failed: %v", err)
	}
	if len(result) != 3 {
		t.Errorf("expected 3 dimensions, got %d", len(result))
	}
}

func TestVoyageEmbedder_EmbedBatchSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data": [{"index": 0, "embedding": [0.1, 0.2]}, {"index": 1, "embedding": [0.3, 0.4]}]}`))
	}))
	defer server.Close()

	original := os.Getenv("VOYAGE_API_KEY")
	defer os.Setenv("VOYAGE_API_KEY", original)
	os.Setenv("VOYAGE_API_KEY", "test-key")

	e := NewVoyageEmbedder(WithVoyageURL(server.URL))
	results, err := e.EmbedBatch(context.Background(), []string{"text1", "text2"})

	if err != nil {
		t.Fatalf("EmbedBatch failed: %v", err)
	}
	if len(results) != 2 {
		t.Errorf("expected 2 results, got %d", len(results))
	}
}

func TestVoyageEmbedder_APIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte(`{"detail": "rate limit exceeded"}`))
	}))
	defer server.Close()

	original := os.Getenv("VOYAGE_API_KEY")
	defer os.Setenv("VOYAGE_API_KEY", original)
	os.Setenv("VOYAGE_API_KEY", "test-key")

	e := NewVoyageEmbedder(WithVoyageURL(server.URL))
	_, err := e.Embed(context.Background(), "test")

	if err == nil {
		t.Error("expected error for API error response")
	}
}

func TestVoyageEmbedder_InvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{invalid json`))
	}))
	defer server.Close()

	original := os.Getenv("VOYAGE_API_KEY")
	defer os.Setenv("VOYAGE_API_KEY", original)
	os.Setenv("VOYAGE_API_KEY", "test-key")

	e := NewVoyageEmbedder(WithVoyageURL(server.URL))
	_, err := e.Embed(context.Background(), "test")

	if err == nil {
		t.Error("expected error for invalid JSON response")
	}
}



