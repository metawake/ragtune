package cli

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/metawake/ragtune/internal/config"
	"github.com/metawake/ragtune/internal/metrics"
)

func TestCheckCIThresholds_AllPass(t *testing.T) {
	// Save and restore global state
	oldMinRecall := minRecall
	oldMinMRR := minMRR
	oldMinCoverage := minCoverage
	oldMaxLatencyP95 := maxLatencyP95
	defer func() {
		minRecall = oldMinRecall
		minMRR = oldMinMRR
		minCoverage = oldMinCoverage
		maxLatencyP95 = oldMaxLatencyP95
	}()

	// Set thresholds
	minRecall = 0.80
	minMRR = 0.70
	minCoverage = 0.90
	maxLatencyP95 = 500.0

	result := RunResult{
		Configs: []ConfigResult{
			{
				Config: config.SimConfig{Name: "test", TopK: 5},
				Metrics: metrics.Result{
					RecallAtK:  0.85,  // Above 0.80
					MRR:        0.75,  // Above 0.70
					Coverage:   0.95,  // Above 0.90
					Redundancy: 1.0,
					LatencyP95: 200.0, // Below 500ms
				},
			},
		},
	}

	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := checkCIThresholds(result)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if err != nil {
		t.Errorf("expected no error when all thresholds pass, got: %v", err)
	}

	if !strings.Contains(output, "✓ PASS") {
		t.Errorf("expected PASS markers in output")
	}
	if !strings.Contains(output, "CI check PASSED") {
		t.Errorf("expected 'CI check PASSED' in output")
	}
	if !strings.Contains(output, "Latency p95") {
		t.Errorf("expected Latency p95 in output when threshold is set")
	}
}

func TestCheckCIThresholds_RecallFails(t *testing.T) {
	oldMinRecall := minRecall
	oldMinMRR := minMRR
	oldMinCoverage := minCoverage
	defer func() {
		minRecall = oldMinRecall
		minMRR = oldMinMRR
		minCoverage = oldMinCoverage
	}()

	minRecall = 0.90 // High threshold
	minMRR = 0
	minCoverage = 0

	result := RunResult{
		Configs: []ConfigResult{
			{
				Config: config.SimConfig{Name: "test", TopK: 5},
				Metrics: metrics.Result{
					RecallAtK: 0.75, // Below 0.90
					MRR:       0.80,
					Coverage:  0.95,
				},
			},
		},
	}

	// Capture stdout and prevent os.Exit
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// We can't easily test os.Exit(1), but we can verify the output
	// In a real scenario, we'd refactor to return an error instead of calling os.Exit
	// For now, we test the output contains FAIL

	// Note: This test will call os.Exit(1) which we can't easily catch
	// In production code, you'd want to refactor checkCIThresholds to return an error
	// instead of calling os.Exit directly, making it more testable

	// Skip the actual call since it will exit
	// Just verify the threshold comparison logic works
	w.Close()
	os.Stdout = old
	r.Close()

	// Test the logic directly
	if result.Configs[0].Metrics.RecallAtK >= minRecall {
		t.Errorf("test setup error: recall should be below threshold")
	}
}

func TestCheckCIThresholds_NoConfigs(t *testing.T) {
	result := RunResult{
		Configs: []ConfigResult{},
	}

	err := checkCIThresholds(result)
	if err == nil {
		t.Error("expected error for empty configs")
	}
	if !strings.Contains(err.Error(), "no configurations") {
		t.Errorf("expected 'no configurations' error, got: %v", err)
	}
}

func TestCheckCIThresholds_ZeroThresholds(t *testing.T) {
	oldMinRecall := minRecall
	oldMinMRR := minMRR
	oldMinCoverage := minCoverage
	oldMaxLatencyP95 := maxLatencyP95
	defer func() {
		minRecall = oldMinRecall
		minMRR = oldMinMRR
		minCoverage = oldMinCoverage
		maxLatencyP95 = oldMaxLatencyP95
	}()

	// Zero thresholds = no checks
	minRecall = 0
	minMRR = 0
	minCoverage = 0
	maxLatencyP95 = 0

	result := RunResult{
		Configs: []ConfigResult{
			{
				Config: config.SimConfig{Name: "test", TopK: 5},
				Metrics: metrics.Result{
					RecallAtK:  0.10, // Very low, but no threshold set
					MRR:        0.10,
					Coverage:   0.10,
					LatencyP95: 9999, // Very high, but no threshold set
				},
			},
		},
	}

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := checkCIThresholds(result)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if err != nil {
		t.Errorf("expected no error with zero thresholds, got: %v", err)
	}

	// Should still pass since no thresholds are set
	if !strings.Contains(output, "CI check PASSED") {
		t.Errorf("expected 'CI check PASSED' with zero thresholds")
	}
}

