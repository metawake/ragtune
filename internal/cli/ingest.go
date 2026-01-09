package cli

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/metawake/ragtune/internal/chunker"
	"github.com/metawake/ragtune/internal/embedder"
	"github.com/metawake/ragtune/internal/vectorstore"
	"github.com/metawake/ragtune/internal/vectorstore/chroma"
	"github.com/metawake/ragtune/internal/vectorstore/mock"
	"github.com/metawake/ragtune/internal/vectorstore/pgvector"
	"github.com/metawake/ragtune/internal/vectorstore/pinecone"
	"github.com/metawake/ragtune/internal/vectorstore/qdrant"
	"github.com/metawake/ragtune/internal/vectorstore/weaviate"
)

var (
	chunkSize    int
	chunkOverlap int
	embeddingDim int
	explainMode  bool
)

var ingestCmd = &cobra.Command{
	Use:   "ingest <docs-path>",
	Short: "Load documents, chunk, embed, and upsert into vector store",
	Long: `Ingest documents into a vector store for RAG retrieval.

Reads .md and .txt files from the specified directory, splits them into chunks,
generates embeddings, and upserts into the configured vector store.

Example:
  ragtune ingest ./data/docs --store qdrant --collection demo --chunk-size 512`,
	Args: cobra.ExactArgs(1),
	RunE: runIngest,
}

func init() {
	ingestCmd.Flags().IntVar(&chunkSize, "chunk-size", 512, "Target chunk size in characters")
	ingestCmd.Flags().IntVar(&chunkOverlap, "chunk-overlap", 64, "Overlap between chunks in characters")
	ingestCmd.Flags().IntVar(&embeddingDim, "embedding-dim", 0, "Embedding dimension (auto-detected from embedder if not set)")
	ingestCmd.Flags().BoolVar(&explainMode, "explain", false, "Explain each step of the ingestion process")
}

func runIngest(cmd *cobra.Command, args []string) error {
	docsPath := args[0]

	if collectionName == "" {
		return fmt.Errorf("--collection is required")
	}

	ctx := context.Background()
	totalStart := time.Now()

	// Initialize embedder first (need dimension for collection)
	emb, err := initEmbedder()
	if err != nil {
		return fmt.Errorf("failed to init embedder: %w", err)
	}

	// Initialize vector store
	store, err := initVectorStore(ctx)
	if err != nil {
		return fmt.Errorf("failed to init vector store: %w", err)
	}
	defer store.Close()

	// Determine embedding dimension (auto-detect or override)
	dim := emb.Dim()
	if embeddingDim > 0 {
		dim = embeddingDim
		fmt.Printf("Using embedding dimension: %d (override)\n", dim)
	} else {
		fmt.Printf("Using embedding dimension: %d (auto-detected from %s)\n", dim, embedderName)
	}
	if explainMode {
		fmt.Println("  💡 Embeddings are vectors (lists of numbers) representing meaning.")
		fmt.Println("     Similar texts have similar vectors, enabling semantic search.")
		fmt.Println()
	}

	if err := store.EnsureCollection(ctx, collectionName, dim); err != nil {
		return fmt.Errorf("failed to ensure collection: %w", err)
	}

	// Read and chunk documents
	fmt.Printf("Reading documents from %s...\n", docsPath)
	readStart := time.Now()
	docs, err := readDocuments(docsPath)
	if err != nil {
		return fmt.Errorf("failed to read documents: %w", err)
	}
	readTime := time.Since(readStart)
	fmt.Printf("Found %d documents (read in %s)\n", len(docs), readTime.Round(time.Millisecond))

	// Chunk documents
	chunkStart := time.Now()
	c := chunker.New(chunkSize, chunkOverlap)
	var allChunks []chunker.Chunk
	for _, doc := range docs {
		chunks := c.Chunk(doc.Content, doc.Path)
		allChunks = append(allChunks, chunks...)
	}
	chunkTime := time.Since(chunkStart)
	fmt.Printf("Created %d chunks (chunked in %s)\n", len(allChunks), chunkTime.Round(time.Millisecond))
	if explainMode {
		avgChunkSize := 0
		if len(allChunks) > 0 {
			totalChars := 0
			for _, ch := range allChunks {
				totalChars += len(ch.Text)
			}
			avgChunkSize = totalChars / len(allChunks)
		}
		fmt.Printf("  💡 Chunking splits documents into smaller pieces for embedding.\n")
		fmt.Printf("     • Target size: %d chars, Overlap: %d chars\n", chunkSize, chunkOverlap)
		fmt.Printf("     • Actual avg: %d chars per chunk\n", avgChunkSize)
		fmt.Println("     • Smaller chunks = precise matching, less context")
		fmt.Println("     • Larger chunks = more context, may include noise")
		fmt.Println()
	}

	// Generate embeddings in batches
	fmt.Println("Generating embeddings...")
	embedStart := time.Now()
	var points []vectorstore.Point

	// Batch size varies by provider (TEI: 32, Cohere: 96, Voyage: 128)
	batchSize := 64
	if embedderName == "tei" {
		batchSize = 32 // TEI default max_client_batch_size
	}
	lastProgress := time.Now()
	for i := 0; i < len(allChunks); i += batchSize {
		end := i + batchSize
		if end > len(allChunks) {
			end = len(allChunks)
		}
		batch := allChunks[i:end]

		// Extract texts for batch embedding
		texts := make([]string, len(batch))
		for j, chunk := range batch {
			texts[j] = chunk.Text
		}

		// Batch embed
		vectors, err := emb.EmbedBatch(ctx, texts)
		if err != nil {
			return fmt.Errorf("failed to embed batch starting at %d: %w", i, err)
		}

		// Create points from results
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

		// Progress update every 5 seconds or every 1000 chunks
		if time.Since(lastProgress) > 5*time.Second || end%1000 == 0 {
			elapsed := time.Since(embedStart)
			rate := float64(end) / elapsed.Seconds()
			eta := time.Duration(float64(len(allChunks)-end)/rate) * time.Second
			fmt.Printf("  Embedded %d/%d chunks (%.1f/sec, ETA: %s)\n", end, len(allChunks), rate, eta.Round(time.Second))
			lastProgress = time.Now()
		}
	}
	embedTime := time.Since(embedStart)
	embedRate := float64(len(allChunks)) / embedTime.Seconds()

	// Upsert into vector store
	fmt.Printf("Upserting into %s...\n", storeName)
	upsertStart := time.Now()
	if err := store.Upsert(ctx, collectionName, points); err != nil {
		return fmt.Errorf("failed to upsert: %w", err)
	}
	upsertTime := time.Since(upsertStart)

	if explainMode {
		fmt.Printf("  💡 Stored %d vectors in collection '%s' on %s.\n", len(points), collectionName, storeName)
		fmt.Println("     Each vector is stored with metadata (source file, text).")
		fmt.Println("     Queries will find vectors with similar meaning to your question.")
		fmt.Println()
	}

	totalTime := time.Since(totalStart)

	// Print summary
	fmt.Printf("\n")
	fmt.Printf("╔══════════════════════════════════════════════════════════════╗\n")
	fmt.Printf("║  ✓ Ingestion Complete                                        ║\n")
	fmt.Printf("╠══════════════════════════════════════════════════════════════╣\n")
	fmt.Printf("║  Documents:     %8d                                      ║\n", len(docs))
	fmt.Printf("║  Chunks:        %8d                                      ║\n", len(points))
	fmt.Printf("║  Collection:    %-42s ║\n", collectionName)
	fmt.Printf("╠══════════════════════════════════════════════════════════════╣\n")
	fmt.Printf("║  Read Time:     %8s                                      ║\n", readTime.Round(time.Millisecond))
	fmt.Printf("║  Chunk Time:    %8s                                      ║\n", chunkTime.Round(time.Millisecond))
	fmt.Printf("║  Embed Time:    %8s  (%.1f chunks/sec)                   ║\n", embedTime.Round(time.Second), embedRate)
	fmt.Printf("║  Upsert Time:   %8s                                      ║\n", upsertTime.Round(time.Millisecond))
	fmt.Printf("║  Total Time:    %8s                                      ║\n", totalTime.Round(time.Second))
	fmt.Printf("╚══════════════════════════════════════════════════════════════╝\n")

	return nil
}

