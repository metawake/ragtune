package metrics

import (
	"math"
	"testing"
)

func TestRecallAtK(t *testing.T) {
	tests := []struct {
		name      string
		retrieved []string
		relevant  []string
		k         int
		expected  float64
	}{
		{
			name:      "perfect recall",
			retrieved: []string{"a", "b", "c"},
			relevant:  []string{"a", "b"},
			k:         3,
			expected:  1.0,
		},
		{
			name:      "partial recall",
			retrieved: []string{"a", "x", "y"},
			relevant:  []string{"a", "b"},
			k:         3,
			expected:  0.5,
		},
		{
			name:      "no recall",
			retrieved: []string{"x", "y", "z"},
			relevant:  []string{"a", "b"},
			k:         3,
			expected:  0.0,
		},
		{
			name:      "k limits results",
			retrieved: []string{"x", "y", "a", "b"},
			relevant:  []string{"a", "b"},
			k:         2,
			expected:  0.0, // a and b are outside top-2
		},
		{
			name:      "no relevant docs",
			retrieved: []string{"a", "b"},
			relevant:  []string{},
			k:         2,
			expected:  1.0,
		},
		{
			name:      "duplicates in retrieved",
			retrieved: []string{"a", "a", "a"},
			relevant:  []string{"a"},
			k:         3,
			expected:  1.0, // Should be 1.0, not 3.0
		},
		{
			name:      "duplicates with partial match",
			retrieved: []string{"a", "a", "b", "b"},
			relevant:  []string{"a", "c"},
			k:         4,
			expected:  0.5, // Only "a" found, not "c"
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RecallAtK(tt.retrieved, tt.relevant, tt.k)
			if math.Abs(got-tt.expected) > 0.001 {
				t.Errorf("RecallAtK() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestReciprocalRank(t *testing.T) {
	tests := []struct {
		name      string
		retrieved []string
		relevant  []string
		expected  float64
	}{
		{
			name:      "first position",
			retrieved: []string{"a", "b", "c"},
			relevant:  []string{"a"},
			expected:  1.0,
		},
		{
			name:      "second position",
			retrieved: []string{"x", "a", "c"},
			relevant:  []string{"a"},
			expected:  0.5,
		},
		{
			name:      "third position",
			retrieved: []string{"x", "y", "a"},
			relevant:  []string{"a"},
			expected:  1.0 / 3.0,
		},
		{
			name:      "not found",
			retrieved: []string{"x", "y", "z"},
			relevant:  []string{"a"},
			expected:  0.0,
		},
		{
			name:      "multiple relevant - first wins",
			retrieved: []string{"x", "a", "b"},
			relevant:  []string{"a", "b"},
			expected:  0.5, // "a" at position 2
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReciprocalRank(tt.retrieved, tt.relevant)
			if math.Abs(got-tt.expected) > 0.001 {
				t.Errorf("ReciprocalRank() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCoverage(t *testing.T) {
	tests := []struct {
		name      string
		retrieved map[string]int
		relevant  map[string]struct{}
		expected  float64
	}{
		{
			name:      "full coverage",
			retrieved: map[string]int{"a": 1, "b": 2},
			relevant:  map[string]struct{}{"a": {}, "b": {}},
			expected:  1.0,
		},
		{
			name:      "partial coverage",
			retrieved: map[string]int{"a": 1, "c": 1},
			relevant:  map[string]struct{}{"a": {}, "b": {}},
			expected:  0.5,
		},
		{
			name:      "no coverage",
			retrieved: map[string]int{"x": 1, "y": 1},
			relevant:  map[string]struct{}{"a": {}, "b": {}},
			expected:  0.0,
		},
		{
			name:      "no relevant docs",
			retrieved: map[string]int{"a": 1},
			relevant:  map[string]struct{}{},
			expected:  1.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Coverage(tt.retrieved, tt.relevant)
			if math.Abs(got-tt.expected) > 0.001 {
				t.Errorf("Coverage() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestRedundancy(t *testing.T) {
	tests := []struct {
		name      string
		retrieved map[string]int
		expected  float64
	}{
		{
			name:      "no redundancy",
			retrieved: map[string]int{"a": 1, "b": 1, "c": 1},
			expected:  1.0,
		},
		{
			name:      "some redundancy",
			retrieved: map[string]int{"a": 3, "b": 1},
			expected:  2.0, // (3+1)/2
		},
		{
			name:      "high redundancy",
			retrieved: map[string]int{"a": 10},
			expected:  10.0,
		},
		{
			name:      "empty",
			retrieved: map[string]int{},
			expected:  0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Redundancy(tt.retrieved)
			if math.Abs(got-tt.expected) > 0.001 {
				t.Errorf("Redundancy() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCompute(t *testing.T) {
	results := []QueryResult{
		{
			QueryID:      "q1",
			RetrievedIDs: []string{"a", "b", "c"},
			RelevantIDs:  []string{"a", "b"},
		},
		{
			QueryID:      "q2",
			RetrievedIDs: []string{"a", "d", "e"},
			RelevantIDs:  []string{"a", "c"},
		},
	}

	m := Compute(results, 3)

	// Recall: q1=1.0, q2=0.5 -> avg 0.75
	if math.Abs(m.RecallAtK-0.75) > 0.001 {
		t.Errorf("RecallAtK = %v, want 0.75", m.RecallAtK)
	}

	// MRR: q1=1.0 (a at pos 1), q2=1.0 (a at pos 1) -> avg 1.0
	if math.Abs(m.MRR-1.0) > 0.001 {
		t.Errorf("MRR = %v, want 1.0", m.MRR)
	}

	// Coverage: relevant={a,b,c}, retrieved={a,b,c,d,e} -> all relevant found = 1.0
	if math.Abs(m.Coverage-1.0) > 0.001 {
		t.Errorf("Coverage = %v, want 1.0", m.Coverage)
	}
}

func TestComputeLatencyStats(t *testing.T) {
	tests := []struct {
		name       string
		latencies  []float64
		wantP50    float64
		wantP95    float64
		wantP99    float64
		wantAvg    float64
	}{
		{
			name:       "empty results",
			latencies:  []float64{},
			wantP50:    0,
			wantP95:    0,
			wantP99:    0,
			wantAvg:    0,
		},
		{
			name:       "single value",
			latencies:  []float64{100},
			wantP50:    100,
			wantP95:    100,
			wantP99:    100,
			wantAvg:    100,
		},
		{
			name:       "two values",
			latencies:  []float64{100, 200},
			wantP50:    150, // interpolated
			wantP95:    195, // near 200
			wantP99:    199, // very near 200
			wantAvg:    150,
		},
		{
			name:       "ten values",
			latencies:  []float64{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
			wantP50:    55,  // between 50 and 60
			wantP95:    96.5, // near 100
			wantP99:    99.1, // very near 100
			wantAvg:    55,
		},
		{
			name:       "with zeros (skipped)",
			latencies:  []float64{0, 100, 0, 200, 0},
			wantP50:    150,
			wantP95:    195,
			wantP99:    199,
			wantAvg:    150,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Build query results from latencies
			var results []QueryResult
			for i, lat := range tt.latencies {
				results = append(results, QueryResult{
					QueryID:   "q" + string(rune('0'+i)),
					LatencyMs: lat,
				})
			}

			p50, p95, p99, avg := ComputeLatencyStats(results)

			if math.Abs(p50-tt.wantP50) > 1.0 {
				t.Errorf("p50 = %v, want %v", p50, tt.wantP50)
			}
			if math.Abs(p95-tt.wantP95) > 1.0 {
				t.Errorf("p95 = %v, want %v", p95, tt.wantP95)
			}
			if math.Abs(p99-tt.wantP99) > 1.0 {
				t.Errorf("p99 = %v, want %v", p99, tt.wantP99)
			}
			if math.Abs(avg-tt.wantAvg) > 1.0 {
				t.Errorf("avg = %v, want %v", avg, tt.wantAvg)
			}
		})
	}
}

func TestComputeWithLatency(t *testing.T) {
	results := []QueryResult{
		{
			QueryID:      "q1",
			RetrievedIDs: []string{"a"},
			RelevantIDs:  []string{"a"},
			LatencyMs:    50,
		},
		{
			QueryID:      "q2",
			RetrievedIDs: []string{"b"},
			RelevantIDs:  []string{"b"},
			LatencyMs:    100,
		},
		{
			QueryID:      "q3",
			RetrievedIDs: []string{"c"},
			RelevantIDs:  []string{"c"},
			LatencyMs:    150,
		},
	}

	m := Compute(results, 1)

	// Check latency stats are populated
	if m.LatencyAvg == 0 {
		t.Error("LatencyAvg should not be 0")
	}
	if math.Abs(m.LatencyAvg-100) > 1.0 {
		t.Errorf("LatencyAvg = %v, want ~100", m.LatencyAvg)
	}
	if m.LatencyP50 == 0 {
		t.Error("LatencyP50 should not be 0")
	}
	if m.LatencyP95 == 0 {
		t.Error("LatencyP95 should not be 0")
	}
	if m.LatencyP99 == 0 {
		t.Error("LatencyP99 should not be 0")
	}
}
func TestNDCGAtK(t *testing.T) {
	tests := []struct {
		name      string
		retrieved []string
		relevant  []string
		k         int
		expected  float64
	}{
		{
			name:      "perfect ranking - all relevant first",
			retrieved: []string{"a", "b", "x", "y"},
			relevant:  []string{"a", "b"},
			k:         4,
			expected:  1.0, // Ideal ranking
		},
		{
			name:      "no relevant docs",
			retrieved: []string{"a", "b"},
			relevant:  []string{},
			k:         2,
			expected:  1.0, // Perfect NDCG when no relevant docs
		},
		{
			name:      "no matches",
			retrieved: []string{"x", "y", "z"},
			relevant:  []string{"a", "b"},
			k:         3,
			expected:  0.0, // No relevant docs retrieved
		},
		{
			name:      "relevant at position 2",
			retrieved: []string{"x", "a", "y"},
			relevant:  []string{"a"},
			k:         3,
			// DCG = 1/log2(3) = 0.631
			// IDCG = 1/log2(2) = 1.0
			// NDCG = 0.631
			expected:  0.631,
		},
		{
			name:      "k limits results",
			retrieved: []string{"x", "y", "a", "b"},
			relevant:  []string{"a", "b"},
			k:         2,
			expected:  0.0, // a and b outside top-2
		},
		{
			name:      "suboptimal ranking",
			retrieved: []string{"x", "a", "b", "y"},
			relevant:  []string{"a", "b"},
			k:         4,
			// DCG = 1/log2(3) + 1/log2(4) = 0.631 + 0.5 = 1.131
			// IDCG = 1/log2(2) + 1/log2(3) = 1.0 + 0.631 = 1.631
			// NDCG = 1.131/1.631 = 0.693
			expected:  0.693,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NDCGAtK(tt.retrieved, tt.relevant, tt.k)
			if math.Abs(got-tt.expected) > 0.01 {
				t.Errorf("NDCGAtK() = %v, want %v", got, tt.expected)
			}
		})
	}
}