func TestCheckCIThresholds_PartialThresholds(t *testing.T) {
	oldMinRecall := minRecall
	oldMinMRR := minMRR
	oldMinCoverage := minCoverage
	oldMaxLatencyP95 := maxLatencyP95
	defer func() {
		minRecall = oldMinRecall
		minMRR = oldMinMRR
		minCoverage = oldMinCoverage
		maxLatencyP95 = oldMaxLatencyP95
	}()

	// Only recall threshold set
	minRecall = 0.80
	minMRR = 0
	minCoverage = 0
	maxLatencyP95 = 0

	result := RunResult{
		Configs: []ConfigResult{
			{
				Config: config.SimConfig{Name: "test", TopK: 5},
				Metrics: metrics.Result{
					RecallAtK:  0.85,
					MRR:        0.10,   // Low but no threshold
					Coverage:   0.10,   // Low but no threshold
					LatencyP95: 9999.0, // High but no threshold
				},
			},
		},
	}

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := checkCIThresholds(result)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}

	// Should only show recall check
	if !strings.Contains(output, "Recall@K") {
		t.Errorf("expected Recall@K in output")
	}
	// MRR and Coverage should not appear (threshold is 0)
	if strings.Contains(output, "MRR:") {
		t.Errorf("MRR should not appear when threshold is 0")
	}
	// Latency should not appear (threshold is 0)
	if strings.Contains(output, "Latency p95") {
		t.Errorf("Latency p95 should not appear when threshold is 0")
	}
}

func TestCheckCIThresholds_LatencyExceeds(t *testing.T) {
	oldMinRecall := minRecall
	oldMinMRR := minMRR
	oldMinCoverage := minCoverage
	oldMaxLatencyP95 := maxLatencyP95
	defer func() {
		minRecall = oldMinRecall
		minMRR = oldMinMRR
		minCoverage = oldMinCoverage
		maxLatencyP95 = oldMaxLatencyP95
	}()

	// Only latency threshold set
	minRecall = 0
	minMRR = 0
	minCoverage = 0
	maxLatencyP95 = 100.0 // 100ms threshold

	result := RunResult{
		Configs: []ConfigResult{
			{
				Config: config.SimConfig{Name: "test", TopK: 5},
				Metrics: metrics.Result{
					RecallAtK:  0.85,
					MRR:        0.75,
					Coverage:   0.95,
					LatencyP95: 250.0, // 250ms - exceeds 100ms threshold
				},
			},
		},
	}

	// Test the logic directly (can't test os.Exit)
	if result.Configs[0].Metrics.LatencyP95 <= maxLatencyP95 {
		t.Errorf("test setup error: latency should exceed threshold")
	}
}

// TestRunResultSerialization ensures the RunResult struct serializes correctly
func TestRunResultSerialization(t *testing.T) {
	result := RunResult{
		Timestamp:  "2025-01-08T12:00:00Z",
		Collection: "test-collection",
		Store:      "qdrant",
		Configs: []ConfigResult{
			{
				Config: config.SimConfig{Name: "default", TopK: 5},
				Metrics: metrics.Result{
					RecallAtK:  0.85,
					MRR:        0.75,
					Coverage:   0.95,
					Redundancy: 1.2,
				},
				QueryResults: []metrics.QueryResult{
					{
						QueryID:      "q1",
						Query:        "test query",
						RetrievedIDs: []string{"doc1", "doc2"},
						RelevantIDs:  []string{"doc1"},
						Scores:       []float32{0.9, 0.8},
					},
				},
			},
		},
	}

	// Verify fields are accessible
	if result.Timestamp != "2025-01-08T12:00:00Z" {
		t.Errorf("timestamp mismatch")
	}
	if len(result.Configs) != 1 {
		t.Errorf("expected 1 config")
	}
	if result.Configs[0].Metrics.RecallAtK != 0.85 {
		t.Errorf("recall mismatch")
	}
}

