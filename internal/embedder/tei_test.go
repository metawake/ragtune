package embedder

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewTEIEmbedder(t *testing.T) {
	t.Run("default configuration", func(t *testing.T) {
		e := NewTEIEmbedder()

		if e.baseURL != "http://localhost:8080" {
			t.Errorf("baseURL = %v, want http://localhost:8080", e.baseURL)
		}
		if e.model != "BAAI/bge-base-en-v1.5" {
			t.Errorf("model = %v, want BAAI/bge-base-en-v1.5", e.model)
		}
		if e.dim != 768 {
			t.Errorf("dim = %v, want 768", e.dim)
		}
	})

	t.Run("with custom URL", func(t *testing.T) {
		e := NewTEIEmbedder(WithTEIURL("http://custom:9090"))

		if e.baseURL != "http://custom:9090" {
			t.Errorf("baseURL = %v, want http://custom:9090", e.baseURL)
		}
	})

	t.Run("with custom model", func(t *testing.T) {
		e := NewTEIEmbedder(WithTEIModel("BAAI/bge-large-en-v1.5"))

		if e.model != "BAAI/bge-large-en-v1.5" {
			t.Errorf("model = %v, want BAAI/bge-large-en-v1.5", e.model)
		}
		if e.dim != 1024 {
			t.Errorf("dim = %v, want 1024 for bge-large", e.dim)
		}
	})

	t.Run("with explicit dimension", func(t *testing.T) {
		e := NewTEIEmbedder(WithTEIDim(2048))

		if e.dim != 2048 {
			t.Errorf("dim = %v, want 2048", e.dim)
		}
	})

	t.Run("explicit dim overrides model inference", func(t *testing.T) {
		// First set model, then override with explicit dim
		e := NewTEIEmbedder(
			WithTEIModel("BAAI/bge-small-en-v1.5"), // Would be 384
			WithTEIDim(512),                        // Override to 512
		)

		// Explicit dim takes precedence over model-based detection
		if e.dim != 512 {
			t.Errorf("dim = %v, want 512 (explicit dim should override)", e.dim)
		}
	})
}

func TestTEIEmbedder_Dim(t *testing.T) {
	tests := []struct {
		model    string
		expected int
	}{
		// BGE models
		{"BAAI/bge-small-en-v1.5", 384},
		{"BAAI/bge-base-en-v1.5", 768},
		{"BAAI/bge-large-en-v1.5", 1024},
		// Sentence Transformers
		{"sentence-transformers/all-MiniLM-L6-v2", 384},
		{"sentence-transformers/all-mpnet-base-v2", 768},
		// Nomic
		{"nomic-ai/nomic-embed-text-v1.5", 768},
		// GTE models
		{"thenlper/gte-small", 384},
		{"thenlper/gte-base", 768},
		{"thenlper/gte-large", 1024},
		{"Alibaba-NLP/gte-Qwen2-1.5B-instruct", 1536},
		// E5 models
		{"intfloat/e5-small-v2", 384},
		{"intfloat/e5-base-v2", 768},
		{"intfloat/e5-large-v2", 1024},
		// Unknown model - defaults to 768
		{"unknown/custom-model", 768},
	}

	for _, tt := range tests {
		t.Run(tt.model, func(t *testing.T) {
			e := NewTEIEmbedder(WithTEIModel(tt.model))
			if e.Dim() != tt.expected {
				t.Errorf("Dim() = %v, want %v", e.Dim(), tt.expected)
			}
		})
	}
}

func TestTEIEmbedder_EmbedBatch(t *testing.T) {
	// Mock TEI server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/embed" {
			t.Errorf("unexpected path: %s", r.URL.Path)
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		if r.Method != "POST" {
			t.Errorf("unexpected method: %s", r.Method)
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req teiEmbedRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Return mock embeddings (3 dimensions for simplicity)
		embeddings := make([][]float64, len(req.Inputs))
		for i := range req.Inputs {
			embeddings[i] = []float64{0.1 * float64(i+1), 0.2 * float64(i+1), 0.3 * float64(i+1)}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(embeddings)
	}))
	defer server.Close()

	e := NewTEIEmbedder(WithTEIURL(server.URL))

	t.Run("batch embedding", func(t *testing.T) {
		texts := []string{"hello", "world", "test"}
		results, err := e.EmbedBatch(context.Background(), texts)
		if err != nil {
			t.Fatalf("EmbedBatch failed: %v", err)
		}

		if len(results) != 3 {
			t.Errorf("got %d results, want 3", len(results))
		}

		// Check first embedding values
		if len(results[0]) != 3 {
			t.Errorf("embedding dimension = %d, want 3", len(results[0]))
		}
		if results[0][0] != 0.1 {
			t.Errorf("results[0][0] = %v, want 0.1", results[0][0])
		}
	})

	t.Run("single embedding", func(t *testing.T) {
		result, err := e.Embed(context.Background(), "hello")
		if err != nil {
			t.Fatalf("Embed failed: %v", err)
		}

		if len(result) != 3 {
			t.Errorf("embedding dimension = %d, want 3", len(result))
		}
	})

	t.Run("empty batch", func(t *testing.T) {
		results, err := e.EmbedBatch(context.Background(), []string{})
		if err != nil {
			t.Fatalf("EmbedBatch failed for empty input: %v", err)
		}
		if results != nil {
			t.Errorf("expected nil for empty input, got %v", results)
		}
	})
}

func TestTEIEmbedder_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}))
	defer server.Close()

	e := NewTEIEmbedder(WithTEIURL(server.URL))

	_, err := e.Embed(context.Background(), "test")
	if err == nil {
		t.Error("expected error for server error response")
	}
}

func TestTEIEmbedder_ConnectionError(t *testing.T) {
	e := NewTEIEmbedder(WithTEIURL("http://localhost:99999"))

	_, err := e.Embed(context.Background(), "test")
	if err == nil {
		t.Error("expected error for connection failure")
	}
}

