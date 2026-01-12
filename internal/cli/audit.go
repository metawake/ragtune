package cli

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"github.com/metawake/ragtune/internal/config"
	"github.com/metawake/ragtune/internal/metrics"
)

var (
	auditMinRecall     float64
	auditMinMRR        float64
	auditMinCoverage   float64
	auditMaxLatencyP95 float64
)

var auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Quick health check with pass/fail and recommendations",
	Long: `Run a RAG health check and get a pass/fail report with recommendations.

Audit is designed for quick assessment — it runs the same analysis as simulate
but presents results as a health report with actionable recommendations.

Default thresholds (override with flags):
  --min-recall 0.85       Recall@K threshold
  --min-mrr 0.70          MRR threshold  
  --min-coverage 0.90     Coverage threshold
  --max-latency-p95 0     p95 latency threshold in ms (0 = no limit)

Exit codes:
  0 = All checks passed
  1 = One or more checks failed

Examples:
  # Basic health check with default thresholds
  ragtune audit --collection prod --queries golden.json

  # Custom thresholds
  ragtune audit --collection prod --queries golden.json \
    --min-recall 0.90 --min-coverage 0.95

  # Use in CI (exit code indicates pass/fail)
  ragtune audit --collection prod --queries golden.json || exit 1`,
	RunE: runAudit,
}

func init() {
	auditCmd.Flags().StringVar(&queriesPath, "queries", "", "Path to queries JSON file (required)")
	auditCmd.Flags().Float64Var(&auditMinRecall, "min-recall", 0.85, "Minimum Recall@K threshold")
	auditCmd.Flags().Float64Var(&auditMinMRR, "min-mrr", 0.70, "Minimum MRR threshold")
	auditCmd.Flags().Float64Var(&auditMinCoverage, "min-coverage", 0.90, "Minimum Coverage threshold")
	auditCmd.Flags().Float64Var(&auditMaxLatencyP95, "max-latency-p95", 0, "Maximum p95 latency in ms (0 = no limit)")
	auditCmd.MarkFlagRequired("queries")

	rootCmd.AddCommand(auditCmd)
}

func runAudit(cmd *cobra.Command, args []string) error {
	if collectionName == "" {
		return fmt.Errorf("--collection is required")
	}

	ctx := context.Background()

	// Initialize store
	store, err := initVectorStore(ctx)
	if err != nil {
		return fmt.Errorf("failed to init vector store: %w", err)
	}
	defer store.Close()

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

	fmt.Printf("Running health check on '%s' (%s) with %d queries...\n\n", collectionName, storeName, len(queries))

	// Run queries and collect results
	var queryResults []metrics.QueryResult

	for _, q := range queries {
		queryStart := time.Now()

		vec, err := emb.Embed(ctx, q.Text)
		if err != nil {
			return fmt.Errorf("failed to embed query %s: %w", q.ID, err)
		}

		results, err := store.Search(ctx, collectionName, vec, topK)
		if err != nil {
			return fmt.Errorf("search failed for query %s: %w", q.ID, err)
		}

		latencyMs := float64(time.Since(queryStart).Microseconds()) / 1000.0

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
			LatencyMs:    latencyMs,
		})
	}

	// Compute metrics
	m := metrics.Compute(queryResults, topK)

	// Generate and print report
	return printAuditReport(collectionName, len(queries), m, auditMaxLatencyP95)
}

// auditCheck represents a single metric check result
type auditCheck struct {
	name           string
	value          float64
	threshold      float64
	passed         bool
	warning        bool
	isLatency      bool
	recommendation string
}

