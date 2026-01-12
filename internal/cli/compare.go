package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/metawake/ragtune/internal/chunker"
	"github.com/metawake/ragtune/internal/config"
	"github.com/metawake/ragtune/internal/embedder"
	"github.com/metawake/ragtune/internal/metrics"
	"github.com/metawake/ragtune/internal/vectorstore"
)

var (
	collections      string
	compareEmbedders string
	compareDocs      string
	compareChunkSize int
	compareKeep      bool
)

var compareCmd = &cobra.Command{
	Use:   "compare",
	Short: "Compare retrieval quality across multiple collections or embedders",
	Long: `Compare retrieval metrics across multiple collections or embedding models.

MODE 1: Compare existing collections (different chunk sizes)
  ragtune compare --collections demo-256,demo-512,demo-1024 --queries queries.json

MODE 2: Compare embedding models (auto-ingest and compare)
  ragtune compare --embedders openai,ollama --docs ./docs --queries queries.json

The --embedders mode automatically:
  1. Creates temporary collections for each embedder
  2. Ingests documents with each embedder
  3. Runs the comparison
  4. Cleans up (unless --keep is specified)

Examples:
  # Compare chunk sizes (manual ingest first)
  ragtune ingest ./docs --collection demo-256 --chunk-size 256
  ragtune ingest ./docs --collection demo-512 --chunk-size 512
  ragtune compare --collections demo-256,demo-512 --queries queries.json

  # Compare embedders (automatic)
  ragtune compare --embedders openai,ollama --docs ./docs --queries queries.json --chunk-size 512`,
	RunE: runCompare,
}

var compareTopK int

func init() {
	compareCmd.Flags().StringVar(&collections, "collections", "", "Comma-separated collection names to compare")
	compareCmd.Flags().StringVar(&compareEmbedders, "embedders", "", "Comma-separated embedder names to compare (openai, ollama)")
	compareCmd.Flags().StringVar(&compareDocs, "docs", "", "Path to documents (required with --embedders)")
	compareCmd.Flags().IntVar(&compareChunkSize, "chunk-size", 512, "Chunk size for --embedders mode")
	compareCmd.Flags().BoolVar(&compareKeep, "keep", false, "Keep auto-created collections (don't delete after comparison)")
	compareCmd.Flags().StringVar(&queriesPath, "queries", "", "Path to queries JSON file (required)")
	compareCmd.Flags().StringVar(&outputDir, "output", "runs", "Output directory for run artifacts")
	compareCmd.Flags().IntVar(&compareTopK, "top-k", 5, "Number of results to retrieve")
	compareCmd.MarkFlagRequired("queries")

	rootCmd.AddCommand(compareCmd)
}

// CompareResult holds comparison results across collections.
type CompareResult struct {
	Timestamp   string              `json:"timestamp"`
	Collections []CollectionResult  `json:"collections"`
}

// CollectionResult holds results for a single collection.
type CollectionResult struct {
	Collection   string               `json:"collection"`
	Metrics      metrics.Result       `json:"metrics"`
	QueryResults []metrics.QueryResult `json:"query_results,omitempty"`
}

