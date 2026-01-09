//go:build integration

package pgvector

import (
	"context"
	"os"
	"testing"

	"github.com/metawake/ragtune/internal/vectorstore"
)

// TestPgvectorIntegration tests the full CRUD cycle against a real PostgreSQL instance.
// Run with: go test -tags=integration ./internal/vectorstore/pgvector/...
//
// Requires pgvector container:
//   docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=test pgvector/pgvector:pg16
func TestPgvectorIntegration(t *testing.T) {
	connStr := os.Getenv("PGVECTOR_URL")
	if connStr == "" {
		connStr = "postgres://postgres:test@localhost:5432/postgres?sslmode=disable"
	}

	ctx := context.Background()

	// Connect
	client, err := New(ctx, connStr)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer client.Close()

	collection := "test_integration"
	dim := 4 // Small dimension for testing

	// Cleanup from previous runs
	_ = client.DeleteCollection(ctx, collection)

	// Test EnsureCollection
	t.Run("EnsureCollection", func(t *testing.T) {
		err := client.EnsureCollection(ctx, collection, dim)
		if err != nil {
			t.Fatalf("EnsureCollection failed: %v", err)
		}

		// Idempotent - should succeed again
		err = client.EnsureCollection(ctx, collection, dim)
		if err != nil {
			t.Fatalf("EnsureCollection (idempotent) failed: %v", err)
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
			{
				ID:      "doc3",
				Vector:  []float32{0.9, 0.1, 0.0, 0.0}, // Similar to doc1
				Payload: map[string]interface{}{"source": "test3.md", "text": "hello again"},
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
		if count != 3 {
			t.Errorf("Count = %d, want 3", count)
		}
	})

	// Test Search
	t.Run("Search", func(t *testing.T) {
		// Search for vector similar to doc1
		query := []float32{1.0, 0.0, 0.0, 0.0}
		results, err := client.Search(ctx, collection, query, 3)
		if err != nil {
			t.Fatalf("Search failed: %v", err)
		}

		if len(results) != 3 {
			t.Fatalf("Expected 3 results, got %d", len(results))
		}

		// doc1 should be first (exact match)
		if results[0].ID != "doc1" {
			t.Errorf("Expected doc1 first, got %s", results[0].ID)
		}

		// Score should be ~1.0 for exact match
		if results[0].Score < 0.99 {
			t.Errorf("Expected score ~1.0 for exact match, got %f", results[0].Score)
		}

		// doc3 should be second (similar to doc1)
		if results[1].ID != "doc3" {
			t.Errorf("Expected doc3 second, got %s", results[1].ID)
		}

		// doc2 should be last (orthogonal)
		if results[2].ID != "doc2" {
			t.Errorf("Expected doc2 last, got %s", results[2].ID)
		}

		// Verify payload is returned
		if source, ok := results[0].Payload["source"].(string); !ok || source != "test1.md" {
			t.Errorf("Payload source = %v, want test1.md", results[0].Payload["source"])
		}
	})

	// Test Upsert (update existing)
	t.Run("Upsert_Update", func(t *testing.T) {
		points := []vectorstore.Point{
			{
				ID:      "doc1",
				Vector:  []float32{0.0, 0.0, 1.0, 0.0}, // Changed vector
				Payload: map[string]interface{}{"source": "test1-updated.md", "text": "updated"},
			},
		}

		err := client.Upsert(ctx, collection, points)
		if err != nil {
			t.Fatalf("Upsert (update) failed: %v", err)
		}

		// Count should still be 3
		count, err := client.Count(ctx, collection)
		if err != nil {
			t.Fatalf("Count failed: %v", err)
		}
		if count != 3 {
			t.Errorf("Count after update = %d, want 3", count)
		}

		// Search should reflect update
		query := []float32{0.0, 0.0, 1.0, 0.0}
		results, err := client.Search(ctx, collection, query, 1)
		if err != nil {
			t.Fatalf("Search failed: %v", err)
		}
		if results[0].ID != "doc1" {
			t.Errorf("Expected updated doc1 to match, got %s", results[0].ID)
		}
	})

	// Test DeleteCollection
	t.Run("DeleteCollection", func(t *testing.T) {
		err := client.DeleteCollection(ctx, collection)
		if err != nil {
			t.Fatalf("DeleteCollection failed: %v", err)
		}

		// Search should fail now
		query := []float32{1.0, 0.0, 0.0, 0.0}
		_, err = client.Search(ctx, collection, query, 1)
		if err == nil {
			t.Error("Expected error searching deleted collection")
		}
	})
}
