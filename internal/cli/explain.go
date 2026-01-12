package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// Diagnostic thresholds for score analysis.
const (
	// scoreThresholdLow indicates queries may be out-of-domain
	scoreThresholdLow = 0.5

	// scoreThresholdStrong indicates high-confidence retrieval
	scoreThresholdStrong = 0.85

	// spreadThresholdTight indicates results are nearly indistinguishable
	spreadThresholdTight = 0.05

	// spreadThresholdHigh indicates significant relevance variance
	spreadThresholdHigh = 0.3

	// stdDevThresholdTight indicates poor discrimination between chunks
	stdDevThresholdTight = 0.02

	// topGapThresholdLarge may indicate outlier top result
	topGapThresholdLarge = 0.15

	// stdDevThresholdShape classifies distribution as "tight"
	stdDevThresholdShape = 0.03

	// spreadThresholdShape classifies distribution as "spread"
	spreadThresholdShape = 0.35
)

var (
	saveQuery   bool
	goldenFile  string
	relevantDoc string
)

var explainCmd = &cobra.Command{
	Use:   "explain <query>",
	Short: "Show top-k retrieved chunks for a query with score distribution analysis",
	Long: `Explain RAG retrieval for a single query.

Shows the top-k retrieved chunks, their similarity scores, source documents,
and comprehensive diagnostic analysis.

Diagnostics include:
  • Score statistics  - Range, mean, standard deviation
  • Quartiles         - Q1, median (Q2), Q3 for distribution shape
  • Top gap           - Distance between #1 and #2 results
  • Distribution type - Classified as tight/spread/bimodal/normal
  • Insights          - Positive observations (strong match, good separation)
  • Warnings          - Issues requiring attention (low scores, poor discrimination)

Use --save to add this query to your golden queries file for regression testing.
Think of it like bookmarking queries in Postman.

Examples:
  ragtune explain "How to rotate API key?" --collection demo --top-k 5

  # Save as golden query (infers relevant doc from top result)
  ragtune explain "How to reset password?" --collection prod --save

  # Save with explicit relevant doc
  ragtune explain "How to reset password?" --collection prod --save --relevant docs/auth.md`,
	Args: cobra.ExactArgs(1),
	RunE: runExplain,
}

func init() {
	explainCmd.Flags().BoolVar(&saveQuery, "save", false, "Save query to golden queries file")
	explainCmd.Flags().StringVar(&goldenFile, "golden-file", "golden-queries.json", "Path to golden queries file")
	explainCmd.Flags().StringVar(&relevantDoc, "relevant", "", "Relevant doc (inferred from top result if not specified)")
}

