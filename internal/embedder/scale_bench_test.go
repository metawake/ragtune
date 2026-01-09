package embedder

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"testing"
	"time"
)

// TestEnterpriseScaleBenchmark benchmarks embedding throughput on the 50k synthetic corpus.
// This tests real-world enterprise scenarios with actual document content.
//
// Run with:
//
//	INTEGRATION_TEST=1 go test -v -run TestEnterpriseScaleBenchmark ./internal/embedder/ -timeout 30m
//
// Prerequisites:
//   - Ollama running locally with nomic-embed-text model
//   - benchmarks/synthetic-50k/corpus/ populated (run prepare.py first)
func TestEnterpriseScaleBenchmark(t *testing.T) {
	if os.Getenv("INTEGRATION_TEST") == "" {
		t.Skip("Skipping integration test (set INTEGRATION_TEST=1 to run)")
	}

	// Find corpus directory
	corpusDir := findCorpusDir(t)
	if corpusDir == "" {
		t.Fatal("Could not find benchmarks/synthetic-50k/corpus directory")
	}

	// Load documents
	docs := loadCorpus(t, corpusDir)
	if len(docs) == 0 {
		t.Fatal("No documents found in corpus. Run: python benchmarks/synthetic-50k/prepare.py")
	}

	// Test configurations
	testSizes := []int{1000, 5000, 10000, 50000}
	concurrencyLevels := []int{1, 4, 8, 16}

	// Filter test sizes based on available docs
	var validSizes []int
	for _, size := range testSizes {
		if size <= len(docs) {
			validSizes = append(validSizes, size)
		}
	}

	fmt.Printf(`
╔══════════════════════════════════════════════════════════════════════════════╗
║  RagTune Enterprise Scale Benchmark                                          ║
║  Testing embedding throughput for enterprise viability                       ║
╠══════════════════════════════════════════════════════════════════════════════╣
║  Corpus Size:     %6d documents                                             ║
║  Embedder:        Ollama (nomic-embed-text)                                  ║
║  System:          %s (%d cores)                                       ║
╚══════════════════════════════════════════════════════════════════════════════╝

`, len(docs), runtime.GOOS, runtime.NumCPU())

	// Print header
	fmt.Printf("%-12s %-12s %-14s %-14s %-12s %-10s\n",
		"Docs", "Concurrency", "Total Time", "Per Doc", "Docs/Sec", "Status")
	fmt.Println(strings.Repeat("─", 80))

	// Track best results
	type result struct {
		docs        int
		concurrency int
		duration    time.Duration
		docsPerSec  float64
	}
	var results []result

	for _, numDocs := range validSizes {
		batch := extractTexts(docs[:numDocs])

		for _, concurrency := range concurrencyLevels {
			// Skip high concurrency for small batches (not meaningful)
			if numDocs < 1000 && concurrency > 4 {
				continue
			}

			e := NewOllamaEmbedder(WithOllamaConcurrency(concurrency))
			ctx := context.Background()

			start := time.Now()
			embeddings, err := e.EmbedBatch(ctx, batch)
			elapsed := time.Since(start)

			status := "✓"
			if err != nil {
				status = "✗ " + truncateError(err)
				fmt.Printf("%-12d %-12d %-14s %-14s %-12s %-10s\n",
					numDocs, concurrency, "-", "-", "-", status)
				continue
			}

			if len(embeddings) != numDocs {
				status = fmt.Sprintf("✗ got %d", len(embeddings))
				fmt.Printf("%-12d %-12d %-14s %-14s %-12s %-10s\n",
					numDocs, concurrency, "-", "-", "-", status)
				continue
			}

			perDoc := elapsed / time.Duration(numDocs)
			docsPerSec := float64(numDocs) / elapsed.Seconds()

			results = append(results, result{
				docs:        numDocs,
				concurrency: concurrency,
				duration:    elapsed,
				docsPerSec:  docsPerSec,
			})

			// Color code based on performance
			if docsPerSec >= 100 {
				status = "✓ FAST"
			} else if docsPerSec >= 50 {
				status = "✓ OK"
			} else {
				status = "✓ SLOW"
			}

			fmt.Printf("%-12d %-12d %-14s %-14s %-12.1f %-10s\n",
				numDocs, concurrency,
				formatDuration(elapsed),
				formatDuration(perDoc),
				docsPerSec,
				status)
		}
		fmt.Println()
	}

	// Summary
	if len(results) > 0 {
		// Find best result for 50k (or largest tested)
		var best50k *result
		for i := range results {
			if results[i].docs == 50000 {
				if best50k == nil || results[i].docsPerSec > best50k.docsPerSec {
					best50k = &results[i]
				}
			}
		}

		// Find overall best
		sort.Slice(results, func(i, j int) bool {
			return results[i].docsPerSec > results[j].docsPerSec
		})
		best := results[0]

		fmt.Println(strings.Repeat("═", 80))
		fmt.Println()
		fmt.Printf("Best Throughput: %.1f docs/sec (concurrency=%d)\n", best.docsPerSec, best.concurrency)

		if best50k != nil {
			fmt.Printf("50K Benchmark:   %.1f docs/sec in %s (concurrency=%d)\n",
				best50k.docsPerSec, formatDuration(best50k.duration), best50k.concurrency)

			// Enterprise viability check
			if best50k.duration < 5*time.Minute {
				fmt.Println("\n✅ ENTERPRISE READY: 50k documents embedded in < 5 minutes")
			} else if best50k.duration < 10*time.Minute {
				fmt.Println("\n⚠️  ACCEPTABLE: 50k documents embedded in < 10 minutes")
			} else {
				fmt.Println("\n❌ NEEDS OPTIMIZATION: Consider TEI or GPU acceleration")
			}
		}

		// Recommendations
		fmt.Println("\n📊 Recommendations:")
		if best.concurrency >= 8 {
			fmt.Println("   • High concurrency (8+) provides best throughput")
		}
		if best.docsPerSec < 100 {
			fmt.Println("   • Consider TEI with GPU for 10x+ speedup")
			fmt.Println("   • Consider chunking documents for better batching")
		}
		fmt.Println()
	}
}

