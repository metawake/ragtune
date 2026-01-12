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
var _ Embedder = (*VoyageEmbedder)(nil)

// VoyageEmbedder uses Voyage AI's embedding API.
type VoyageEmbedder struct {
	apiKey    string
	model     string
	dim       int
	inputType string // "document" or "query"
	client    *http.Client
}

// VoyageOption configures the Voyage embedder.
type VoyageOption func(*VoyageEmbedder)

// WithVoyageModel sets the Voyage embedding model.
func WithVoyageModel(model string) VoyageOption {
	return func(e *VoyageEmbedder) {
		e.model = model
	}
}

// WithVoyageInputType sets the input type for embeddings.
// Use "document" for corpus documents, "query" for queries.
func WithVoyageInputType(inputType string) VoyageOption {
	return func(e *VoyageEmbedder) {
		e.inputType = inputType
	}
}

// NewVoyageEmbedder creates a new Voyage AI embedder.
// Uses VOYAGE_API_KEY environment variable.
// Default model: voyage-2 (1024 dimensions)
func NewVoyageEmbedder(opts ...VoyageOption) *VoyageEmbedder {
	e := &VoyageEmbedder{
		apiKey:    os.Getenv("VOYAGE_API_KEY"),
		model:     "voyage-2",
		dim:       1024,
		inputType: "document",
		client:    &http.Client{Timeout: 30 * time.Second},
	}
	for _, opt := range opts {
		opt(e)
	}

	// Adjust dimensions based on model
	// All current Voyage models use 1024 dimensions
	switch e.model {
	case "voyage-2", "voyage-large-2", "voyage-law-2", "voyage-code-2", "voyage-finance-2":
		e.dim = 1024
	case "voyage-lite-02-instruct":
		e.dim = 1024
	}

	return e
}

// Dim returns the embedding dimension.
func (e *VoyageEmbedder) Dim() int {
	return e.dim
}

// Embed generates an embedding for a single text.
func (e *VoyageEmbedder) Embed(ctx context.Context, text string) ([]float32, error) {
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
func (e *VoyageEmbedder) EmbedBatch(ctx context.Context, texts []string) ([][]float32, error) {
	if e.apiKey == "" {
		return nil, fmt.Errorf("VOYAGE_API_KEY environment variable not set")
	}

	reqBody := voyageEmbedRequest{
		Model:     e.model,
		Input:     texts,
		InputType: e.inputType,
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", "https://api.voyageai.com/v1/embeddings", bytes.NewReader(bodyBytes))
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
		var errResp voyageErrorResponse
		json.NewDecoder(resp.Body).Decode(&errResp)
		return nil, fmt.Errorf("Voyage API error (status %d): %s", resp.StatusCode, errResp.Detail)
	}

	var embResp voyageEmbedResponse
	if err := json.NewDecoder(resp.Body).Decode(&embResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Sort by index and extract embeddings
	embeddings := make([][]float32, len(texts))
	for _, data := range embResp.Data {
		if data.Index < len(embeddings) {
			// Convert float64 to float32
			embedding := make([]float32, len(data.Embedding))
			for j, v := range data.Embedding {
				embedding[j] = float32(v)
			}
			embeddings[data.Index] = embedding
		}
	}

	return embeddings, nil
}

// API types

type voyageEmbedRequest struct {
	Model     string   `json:"model"`
	Input     []string `json:"input"`
	InputType string   `json:"input_type,omitempty"`
}

type voyageEmbedResponse struct {
	Data []voyageEmbedData `json:"data"`
}

type voyageEmbedData struct {
	Index     int       `json:"index"`
	Embedding []float64 `json:"embedding"`
}

type voyageErrorResponse struct {
	Detail string `json:"detail"`
}