func runExplain(cmd *cobra.Command, args []string) error {
	query := args[0]

	if collectionName == "" {
		return fmt.Errorf("--collection is required")
	}

	ctx := context.Background()

	// Initialize vector store
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

	// Embed query
	fmt.Printf("Query: %q\n", query)
	fmt.Println("Generating query embedding...")
	queryVec, err := emb.Embed(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to embed query: %w", err)
	}

	// Search
	fmt.Printf("Searching collection '%s' on %s (top-k=%d)...\n", collectionName, storeName, topK)
	results, err := store.Search(ctx, collectionName, queryVec, topK)
	if err != nil {
		return fmt.Errorf("search failed: %w", err)
	}

	// Display results
	fmt.Println()
	fmt.Println(strings.Repeat("=", 80))
	fmt.Printf("EXPLAIN RETRIEVAL: %d results\n", len(results))
	fmt.Println(strings.Repeat("=", 80))

	if len(results) == 0 {
		fmt.Println("⚠ No results found. Collection may be empty or query embedding failed.")
		return nil
	}

	// Compute diagnostics
	var scores []float32
	for _, r := range results {
		scores = append(scores, r.Score)
	}
	diag := computeDiagnostics(scores)

	for i, r := range results {
		fmt.Println()
		fmt.Printf("[%d] Score: %.4f | ID: %s\n", i+1, r.Score, r.ID)
		fmt.Printf("    Source: %s\n", getPayloadString(r.Payload, "source"))

		text := getPayloadString(r.Payload, "text")
		fmt.Printf("    Text: %s\n", truncate(text, 200))
	}

	// Diagnostics section
	fmt.Println()
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("DIAGNOSTICS")
	fmt.Println(strings.Repeat("-", 80))

	// Basic stats
	fmt.Printf("Score range:  %.4f - %.4f (spread: %.4f)\n", diag.minScore, diag.maxScore, diag.spread)
	fmt.Printf("Mean score:   %.4f\n", diag.meanScore)

	// Enhanced distribution stats
	fmt.Printf("Std dev:      %.4f\n", diag.stdDev)
	fmt.Printf("Quartiles:    Q1=%.4f  Median=%.4f  Q3=%.4f\n", diag.q1, diag.median, diag.q3)
	if diag.topGap > 0 {
		fmt.Printf("Top gap:      %.4f (distance between #1 and #2)\n", diag.topGap)
	}
	fmt.Printf("Distribution: %s\n", diag.scoreShape)

	// Insights (positive observations)
	if len(diag.insights) > 0 {
		fmt.Println()
		fmt.Println("Insights:")
		for _, i := range diag.insights {
			fmt.Printf("  ✓ %s\n", i)
		}
	}

	// Warnings
	if len(diag.warnings) > 0 {
		fmt.Println()
		fmt.Println("Warnings:")
		for _, w := range diag.warnings {
			fmt.Printf("  ⚠ %s\n", w)
		}
	} else if len(diag.insights) == 0 {
		fmt.Println()
		fmt.Println("✓ No warnings")
	}

	// Save as golden query if requested
	if saveQuery {
		if len(results) == 0 {
			return fmt.Errorf("cannot save: no results to infer relevant doc from")
		}

		// Determine relevant doc
		relDoc := relevantDoc
		if relDoc == "" {
			// Infer from top result
			relDoc = getPayloadString(results[0].Payload, "source")
		}

		if err := appendGoldenQuery(goldenFile, query, relDoc); err != nil {
			return fmt.Errorf("failed to save golden query: %w", err)
		}
		fmt.Printf("\n✓ Saved to %s (relevant: %s)\n", goldenFile, relDoc)
	}

	return nil
}

type diagnostics struct {
	minScore   float32
	maxScore   float32
	meanScore  float32
	spread     float32
	stdDev     float32
	median     float32
	q1         float32 // 25th percentile
	q3         float32 // 75th percentile
	topGap     float32 // Gap between #1 and #2 scores
	scoreShape string  // "tight", "spread", "bimodal", "normal"
	warnings   []string
	insights   []string // Positive observations
}

func computeDiagnostics(scores []float32) diagnostics {
	if len(scores) == 0 {
		return diagnostics{}
	}

	var d diagnostics
	d.minScore = scores[0]
	d.maxScore = scores[0]
	var sum float32

	for _, s := range scores {
		if s < d.minScore {
			d.minScore = s
		}
		if s > d.maxScore {
			d.maxScore = s
		}
		sum += s
	}

	d.meanScore = sum / float32(len(scores))
	d.spread = d.maxScore - d.minScore

	// Compute standard deviation
	var variance float32
	for _, s := range scores {
		diff := s - d.meanScore
		variance += diff * diff
	}
	variance /= float32(len(scores))
	d.stdDev = float32(math.Sqrt(float64(variance)))

	// Compute percentiles (need sorted copy)
	sorted := make([]float32, len(scores))
	copy(sorted, scores)
	slices.Sort(sorted)

	d.median = percentileFloat32(sorted, 50)
	d.q1 = percentileFloat32(sorted, 25)
	d.q3 = percentileFloat32(sorted, 75)

	// Gap between top two results
	if len(scores) >= 2 {
		d.topGap = scores[0] - scores[1]
	}

	// Determine score distribution shape
	d.scoreShape = classifyScoreShape(d)

	// Generate warnings
	if d.maxScore < scoreThresholdLow {
		d.warnings = append(d.warnings, "Low top score (<0.5): query may be out-of-domain or embeddings mismatched")
	}
	if d.spread > spreadThresholdHigh {
		d.warnings = append(d.warnings, "High score spread (>0.3): results vary significantly in relevance")
	}
	if d.spread < spreadThresholdTight && len(scores) > 1 {
		d.warnings = append(d.warnings, "Very low spread (<0.05): results are nearly indistinguishable, consider reviewing chunking")
	}
	if d.topGap > topGapThresholdLarge && len(scores) >= 2 {
		d.warnings = append(d.warnings, fmt.Sprintf("Large gap (%.2f) between #1 and #2: verify top result is truly best match", d.topGap))
	}
	if d.stdDev < stdDevThresholdTight && len(scores) > 2 {
		d.warnings = append(d.warnings, "Very tight distribution (σ<0.02): retrieval may not discriminate well between chunks")
	}

	// Generate positive insights
	if d.maxScore > scoreThresholdStrong {
		d.insights = append(d.insights, "Strong top match (>0.85): likely high-quality retrieval")
	}
	if d.spread > 0.1 && d.spread < 0.25 && d.maxScore > 0.7 {
		d.insights = append(d.insights, "Good score separation: retrieval is discriminating effectively")
	}
	if d.topGap > 0.05 && d.topGap < topGapThresholdLarge && d.maxScore > 0.75 {
		d.insights = append(d.insights, "Clear top result with gradual falloff: healthy ranking")
	}

	return d
}

