// Package embedder provides embedding generation for RAG.
package embedder

import (
	"context"
)

// Embedder generates vector embeddings from text.
type Embedder interface {
	// Embed generates an embedding for a single text.
	Embed(ctx context.Context, text string) ([]float32, error)

	// EmbedBatch generates embeddings for multiple texts.
	EmbedBatch(ctx context.Context, texts []string) ([][]float32, error)

	// Dim returns the embedding dimension.
	Dim() int
}
