package embedder

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// TEIEmbedder uses Hugging Face Text Embeddings Inference server.
// TEI provides native batching and is optimized for high-throughput embedding.
// Run with: docker run -p 8080:80 ghcr.io/huggingface/text-embeddings-inference:latest --model-id BAAI/bge-base-en-v1.5
type TEIEmbedder struct {
	baseURL     string
	model       string // For dimension lookup; TEI model is set at server start
	dim         int
	dimExplicit bool // True if dim was explicitly set via WithTEIDim
	client      *http.Client
}

// TEIOption configures the TEI embedder.
type TEIOption func(*TEIEmbedder)

// WithTEIURL sets the TEI server URL.
func WithTEIURL(url string) TEIOption {
	return func(e *TEIEmbedder) {
		e.baseURL = url
	}
}

// WithTEIModel sets the model name for dimension inference.
// Note: The actual model is configured when starting the TEI server.
func WithTEIModel(model string) TEIOption {
	return func(e *TEIEmbedder) {
		e.model = model
	}
}

// WithTEIDim explicitly sets the embedding dimension.
// Use this if your model isn't in the auto-detect list.
// Explicit dimension takes precedence over model-based auto-detection.
func WithTEIDim(dim int) TEIOption {
	return func(e *TEIEmbedder) {
		if dim > 0 {
			e.dim = dim
			e.dimExplicit = true
		}
	}
}

// NewTEIEmbedder creates a new Hugging Face TEI embedder.
// Default URL: http://localhost:8080
// Default model: BAAI/bge-base-en-v1.5 (768 dimensions)
func NewTEIEmbedder(opts ...TEIOption) *TEIEmbedder {
	e := &TEIEmbedder{
		baseURL: "http://localhost:8080",
		model:   "BAAI/bge-base-en-v1.5",
		dim:     768,
		client:  &http.Client{},
	}
	for _, opt := range opts {
		opt(e)
	}

	// Auto-detect dimensions based on known models (skip if explicitly set)
	if !e.dimExplicit {
		switch e.model {
		// BGE models
		case "BAAI/bge-small-en-v1.5", "BAAI/bge-small-en":
			e.dim = 384
		case "BAAI/bge-base-en-v1.5", "BAAI/bge-base-en":
			e.dim = 768
		case "BAAI/bge-large-en-v1.5", "BAAI/bge-large-en":
			e.dim = 1024
		// Sentence Transformers
		case "sentence-transformers/all-MiniLM-L6-v2":
			e.dim = 384
		case "sentence-transformers/all-mpnet-base-v2":
			e.dim = 768
		// Nomic
		case "nomic-ai/nomic-embed-text-v1.5", "nomic-ai/nomic-embed-text-v1":
			e.dim = 768
		// GTE models
		case "thenlper/gte-small":
			e.dim = 384
		case "thenlper/gte-base":
			e.dim = 768
		case "thenlper/gte-large":
			e.dim = 1024
		case "Alibaba-NLP/gte-Qwen2-1.5B-instruct":
			e.dim = 1536
		// E5 models
		case "intfloat/e5-small-v2":
			e.dim = 384
		case "intfloat/e5-base-v2":
			e.dim = 768
		case "intfloat/e5-large-v2":
			e.dim = 1024
		}
	}

	return e
}

// Dim returns the embedding dimension.
func (e *TEIEmbedder) Dim() int {
	return e.dim
}

// Embed generates an embedding for a single text.
func (e *TEIEmbedder) Embed(ctx context.Context, text string) ([]float32, error) {
	embeddings, err := e.EmbedBatch(ctx, []string{text})
	if err != nil {
		return nil, err
	}
	if len(embeddings) == 0 {
		return nil, fmt.Errorf("no embeddings returned")
	}
	return embeddings[0], nil
}

// EmbedBatch generates embeddings for multiple texts.
// TEI supports native batching for high throughput.
func (e *TEIEmbedder) EmbedBatch(ctx context.Context, texts []string) ([][]float32, error) {
	if len(texts) == 0 {
		return nil, nil
	}

	reqBody := teiEmbedRequest{
		Inputs: texts,
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", e.baseURL+"/embed", bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := e.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed (is TEI running at %s?): %w", e.baseURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("TEI API error (status %d): %s", resp.StatusCode, string(body))
	}

	// TEI returns array of arrays: [[0.1, 0.2, ...], [0.3, 0.4, ...]]
	var embeddings [][]float64
	if err := json.NewDecoder(resp.Body).Decode(&embeddings); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Convert float64 to float32
	results := make([][]float32, len(embeddings))
	for i, emb := range embeddings {
		results[i] = make([]float32, len(emb))
		for j, v := range emb {
			results[i][j] = float32(v)
		}
	}

	return results, nil
}

// API types

type teiEmbedRequest struct {
	Inputs []string `json:"inputs"`
}

