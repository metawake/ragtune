package embedder

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func TestNewOllamaEmbedder(t *testing.T) {
	t.Run("default configuration", func(t *testing.T) {
		e := NewOllamaEmbedder()

		if e.baseURL != "http://localhost:11434" {
			t.Errorf("baseURL = %v, want http://localhost:11434", e.baseURL)
		}
		if e.model != "nomic-embed-text" {
			t.Errorf("model = %v, want nomic-embed-text", e.model)
		}
		if e.dim != 768 {
			t.Errorf("dim = %v, want 768", e.dim)
		}
		if e.concurrency != 8 {
			t.Errorf("concurrency = %v, want 8", e.concurrency)
		}
	})

	t.Run("with custom URL", func(t *testing.T) {
		e := NewOllamaEmbedder(WithOllamaURL("http://custom:8080"))

		if e.baseURL != "http://custom:8080" {
			t.Errorf("baseURL = %v, want http://custom:8080", e.baseURL)
		}
	})

	t.Run("with custom model", func(t *testing.T) {
		e := NewOllamaEmbedder(WithOllamaModel("mxbai-embed-large"))

		if e.model != "mxbai-embed-large" {
			t.Errorf("model = %v, want mxbai-embed-large", e.model)
		}
		if e.dim != 1024 {
			t.Errorf("dim = %v, want 1024 for mxbai-embed-large", e.dim)
		}
	})

	t.Run("with custom concurrency", func(t *testing.T) {
		e := NewOllamaEmbedder(WithOllamaConcurrency(16))

		if e.concurrency != 16 {
			t.Errorf("concurrency = %v, want 16", e.concurrency)
		}
	})

	t.Run("concurrency ignores invalid values", func(t *testing.T) {
		e := NewOllamaEmbedder(WithOllamaConcurrency(0))

		if e.concurrency != 8 { // Should keep default
			t.Errorf("concurrency = %v, want 8 (default) for invalid input", e.concurrency)
		}

		e2 := NewOllamaEmbedder(WithOllamaConcurrency(-5))
		if e2.concurrency != 8 {
			t.Errorf("concurrency = %v, want 8 (default) for negative input", e2.concurrency)
		}
	})
}

func TestOllamaEmbedder_Dim(t *testing.T) {
	tests := []struct {
		model    string
		expected int
	}{
		{"nomic-embed-text", 768},
		{"mxbai-embed-large", 1024},
		{"all-minilm", 384},
		{"unknown-model", 768}, // defaults to 768
	}

	for _, tt := range tests {
		t.Run(tt.model, func(t *testing.T) {
			e := NewOllamaEmbedder(WithOllamaModel(tt.model))
			if e.Dim() != tt.expected {
				t.Errorf("Dim() = %v, want %v", e.Dim(), tt.expected)
			}
		})
	}
}

func TestOllamaEmbedder_EmbedBatchConcurrency(t *testing.T) {
	// Track concurrent requests
	var currentConcurrent int32
	var maxConcurrent int32

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Track concurrent requests
		current := atomic.AddInt32(&currentConcurrent, 1)
		defer atomic.AddInt32(&currentConcurrent, -1)

		// Update max seen
		for {
			max := atomic.LoadInt32(&maxConcurrent)
			if current <= max || atomic.CompareAndSwapInt32(&maxConcurrent, max, current) {
				break
			}
		}

		// Simulate some processing time
		time.Sleep(10 * time.Millisecond)

		// Return mock embedding
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"embedding": [0.1, 0.2, 0.3]}`))
	}))
	defer server.Close()

	e := NewOllamaEmbedder(
		WithOllamaURL(server.URL),
		WithOllamaConcurrency(4),
	)

	// Embed 20 texts - should use up to 4 concurrent requests
	texts := make([]string, 20)
	for i := range texts {
		texts[i] = "test text"
	}

	results, err := e.EmbedBatch(context.Background(), texts)
	if err != nil {
		t.Fatalf("EmbedBatch failed: %v", err)
	}

	if len(results) != 20 {
		t.Errorf("got %d results, want 20", len(results))
	}

	// Verify concurrency was used (should be close to 4)
	maxSeen := atomic.LoadInt32(&maxConcurrent)
	if maxSeen < 2 {
		t.Errorf("max concurrent requests = %d, expected at least 2 (concurrency=4)", maxSeen)
	}
	if maxSeen > 4 {
		t.Errorf("max concurrent requests = %d, exceeded limit of 4", maxSeen)
	}
}

func TestOllamaEmbedder_EmbedBatchEmpty(t *testing.T) {
	e := NewOllamaEmbedder()

	results, err := e.EmbedBatch(context.Background(), []string{})
	if err != nil {
		t.Fatalf("EmbedBatch failed for empty input: %v", err)
	}
	if results != nil {
		t.Errorf("expected nil for empty input, got %v", results)
	}
}

func TestOllamaEmbedder_EmbedBatchCancellation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Slow response to allow cancellation
		time.Sleep(100 * time.Millisecond)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"embedding": [0.1, 0.2, 0.3]}`))
	}))
	defer server.Close()

	e := NewOllamaEmbedder(
		WithOllamaURL(server.URL),
		WithOllamaConcurrency(2),
	)

	ctx, cancel := context.WithCancel(context.Background())

	// Cancel after a short delay
	go func() {
		time.Sleep(20 * time.Millisecond)
		cancel()
	}()

	texts := make([]string, 10)
	for i := range texts {
		texts[i] = "test text"
	}

	_, err := e.EmbedBatch(ctx, texts)
	if err == nil {
		t.Error("expected error after context cancellation")
	}
}