// TestScaleThroughputReport generates a detailed throughput report.
// Run with: INTEGRATION_TEST=1 go test -v -run TestScaleThroughputReport ./internal/embedder/
func TestScaleThroughputReport(t *testing.T) {
	if os.Getenv("INTEGRATION_TEST") == "" {
		t.Skip("Skipping integration test (set INTEGRATION_TEST=1 to run)")
	}

	corpusDir := findCorpusDir(t)
	if corpusDir == "" {
		t.Fatal("Could not find benchmarks/synthetic-50k/corpus directory")
	}

	docs := loadCorpus(t, corpusDir)
	if len(docs) < 1000 {
		t.Fatalf("Need at least 1000 docs, found %d", len(docs))
	}

	// Quick throughput test at optimal concurrency
	testSize := min(10000, len(docs))
	batch := extractTexts(docs[:testSize])

	e := NewOllamaEmbedder(WithOllamaConcurrency(8))
	ctx := context.Background()

	fmt.Printf("\n🚀 Quick Throughput Test (%d documents)\n\n", testSize)

	start := time.Now()
	embeddings, err := e.EmbedBatch(ctx, batch)
	elapsed := time.Since(start)

	if err != nil {
		t.Fatalf("Embedding failed: %v", err)
	}

	docsPerSec := float64(len(embeddings)) / elapsed.Seconds()
	estimated50k := time.Duration(float64(50000) / docsPerSec * float64(time.Second))

	fmt.Printf("Documents:        %d\n", len(embeddings))
	fmt.Printf("Time:             %s\n", formatDuration(elapsed))
	fmt.Printf("Throughput:       %.1f docs/sec\n", docsPerSec)
	fmt.Printf("Embedding Dim:    %d\n", len(embeddings[0]))
	fmt.Printf("\nEstimated 50K:    %s\n", formatDuration(estimated50k))

	if estimated50k < 5*time.Minute {
		fmt.Println("\n✅ Projected: Enterprise-ready performance")
	} else {
		fmt.Println("\n⚠️  Projected: May need optimization for production")
	}
}

// Helper functions

func findCorpusDir(t *testing.T) string {
	// Try common paths
	paths := []string{
		"../../benchmarks/synthetic-50k/corpus",
		"benchmarks/synthetic-50k/corpus",
		"../../../benchmarks/synthetic-50k/corpus",
	}

	// Also try from WORKSPACE env if set
	if ws := os.Getenv("WORKSPACE"); ws != "" {
		paths = append([]string{filepath.Join(ws, "benchmarks/synthetic-50k/corpus")}, paths...)
	}

	for _, p := range paths {
		if info, err := os.Stat(p); err == nil && info.IsDir() {
			absPath, _ := filepath.Abs(p)
			t.Logf("Found corpus at: %s", absPath)
			return p
		}
	}
	return ""
}

type document struct {
	path    string
	content string
}

func loadCorpus(t *testing.T, dir string) []document {
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("Failed to read corpus directory: %v", err)
	}

	var docs []document
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".txt") {
			continue
		}

		path := filepath.Join(dir, entry.Name())
		content, err := os.ReadFile(path)
		if err != nil {
			continue
		}

		docs = append(docs, document{
			path:    path,
			content: string(content),
		})
	}

	t.Logf("Loaded %d documents from corpus", len(docs))
	return docs
}

func extractTexts(docs []document) []string {
	texts := make([]string, len(docs))
	for i, doc := range docs {
		texts[i] = doc.content
	}
	return texts
}

func formatDuration(d time.Duration) string {
	if d >= time.Minute {
		return fmt.Sprintf("%.1fm", d.Minutes())
	}
	if d >= time.Second {
		return fmt.Sprintf("%.1fs", d.Seconds())
	}
	return fmt.Sprintf("%dms", d.Milliseconds())
}

func truncateError(err error) string {
	s := err.Error()
	if len(s) > 30 {
		return s[:27] + "..."
	}
	return s
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}



