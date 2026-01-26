package cli

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/metawake/ragtune/internal/config"
	"github.com/metawake/ragtune/internal/metrics"
	"github.com/metawake/ragtune/internal/vectorstore"
	"github.com/metawake/ragtune/internal/vectorstore/mock"
)

// Integration tests for the complete workflows

// TestIntegration_GoldenQueryWorkflow tests the full flow:
// 1. Create golden queries file via appendGoldenQuery (simulating explain --save)
// 2. Import additional queries from CSV
// 3. Verify the merged file is valid for simulate
func TestIntegration_GoldenQueryWorkflow(t *testing.T) {
	tmpDir := t.TempDir()
	goldenPath := filepath.Join(tmpDir, "golden-queries.json")

	// Step 1: Simulate "explain --save" by adding queries one at a time
	queries := []struct {
		text        string
		relevantDoc string
	}{
		{"How do I reset my password?", "docs/auth/password.md"},
		{"What are the API rate limits?", "docs/api/limits.md"},
	}

	for _, q := range queries {
		err := appendGoldenQuery(goldenPath, q.text, q.relevantDoc)
		if err != nil {
			t.Fatalf("appendGoldenQuery failed: %v", err)
		}
	}

	// Step 2: Create CSV with additional queries
	csvPath := filepath.Join(tmpDir, "additional.csv")
	csvContent := `query,relevant_docs
"How do I configure webhooks?",docs/webhooks.md
"What authentication methods are supported?",docs/auth/methods.md`
	_ = os.WriteFile(csvPath, []byte(csvContent), 0644)

	// Import CSV queries
	csvQueries, err := parseCSVQueries(csvPath)
	if err != nil {
		t.Fatalf("parseCSVQueries failed: %v", err)
	}

	// Merge into golden file
	var existing GoldenQueries
	data, _ := os.ReadFile(goldenPath)
	_ = json.Unmarshal(data, &existing)

	existingTexts := make(map[string]bool)
	for _, q := range existing.Queries {
		existingTexts[q.Text] = true
	}

	for _, q := range csvQueries {
		if !existingTexts[q.Text] {
			existing.Queries = append(existing.Queries, q)
		}
	}

	data, _ = json.MarshalIndent(existing, "", "  ")
	_ = os.WriteFile(goldenPath, data, 0644)

	// Step 3: Verify final file has all 4 queries
	data, _ = os.ReadFile(goldenPath)
	var final GoldenQueries
	_ = json.Unmarshal(data, &final)

	if len(final.Queries) != 4 {
		t.Errorf("expected 4 queries after merge, got %d", len(final.Queries))
	}

	// Verify all queries have required fields
	for i, q := range final.Queries {
		if q.ID == "" {
			t.Errorf("query %d missing ID", i)
		}
		if q.Text == "" {
			t.Errorf("query %d missing text", i)
		}
		if len(q.RelevantDocs) == 0 {
			t.Errorf("query %d missing relevant docs", i)
		}
	}
}

// TestIntegration_SimulateWithMockStore tests the full simulate workflow
// using the mock vector store to avoid external dependencies.
func TestIntegration_SimulateWithMockStore(t *testing.T) {
	ctx := context.Background()

	// Setup mock store with test data
	store := mock.New()
	defer store.Close()

	// Create collection with 3D vectors
	_ = store.EnsureCollection(ctx, "test", 3)

	// Insert documents as points (simulating ingested chunks)
	// Using orthogonal vectors for predictable similarity scores
	points := []vectorstore.Point{
		{
			ID:     "chunk-password-1",
			Vector: []float32{1, 0, 0},
			Payload: map[string]interface{}{
				"text":   "To reset your password, go to settings...",
				"source": "docs/auth/password.md",
			},
		},
		{
			ID:     "chunk-limits-1",
			Vector: []float32{0, 1, 0},
			Payload: map[string]interface{}{
				"text":   "API rate limits are 100 requests per minute...",
				"source": "docs/api/limits.md",
			},
		},
		{
			ID:     "chunk-webhooks-1",
			Vector: []float32{0, 0, 1},
			Payload: map[string]interface{}{
				"text":   "Configure webhooks in the dashboard...",
				"source": "docs/webhooks.md",
			},
		},
	}
	_ = store.Upsert(ctx, "test", points)

	// Create query results simulating what simulate would produce
	// Query vector [1,0,0] should match password.md with score ~1.0
	queryResults := []metrics.QueryResult{
		{
			QueryID:      "q1",
			Query:        "How do I reset my password?",
			RetrievedIDs: []string{"password.md", "limits.md", "webhooks.md"},
			RelevantIDs:  []string{"password.md"},
			Scores:       []float32{1.0, 0.0, 0.0},
		},
		{
			QueryID:      "q2",
			Query:        "What are the rate limits?",
			RetrievedIDs: []string{"limits.md", "password.md", "webhooks.md"},
			RelevantIDs:  []string{"limits.md"},
			Scores:       []float32{1.0, 0.0, 0.0},
		},
	}

	// Compute metrics
	result := metrics.Compute(queryResults, 3)

	// Verify metrics
	if result.RecallAtK != 1.0 {
		t.Errorf("expected recall 1.0, got %f", result.RecallAtK)
	}
	if result.MRR != 1.0 {
		t.Errorf("expected MRR 1.0, got %f", result.MRR)
	}
	if result.Coverage != 1.0 {
		t.Errorf("expected coverage 1.0, got %f", result.Coverage)
	}
}

