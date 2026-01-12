// Package vectorstore defines the interface for vector store backends.
// RagTune is vector-store agnostic by design. This interface enables
// pluggable backends (Qdrant, pgvector, Weaviate, etc.).
package vectorstore

import (
	"context"
	"errors"
)

// Sentinel errors for vector store operations.
var (
	// ErrCollectionNotFound is returned when the requested collection doesn't exist.
	ErrCollectionNotFound = errors.New("collection not found")

	// ErrConnectionFailed is returned when unable to connect to the vector store.
	ErrConnectionFailed = errors.New("connection failed")

	// ErrDimensionMismatch is returned when vector dimensions don't match the collection.
	ErrDimensionMismatch = errors.New("vector dimension mismatch")
)

// Store defines the minimal interface for vector store operations.
// Implementations should handle connection management internally.
type Store interface {
	// EnsureCollection creates a collection if it doesn't exist.
	// dim specifies the vector dimension.
	EnsureCollection(ctx context.Context, name string, dim int) error

	// Upsert inserts or updates points in a collection.
	Upsert(ctx context.Context, collection string, points []Point) error

	// Search performs similarity search and returns top-k results.
	Search(ctx context.Context, collection string, vector []float32, topK int) ([]Result, error)

	// Count returns the number of points in a collection (best-effort).
	Count(ctx context.Context, collection string) (int64, error)

	// DeleteCollection removes a collection and all its data.
	DeleteCollection(ctx context.Context, name string) error

	// Close releases any resources held by the store.
	Close() error
}

// Point represents a vector with metadata to be stored.
type Point struct {
	ID      string
	Vector  []float32
	Payload map[string]interface{}
}

// Result represents a search result with score.
type Result struct {
	ID      string
	Score   float32
	Payload map[string]interface{}
}