// Document represents a loaded document
type Document struct {
	Path    string
	Content string
}

// readDocuments reads all .md and .txt files from a directory
func readDocuments(dir string) ([]Document, error) {
	var docs []Document

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		ext := strings.ToLower(filepath.Ext(path))
		if ext != ".md" && ext != ".txt" {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", path, err)
		}

		docs = append(docs, Document{
			Path:    path,
			Content: string(content),
		})
		return nil
	})

	return docs, err
}

// initVectorStore creates the appropriate vector store based on flags
func initVectorStore(ctx context.Context) (vectorstore.Store, error) {
	switch storeName {
	case "qdrant":
		return qdrant.New(ctx, qdrantAddr)
	case "pgvector":
		if pgvectorConnStr == "" {
			return nil, fmt.Errorf("pgvector store requires --pgvector-url flag")
		}
		return pgvector.New(ctx, pgvectorConnStr)
	case "weaviate":
		return weaviate.New(ctx, weaviateHost, weaviateScheme)
	case "pinecone":
		if pineconeHost == "" {
			return nil, fmt.Errorf("pinecone store requires --pinecone-host flag")
		}
		return pinecone.New(ctx, pineconeHost, pineconeAPIKey)
	case "chroma":
		return chroma.New(ctx, chromaURL)
	case "mock":
		return mock.New(), nil
	default:
		return nil, fmt.Errorf("unsupported store: %s (supported: qdrant, pgvector, weaviate, pinecone, chroma)", storeName)
	}
}

// initEmbedder creates the appropriate embedder based on flags
func initEmbedder() (embedder.Embedder, error) {
	switch embedderName {
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
		return nil, fmt.Errorf("unsupported embedder: %s (supported: openai, ollama, tei, cohere, voyage)", embedderName)
	}
}

// sanitizeString removes invalid UTF-8 and control characters for gRPC compatibility.
func sanitizeString(s string) string {
	var b strings.Builder
	for _, r := range s {
		if r == '\n' || r == '\t' || r >= 32 && r != 0xFFFD {
			b.WriteRune(r)
		}
	}
	return b.String()
}

