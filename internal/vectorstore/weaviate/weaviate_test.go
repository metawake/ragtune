//go:build integration

package weaviate

import (
	"context"
	"os"
	"testing"

	"github.com/metawake/ragtune/internal/vectorstore"
)

// TestWeaviateIntegration tests the full CRUD cycle.
// Run with: go test -tags=integration ./internal/vectorstore/weaviate/... -v
//
// Requires: docker run -d -p 8080:8080 semitechnologies/weaviate:latest
func TestWeaviateIntegration(t *testing.T) {
	host := os.Getenv("WEAVIATE_HOST")
	if host == "" {
		host = "localhost:8080"
	}

	ctx := context.Background()

	client, err := New(ctx, host, "http")
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
				ID:      "00000000-0000-4000-8000-000000000001",
				Vector:  []float32{1.0, 0.0, 0.0, 0.0},
				Payload: map[string]interface{}{"source": "test1.md", "text": "hello"},
			},
			{
				ID:      "00000000-0000-4000-8000-000000000002",
				Vector:  []float32{0.0, 1.0, 0.0, 0.0},
				Payload: map[string]interface{}{"source": "test2.md", "text": "world"},
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

		// First result should be the exact match
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

	t.Log("✓ Weaviate integration test passed")
}

func TestClassName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"demo", "Ragtune_Demo"},
		{"my-collection", "Ragtune_My_collection"},
		{"test_123", "Ragtune_Test_123"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := className(tt.input)
			if got != tt.expected {
				t.Errorf("className(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}
