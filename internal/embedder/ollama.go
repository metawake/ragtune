package embedder

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Compile-time interface compliance check.
var _ Embedder = (*OllamaEmbedder)(nil)

// OllamaEmbedder uses Ollama's local embedding API.
type OllamaEmbedder struct {
	baseURL     string
	model       string
	dim         int
	concurrency int
	client      *http.Client
}

// OllamaOption configures the Ollama embedder.
type OllamaOption func(*OllamaEmbedder)

// WithOllamaURL sets the Ollama server URL.
func WithOllamaURL(url string) OllamaOption {
	return func(e *OllamaEmbedder) {
		e.baseURL = url
	}
}

// WithOllamaModel sets the embedding model.
func WithOllamaModel(model string) OllamaOption {
	return func(e *OllamaEmbedder) {
		e.model = model
	}
}

// WithOllamaConcurrency sets the number of concurrent embedding requests.
// Default is 8. Higher values speed up large batches but increase memory/CPU usage.
// For 100k+ datasets, 8-16 is recommended.
func WithOllamaConcurrency(n int) OllamaOption {
	return func(e *OllamaEmbedder) {
		if n > 0 {
			e.concurrency = n
		}
	}
}

// NewOllamaEmbedder creates a new Ollama embedder.
// Default model: nomic-embed-text (768 dimensions)
// Default URL: http://localhost:11434
// Default concurrency: 8 (for parallel batch embedding)
func NewOllamaEmbedder(opts ...OllamaOption) *OllamaEmbedder {
	e := &OllamaEmbedder{
		baseURL:     "http://localhost:11434",
		model:       "nomic-embed-text",
		dim:         768,
		concurrency: 8,
		client:      &http.Client{Timeout: 60 * time.Second}, // Longer timeout for local inference
	}
	for _, opt := range opts {
		opt(e)
	}

	// Adjust dimensions based on model
	switch e.model {
	case "nomic-embed-text":
		e.dim = 768
	case "mxbai-embed-large":
		e.dim = 1024
	case "all-minilm":
		e.dim = 384
	}

	return e
}

// Dim returns the embedding dimension.
func (e *OllamaEmbedder) Dim() int {
	return e.dim
}

// Embed generates an embedding for a single text.
func (e *OllamaEmbedder) Embed(ctx context.Context, text string) ([]float32, error) {
	reqBody := ollamaEmbedRequest{
		Model:  e.model,
		Prompt: text,
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", e.baseURL+"/api/embeddings", bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := e.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed (is Ollama running at %s?): %w", e.baseURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Ollama API error (status %d)", resp.StatusCode)
	}

	var embResp ollamaEmbedResponse
	if err := json.NewDecoder(resp.Body).Decode(&embResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Convert float64 to float32
	result := make([]float32, len(embResp.Embedding))
	for i, v := range embResp.Embedding {
		result[i] = float32(v)
	}

	return result, nil
}

// EmbedBatch generates embeddings for multiple texts concurrently.
// Ollama doesn't support native batching, so we parallelize with a worker pool.
func (e *OllamaEmbedder) EmbedBatch(ctx context.Context, texts []string) ([][]float32, error) {
	if len(texts) == 0 {
		return nil, nil
	}

	results := make([][]float32, len(texts))
	var mu sync.Mutex
	var firstErr error

	// Semaphore channel for concurrency control
	sem := make(chan struct{}, e.concurrency)
	var wg sync.WaitGroup

	for i, text := range texts {
		// Check for cancellation or previous error
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		mu.Lock()
		if firstErr != nil {
			mu.Unlock()
			break
		}
		mu.Unlock()

		wg.Add(1)
		sem <- struct{}{} // Acquire semaphore

		go func(idx int, txt string) {
			defer wg.Done()
			defer func() { <-sem }() // Release semaphore

			vec, err := e.Embed(ctx, txt)

			mu.Lock()
			defer mu.Unlock()
			if err != nil && firstErr == nil {
				firstErr = fmt.Errorf("failed to embed text %d: %w", idx, err)
				return
			}
			if firstErr == nil {
				results[idx] = vec
			}
		}(i, text)
	}

	wg.Wait()

	if firstErr != nil {
		return nil, firstErr
	}
	return results, nil
}

// API types

type ollamaEmbedRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type ollamaEmbedResponse struct {
	Embedding []float64 `json:"embedding"`
}


