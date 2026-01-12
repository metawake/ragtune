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
	"github.com/metawake/ragtune/internal/config"
	"github.com/metawake/ragtune/internal/metrics"
)

var (
	queriesPath string
	configsPath string
	outputDir   string
	// CI mode flags
	ciMode        bool
	minRecall     float64
	minMRR        float64
	minCoverage   float64
	maxLatencyP95 float64
)

var simulateCmd = &cobra.Command{
	Use:   "simulate",
	Short: "Run batch queries and compute retrieval metrics",
	Long: `Simulate RAG retrieval across multiple queries and configurations.

Runs each query against the vector store and computes metrics:
  • Recall@K   - Fraction of relevant docs in top-K results
  • MRR        - Mean Reciprocal Rank (how high first relevant result ranks)
  • NDCG@K     - Normalized Discounted Cumulative Gain (ranking quality)
  • Coverage   - Fraction of relevant docs ever retrieved
  • Redundancy - How often same docs appear across queries
  • Latency    - p50/p95/p99 percentiles

Results are saved to a timestamped JSON file for tracking over time.

Failure Analysis:
  After metrics, shows queries with Recall@K = 0 (complete failures) with
  debugging hints. Helps identify specific queries that need attention.

CI Mode:
  Use --ci with threshold flags for automated quality gates.
  Exit code 1 if any threshold is not met.

Examples:
  ragtune simulate --collection demo --queries data/queries.json

  # CI mode with thresholds
  ragtune simulate --collection prod --queries golden.json \
    --ci --min-recall 0.85 --min-coverage 0.90 --max-latency-p95 500`,
	RunE: runSimulate,
}

func init() {
	simulateCmd.Flags().StringVar(&queriesPath, "queries", "", "Path to queries JSON file (required)")
	simulateCmd.Flags().StringVar(&configsPath, "configs", "", "Path to configs YAML/JSON file (optional)")
	simulateCmd.Flags().StringVar(&outputDir, "output", "runs", "Output directory for run artifacts")
	simulateCmd.MarkFlagRequired("queries")

	// CI mode flags
	simulateCmd.Flags().BoolVar(&ciMode, "ci", false, "CI mode: exit 1 if thresholds not met")
	simulateCmd.Flags().Float64Var(&minRecall, "min-recall", 0, "Minimum Recall@K threshold (CI mode)")
	simulateCmd.Flags().Float64Var(&minMRR, "min-mrr", 0, "Minimum MRR threshold (CI mode)")
	simulateCmd.Flags().Float64Var(&minCoverage, "min-coverage", 0, "Minimum Coverage threshold (CI mode)")
	simulateCmd.Flags().Float64Var(&maxLatencyP95, "max-latency-p95", 0, "Maximum p95 latency in ms (CI mode, 0 = no limit)")

	rootCmd.AddCommand(simulateCmd)
}

// RunResult represents the complete simulation run output.
type RunResult struct {
	Timestamp   string                  `json:"timestamp"`
	Collection  string                  `json:"collection"`
	Store       string                  `json:"store"`
	Configs     []ConfigResult          `json:"configs"`
}

// ConfigResult represents results for a single configuration.
type ConfigResult struct {
	Config       config.SimConfig       `json:"config"`
	Metrics      metrics.Result         `json:"metrics"`
	QueryResults []metrics.QueryResult  `json:"query_results"`
}

