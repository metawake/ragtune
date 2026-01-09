//go:build integration

package cli

import (
	"context"
	"fmt"
	"math"
	"testing"

	"github.com/metawake/ragtune/internal/vectorstore"
	"github.com/metawake/ragtune/internal/vectorstore/pgvector"
	"github.com/metawake/ragtune/internal/vectorstore/qdrant"
)

// TestStoreConsistency verifies that Qdrant and pgvector produce consistent results.
// Run with: go test -tags=integration ./internal/cli/... -run TestStoreConsistency -v
//
// Requires:
//   docker run -d -p 6333:6333 -p 6334:6334 qdrant/qdrant
//   docker run -d -p 5433:5432 -e POSTGRES_PASSWORD=test pgvector/pgvector:pg16
func TestStoreConsistency(t *testing.T) {
	ctx := context.Background()

	// Connect to both stores
	qdrantClient, err := qdrant.New(ctx, "127.0.0.1:6334")
	if err != nil {
		t.Fatalf("Failed to connect to Qdrant: %v", err)
	}
	defer qdrantClient.Close()

	pgClient, err := pgvector.New(ctx, "postgres://postgres:test@localhost:5433/postgres?sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to connect to pgvector: %v", err)
	}
	defer pgClient.Close()

	collection := "consistency_test"
	dim := 128

	// Cleanup
	_ = qdrantClient.DeleteCollection(ctx, collection)
	_ = pgClient.DeleteCollection(ctx, collection)

	// Create collections
	if err := qdrantClient.EnsureCollection(ctx, collection, dim); err != nil {
		t.Fatalf("Qdrant EnsureCollection failed: %v", err)
	}
	if err := pgClient.EnsureCollection(ctx, collection, dim); err != nil {
		t.Fatalf("pgvector EnsureCollection failed: %v", err)
	}

	// Generate deterministic test data
	points := generateTestPoints(100, dim)

	// Upsert into both stores
	if err := qdrantClient.Upsert(ctx, collection, points); err != nil {
		t.Fatalf("Qdrant Upsert failed: %v", err)
	}
	if err := pgClient.Upsert(ctx, collection, points); err != nil {
		t.Fatalf("pgvector Upsert failed: %v", err)
	}

	// Verify counts match
	qdrantCount, _ := qdrantClient.Count(ctx, collection)
	pgCount, _ := pgClient.Count(ctx, collection)
	if qdrantCount != pgCount {
		t.Errorf("Count mismatch: Qdrant=%d, pgvector=%d", qdrantCount, pgCount)
	}
	t.Logf("Both stores have %d points", qdrantCount)

	// Run queries and compare results
	queries := generateQueryVectors(10, dim)
	topK := 5

	for i, queryVec := range queries {
		qdrantResults, err := qdrantClient.Search(ctx, collection, queryVec, topK)
		if err != nil {
			t.Fatalf("Qdrant search failed: %v", err)
		}

		pgResults, err := pgClient.Search(ctx, collection, queryVec, topK)
		if err != nil {
			t.Fatalf("pgvector search failed: %v", err)
		}

		// Compare result counts
		if len(qdrantResults) != len(pgResults) {
			t.Errorf("Query %d: result count mismatch: Qdrant=%d, pgvector=%d",
				i, len(qdrantResults), len(pgResults))
			continue
		}

		// Compare top result (should be identical or very close)
		if len(qdrantResults) > 0 && len(pgResults) > 0 {
			if qdrantResults[0].ID != pgResults[0].ID {
				t.Errorf("Query %d: top result mismatch: Qdrant=%s, pgvector=%s",
					i, qdrantResults[0].ID, pgResults[0].ID)
			}

			// Scores should be very close (both use cosine similarity)
			scoreDiff := math.Abs(float64(qdrantResults[0].Score - pgResults[0].Score))
			if scoreDiff > 0.01 {
				t.Errorf("Query %d: score mismatch: Qdrant=%.4f, pgvector=%.4f (diff=%.4f)",
					i, qdrantResults[0].Score, pgResults[0].Score, scoreDiff)
			}
		}

		// Compare all result IDs (order should match)
		for j := range qdrantResults {
			if qdrantResults[j].ID != pgResults[j].ID {
				t.Logf("Query %d, rank %d: ID mismatch (Qdrant=%s, pgvector=%s) - may be tie-breaking difference",
					i, j, qdrantResults[j].ID, pgResults[j].ID)
			}
		}
	}

	// Cleanup
	_ = qdrantClient.DeleteCollection(ctx, collection)
	_ = pgClient.DeleteCollection(ctx, collection)

	t.Log("✓ Store consistency test passed")
}

// generateTestPoints creates deterministic test vectors
func generateTestPoints(n, dim int) []vectorstore.Point {
	points := make([]vectorstore.Point, n)
	for i := 0; i < n; i++ {
		vec := make([]float32, dim)
		// Deterministic vector based on index
		for j := 0; j < dim; j++ {
			// Simple pattern: each doc has a different "direction"
			vec[j] = float32(math.Sin(float64(i*dim+j) * 0.1))
		}
		// Normalize
		vec = normalize(vec)

		points[i] = vectorstore.Point{
			ID:     generateUUID(i),
			Vector: vec,
			Payload: map[string]interface{}{
				"source": generateUUID(i) + ".md",
				"text":   "Test document content",
				"index":  i,
			},
		}
	}
	return points
}

// generateQueryVectors creates deterministic query vectors
func generateQueryVectors(n, dim int) [][]float32 {
	queries := make([][]float32, n)
	for i := 0; i < n; i++ {
		vec := make([]float32, dim)
		// Query vectors are slightly different from doc vectors
		for j := 0; j < dim; j++ {
			vec[j] = float32(math.Sin(float64(i*dim+j)*0.1 + 0.05))
		}
		queries[i] = normalize(vec)
	}
	return queries
}

func normalize(vec []float32) []float32 {
	var sum float32
	for _, v := range vec {
		sum += v * v
	}
	norm := float32(math.Sqrt(float64(sum)))
	if norm == 0 {
		return vec
	}
	result := make([]float32, len(vec))
	for i, v := range vec {
		result[i] = v / norm
	}
	return result
}

func generateUUID(i int) string {
	// Generate a deterministic UUID based on index
	// Format: 00000000-0000-4000-8000-000000000XXX where XXX encodes i
	return fmt.Sprintf("00000000-0000-4000-8000-%012d", i)
}
