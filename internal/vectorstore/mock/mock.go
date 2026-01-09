// Package mock provides a mock implementation of vectorstore.Store for testing.
package mock

import (
	"context"
	"fmt"
	"sort"
	"sync"

	"github.com/metawake/ragtune/internal/vectorstore"
)

// Store is an in-memory mock implementation of vectorstore.Store.
// Useful for testing without a real vector database.
type Store struct {
	mu          sync.RWMutex
	collections map[string]*collection
	closed      bool

	// Hooks for testing behavior
	EnsureCollectionFunc func(ctx context.Context, name string, dim int) error
	UpsertFunc           func(ctx context.Context, collection string, points []vectorstore.Point) error
	SearchFunc           func(ctx context.Context, collection string, vector []float32, topK int) ([]vectorstore.Result, error)
}

type collection struct {
	dim    int
	points map[string]vectorstore.Point
}

// New creates a new mock store.
func New() *Store {
	return &Store{
		collections: make(map[string]*collection),
	}
}

// EnsureCollection creates a collection if it doesn't exist.
func (s *Store) EnsureCollection(ctx context.Context, name string, dim int) error {
	if s.EnsureCollectionFunc != nil {
		return s.EnsureCollectionFunc(ctx, name, dim)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.closed {
		return fmt.Errorf("store is closed")
	}

	if _, exists := s.collections[name]; !exists {
		s.collections[name] = &collection{
			dim:    dim,
			points: make(map[string]vectorstore.Point),
		}
	}
	return nil
}

// Upsert inserts or updates points in a collection.
func (s *Store) Upsert(ctx context.Context, collectionName string, points []vectorstore.Point) error {
	if s.UpsertFunc != nil {
		return s.UpsertFunc(ctx, collectionName, points)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.closed {
		return fmt.Errorf("store is closed")
	}

	coll, exists := s.collections[collectionName]
	if !exists {
		return fmt.Errorf("collection %q does not exist", collectionName)
	}

	for _, p := range points {
		if len(p.Vector) != coll.dim {
			return fmt.Errorf("vector dimension mismatch: expected %d, got %d", coll.dim, len(p.Vector))
		}
		coll.points[p.ID] = p
	}
	return nil
}

// Search performs similarity search using cosine similarity.
func (s *Store) Search(ctx context.Context, collectionName string, vector []float32, topK int) ([]vectorstore.Result, error) {
	if s.SearchFunc != nil {
		return s.SearchFunc(ctx, collectionName, vector, topK)
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.closed {
		return nil, fmt.Errorf("store is closed")
	}

	coll, exists := s.collections[collectionName]
	if !exists {
		return nil, fmt.Errorf("collection %q does not exist", collectionName)
	}

	// Calculate cosine similarity for all points
	type scored struct {
		id      string
		score   float32
		payload map[string]interface{}
	}

	var results []scored
	for _, p := range coll.points {
		score := cosineSimilarity(vector, p.Vector)
		results = append(results, scored{
			id:      p.ID,
			score:   score,
			payload: p.Payload,
		})
	}

	// Sort by score descending
	sort.Slice(results, func(i, j int) bool {
		return results[i].score > results[j].score
	})

	// Return top-k
	if topK > len(results) {
		topK = len(results)
	}

	out := make([]vectorstore.Result, topK)
	for i := 0; i < topK; i++ {
		out[i] = vectorstore.Result{
			ID:      results[i].id,
			Score:   results[i].score,
			Payload: results[i].payload,
		}
	}
	return out, nil
}

// Count returns the number of points in a collection.
func (s *Store) Count(ctx context.Context, collectionName string) (int64, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.closed {
		return 0, fmt.Errorf("store is closed")
	}

	coll, exists := s.collections[collectionName]
	if !exists {
		return 0, fmt.Errorf("collection %q does not exist", collectionName)
	}

	return int64(len(coll.points)), nil
}

// DeleteCollection removes a collection and all its data.
func (s *Store) DeleteCollection(ctx context.Context, name string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.closed {
		return fmt.Errorf("store is closed")
	}

	if _, exists := s.collections[name]; !exists {
		return fmt.Errorf("collection %q does not exist", name)
	}

	delete(s.collections, name)
	return nil
}

// Close marks the store as closed.
func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.closed = true
	return nil
}

// GetPoints returns all points in a collection (for test assertions).
func (s *Store) GetPoints(collectionName string) ([]vectorstore.Point, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	coll, exists := s.collections[collectionName]
	if !exists {
		return nil, fmt.Errorf("collection %q does not exist", collectionName)
	}

	points := make([]vectorstore.Point, 0, len(coll.points))
	for _, p := range coll.points {
		points = append(points, p)
	}
	return points, nil
}

// cosineSimilarity computes cosine similarity between two vectors.
func cosineSimilarity(a, b []float32) float32 {
	if len(a) != len(b) {
		return 0
	}

	var dotProduct, normA, normB float32
	for i := range a {
		dotProduct += a[i] * b[i]
		normA += a[i] * a[i]
		normB += b[i] * b[i]
	}

	if normA == 0 || normB == 0 {
		return 0
	}

	return dotProduct / (sqrt(normA) * sqrt(normB))
}

// sqrt computes square root for float32.
func sqrt(x float32) float32 {
	if x <= 0 {
		return 0
	}
	// Newton's method
	z := x / 2
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z
}