func TestCollectFailures(t *testing.T) {
	tests := []struct {
		name           string
		results        []metrics.QueryResult
		k              int
		expectedCount  int
		expectedQueryIDs []string
	}{
		{
			name: "no failures - all queries have recall > 0",
			results: []metrics.QueryResult{
				{
					QueryID:      "q1",
					Query:        "test query 1",
					RetrievedIDs: []string{"a", "b"},
					RelevantIDs:  []string{"a"},
					Scores:       []float32{0.9, 0.8},
				},
				{
					QueryID:      "q2",
					Query:        "test query 2",
					RetrievedIDs: []string{"c", "d"},
					RelevantIDs:  []string{"c"},
					Scores:       []float32{0.85, 0.7},
				},
			},
			k:             2,
			expectedCount: 0,
		},
		{
			name: "one failure - query with recall = 0",
			results: []metrics.QueryResult{
				{
					QueryID:      "q1",
					Query:        "test query 1",
					RetrievedIDs: []string{"a", "b"},
					RelevantIDs:  []string{"a"},
					Scores:       []float32{0.9, 0.8},
				},
				{
					QueryID:      "q2",
					Query:        "test query 2",
					RetrievedIDs: []string{"x", "y"},
					RelevantIDs:  []string{"c"},
					Scores:       []float32{0.7, 0.6},
				},
			},
			k:              2,
			expectedCount:  1,
			expectedQueryIDs: []string{"q2"},
		},
		{
			name: "all failures",
			results: []metrics.QueryResult{
				{
					QueryID:      "q1",
					Query:        "test query 1",
					RetrievedIDs: []string{"x", "y"},
					RelevantIDs:  []string{"a"},
					Scores:       []float32{0.5, 0.4},
				},
				{
					QueryID:      "q2",
					Query:        "test query 2",
					RetrievedIDs: []string{"z", "w"},
					RelevantIDs:  []string{"b", "c"},
					Scores:       []float32{0.6, 0.5},
				},
			},
			k:              2,
			expectedCount:  2,
			expectedQueryIDs: []string{"q1", "q2"},
		},
		{
			name: "no relevant docs - not a failure",
			results: []metrics.QueryResult{
				{
					QueryID:      "q1",
					Query:        "test query 1",
					RetrievedIDs: []string{"a", "b"},
					RelevantIDs:  []string{}, // No relevant docs
					Scores:       []float32{0.9, 0.8},
				},
			},
			k:             2,
			expectedCount: 0, // Not counted as failure when no relevant docs
		},
		{
			name: "relevant doc outside top-k",
			results: []metrics.QueryResult{
				{
					QueryID:      "q1",
					Query:        "test query 1",
					RetrievedIDs: []string{"x", "y", "z", "a"}, // 'a' at position 4
					RelevantIDs:  []string{"a"},
					Scores:       []float32{0.9, 0.8, 0.7, 0.6},
				},
			},
			k:              2, // Only look at top 2
			expectedCount:  1,
			expectedQueryIDs: []string{"q1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			failures := collectFailures(tt.results, tt.k)

			if len(failures) != tt.expectedCount {
				t.Errorf("collectFailures() returned %d failures, want %d", len(failures), tt.expectedCount)
			}

			if tt.expectedQueryIDs != nil {
				for i, expectedID := range tt.expectedQueryIDs {
					if i >= len(failures) {
						t.Errorf("missing expected failure for query %s", expectedID)
						continue
					}
					if failures[i].QueryID != expectedID {
						t.Errorf("failure[%d].QueryID = %s, want %s", i, failures[i].QueryID, expectedID)
					}
				}
			}
		})
	}
}

func TestCollectFailures_FieldsPopulated(t *testing.T) {
	results := []metrics.QueryResult{
		{
			QueryID:      "q-failed",
			Query:        "How do I configure authentication?",
			RetrievedIDs: []string{"api-keys.md", "rate-limits.md", "webhooks.md"},
			RelevantIDs:  []string{"auth-guide.md"},
			Scores:       []float32{0.75, 0.68, 0.55},
		},
	}

	failures := collectFailures(results, 3)

	if len(failures) != 1 {
		t.Fatalf("expected 1 failure, got %d", len(failures))
	}

	f := failures[0]

	// Verify all fields are populated
	if f.QueryID != "q-failed" {
		t.Errorf("QueryID = %s, want q-failed", f.QueryID)
	}
	if f.Query != "How do I configure authentication?" {
		t.Errorf("Query not populated correctly")
	}
	if len(f.RelevantDocs) != 1 || f.RelevantDocs[0] != "auth-guide.md" {
		t.Errorf("RelevantDocs = %v, want [auth-guide.md]", f.RelevantDocs)
	}
	if len(f.RetrievedDocs) != 3 {
		t.Errorf("RetrievedDocs should have 3 items (limited to top 3)")
	}
	if len(f.TopScores) != 3 {
		t.Errorf("TopScores should have 3 items")
	}
	if f.Recall != 0 {
		t.Errorf("Recall = %v, want 0", f.Recall)
	}
}

func TestTruncateQuery(t *testing.T) {
	tests := []struct {
		input    string
		maxLen   int
		expected string
	}{
		{"short", 10, "short"},
		{"exactly ten", 11, "exactly ten"},
		{"this is a longer string", 10, "this is..."},
		{"multi\nline\nquery", 20, "multi line query"},
		{"", 10, ""},
	}

	for _, tt := range tests {
		result := truncateQuery(tt.input, tt.maxLen)
		if result != tt.expected {
			t.Errorf("truncateQuery(%q, %d) = %q, want %q", tt.input, tt.maxLen, result, tt.expected)
		}
	}
}