// classifyScoreShape determines the distribution pattern of scores
func classifyScoreShape(d diagnostics) string {
	if d.stdDev < stdDevThresholdShape {
		return "tight" // All scores clustered together
	}
	if d.spread > spreadThresholdShape {
		return "spread" // Wide range of scores
	}
	// Check for bimodal (gap in middle)
	iqr := d.q3 - d.q1
	if iqr < d.spread*0.3 {
		return "bimodal" // Scores clustered at extremes
	}
	return "normal" // Typical distribution
}

// percentileFloat32 calculates the p-th percentile of a sorted float32 slice
func percentileFloat32(sorted []float32, p int) float32 {
	if len(sorted) == 0 {
		return 0
	}
	if len(sorted) == 1 {
		return sorted[0]
	}

	rank := float64(p) / 100.0 * float64(len(sorted)-1)
	lower := int(rank)
	upper := lower + 1
	if upper >= len(sorted) {
		return sorted[len(sorted)-1]
	}

	weight := float32(rank - float64(lower))
	return sorted[lower]*(1-weight) + sorted[upper]*weight
}

func getPayloadString(payload map[string]interface{}, key string) string {
	if v, ok := payload[key]; ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return "<unknown>"
}

func truncate(s string, maxLen int) string {
	// Replace newlines with spaces for display
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.Join(strings.Fields(s), " ") // Normalize whitespace

	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// GoldenQueries represents the golden queries file format.
type GoldenQueries struct {
	Queries []GoldenQuery `json:"queries"`
}

// GoldenQuery represents a single golden query.
type GoldenQuery struct {
	ID           string   `json:"id"`
	Text         string   `json:"text"`
	RelevantDocs []string `json:"relevant_docs"`
}

// appendGoldenQuery adds a query to the golden queries file.
func appendGoldenQuery(filePath, queryText, relevantDoc string) error {
	var gq GoldenQueries

	// Load existing file if it exists
	if data, err := os.ReadFile(filePath); err == nil {
		if err := json.Unmarshal(data, &gq); err != nil {
			return fmt.Errorf("failed to parse existing golden queries: %w", err)
		}
	}

	// Check for duplicate query text
	for _, q := range gq.Queries {
		if q.Text == queryText {
			return fmt.Errorf("query already exists: %q", queryText)
		}
	}

	// Generate unique ID
	id := fmt.Sprintf("q-%d", time.Now().UnixNano())

	// Normalize relevant doc path
	relevantDoc = filepath.Base(relevantDoc)

	// Add new query
	gq.Queries = append(gq.Queries, GoldenQuery{
		ID:           id,
		Text:         queryText,
		RelevantDocs: []string{relevantDoc},
	})

	// Write back
	data, err := json.MarshalIndent(gq, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal golden queries: %w", err)
	}

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write golden queries file: %w", err)
	}

	return nil
}