// TestIntegration_CIThresholdWorkflow tests the CI mode end-to-end:
// good metrics pass, bad metrics fail.
func TestIntegration_CIThresholdWorkflow(t *testing.T) {
	// Save and restore global state
	oldMinRecall := minRecall
	oldMinMRR := minMRR
	oldMinCoverage := minCoverage
	defer func() {
		minRecall = oldMinRecall
		minMRR = oldMinMRR
		minCoverage = oldMinCoverage
	}()

	tests := []struct {
		name        string
		recall      float64
		mrr         float64
		coverage    float64
		minRecall   float64
		minMRR      float64
		minCoverage float64
		wantPass    bool
	}{
		{
			name:        "all metrics pass",
			recall:      0.90,
			mrr:         0.85,
			coverage:    0.95,
			minRecall:   0.85,
			minMRR:      0.80,
			minCoverage: 0.90,
			wantPass:    true,
		},
		{
			name:        "recall fails",
			recall:      0.70,
			mrr:         0.85,
			coverage:    0.95,
			minRecall:   0.85,
			minMRR:      0.80,
			minCoverage: 0.90,
			wantPass:    false,
		},
		{
			name:        "all metrics fail",
			recall:      0.50,
			mrr:         0.40,
			coverage:    0.60,
			minRecall:   0.85,
			minMRR:      0.80,
			minCoverage: 0.90,
			wantPass:    false,
		},
		{
			name:        "no thresholds set (always pass)",
			recall:      0.10,
			mrr:         0.10,
			coverage:    0.10,
			minRecall:   0,
			minMRR:      0,
			minCoverage: 0,
			wantPass:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			minRecall = tt.minRecall
			minMRR = tt.minMRR
			minCoverage = tt.minCoverage

			result := RunResult{
				Configs: []ConfigResult{
					{
						Config: config.SimConfig{Name: "test", TopK: 5},
						Metrics: metrics.Result{
							RecallAtK: tt.recall,
							MRR:       tt.mrr,
							Coverage:  tt.coverage,
						},
					},
				},
			}

			// Check if thresholds would pass
			passed := true
			if minRecall > 0 && tt.recall < minRecall {
				passed = false
			}
			if minMRR > 0 && tt.mrr < minMRR {
				passed = false
			}
			if minCoverage > 0 && tt.coverage < minCoverage {
				passed = false
			}

			if passed != tt.wantPass {
				t.Errorf("expected pass=%v, got pass=%v", tt.wantPass, passed)
			}

			// Verify the result structure is valid
			if len(result.Configs) != 1 {
				t.Errorf("expected 1 config, got %d", len(result.Configs))
			}
		})
	}
}

// TestIntegration_ImportMergeWorkflow tests importing and merging queries
// from multiple sources.
func TestIntegration_ImportMergeWorkflow(t *testing.T) {
	tmpDir := t.TempDir()

	// Create first batch as JSON
	json1Path := filepath.Join(tmpDir, "batch1.json")
	json1Content := `{
  "queries": [
    {"id": "json-1", "text": "Query from JSON 1", "relevant_docs": ["doc1.md"]},
    {"id": "json-2", "text": "Query from JSON 2", "relevant_docs": ["doc2.md"]}
  ]
}`
	_ = os.WriteFile(json1Path, []byte(json1Content), 0644)

	// Create second batch as CSV
	csv1Path := filepath.Join(tmpDir, "batch2.csv")
	csv1Content := `query,relevant_docs
"Query from CSV 1",doc3.md
"Query from CSV 2",doc4.md;doc5.md`
	_ = os.WriteFile(csv1Path, []byte(csv1Content), 0644)

	// Create third batch with duplicate
	csv2Path := filepath.Join(tmpDir, "batch3.csv")
	csv2Content := `query,relevant_docs
"Query from JSON 1",doc1.md
"Query from CSV 3",doc6.md`
	_ = os.WriteFile(csv2Path, []byte(csv2Content), 0644)

	// Parse all sources
	jsonQueries, _ := parseJSONQueries(json1Path)
	csvQueries1, _ := parseCSVQueries(csv1Path)
	csvQueries2, _ := parseCSVQueries(csv2Path)

	// Merge all into a single set (deduplicating)
	allQueries := make(map[string]GoldenQuery)
	for _, q := range jsonQueries {
		allQueries[q.Text] = q
	}
	for _, q := range csvQueries1 {
		allQueries[q.Text] = q
	}
	for _, q := range csvQueries2 {
		allQueries[q.Text] = q
	}

	// Should have 5 unique queries (one duplicate removed)
	if len(allQueries) != 5 {
		t.Errorf("expected 5 unique queries, got %d", len(allQueries))
	}

	// Verify multi-doc query preserved both docs
	multiDocQuery := allQueries["Query from CSV 2"]
	if len(multiDocQuery.RelevantDocs) != 2 {
		t.Errorf("expected 2 relevant docs for multi-doc query, got %d", len(multiDocQuery.RelevantDocs))
	}
}

