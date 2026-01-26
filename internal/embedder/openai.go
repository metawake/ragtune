package embedder

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

const (
	// defaultHTTPTimeout is the timeout for HTTP requests to embedding APIs.
	// 30 seconds allows for large batch requests while preventing indefinite hangs.
	defaultHTTPTimeout = 30 * time.Second
)

// Compile-time interface compliance check.
var _ Embedder = (*OpenAIEmbedder)(nil)

// OpenAIEmbedder uses OpenAI's embedding API.
type OpenAIEmbedder struct {
	apiKey  string
	baseURL string
	model   string
	dim     int
	client  *http.Client
}

// OpenAIOption configures the OpenAI embedder.
type OpenAIOption func(*OpenAIEmbedder)

// WithOpenAIURL sets a custom API URL (for testing or proxies).
func WithOpenAIURL(url string) OpenAIOption {
	return func(e *OpenAIEmbedder) {
		e.baseURL = url
	}
}

// NewOpenAIEmbedder creates a new OpenAI embedder.
// Uses OPENAI_API_KEY environment variable.
func NewOpenAIEmbedder(opts ...OpenAIOption) *OpenAIEmbedder {
	e := &OpenAIEmbedder{
		apiKey:  os.Getenv("OPENAI_API_KEY"),
		baseURL: "https://api.openai.com/v1/embeddings",
		model:   "text-embedding-3-small",
		dim:     1536,
		client:  &http.Client{Timeout: defaultHTTPTimeout},
	}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

// Dim returns the embedding dimension.
func (e *OpenAIEmbedder) Dim() int {
	return e.dim
}

// Embed generates an embedding for a single text.
func (e *OpenAIEmbedder) Embed(ctx context.Context, text string) ([]float32, error) {
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
func (e *OpenAIEmbedder) EmbedBatch(ctx context.Context, texts []string) ([][]float32, error) {
	if e.apiKey == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY environment variable not set")
	}

	reqBody := openaiEmbeddingRequest{
		Model: e.model,
		Input: texts,
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", e.baseURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+e.apiKey)

	resp, err := e.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errResp openaiErrorResponse
		_ = json.NewDecoder(resp.Body).Decode(&errResp)
		return nil, fmt.Errorf("API error (status %d): %s", resp.StatusCode, errResp.Error.Message)
	}

	var embResp openaiEmbeddingResponse
	if err := json.NewDecoder(resp.Body).Decode(&embResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Sort by index and extract embeddings
	embeddings := make([][]float32, len(texts))
	for _, data := range embResp.Data {
		if data.Index < len(embeddings) {
			embeddings[data.Index] = data.Embedding
		}
	}

	return embeddings, nil
}

// API types

type openaiEmbeddingRequest struct {
	Model string   `json:"model"`
	Input []string `json:"input"`
}

type openaiEmbeddingResponse struct {
	Data []openaiEmbeddingData `json:"data"`
}

type openaiEmbeddingData struct {
	Index     int       `json:"index"`
	Embedding []float32 `json:"embedding"`
}

type openaiErrorResponse struct {
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}




