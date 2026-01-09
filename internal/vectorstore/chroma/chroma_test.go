//go:build integration

package chroma

import (
	"context"
	"os"
	"testing"

	"github.com/metawake/ragtune/internal/vectorstore"
)

// TestChromaIntegration tests the full CRUD cycle.
// Run with: go test -tags=integration ./internal/vectorstore/chroma/... -v
//
// Requires: docker run -d -p 8000:8000 chromadb/chroma:latest
func TestChromaIntegration(t *testing.T) {
	url := os.Getenv("CHROMA_URL")
	if url == "" {
		url = "http://localhost:8000"
	}

	ctx := context.Background()

	client, err := New(ctx, url)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer client.Close()

	collection := "integration_test"
	dim := 4

	// Cleanup
	_ = client.DeleteCollection(ctx, collection)

	// Test EnsureCollection
	t.Run("EnsureCollection", func(t *testing.T) {
		err := client.EnsureCollection(ctx, collection, dim)
		if err != nil {
			t.Fatalf("EnsureCollection failed: %v", err)
		}
	})

	// Test Upsert
	t.Run("Upsert", func(t *testing.T) {
		points := []vectorstore.Point{
			{
				ID:      "doc1",
				Vector:  []float32{1.0, 0.0, 0.0, 0.0},
				Payload: map[string]interface{}{"source": "test1.md", "text": "hello world"},
			},
			{
				ID:      "doc2",
				Vector:  []float32{0.0, 1.0, 0.0, 0.0},
				Payload: map[string]interface{}{"source": "test2.md", "text": "goodbye world"},
			},
		}

		err := client.Upsert(ctx, collection, points)
		if err != nil {
			t.Fatalf("Upsert failed: %v", err)
		}
	})

	// Test Count
	t.Run("Count", func(t *testing.T) {
		count, err := client.Count(ctx, collection)
		if err != nil {
			t.Fatalf("Count failed: %v", err)
		}
		if count != 2 {
			t.Errorf("Count = %d, want 2", count)
		}
	})

	// Test Search
	t.Run("Search", func(t *testing.T) {
		results, err := client.Search(ctx, collection, []float32{1.0, 0.0, 0.0, 0.0}, 2)
		if err != nil {
			t.Fatalf("Search failed: %v", err)
		}

		if len(results) != 2 {
			t.Fatalf("Expected 2 results, got %d", len(results))
		}

		// First result should be doc1 (exact match)
		if results[0].ID != "doc1" {
			t.Errorf("Expected doc1 first, got %s", results[0].ID)
		}

		// Score should be high for exact match
		if results[0].Score < 0.99 {
			t.Errorf("Expected high score for exact match, got %f", results[0].Score)
		}
	})

	// Cleanup
	t.Run("DeleteCollection", func(t *testing.T) {
		err := client.DeleteCollection(ctx, collection)
		if err != nil {
			t.Fatalf("DeleteCollection failed: %v", err)
		}
	})

	t.Log("✓ Chroma integration test passed")
}