func printAuditReport(collection string, queryCount int, m metrics.Result, maxLatencyP95 float64) error {
	// Build checks
	checks := []auditCheck{
		buildCheck("Recall@K", m.RecallAtK, auditMinRecall, false,
			"Relevant documents not appearing in top results",
			"Try: chunk size, embedder, or top-k — use compare to measure"),
		buildCheck("MRR", m.MRR, auditMinMRR, false,
			"First relevant result not ranking high enough",
			"Try: reranking, chunk size — run compare to see what helps"),
		buildCheck("Coverage", m.Coverage, auditMinCoverage, false,
			"Some relevant documents never retrieved",
			"Check: is the content indexed? Try a different embedder"),
	}

	// Add latency check if threshold is set
	if maxLatencyP95 > 0 {
		checks = append(checks, buildCheck("Latency p95", m.LatencyP95, maxLatencyP95, true,
			"p95 latency exceeds threshold",
			"Try: faster embedder (TEI), GPU acceleration, or reduce corpus size"))
	}

	// Determine overall status
	status := "HEALTHY"
	statusIcon := "✅"
	failCount := 0
	warnCount := 0

	for _, c := range checks {
		if !c.passed {
			failCount++
		}
		if c.warning {
			warnCount++
		}
	}

	if failCount > 0 {
		status = "FAILING"
		statusIcon = "❌"
	} else if warnCount > 0 {
		status = "WARNING"
		statusIcon = "⚠️"
	}

	// Print report
	fmt.Println("╔════════════════════════════════════════════════════════════════╗")
	fmt.Printf("║  RAG HEALTH REPORT: %-42s ║\n", collection)
	fmt.Println("╠════════════════════════════════════════════════════════════════╣")

	// Query count with guidance
	queryGuidance := ""
	if queryCount < 20 {
		queryGuidance = "⚠ very low, results unreliable"
	} else if queryCount < 50 {
		queryGuidance = "⚠ low, recommend 50+"
	} else if queryCount < 100 {
		queryGuidance = "okay for CI"
	} else {
		queryGuidance = "✓ good sample size"
	}
	fmt.Printf("║  Queries:     %-4d  %-43s ║\n", queryCount, queryGuidance)
	
	fmt.Println("╠════════════════════════════════════════════════════════════════╣")

	// Metrics
	for _, c := range checks {
		statusStr := "✓ PASS"
		if !c.passed {
			statusStr = "✗ FAIL"
		} else if c.warning {
			statusStr = "⚠ WARN"
		}
		if c.isLatency {
			fmt.Printf("║  %-12s %.0fms  %-6s  (threshold: %.0fms)                  ║\n",
				c.name+":", c.value, statusStr, c.threshold)
		} else {
			fmt.Printf("║  %-12s %.3f  %-6s  (threshold: %.2f)                   ║\n",
				c.name+":", c.value, statusStr, c.threshold)
		}
	}

	// Latency stats (informational, only if not already showing latency check)
	if m.LatencyAvg > 0 && maxLatencyP95 == 0 {
		fmt.Println("╠════════════════════════════════════════════════════════════════╣")
		fmt.Printf("║  Latency:     p50=%.0fms  p95=%.0fms  p99=%.0fms  avg=%.0fms         ║\n",
			m.LatencyP50, m.LatencyP95, m.LatencyP99, m.LatencyAvg)
	}

	fmt.Println("╠════════════════════════════════════════════════════════════════╣")
	fmt.Printf("║  STATUS: %s %-53s ║\n", statusIcon, status)

	// Recommendations if any failures
	if failCount > 0 || warnCount > 0 {
		fmt.Println("╠════════════════════════════════════════════════════════════════╣")
		fmt.Println("║  RECOMMENDATIONS:                                              ║")
		for _, c := range checks {
			if !c.passed || c.warning {
				// Wrap recommendation text
				fmt.Printf("║  • %-60s ║\n", truncateStr(c.recommendation, 60))
			}
		}
	}

	fmt.Println("╚════════════════════════════════════════════════════════════════╝")

	// Return error if failing (let main.go handle exit code)
	if failCount > 0 {
		fmt.Println()
		fmt.Printf("Audit failed: %d metric(s) below threshold\n", failCount)
		return &AuditError{FailCount: failCount}
	}

	if warnCount > 0 {
		fmt.Println()
		fmt.Printf("Audit passed with %d warning(s)\n", warnCount)
	} else {
		fmt.Println()
		fmt.Println("Audit passed: all metrics within thresholds")
	}

	return nil
}

func buildCheck(name string, value, threshold float64, isLatency bool, issue, recommendation string) auditCheck {
	var passed, warning bool

	if isLatency {
		// For latency, lower is better (value should be <= threshold)
		passed = value <= threshold
		warning = passed && value > threshold*0.9 // Within 10% of threshold = warning
	} else {
		// For other metrics, higher is better (value should be >= threshold)
		passed = value >= threshold
		warning = passed && value < threshold*1.1 // Within 10% of threshold = warning
	}

	rec := recommendation
	if !passed {
		rec = issue + ". " + recommendation
	}

	return auditCheck{
		name:           name,
		value:          value,
		threshold:      threshold,
		passed:         passed,
		warning:        warning,
		isLatency:      isLatency,
		recommendation: rec,
	}
}

func truncateStr(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}
