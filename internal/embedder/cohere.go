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

// Compile-time interface compliance check.
var _ Embedder = (*CohereEmbedder)(nil)

// CohereEmbedder uses Cohere's embedding API.
type CohereEmbedder struct {
	apiKey    string
	model     string
	dim       int
	inputType string // "search_document" or "search_query"
	client    *http.Client
}

// CohereOption configures the Cohere embedder.
type CohereOption func(*CohereEmbedder)

// WithCohereModel sets the Cohere embedding model.
func WithCohereModel(model string) CohereOption {
	return func(e *CohereEmbedder) {
		e.model = model
	}
}

// WithCohereInputType sets the input type for embeddings.
// Use "search_document" for corpus documents, "search_query" for queries.
func WithCohereInputType(inputType string) CohereOption {
	return func(e *CohereEmbedder) {
		e.inputType = inputType
	}
}

// NewCohereEmbedder creates a new Cohere embedder.
// Uses COHERE_API_KEY environment variable.
// Default model: embed-english-v3.0 (1024 dimensions)
func NewCohereEmbedder(opts ...CohereOption) *CohereEmbedder {
	e := &CohereEmbedder{
		apiKey:    os.Getenv("COHERE_API_KEY"),
		model:     "embed-english-v3.0",
		dim:       1024,
		inputType: "search_document",
		client:    &http.Client{Timeout: 30 * time.Second},
	}
	for _, opt := range opts {
		opt(e)
	}

	// Adjust dimensions based on model
	switch e.model {
	case "embed-english-v3.0", "embed-multilingual-v3.0":
		e.dim = 1024
	case "embed-english-light-v3.0", "embed-multilingual-light-v3.0":
		e.dim = 384
	}

	return e
}

// Dim returns the embedding dimension.
func (e *CohereEmbedder) Dim() int {
	return e.dim
}

// Embed generates an embedding for a single text.
func (e *CohereEmbedder) Embed(ctx context.Context, text string) ([]float32, error) {
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
func (e *CohereEmbedder) EmbedBatch(ctx context.Context, texts []string) ([][]float32, error) {
	if e.apiKey == "" {
		return nil, fmt.Errorf("COHERE_API_KEY environment variable not set")
	}

	reqBody := cohereEmbedRequest{
		Model:     e.model,
		Texts:     texts,
		InputType: e.inputType,
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", "https://api.cohere.ai/v1/embed", bytes.NewReader(bodyBytes))
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
		var errResp cohereErrorResponse
		json.NewDecoder(resp.Body).Decode(&errResp)
		return nil, fmt.Errorf("Cohere API error (status %d): %s", resp.StatusCode, errResp.Message)
	}

	var embResp cohereEmbedResponse
	if err := json.NewDecoder(resp.Body).Decode(&embResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Convert float64 to float32
	embeddings := make([][]float32, len(embResp.Embeddings))
	for i, emb := range embResp.Embeddings {
		embeddings[i] = make([]float32, len(emb))
		for j, v := range emb {
			embeddings[i][j] = float32(v)
		}
	}

	return embeddings, nil
}

// API types

type cohereEmbedRequest struct {
	Model     string   `json:"model"`
	Texts     []string `json:"texts"`
	InputType string   `json:"input_type"`
}

type cohereEmbedResponse struct {
	Embeddings [][]float64 `json:"embeddings"`
}

type cohereErrorResponse struct {
	Message string `json:"message"`
}