func runCompare(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	// Validate flags: need either --collections or --embedders
	if collections == "" && compareEmbedders == "" {
		return fmt.Errorf("either --collections or --embedders is required")
	}
	if collections != "" && compareEmbedders != "" {
		return fmt.Errorf("use either --collections or --embedders, not both")
	}

	// Initialize store
	store, err := initVectorStore(ctx)
	if err != nil {
		return fmt.Errorf("failed to init vector store: %w", err)
	}
	defer store.Close()

	var collectionList []string
	var embeddersUsed []embedder.Embedder
	var autoCreated bool

	if compareEmbedders != "" {
		// Mode 2: Auto-ingest with different embedders
		if compareDocs == "" {
			return fmt.Errorf("--docs is required when using --embedders")
		}

		embedderNames := strings.Split(compareEmbedders, ",")
		for i := range embedderNames {
			embedderNames[i] = strings.TrimSpace(embedderNames[i])
		}
		if len(embedderNames) < 2 {
			return fmt.Errorf("need at least 2 embedders to compare")
		}

		// Generate timestamp for collection names
		ts := time.Now().Format("20060102-150405")

		fmt.Printf("Auto-ingesting with %d embedders: %s\n\n", len(embedderNames), strings.Join(embedderNames, ", "))

		// Read documents once
		docs, err := readDocuments(compareDocs)
		if err != nil {
			return fmt.Errorf("failed to read documents: %w", err)
		}
		fmt.Printf("Found %d documents\n", len(docs))

		// Chunk documents once
		c, err := chunker.New(compareChunkSize, compareChunkSize/8) // 12.5% overlap
		if err != nil {
			return fmt.Errorf("invalid chunker config: %w", err)
		}
		var allChunks []chunker.Chunk
		for _, doc := range docs {
			chunks := c.Chunk(doc.Content, doc.Path)
			allChunks = append(allChunks, chunks...)
		}
		fmt.Printf("Created %d chunks (size=%d)\n\n", len(allChunks), compareChunkSize)

		// Ingest with each embedder
		for _, embName := range embedderNames {
			collName := fmt.Sprintf("compare-%s-%s", embName, ts)
			collectionList = append(collectionList, collName)

			fmt.Printf("--- Ingesting with embedder: %s ---\n", embName)

			// Create embedder for this iteration
			emb, err := createEmbedder(embName)
			if err != nil {
				return fmt.Errorf("failed to create embedder %s: %w", embName, err)
			}
			embeddersUsed = append(embeddersUsed, emb)

			// Create collection
			if err := store.EnsureCollection(ctx, collName, emb.Dim()); err != nil {
				return fmt.Errorf("failed to create collection %s: %w", collName, err)
			}

			// Generate embeddings
			var points []vectorstore.Point
			batchSize := 64 // Conservative default safe for most providers
			for i := 0; i < len(allChunks); i += batchSize {
				end := i + batchSize
				if end > len(allChunks) {
					end = len(allChunks)
				}
				batch := allChunks[i:end]

				texts := make([]string, len(batch))
				for j, chunk := range batch {
					texts[j] = chunk.Text
				}

				vectors, err := emb.EmbedBatch(ctx, texts)
				if err != nil {
					return fmt.Errorf("failed to embed batch with %s: %w", embName, err)
				}

				for j, chunk := range batch {
					points = append(points, vectorstore.Point{
						ID:     chunk.ID,
						Vector: vectors[j],
						Payload: map[string]interface{}{
							"text":     sanitizeString(chunk.Text),
							"source":   sanitizeString(chunk.Source),
							"chunk_id": chunk.Index,
						},
					})
				}

				if end%200 == 0 || end == len(allChunks) {
					fmt.Printf("  Embedded %d/%d chunks\n", end, len(allChunks))
				}
			}

			// Upsert
			if err := store.Upsert(ctx, collName, points); err != nil {
				return fmt.Errorf("failed to upsert to %s: %w", collName, err)
			}
			fmt.Printf("  ✓ Ingested %d chunks into '%s'\n\n", len(points), collName)
			points = nil // Reset for next embedder
		}

		autoCreated = true
	} else {
		// Mode 1: Use existing collections
		collectionList = strings.Split(collections, ",")
		if len(collectionList) < 2 {
			return fmt.Errorf("need at least 2 collections to compare")
		}

		// Trim whitespace
		for i := range collectionList {
			collectionList[i] = strings.TrimSpace(collectionList[i])
		}
	}

	// Load queries
	queries, err := config.LoadQueries(queriesPath)
	if err != nil {
		return fmt.Errorf("failed to load queries: %w", err)
	}
	fmt.Printf("Loaded %d queries\n", len(queries))
	fmt.Printf("Comparing %d collections on %s: %s\n\n", len(collectionList), storeName, strings.Join(collectionList, ", "))

	result := CompareResult{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	// Use compare-specific topK
	k := compareTopK

	// Run against each collection
	for i, coll := range collectionList {
		fmt.Printf("--- Collection: %s ---\n", coll)

		// Get the right embedder for this collection
		var emb embedder.Embedder
		if autoCreated && i < len(embeddersUsed) {
			emb = embeddersUsed[i]
		} else {
			// Use the global embedder setting
			emb, err = initEmbedder()
			if err != nil {
				return fmt.Errorf("failed to init embedder: %w", err)
			}
		}

		var queryResults []metrics.QueryResult

		for j, q := range queries {
			// Embed query
			vec, err := emb.Embed(ctx, q.Text)
			if err != nil {
				return fmt.Errorf("failed to embed query %s: %w", q.ID, err)
			}

			// Search
			results, err := store.Search(ctx, coll, vec, k)
			if err != nil {
				return fmt.Errorf("search failed for query %s in %s: %w", q.ID, coll, err)
			}

			// Extract IDs
			var retrievedIDs []string
			var scores []float32
			for _, r := range results {
				source := getPayloadString(r.Payload, "source")
				source = filepath.Base(source)
				retrievedIDs = append(retrievedIDs, source)
				scores = append(scores, r.Score)
			}

			queryResults = append(queryResults, metrics.QueryResult{
				QueryID:      q.ID,
				Query:        q.Text,
				RetrievedIDs: retrievedIDs,
				RelevantIDs:  q.RelevantDocs,
				Scores:       scores,
			})

			if (j+1)%50 == 0 {
				fmt.Printf("  Processed %d/%d queries\n", j+1, len(queries))
			}
		}

		// Compute metrics
		m := metrics.Compute(queryResults, k)

		fmt.Printf("  Recall@%d: %.3f | MRR: %.3f | Coverage: %.3f\n\n", k, m.RecallAtK, m.MRR, m.Coverage)

		result.Collections = append(result.Collections, CollectionResult{
			Collection:   coll,
			Metrics:      m,
			QueryResults: queryResults,
		})
	}

	// Print comparison table
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println("COMPARISON SUMMARY")
	fmt.Println(strings.Repeat("=", 70))
	fmt.Printf("\n| %-20s | Recall@%d | MRR    | Coverage | Redundancy |\n", "Collection", k)
	fmt.Println("|" + strings.Repeat("-", 22) + "|----------|--------|----------|------------|")

	var bestRecall float64
	var bestCollection string
	for _, cr := range result.Collections {
		fmt.Printf("| %-20s | %.3f    | %.3f  | %.3f    | %.2f       |\n",
			cr.Collection,
			cr.Metrics.RecallAtK,
			cr.Metrics.MRR,
			cr.Metrics.Coverage,
			cr.Metrics.Redundancy,
		)
		if cr.Metrics.RecallAtK > bestRecall {
			bestRecall = cr.Metrics.RecallAtK
			bestCollection = cr.Collection
		}
	}

	fmt.Println()
	fmt.Printf("✓ Highest recall: %s (%.3f)\n", bestCollection, bestRecall)

	// Add context if difference is small
	if len(result.Collections) > 1 {
		var minRecall float64 = 1.0
		for _, cr := range result.Collections {
			if cr.Metrics.RecallAtK < minRecall {
				minRecall = cr.Metrics.RecallAtK
			}
		}
		diff := bestRecall - minRecall
		if diff < 0.05 {
			fmt.Println("  (Difference is small — may not be significant for your use case)")
		}
	}

	// Save results
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output dir: %w", err)
	}

	ts := strings.ReplaceAll(result.Timestamp, ":", "-")
	runPath := filepath.Join(outputDir, fmt.Sprintf("compare-%s.json", ts))

	runData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal result: %w", err)
	}

	if err := os.WriteFile(runPath, runData, 0644); err != nil {
		return fmt.Errorf("failed to write run file: %w", err)
	}

	fmt.Printf("✓ Comparison saved to %s\n", runPath)

	// Cleanup auto-created collections unless --keep is specified
	if autoCreated && !compareKeep {
		fmt.Println("\nCleaning up auto-created collections...")
		for _, coll := range collectionList {
			if err := store.DeleteCollection(ctx, coll); err != nil {
				fmt.Printf("  Warning: failed to delete %s: %v\n", coll, err)
			} else {
				fmt.Printf("  Deleted %s\n", coll)
			}
		}
		fmt.Println("✓ Cleanup complete (use --keep to preserve collections)")
	} else if autoCreated && compareKeep {
		fmt.Println("\nCollections preserved (--keep specified):")
		for _, coll := range collectionList {
			fmt.Printf("  - %s\n", coll)
		}
	}

	return nil
}

// createEmbedder creates an embedder by name
func createEmbedder(name string) (embedder.Embedder, error) {
	switch name {
	case "openai":
		return embedder.NewOpenAIEmbedder(), nil
	case "ollama":
		return embedder.NewOllamaEmbedder(
			embedder.WithOllamaURL(ollamaAddr),
			embedder.WithOllamaModel(ollamaModel),
			embedder.WithOllamaConcurrency(ollamaConcurrency),
		), nil
	case "tei":
		return embedder.NewTEIEmbedder(
			embedder.WithTEIURL(teiAddr),
			embedder.WithTEIModel(teiModel),
		), nil
	case "cohere":
		return embedder.NewCohereEmbedder(
			embedder.WithCohereModel(cohereModel),
		), nil
	case "voyage":
		return embedder.NewVoyageEmbedder(
			embedder.WithVoyageModel(voyageModel),
		), nil
	default:
		return nil, fmt.Errorf("unsupported embedder: %s (supported: openai, ollama, tei, cohere, voyage)", name)
	}
}

