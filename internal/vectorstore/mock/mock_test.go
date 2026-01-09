package mock

import (
	"context"
	"testing"

	"github.com/metawake/ragtune/internal/vectorstore"
)

func TestStore_EnsureCollection(t *testing.T) {
	s := New()
	defer s.Close()

	ctx := context.Background()

	// Create collection
	err := s.EnsureCollection(ctx, "test", 3)
	if err != nil {
		t.Fatalf("EnsureCollection failed: %v", err)
	}

	// Idempotent - should not error
	err = s.EnsureCollection(ctx, "test", 3)
	if err != nil {
		t.Fatalf("EnsureCollection (idempotent) failed: %v", err)
	}
}

func TestStore_Upsert(t *testing.T) {
	s := New()
	defer s.Close()

	ctx := context.Background()

	// Setup
	err := s.EnsureCollection(ctx, "test", 3)
	if err != nil {
		t.Fatalf("EnsureCollection failed: %v", err)
	}

	// Upsert points
	points := []vectorstore.Point{
		{ID: "p1", Vector: []float32{1, 0, 0}, Payload: map[string]interface{}{"text": "first"}},
		{ID: "p2", Vector: []float32{0, 1, 0}, Payload: map[string]interface{}{"text": "second"}},
	}

	err = s.Upsert(ctx, "test", points)
	if err != nil {
		t.Fatalf("Upsert failed: %v", err)
	}

	// Verify count
	count, err := s.Count(ctx, "test")
	if err != nil {
		t.Fatalf("Count failed: %v", err)
	}
	if count != 2 {
		t.Errorf("expected count 2, got %d", count)
	}
}

func TestStore_UpsertDimensionMismatch(t *testing.T) {
	s := New()
	defer s.Close()

	ctx := context.Background()

	err := s.EnsureCollection(ctx, "test", 3)
	if err != nil {
		t.Fatalf("EnsureCollection failed: %v", err)
	}

	// Wrong dimension
	points := []vectorstore.Point{
		{ID: "p1", Vector: []float32{1, 0}, Payload: nil}, // 2D instead of 3D
	}

	err = s.Upsert(ctx, "test", points)
	if err == nil {
		t.Error("expected error for dimension mismatch, got nil")
	}
}

func TestStore_UpsertNonexistentCollection(t *testing.T) {
	s := New()
	defer s.Close()

	ctx := context.Background()

	points := []vectorstore.Point{
		{ID: "p1", Vector: []float32{1, 0, 0}, Payload: nil},
	}

	err := s.Upsert(ctx, "nonexistent", points)
	if err == nil {
		t.Error("expected error for nonexistent collection, got nil")
	}
}

func TestStore_Search(t *testing.T) {
	s := New()
	defer s.Close()

	ctx := context.Background()

	err := s.EnsureCollection(ctx, "test", 3)
	if err != nil {
		t.Fatalf("EnsureCollection failed: %v", err)
	}

	// Insert orthogonal vectors for easy verification
	points := []vectorstore.Point{
		{ID: "x", Vector: []float32{1, 0, 0}, Payload: map[string]interface{}{"axis": "x"}},
		{ID: "y", Vector: []float32{0, 1, 0}, Payload: map[string]interface{}{"axis": "y"}},
		{ID: "z", Vector: []float32{0, 0, 1}, Payload: map[string]interface{}{"axis": "z"}},
	}

	err = s.Upsert(ctx, "test", points)
	if err != nil {
		t.Fatalf("Upsert failed: %v", err)
	}

	// Search for x-axis vector
	results, err := s.Search(ctx, "test", []float32{1, 0, 0}, 3)
	if err != nil {
		t.Fatalf("Search failed: %v", err)
	}

	if len(results) != 3 {
		t.Fatalf("expected 3 results, got %d", len(results))
	}

	// First result should be x (score = 1.0)
	if results[0].ID != "x" {
		t.Errorf("expected first result to be 'x', got %q", results[0].ID)
	}
	if results[0].Score < 0.99 {
		t.Errorf("expected score ~1.0, got %f", results[0].Score)
	}

	// Other results should have score 0 (orthogonal)
	if results[1].Score > 0.01 {
		t.Errorf("expected orthogonal vectors to have score ~0, got %f", results[1].Score)
	}
}

