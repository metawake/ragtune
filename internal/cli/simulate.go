package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/metawake/ragtune/internal/config"
	"github.com/metawake/ragtune/internal/metrics"
	"github.com/spf13/cobra"
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
	// Baseline comparison flags
	baselinePath     string
	failOnRegression bool
	// Output format flags
	jsonOutput bool
	// Bootstrap flags
	bootstrapN    int
	bootstrapSeed int64
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

Baseline Comparison:
  Use --baseline to compare against a previous run. Shows deltas for each
  metric. Use --fail-on-regression to fail CI if any metric decreased.

Bootstrap Confidence Intervals:
  Use --bootstrap N to run N bootstrap samples and report mean ± std for
  each metric. This enables distinguishing real changes from random variance.
  Example: "Recall@5: 0.664 ± 0.012" means the true value is likely in that range.

Examples:
  ragtune simulate --collection demo --queries data/queries.json

  # With bootstrap confidence intervals
  ragtune simulate --collection prod --queries golden.json --bootstrap 20

  # CI mode with thresholds
  ragtune simulate --collection prod --queries golden.json \
    --ci --min-recall 0.85 --min-coverage 0.90 --max-latency-p95 500

  # Compare against baseline (regression testing)
  ragtune simulate --collection prod --queries golden.json \
    --baseline runs/latest.json --fail-on-regression`,
	RunE: runSimulate,
}

func init() {
	simulateCmd.Flags().StringVar(&queriesPath, "queries", "", "Path to queries JSON file (required)")
	simulateCmd.Flags().StringVar(&configsPath, "configs", "", "Path to configs YAML/JSON file (optional)")
	simulateCmd.Flags().StringVar(&outputDir, "output", "runs", "Output directory for run artifacts")
	_ = simulateCmd.MarkFlagRequired("queries")

	// CI mode flags
	simulateCmd.Flags().BoolVar(&ciMode, "ci", false, "CI mode: exit 1 if thresholds not met")
	simulateCmd.Flags().Float64Var(&minRecall, "min-recall", 0, "Minimum Recall@K threshold (CI mode)")
	simulateCmd.Flags().Float64Var(&minMRR, "min-mrr", 0, "Minimum MRR threshold (CI mode)")
	simulateCmd.Flags().Float64Var(&minCoverage, "min-coverage", 0, "Minimum Coverage threshold (CI mode)")
	simulateCmd.Flags().Float64Var(&maxLatencyP95, "max-latency-p95", 0, "Maximum p95 latency in ms (CI mode, 0 = no limit)")

	// Baseline comparison flags
	simulateCmd.Flags().StringVar(&baselinePath, "baseline", "", "Path to baseline run JSON for comparison (e.g., runs/latest.json)")
	simulateCmd.Flags().BoolVar(&failOnRegression, "fail-on-regression", false, "Exit 1 if any metric regressed vs baseline (requires --baseline)")

	// Output format flags
	simulateCmd.Flags().BoolVar(&jsonOutput, "json", false, "Output results as JSON (for CI parsing)")

	// Bootstrap flags for statistical confidence
	simulateCmd.Flags().IntVar(&bootstrapN, "bootstrap", 0, "Number of bootstrap samples for confidence intervals (0 = disabled)")
	simulateCmd.Flags().Int64Var(&bootstrapSeed, "bootstrap-seed", 42, "Random seed for bootstrap reproducibility")

	rootCmd.AddCommand(simulateCmd)
}

// RunResult represents the complete simulation run output.
type RunResult struct {
	Timestamp  string         `json:"timestamp"`
	Collection string         `json:"collection"`
	Store      string         `json:"store"`
	Configs    []ConfigResult `json:"configs"`
}

// ConfigResult represents results for a single configuration.
type ConfigResult struct {
	Config       config.SimConfig         `json:"config"`
	Metrics      metrics.Result           `json:"metrics"`
	Bootstrap    *metrics.BootstrapResult `json:"bootstrap,omitempty"`
	QueryResults []metrics.QueryResult    `json:"query_results"`
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
	if !jsonOutput {
		fmt.Printf("Loaded %d queries\n", len(queries))
	}

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
	if !jsonOutput {
		fmt.Printf("Running %d configurations\n", len(configs))
	}

	// Run simulation
	runResult := RunResult{
		Timestamp:  time.Now().UTC().Format(time.RFC3339),
		Collection: collectionName,
		Store:      storeName,
	}

	for _, cfg := range configs {
		if !jsonOutput {
			fmt.Printf("\n--- Config: %s (top_k=%d) ---\n", cfg.Name, cfg.TopK)
		}

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
				// Normalize: map chunk-level sources back to parent doc.
				// e.g. "rfc6749_oauth2_cs0227.txt" -> "rfc6749_oauth2.txt"
				// if "rfc6749_oauth2.txt" is in the relevant docs.
				source = normalizeSource(source, q.RelevantDocs)
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

			if !jsonOutput {
				fmt.Printf("  [%d/%d] %s (%.1fms)\n", i+1, len(queries), q.ID, latencyMs)
			}
		}

		// Compute metrics
		m := metrics.Compute(queryResults, cfg.TopK)

		// Compute bootstrap confidence intervals if requested
		var bs *metrics.BootstrapResult
		if bootstrapN > 0 {
			bsResult := metrics.Bootstrap(queryResults, cfg.TopK, bootstrapN, 0, bootstrapSeed, false)
			bs = &bsResult
		}

		if !jsonOutput {
			fmt.Printf("\n  Metrics:\n")
			if bs != nil {
				// Show metrics with confidence intervals
				fmt.Printf("    Recall@%d:  %.3f ± %.3f  (n=%d)\n", cfg.TopK, bs.RecallMean, bs.RecallStd, bs.N)
				fmt.Printf("    MRR:        %.3f ± %.3f\n", bs.MRRMean, bs.MRRStd)
				fmt.Printf("    NDCG@%d:    %.3f ± %.3f\n", cfg.TopK, bs.NDCGMean, bs.NDCGStd)
				fmt.Printf("    Coverage:   %.3f ± %.3f\n", bs.CoverageMean, bs.CoverageStd)
			} else {
				// Show point estimates only
				fmt.Printf("    Recall@%d:  %.3f\n", cfg.TopK, m.RecallAtK)
				fmt.Printf("    MRR:        %.3f\n", m.MRR)
				fmt.Printf("    NDCG@%d:    %.3f\n", cfg.TopK, m.NDCGAtK)
				fmt.Printf("    Coverage:   %.3f\n", m.Coverage)
			}
			fmt.Printf("    Redundancy: %.2f\n", m.Redundancy)
			if m.LatencyAvg > 0 {
				fmt.Printf("    Latency:    p50=%.1fms  p95=%.1fms  p99=%.1fms  avg=%.1fms\n",
					m.LatencyP50, m.LatencyP95, m.LatencyP99, m.LatencyAvg)
			}
		}

		// Per-query failure analysis
		failures := collectFailures(queryResults, cfg.TopK)
		if len(failures) > 0 && !jsonOutput {
			printFailureReport(failures, cfg.TopK)
		}

		runResult.Configs = append(runResult.Configs, ConfigResult{
			Config:       cfg,
			Metrics:      m,
			Bootstrap:    bs,
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

	if !jsonOutput {
		fmt.Printf("\n✓ Run saved to %s\n", runPath)
		fmt.Printf("✓ Also saved as %s\n", latestPath)
	}

	// Variables for JSON output
	var deltas []MetricDelta
	var hasRegression bool
	var baselineTs string
	var thresholdChecks []JSONThresholdCheck
	thresholdsPassed := true

	// Collect failures for JSON output
	var allFailures []QueryFailure
	if len(runResult.Configs) > 0 {
		allFailures = collectFailures(runResult.Configs[0].QueryResults, runResult.Configs[0].Config.TopK)
	}

	// Baseline comparison (if specified)
	if baselinePath != "" {
		baseline, err := loadBaseline(baselinePath)
		if err != nil {
			return fmt.Errorf("failed to load baseline: %w", err)
		}

		// Compare first config's metrics
		currentMetrics := runResult.Configs[0].Metrics
		baselineMetrics := baseline.Configs[0].Metrics
		k := runResult.Configs[0].Config.TopK
		baselineTs = baseline.Timestamp

		deltas, hasRegression = compareWithBaseline(currentMetrics, baselineMetrics, k)
		if !jsonOutput {
			printBaselineComparison(deltas, baseline.Timestamp)
		}

		// Check for regressions if flag is set
		if failOnRegression && hasRegression {
			if jsonOutput {
				output := buildJSONOutput(runResult, runPath, allFailures, deltas, hasRegression, baselineTs, thresholdChecks, thresholdsPassed)
				printJSONOutput(output)
			}
			return checkRegressions(deltas)
		}
	}

	// CI mode: check thresholds
	if ciMode {
		m := runResult.Configs[0].Metrics

		if minRecall > 0 {
			passed := m.RecallAtK >= minRecall
			thresholdChecks = append(thresholdChecks, JSONThresholdCheck{
				Metric: "Recall@K", Value: m.RecallAtK, Threshold: minRecall, Passed: passed,
			})
			if !passed {
				thresholdsPassed = false
			}
		}
		if minMRR > 0 {
			passed := m.MRR >= minMRR
			thresholdChecks = append(thresholdChecks, JSONThresholdCheck{
				Metric: "MRR", Value: m.MRR, Threshold: minMRR, Passed: passed,
			})
			if !passed {
				thresholdsPassed = false
			}
		}
		if minCoverage > 0 {
			passed := m.Coverage >= minCoverage
			thresholdChecks = append(thresholdChecks, JSONThresholdCheck{
				Metric: "Coverage", Value: m.Coverage, Threshold: minCoverage, Passed: passed,
			})
			if !passed {
				thresholdsPassed = false
			}
		}
		if maxLatencyP95 > 0 {
			passed := m.LatencyP95 <= maxLatencyP95
			thresholdChecks = append(thresholdChecks, JSONThresholdCheck{
				Metric: "Latency p95", Value: m.LatencyP95, Threshold: maxLatencyP95, Passed: passed,
			})
			if !passed {
				thresholdsPassed = false
			}
		}

		if jsonOutput {
			output := buildJSONOutput(runResult, runPath, allFailures, deltas, hasRegression, baselineTs, thresholdChecks, thresholdsPassed)
			printJSONOutput(output)
			if !thresholdsPassed {
				return &CICheckError{FailedChecks: []string{"threshold"}}
			}
			return nil
		}
		return checkCIThresholds(runResult)
	}

	// JSON output for non-CI mode
	if jsonOutput {
		output := buildJSONOutput(runResult, runPath, allFailures, deltas, hasRegression, baselineTs, thresholdChecks, thresholdsPassed)
		printJSONOutput(output)
	}

	return nil
}

// QueryFailure represents a query that failed to retrieve its relevant documents.
type QueryFailure struct {
	QueryID       string
	Query         string
	RelevantDocs  []string
	RetrievedDocs []string
	TopScores     []float32
	Recall        float64
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
				QueryID:       qr.QueryID,
				Query:         qr.Query,
				RelevantDocs:  qr.RelevantIDs,
				RetrievedDocs: topDocs,
				TopScores:     topScores,
				Recall:        recall,
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
		var status string
		if m.RecallAtK >= minRecall {
			status = "✓ PASS"
		} else {
			status = "✗ FAIL"
			failedChecks = append(failedChecks, "Recall@K")
		}
		fmt.Printf("  Recall@K:    %.3f  %s  (threshold: %.3f)\n", m.RecallAtK, status, minRecall)
	}

	if minMRR > 0 {
		var status string
		if m.MRR >= minMRR {
			status = "✓ PASS"
		} else {
			status = "✗ FAIL"
			failedChecks = append(failedChecks, "MRR")
		}
		fmt.Printf("  MRR:         %.3f  %s  (threshold: %.3f)\n", m.MRR, status, minMRR)
	}

	if minCoverage > 0 {
		var status string
		if m.Coverage >= minCoverage {
			status = "✓ PASS"
		} else {
			status = "✗ FAIL"
			failedChecks = append(failedChecks, "Coverage")
		}
		fmt.Printf("  Coverage:    %.3f  %s  (threshold: %.3f)\n", m.Coverage, status, minCoverage)
	}

	if maxLatencyP95 > 0 {
		var status string
		if m.LatencyP95 <= maxLatencyP95 {
			status = "✓ PASS"
		} else {
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

// loadBaseline reads a previous run result from a JSON file.
func loadBaseline(path string) (*RunResult, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read baseline file: %w", err)
	}

	var result RunResult
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse baseline JSON: %w", err)
	}

	if len(result.Configs) == 0 {
		return nil, fmt.Errorf("baseline has no configs")
	}

	return &result, nil
}

// MetricDelta represents the change in a metric from baseline to current.
type MetricDelta struct {
	Name       string
	Baseline   float64
	Current    float64
	Delta      float64
	DeltaPct   float64
	Regressed  bool
	HigherGood bool // true for recall/mrr/coverage, false for latency
}

// compareWithBaseline compares current metrics against baseline and prints the comparison.
// Returns true if any metric regressed.
func compareWithBaseline(current, baseline metrics.Result, k int) ([]MetricDelta, bool) {
	deltas := []MetricDelta{
		{
			Name:       fmt.Sprintf("Recall@%d", k),
			Baseline:   baseline.RecallAtK,
			Current:    current.RecallAtK,
			HigherGood: true,
		},
		{
			Name:       "MRR",
			Baseline:   baseline.MRR,
			Current:    current.MRR,
			HigherGood: true,
		},
		{
			Name:       "Coverage",
			Baseline:   baseline.Coverage,
			Current:    current.Coverage,
			HigherGood: true,
		},
		{
			Name:       "Latency p95",
			Baseline:   baseline.LatencyP95,
			Current:    current.LatencyP95,
			HigherGood: false,
		},
	}

	hasRegression := false

	for i := range deltas {
		d := &deltas[i]
		d.Delta = d.Current - d.Baseline

		if d.Baseline != 0 {
			d.DeltaPct = (d.Delta / d.Baseline) * 100
		}

		// Determine if this is a regression
		if d.HigherGood {
			d.Regressed = d.Delta < -0.001 // Allow tiny floating point tolerance
		} else {
			d.Regressed = d.Delta > 0.001 // For latency, higher is worse
		}

		if d.Regressed {
			hasRegression = true
		}
	}

	return deltas, hasRegression
}

// printBaselineComparison displays the comparison between current and baseline runs.
func printBaselineComparison(deltas []MetricDelta, baselineTimestamp string) {
	fmt.Println()
	fmt.Println("─────────────────────────────────────────────────────────────")
	fmt.Println("BASELINE COMPARISON")
	fmt.Printf("Comparing against: %s\n", baselineTimestamp)
	fmt.Println("─────────────────────────────────────────────────────────────")

	for _, d := range deltas {
		var arrow string
		var status string

		if d.HigherGood {
			if d.Delta > 0.001 {
				arrow = "↑"
				status = "improved"
			} else if d.Delta < -0.001 {
				arrow = "↓"
				status = "REGRESSED"
			} else {
				arrow = "="
				status = "unchanged"
			}
		} else {
			// For latency, lower is better
			if d.Delta < -0.001 {
				arrow = "↓"
				status = "improved"
			} else if d.Delta > 0.001 {
				arrow = "↑"
				status = "REGRESSED"
			} else {
				arrow = "="
				status = "unchanged"
			}
		}

		// Format based on metric type
		if d.Name == "Latency p95" {
			fmt.Printf("  %-12s %.0fms → %.0fms  %s %.1f%%  (%s)\n",
				d.Name+":", d.Baseline, d.Current, arrow, abs(d.DeltaPct), status)
		} else {
			fmt.Printf("  %-12s %.3f → %.3f  %s %.1f%%  (%s)\n",
				d.Name+":", d.Baseline, d.Current, arrow, abs(d.DeltaPct), status)
		}
	}

	fmt.Println("─────────────────────────────────────────────────────────────")
}

// abs returns the absolute value of a float64.
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// RegressionError is returned when metrics regress and --fail-on-regression is set.
type RegressionError struct {
	Regressions []string
}

func (e *RegressionError) Error() string {
	return fmt.Sprintf("metrics regressed: %v", e.Regressions)
}

// checkRegressions returns an error if any regressions occurred.
func checkRegressions(deltas []MetricDelta) error {
	var regressions []string
	for _, d := range deltas {
		if d.Regressed {
			regressions = append(regressions, d.Name)
		}
	}

	if len(regressions) > 0 {
		if !jsonOutput {
			fmt.Println()
			fmt.Println("❌ REGRESSION DETECTED")
			fmt.Printf("   The following metrics decreased: %v\n", regressions)
		}
		return &RegressionError{Regressions: regressions}
	}

	if !jsonOutput {
		fmt.Println()
		fmt.Println("✅ No regressions detected")
	}
	return nil
}

// JSONOutput represents the machine-readable output format for CI pipelines.
type JSONOutput struct {
	Status     string             `json:"status"` // "pass" or "fail"
	Timestamp  string             `json:"timestamp"`
	Collection string             `json:"collection"`
	Store      string             `json:"store"`
	Metrics    JSONMetrics        `json:"metrics"`
	Bootstrap  *JSONBootstrap     `json:"bootstrap,omitempty"`
	Baseline   *JSONBaseline      `json:"baseline,omitempty"`
	Thresholds *JSONThresholds    `json:"thresholds,omitempty"`
	Failures   []JSONQueryFailure `json:"failures,omitempty"`
	RunFile    string             `json:"run_file"`
}

// JSONMetrics contains the core metrics in JSON format.
type JSONMetrics struct {
	RecallAtK  float64 `json:"recall_at_k"`
	MRR        float64 `json:"mrr"`
	NDCGAtK    float64 `json:"ndcg_at_k"`
	Coverage   float64 `json:"coverage"`
	Redundancy float64 `json:"redundancy"`
	LatencyP50 float64 `json:"latency_p50_ms"`
	LatencyP95 float64 `json:"latency_p95_ms"`
	LatencyP99 float64 `json:"latency_p99_ms"`
	LatencyAvg float64 `json:"latency_avg_ms"`
	QueryCount int     `json:"query_count"`
	TopK       int     `json:"top_k"`
}

// JSONBootstrap contains bootstrap confidence interval data.
type JSONBootstrap struct {
	N            int     `json:"n"`
	RecallMean   float64 `json:"recall_mean"`
	RecallStd    float64 `json:"recall_std"`
	RecallCI95Lo float64 `json:"recall_ci95_lo"`
	RecallCI95Hi float64 `json:"recall_ci95_hi"`
	MRRMean      float64 `json:"mrr_mean"`
	MRRStd       float64 `json:"mrr_std"`
	MRRCI95Lo    float64 `json:"mrr_ci95_lo"`
	MRRCI95Hi    float64 `json:"mrr_ci95_hi"`
	NDCGMean     float64 `json:"ndcg_mean"`
	NDCGStd      float64 `json:"ndcg_std"`
	CoverageMean float64 `json:"coverage_mean"`
	CoverageStd  float64 `json:"coverage_std"`
}

// JSONBaseline contains baseline comparison data.
type JSONBaseline struct {
	Timestamp   string           `json:"timestamp"`
	Comparisons []JSONComparison `json:"comparisons"`
	Regressed   bool             `json:"regressed"`
	Regressions []string         `json:"regressions,omitempty"`
}

// JSONComparison represents a single metric comparison.
type JSONComparison struct {
	Metric    string  `json:"metric"`
	Baseline  float64 `json:"baseline"`
	Current   float64 `json:"current"`
	Delta     float64 `json:"delta"`
	DeltaPct  float64 `json:"delta_pct"`
	Regressed bool    `json:"regressed"`
}

// JSONThresholds contains CI threshold results.
type JSONThresholds struct {
	Checks []JSONThresholdCheck `json:"checks"`
	Passed bool                 `json:"passed"`
}

// JSONThresholdCheck represents a single threshold check.
type JSONThresholdCheck struct {
	Metric    string  `json:"metric"`
	Value     float64 `json:"value"`
	Threshold float64 `json:"threshold"`
	Passed    bool    `json:"passed"`
}

// JSONQueryFailure represents a failed query in JSON output.
type JSONQueryFailure struct {
	QueryID       string   `json:"query_id"`
	Query         string   `json:"query"`
	ExpectedDocs  []string `json:"expected_docs"`
	RetrievedDocs []string `json:"retrieved_docs"`
}

// buildJSONOutput creates the JSON output structure from run results.
func buildJSONOutput(result RunResult, runPath string, failures []QueryFailure,
	deltas []MetricDelta, hasRegression bool, baselineTs string,
	thresholdChecks []JSONThresholdCheck, thresholdsPassed bool) JSONOutput {

	cfg := result.Configs[0]
	m := cfg.Metrics

	output := JSONOutput{
		Status:     "pass",
		Timestamp:  result.Timestamp,
		Collection: result.Collection,
		Store:      result.Store,
		RunFile:    runPath,
		Metrics: JSONMetrics{
			RecallAtK:  m.RecallAtK,
			MRR:        m.MRR,
			NDCGAtK:    m.NDCGAtK,
			Coverage:   m.Coverage,
			Redundancy: m.Redundancy,
			LatencyP50: m.LatencyP50,
			LatencyP95: m.LatencyP95,
			LatencyP99: m.LatencyP99,
			LatencyAvg: m.LatencyAvg,
			QueryCount: len(cfg.QueryResults),
			TopK:       cfg.Config.TopK,
		},
	}

	// Add bootstrap data if available
	if cfg.Bootstrap != nil {
		bs := cfg.Bootstrap
		output.Bootstrap = &JSONBootstrap{
			N:            bs.N,
			RecallMean:   bs.RecallMean,
			RecallStd:    bs.RecallStd,
			RecallCI95Lo: bs.RecallCI95Lo,
			RecallCI95Hi: bs.RecallCI95Hi,
			MRRMean:      bs.MRRMean,
			MRRStd:       bs.MRRStd,
			MRRCI95Lo:    bs.MRRCI95Lo,
			MRRCI95Hi:    bs.MRRCI95Hi,
			NDCGMean:     bs.NDCGMean,
			NDCGStd:      bs.NDCGStd,
			CoverageMean: bs.CoverageMean,
			CoverageStd:  bs.CoverageStd,
		}
	}

	// Add failures
	for _, f := range failures {
		output.Failures = append(output.Failures, JSONQueryFailure{
			QueryID:       f.QueryID,
			Query:         f.Query,
			ExpectedDocs:  f.RelevantDocs,
			RetrievedDocs: f.RetrievedDocs,
		})
	}

	// Add baseline comparison if available
	if len(deltas) > 0 {
		var comparisons []JSONComparison
		var regressions []string
		for _, d := range deltas {
			comparisons = append(comparisons, JSONComparison{
				Metric:    d.Name,
				Baseline:  d.Baseline,
				Current:   d.Current,
				Delta:     d.Delta,
				DeltaPct:  d.DeltaPct,
				Regressed: d.Regressed,
			})
			if d.Regressed {
				regressions = append(regressions, d.Name)
			}
		}
		output.Baseline = &JSONBaseline{
			Timestamp:   baselineTs,
			Comparisons: comparisons,
			Regressed:   hasRegression,
			Regressions: regressions,
		}
	}

	// Add threshold checks if in CI mode
	if len(thresholdChecks) > 0 {
		output.Thresholds = &JSONThresholds{
			Checks: thresholdChecks,
			Passed: thresholdsPassed,
		}
	}

	// Set status based on failures
	if hasRegression && failOnRegression {
		output.Status = "fail"
	}
	if !thresholdsPassed {
		output.Status = "fail"
	}

	return output
}

// printJSONOutput marshals and prints the JSON output.
func printJSONOutput(output JSONOutput) {
	data, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to marshal JSON output: %v\n", err)
		return
	}
	fmt.Println(string(data))
}

// normalizeSource maps a chunk-level source filename back to its parent document
// if the parent is among the relevant docs. This enables pre-chunked corpora
// (where each chunkset is a separate file like "doc_cs0042.txt") to match
// relevant docs specified at the document level (like "doc.txt").
//
// The matching is prefix-based: if a relevant doc's stem (without extension)
// is a prefix of the source's stem, the source is normalized to the relevant doc name.
// If no match is found, the source is returned unchanged.
func normalizeSource(source string, relevantDocs []string) string {
	// Fast path: exact match
	for _, rd := range relevantDocs {
		if source == rd {
			return source
		}
	}

	// Prefix match: strip extensions and check if relevant doc stem is a prefix
	sourceBase := strings.TrimSuffix(source, filepath.Ext(source))
	for _, rd := range relevantDocs {
		rdBase := strings.TrimSuffix(rd, filepath.Ext(rd))
		if strings.HasPrefix(sourceBase, rdBase) {
			return rd
		}
	}

	return source
}