func runSimulate(cmd *cobra.Command, args []string) error {
	if collectionName == "" {
		return fmt.Errorf("--collection is required")
	}

	ctx := context.Background()

	// Initialize store
	store, err := initVectorStore(ctx)
	if err != nil {
		return fmt.Errorf("failed to init vector store: %w", err)
	}
	defer closeWithLog(store, "vector store")

	// Initialize embedder
	emb, err := initEmbedder()
	if err != nil {
		return fmt.Errorf("failed to init embedder: %w", err)
	}

	// Load queries
	queries, err := config.LoadQueries(queriesPath)
	if err != nil {
		return fmt.Errorf("failed to load queries: %w", err)
	}
	fmt.Printf("Loaded %d queries\n", len(queries))

	// Load or create default configs
	var configs []config.SimConfig
	if configsPath != "" {
		configs, err = config.LoadConfigs(configsPath)
		if err != nil {
			return fmt.Errorf("failed to load configs: %w", err)
		}
	} else {
		// Default config
		configs = []config.SimConfig{
			{Name: "default", TopK: topK},
		}
	}
	fmt.Printf("Running %d configurations\n", len(configs))

	// Run simulation
	runResult := RunResult{
		Timestamp:  time.Now().UTC().Format(time.RFC3339),
		Collection: collectionName,
		Store:      storeName,
	}

	for _, cfg := range configs {
		fmt.Printf("\n--- Config: %s (top_k=%d) ---\n", cfg.Name, cfg.TopK)

		var queryResults []metrics.QueryResult

		for i, q := range queries {
			// Track latency
			queryStart := time.Now()

			// Embed query
			vec, err := emb.Embed(ctx, q.Text)
			if err != nil {
				return fmt.Errorf("failed to embed query %s: %w", q.ID, err)
			}

			// Search
			results, err := store.Search(ctx, collectionName, vec, cfg.TopK)
			if err != nil {
				return fmt.Errorf("search failed for query %s: %w", q.ID, err)
			}

			// Calculate latency (embedding + search)
			latencyMs := float64(time.Since(queryStart).Microseconds()) / 1000.0

			// Extract IDs (source files) from results
			var retrievedIDs []string
			var scores []float32
			for _, r := range results {
				source := getPayloadString(r.Payload, "source")
				// Extract just the filename for matching
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
				LatencyMs:    latencyMs,
			})

			fmt.Printf("  [%d/%d] %s (%.1fms)\n", i+1, len(queries), q.ID, latencyMs)
		}

		// Compute metrics
		m := metrics.Compute(queryResults, cfg.TopK)

		fmt.Printf("\n  Metrics:\n")
		fmt.Printf("    Recall@%d:  %.3f\n", cfg.TopK, m.RecallAtK)
		fmt.Printf("    MRR:        %.3f\n", m.MRR)
		fmt.Printf("    NDCG@%d:    %.3f\n", cfg.TopK, m.NDCGAtK)
		fmt.Printf("    Coverage:   %.3f\n", m.Coverage)
		fmt.Printf("    Redundancy: %.2f\n", m.Redundancy)
		if m.LatencyAvg > 0 {
			fmt.Printf("    Latency:    p50=%.1fms  p95=%.1fms  p99=%.1fms  avg=%.1fms\n",
				m.LatencyP50, m.LatencyP95, m.LatencyP99, m.LatencyAvg)
		}

		// Per-query failure analysis
		failures := collectFailures(queryResults, cfg.TopK)
		if len(failures) > 0 {
			printFailureReport(failures, cfg.TopK)
		}

		runResult.Configs = append(runResult.Configs, ConfigResult{
			Config:       cfg,
			Metrics:      m,
			QueryResults: queryResults,
		})
	}

	// Save run artifact
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output dir: %w", err)
	}

	ts := strings.ReplaceAll(runResult.Timestamp, ":", "-")
	runPath := filepath.Join(outputDir, fmt.Sprintf("%s.json", ts))

	runData, err := json.MarshalIndent(runResult, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal run result: %w", err)
	}

	if err := os.WriteFile(runPath, runData, 0644); err != nil {
		return fmt.Errorf("failed to write run file: %w", err)
	}

	// Also save as latest.json for convenience
	latestPath := filepath.Join(outputDir, "latest.json")
	if err := os.WriteFile(latestPath, runData, 0644); err != nil {
		return fmt.Errorf("failed to write latest.json: %w", err)
	}

	fmt.Printf("\n✓ Run saved to %s\n", runPath)
	fmt.Printf("✓ Also saved as %s\n", latestPath)

	// CI mode: check thresholds
	if ciMode {
		return checkCIThresholds(runResult)
	}

	return nil
}

// QueryFailure represents a query that failed to retrieve its relevant documents.
type QueryFailure struct {
	QueryID      string
	Query        string
	RelevantDocs []string
	RetrievedDocs []string
	TopScores    []float32
	Recall       float64
}

// collectFailures identifies queries with zero or low recall.
func collectFailures(results []metrics.QueryResult, k int) []QueryFailure {
	var failures []QueryFailure

	for _, qr := range results {
		recall := metrics.RecallAtK(qr.RetrievedIDs, qr.RelevantIDs, k)
		
		// Report queries with Recall@K = 0 (complete failures)
		if recall == 0 && len(qr.RelevantIDs) > 0 {
			topDocs := qr.RetrievedIDs
			if len(topDocs) > 3 {
				topDocs = topDocs[:3]
			}
			topScores := qr.Scores
			if len(topScores) > 3 {
				topScores = topScores[:3]
			}
			
			failures = append(failures, QueryFailure{
				QueryID:      qr.QueryID,
				Query:        qr.Query,
				RelevantDocs: qr.RelevantIDs,
				RetrievedDocs: topDocs,
				TopScores:    topScores,
				Recall:       recall,
			})
		}
	}

	return failures
}

