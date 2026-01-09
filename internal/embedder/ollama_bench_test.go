package embedder

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"
)

// TestOllamaBatchPerformance benchmarks concurrent vs sequential embedding.
// Run with: go test -v -run TestOllamaBatchPerformance ./internal/embedder/
// Requires Ollama running locally with nomic-embed-text model.
func TestOllamaBatchPerformance(t *testing.T) {
	// Skip if not running integration tests
	if os.Getenv("INTEGRATION_TEST") == "" {
		t.Skip("Skipping integration test (set INTEGRATION_TEST=1 to run)")
	}

	ctx := context.Background()
	batchSizes := []int{10, 50, 100}
	concurrencyLevels := []int{1, 4, 8}

	// Generate test texts
	texts := make([]string, 100)
	for i := range texts {
		texts[i] = fmt.Sprintf("This is test document number %d with some content to embed for benchmarking purposes.", i)
	}

	fmt.Println("\n=== Ollama Batch Embedding Performance ===")
	fmt.Println("Model: nomic-embed-text")
	fmt.Printf("%-12s %-12s %-12s %-12s\n", "Batch Size", "Concurrency", "Time", "Per Item")
	fmt.Println("------------------------------------------------")

	for _, batchSize := range batchSizes {
		batch := texts[:batchSize]

		for _, concurrency := range concurrencyLevels {
			e := NewOllamaEmbedder(WithOllamaConcurrency(concurrency))

			start := time.Now()
			results, err := e.EmbedBatch(ctx, batch)
			elapsed := time.Since(start)

			if err != nil {
				t.Logf("Error with batch=%d, concurrency=%d: %v", batchSize, concurrency, err)
				continue
			}

			if len(results) != batchSize {
				t.Errorf("Expected %d results, got %d", batchSize, len(results))
			}

			perItem := elapsed / time.Duration(batchSize)
			fmt.Printf("%-12d %-12d %-12s %-12s\n", batchSize, concurrency, elapsed.Round(time.Millisecond), perItem.Round(time.Millisecond))
		}
	}

	fmt.Println("\n✓ Benchmark complete")
}

// BenchmarkOllamaEmbedBatch provides standard Go benchmarks.
// Run with: go test -bench=BenchmarkOllamaEmbedBatch ./internal/embedder/ -benchtime=10s
func BenchmarkOllamaEmbedBatch(b *testing.B) {
	if os.Getenv("INTEGRATION_TEST") == "" {
		b.Skip("Skipping integration benchmark (set INTEGRATION_TEST=1 to run)")
	}

	texts := make([]string, 20)
	for i := range texts {
		texts[i] = fmt.Sprintf("Benchmark test document %d with content for embedding.", i)
	}

	ctx := context.Background()

	b.Run("concurrency=1", func(b *testing.B) {
		e := NewOllamaEmbedder(WithOllamaConcurrency(1))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = e.EmbedBatch(ctx, texts)
		}
	})

	b.Run("concurrency=4", func(b *testing.B) {
		e := NewOllamaEmbedder(WithOllamaConcurrency(4))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = e.EmbedBatch(ctx, texts)
		}
	})

	b.Run("concurrency=8", func(b *testing.B) {
		e := NewOllamaEmbedder(WithOllamaConcurrency(8))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = e.EmbedBatch(ctx, texts)
		}
	})
}