func TestStore_SearchTopK(t *testing.T) {
	s := New()
	defer s.Close()

	ctx := context.Background()

	err := s.EnsureCollection(ctx, "test", 2)
	if err != nil {
		t.Fatalf("EnsureCollection failed: %v", err)
	}

	// Insert 5 points
	for i := 0; i < 5; i++ {
		points := []vectorstore.Point{
			{ID: string(rune('a' + i)), Vector: []float32{float32(i), float32(i)}, Payload: nil},
		}
		s.Upsert(ctx, "test", points)
	}

	// Search with top-k = 2
	results, err := s.Search(ctx, "test", []float32{4, 4}, 2)
	if err != nil {
		t.Fatalf("Search failed: %v", err)
	}

	if len(results) != 2 {
		t.Errorf("expected 2 results, got %d", len(results))
	}
}

func TestStore_Count(t *testing.T) {
	s := New()
	defer s.Close()

	ctx := context.Background()

	err := s.EnsureCollection(ctx, "test", 2)
	if err != nil {
		t.Fatalf("EnsureCollection failed: %v", err)
	}

	// Initially empty
	count, err := s.Count(ctx, "test")
	if err != nil {
		t.Fatalf("Count failed: %v", err)
	}
	if count != 0 {
		t.Errorf("expected count 0, got %d", count)
	}

	// Add points
	for i := 0; i < 10; i++ {
		points := []vectorstore.Point{
			{ID: string(rune('a' + i)), Vector: []float32{1, 1}, Payload: nil},
		}
		s.Upsert(ctx, "test", points)
	}

	count, err = s.Count(ctx, "test")
	if err != nil {
		t.Fatalf("Count failed: %v", err)
	}
	if count != 10 {
		t.Errorf("expected count 10, got %d", count)
	}
}

func TestStore_Close(t *testing.T) {
	s := New()

	ctx := context.Background()

	err := s.EnsureCollection(ctx, "test", 2)
	if err != nil {
		t.Fatalf("EnsureCollection failed: %v", err)
	}

	// Close
	err = s.Close()
	if err != nil {
		t.Fatalf("Close failed: %v", err)
	}

	// Operations should fail after close
	err = s.EnsureCollection(ctx, "test2", 2)
	if err == nil {
		t.Error("expected error after close, got nil")
	}
}

func TestStore_CustomHook(t *testing.T) {
	s := New()
	defer s.Close()

	called := false
	s.EnsureCollectionFunc = func(ctx context.Context, name string, dim int) error {
		called = true
		return nil
	}

	ctx := context.Background()
	s.EnsureCollection(ctx, "test", 2)

	if !called {
		t.Error("custom hook was not called")
	}
}

func TestCosineSimilarity(t *testing.T) {
	tests := []struct {
		name     string
		a, b     []float32
		expected float32
		delta    float32
	}{
		{"identical", []float32{1, 0, 0}, []float32{1, 0, 0}, 1.0, 0.01},
		{"orthogonal", []float32{1, 0, 0}, []float32{0, 1, 0}, 0.0, 0.01},
		{"opposite", []float32{1, 0, 0}, []float32{-1, 0, 0}, -1.0, 0.01},
		{"similar", []float32{1, 1, 0}, []float32{1, 0, 0}, 0.707, 0.01},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cosineSimilarity(tt.a, tt.b)
			if got < tt.expected-tt.delta || got > tt.expected+tt.delta {
				t.Errorf("cosineSimilarity(%v, %v) = %f, want %f±%f", tt.a, tt.b, got, tt.expected, tt.delta)
			}
		})
	}
}



