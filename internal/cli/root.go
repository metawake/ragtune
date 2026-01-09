package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Version is set at build time via -ldflags
var Version = "dev"

var rootCmd = &cobra.Command{
	Version: Version,
	Use:   "ragtune",
	Short: "EXPLAIN ANALYZE for RAG retrieval",
	Long: `RagTune — inspect, explain, benchmark, and tune RAG retrieval layers.

A production-minded CLI for RAG diagnostics and tuning.
Vector-store agnostic by design.

Commands:
  ingest   Load documents, chunk, embed, and upsert into vector store
  explain  Show top-k retrieved chunks for a query with distances and warnings

Examples:
  ragtune ingest ./data/docs --store qdrant --collection demo
  ragtune explain "How to rotate API key?" --store qdrant --collection demo --top-k 5

Sponsor:
  If RagTune saves you time, star the repo or support development:
  GitHub: https://github.com/metawake/ragtune
  ETH: 0x0a542565b3615e8fc934cc3cc4921a0c22e5dc5e`,
}

// Global flags
var (
	storeName         string
	collectionName    string
	qdrantAddr        string
	pgvectorConnStr   string
	weaviateHost      string
	weaviateScheme    string
	pineconeHost      string
	pineconeAPIKey    string
	chromaURL         string
	embedderName      string
	ollamaAddr        string
	ollamaModel       string
	ollamaConcurrency int
	cohereModel       string
	voyageModel       string
	teiAddr           string
	teiModel          string
	topK              int
)

func init() {
	// Persistent flags available to all subcommands
	rootCmd.PersistentFlags().StringVar(&storeName, "store", "qdrant", "Vector store backend (qdrant, pgvector, weaviate, pinecone, chroma)")
	rootCmd.PersistentFlags().StringVar(&collectionName, "collection", "", "Collection name (required)")
	rootCmd.PersistentFlags().StringVar(&qdrantAddr, "qdrant-addr", "127.0.0.1:6334", "Qdrant gRPC address")
	rootCmd.PersistentFlags().StringVar(&pgvectorConnStr, "pgvector-url", "", "PostgreSQL connection string for pgvector")
	rootCmd.PersistentFlags().StringVar(&weaviateHost, "weaviate-host", "localhost:8080", "Weaviate server host")
	rootCmd.PersistentFlags().StringVar(&weaviateScheme, "weaviate-scheme", "http", "Weaviate scheme (http or https)")
	rootCmd.PersistentFlags().StringVar(&pineconeHost, "pinecone-host", "", "Pinecone index host (from console)")
	rootCmd.PersistentFlags().StringVar(&pineconeAPIKey, "pinecone-api-key", "", "Pinecone API key (or use PINECONE_API_KEY env)")
	rootCmd.PersistentFlags().StringVar(&chromaURL, "chroma-url", "http://localhost:8000", "Chroma server URL")

	// Embedder flags
	rootCmd.PersistentFlags().StringVar(&embedderName, "embedder", "openai", "Embedding backend (openai, ollama, tei, cohere, voyage)")
	rootCmd.PersistentFlags().StringVar(&ollamaAddr, "ollama-addr", "http://localhost:11434", "Ollama server URL")
	rootCmd.PersistentFlags().StringVar(&ollamaModel, "ollama-model", "nomic-embed-text", "Ollama embedding model")
	rootCmd.PersistentFlags().IntVar(&ollamaConcurrency, "ollama-concurrency", 8, "Ollama concurrent requests (higher = faster, more CPU)")
	rootCmd.PersistentFlags().StringVar(&teiAddr, "tei-addr", "http://localhost:8080", "HuggingFace TEI server URL")
	rootCmd.PersistentFlags().StringVar(&teiModel, "tei-model", "BAAI/bge-base-en-v1.5", "TEI model (for dimension inference)")
	rootCmd.PersistentFlags().StringVar(&cohereModel, "cohere-model", "embed-english-v3.0", "Cohere embedding model")
	rootCmd.PersistentFlags().StringVar(&voyageModel, "voyage-model", "voyage-2", "Voyage embedding model (voyage-2, voyage-law-2, voyage-code-2)")

	// Retrieval flags
	rootCmd.PersistentFlags().IntVar(&topK, "top-k", 5, "Number of results to retrieve")

	rootCmd.AddCommand(ingestCmd)
	rootCmd.AddCommand(explainCmd)
}

// Execute runs the root command
func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	return nil
}