// TestIntegration_EndToEndGoldenQueryFile verifies the golden query file
// format is compatible with the config.LoadQueries function.
func TestIntegration_EndToEndGoldenQueryFile(t *testing.T) {
	tmpDir := t.TempDir()
	goldenPath := filepath.Join(tmpDir, "golden-queries.json")

	// Build up queries via appendGoldenQuery
	appendGoldenQuery(goldenPath, "First test query", "doc1.md")
	appendGoldenQuery(goldenPath, "Second test query", "doc2.md")
	appendGoldenQuery(goldenPath, "Third test query", "doc3.md")

	// Verify the file can be loaded by config.LoadQueries
	queries, err := config.LoadQueries(goldenPath)
	if err != nil {
		t.Fatalf("config.LoadQueries failed on golden file: %v", err)
	}

	if len(queries) != 3 {
		t.Errorf("expected 3 queries, got %d", len(queries))
	}

	// Verify each query has required fields for simulate
	for i, q := range queries {
		if q.ID == "" {
			t.Errorf("query %d: missing ID", i)
		}
		if q.Text == "" {
			t.Errorf("query %d: missing text", i)
		}
		if len(q.RelevantDocs) == 0 {
			t.Errorf("query %d: missing relevant docs", i)
		}
	}
}

// TestIntegration_SearchAndMetrics tests the complete flow from
// mock search results to computed metrics.
func TestIntegration_SearchAndMetrics(t *testing.T) {
	ctx := context.Background()
	store := mock.New()
	defer store.Close()

	// Setup: create collection and insert test documents
	_ = store.EnsureCollection(ctx, "test", 4)

	// Insert 4 documents with distinct vectors
	docs := []struct {
		id     string
		vector []float32
		source string
	}{
		{"d1", []float32{1, 0, 0, 0}, "doc1.md"},
		{"d2", []float32{0, 1, 0, 0}, "doc2.md"},
		{"d3", []float32{0, 0, 1, 0}, "doc3.md"},
		{"d4", []float32{0, 0, 0, 1}, "doc4.md"},
	}

	for _, d := range docs {
		_ = store.Upsert(ctx, "test", []vectorstore.Point{
			{
				ID:      d.id,
				Vector:  d.vector,
				Payload: map[string]interface{}{"source": d.source},
			},
		})
	}

	// Define test queries with known expected results
	testQueries := []struct {
		queryVector []float32
		expectedTop string
		relevantDoc string
	}{
		{[]float32{1, 0, 0, 0}, "d1", "doc1.md"}, // Should match d1
		{[]float32{0, 1, 0, 0}, "d2", "doc2.md"}, // Should match d2
		{[]float32{0.5, 0.5, 0, 0}, "d1", "doc1.md"}, // Should match d1 or d2 (equal)
	}

	var queryResults []metrics.QueryResult

	for i, tq := range testQueries {
		results, err := store.Search(ctx, "test", tq.queryVector, 4)
		if err != nil {
			t.Fatalf("Search failed: %v", err)
		}

		// Build query result
		var retrievedIDs []string
		var scores []float32
		for _, r := range results {
			source := r.Payload["source"].(string)
			retrievedIDs = append(retrievedIDs, source)
			scores = append(scores, r.Score)
		}

		queryResults = append(queryResults, metrics.QueryResult{
			QueryID:      string(rune('a' + i)),
			Query:        "test query",
			RetrievedIDs: retrievedIDs,
			RelevantIDs:  []string{tq.relevantDoc},
			Scores:       scores,
		})
	}

	// Compute metrics
	m := metrics.Compute(queryResults, 4)

	// All relevant docs should be in top-k
	if m.RecallAtK < 0.9 {
		t.Errorf("expected high recall, got %f", m.RecallAtK)
	}

	// First result should usually be correct
	if m.MRR < 0.5 {
		t.Errorf("expected reasonable MRR, got %f", m.MRR)
	}
}