// printFailureReport displays detailed information about failed queries.
func printFailureReport(failures []QueryFailure, k int) {
	fmt.Println()
	fmt.Println("  ─────────────────────────────────────────────────────────────")
	fmt.Printf("  FAILURES: %d queries with Recall@%d = 0\n", len(failures), k)
	fmt.Println("  ─────────────────────────────────────────────────────────────")
	
	// Show up to 5 failures in detail
	showCount := len(failures)
	if showCount > 5 {
		showCount = 5
	}
	
	for i := 0; i < showCount; i++ {
		f := failures[i]
		fmt.Println()
		fmt.Printf("  ✗ [%s] %s\n", f.QueryID, truncateQuery(f.Query, 60))
		fmt.Printf("    Expected:  %v\n", f.RelevantDocs)
		
		if len(f.RetrievedDocs) > 0 {
			// Format retrieved docs with scores
			var docsWithScores []string
			for j, doc := range f.RetrievedDocs {
				if j < len(f.TopScores) {
					docsWithScores = append(docsWithScores, fmt.Sprintf("%s (%.3f)", doc, f.TopScores[j]))
				} else {
					docsWithScores = append(docsWithScores, doc)
				}
			}
			fmt.Printf("    Retrieved: %v\n", docsWithScores)
		} else {
			fmt.Println("    Retrieved: (none)")
		}
	}
	
	if len(failures) > 5 {
		fmt.Printf("\n  ... and %d more failures (see JSON output for full list)\n", len(failures)-5)
	}
	
	// Provide actionable hints
	fmt.Println()
	fmt.Println("  💡 Debugging hints:")
	fmt.Println("     • Run `ragtune explain \"<query>\"` to inspect retrieval for specific queries")
	fmt.Println("     • Check if expected documents are in the corpus")
	fmt.Println("     • Try different chunk sizes or embedders with `ragtune compare`")
}

// truncateQuery shortens a query string for display
func truncateQuery(s string, maxLen int) string {
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.Join(strings.Fields(s), " ")
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// checkCIThresholds verifies metrics against thresholds and returns an error if any fail.
func checkCIThresholds(result RunResult) error {
	if len(result.Configs) == 0 {
		return fmt.Errorf("no configurations to check")
	}

	// Use first config's metrics (typically there's only one in CI mode)
	m := result.Configs[0].Metrics

	fmt.Println()
	fmt.Println("─────────────────────────────────────────")
	fmt.Println("CI THRESHOLD CHECK")
	fmt.Println("─────────────────────────────────────────")

	var failedChecks []string

	if minRecall > 0 {
		status := "✓ PASS"
		if m.RecallAtK < minRecall {
			status = "✗ FAIL"
			failedChecks = append(failedChecks, "Recall@K")
		}
		fmt.Printf("  Recall@K:    %.3f  %s  (threshold: %.3f)\n", m.RecallAtK, status, minRecall)
	}

	if minMRR > 0 {
		status := "✓ PASS"
		if m.MRR < minMRR {
			status = "✗ FAIL"
			failedChecks = append(failedChecks, "MRR")
		}
		fmt.Printf("  MRR:         %.3f  %s  (threshold: %.3f)\n", m.MRR, status, minMRR)
	}

	if minCoverage > 0 {
		status := "✓ PASS"
		if m.Coverage < minCoverage {
			status = "✗ FAIL"
			failedChecks = append(failedChecks, "Coverage")
		}
		fmt.Printf("  Coverage:    %.3f  %s  (threshold: %.3f)\n", m.Coverage, status, minCoverage)
	}

	if maxLatencyP95 > 0 {
		status := "✓ PASS"
		if m.LatencyP95 > maxLatencyP95 {
			status = "✗ FAIL"
			failedChecks = append(failedChecks, "Latency p95")
		}
		fmt.Printf("  Latency p95: %.0fms  %s  (threshold: %.0fms)\n", m.LatencyP95, status, maxLatencyP95)
	}

	failed := len(failedChecks) > 0

	fmt.Println("─────────────────────────────────────────")

	if failed {
		fmt.Println("❌ CI check FAILED: one or more thresholds not met")
		return &CICheckError{FailedChecks: failedChecks}
	}

	fmt.Println("✅ CI check PASSED: all thresholds met")
	return nil
}

