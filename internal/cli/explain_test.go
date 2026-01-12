package cli

import (
	"math"
	"slices"
	"testing"
)

func TestComputeDiagnostics_BasicStats(t *testing.T) {
	scores := []float32{0.9, 0.8, 0.7, 0.6, 0.5}

	diag := computeDiagnostics(scores)

	// Basic stats
	if diag.maxScore != 0.9 {
		t.Errorf("maxScore = %v, want 0.9", diag.maxScore)
	}
	if diag.minScore != 0.5 {
		t.Errorf("minScore = %v, want 0.5", diag.minScore)
	}
	if math.Abs(float64(diag.spread)-0.4) > 0.001 {
		t.Errorf("spread = %v, want 0.4", diag.spread)
	}
	if math.Abs(float64(diag.meanScore)-0.7) > 0.001 {
		t.Errorf("meanScore = %v, want 0.7", diag.meanScore)
	}
}

func TestComputeDiagnostics_StdDev(t *testing.T) {
	// Scores with known variance
	scores := []float32{0.8, 0.8, 0.8, 0.8} // All same = 0 std dev

	diag := computeDiagnostics(scores)

	if diag.stdDev != 0 {
		t.Errorf("stdDev = %v, want 0 for uniform scores", diag.stdDev)
	}
}

func TestComputeDiagnostics_Quartiles(t *testing.T) {
	// Simple ascending scores for easy quartile verification
	scores := []float32{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9}

	diag := computeDiagnostics(scores)

	// Median should be around 0.5
	if math.Abs(float64(diag.median)-0.5) > 0.1 {
		t.Errorf("median = %v, expected ~0.5", diag.median)
	}
	// Q1 should be around 0.25
	if math.Abs(float64(diag.q1)-0.3) > 0.1 {
		t.Errorf("q1 = %v, expected ~0.3", diag.q1)
	}
	// Q3 should be around 0.75
	if math.Abs(float64(diag.q3)-0.7) > 0.1 {
		t.Errorf("q3 = %v, expected ~0.7", diag.q3)
	}
}

func TestComputeDiagnostics_TopGap(t *testing.T) {
	tests := []struct {
		name        string
		scores      []float32
		expectedGap float32
	}{
		{
			name:        "clear top result",
			scores:      []float32{0.95, 0.75, 0.70},
			expectedGap: 0.20,
		},
		{
			name:        "no gap",
			scores:      []float32{0.80, 0.80, 0.75},
			expectedGap: 0.0,
		},
		{
			name:        "single result",
			scores:      []float32{0.90},
			expectedGap: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diag := computeDiagnostics(tt.scores)
			if math.Abs(float64(diag.topGap)-float64(tt.expectedGap)) > 0.001 {
				t.Errorf("topGap = %v, want %v", diag.topGap, tt.expectedGap)
			}
		})
	}
}

func TestComputeDiagnostics_ScoreShape(t *testing.T) {
	tests := []struct {
		name          string
		scores        []float32
		expectedShape string
	}{
		{
			name:          "tight distribution",
			scores:        []float32{0.80, 0.81, 0.79, 0.80, 0.80},
			expectedShape: "tight",
		},
		{
			name:          "spread distribution",
			scores:        []float32{0.95, 0.70, 0.50, 0.30, 0.10},
			expectedShape: "spread",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diag := computeDiagnostics(tt.scores)
			if diag.scoreShape != tt.expectedShape {
				t.Errorf("scoreShape = %s, want %s (stdDev=%v, spread=%v)",
					diag.scoreShape, tt.expectedShape, diag.stdDev, diag.spread)
			}
		})
	}
}

func TestComputeDiagnostics_Warnings(t *testing.T) {
	tests := []struct {
		name            string
		scores          []float32
		expectedWarning string
		shouldContain   bool
	}{
		{
			name:            "low top score warning",
			scores:          []float32{0.40, 0.35, 0.30},
			expectedWarning: "Low top score",
			shouldContain:   true,
		},
		{
			name:            "high spread warning",
			scores:          []float32{0.90, 0.50, 0.10},
			expectedWarning: "High score spread",
			shouldContain:   true,
		},
		{
			name:            "low spread warning",
			scores:          []float32{0.80, 0.79, 0.78},
			expectedWarning: "Very low spread",
			shouldContain:   true,
		},
		{
			name:            "large top gap warning",
			scores:          []float32{0.95, 0.70, 0.65},
			expectedWarning: "Large gap",
			shouldContain:   true,
		},
		{
			name:            "no warnings for good scores",
			scores:          []float32{0.90, 0.82, 0.75, 0.68},
			expectedWarning: "",
			shouldContain:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diag := computeDiagnostics(tt.scores)

			if tt.shouldContain {
				found := false
				for _, w := range diag.warnings {
					if contains(w, tt.expectedWarning) {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("expected warning containing %q, got %v", tt.expectedWarning, diag.warnings)
				}
			} else {
				if len(diag.warnings) > 0 {
					t.Errorf("expected no warnings, got %v", diag.warnings)
				}
			}
		})
	}
}

func TestComputeDiagnostics_Insights(t *testing.T) {
	tests := []struct {
		name            string
		scores          []float32
		expectedInsight string
		shouldContain   bool
	}{
		{
			name:            "strong top match insight",
			scores:          []float32{0.92, 0.78, 0.70},
			expectedInsight: "Strong top match",
			shouldContain:   true,
		},
		{
			name:            "good separation insight",
			scores:          []float32{0.85, 0.78, 0.72, 0.68},
			expectedInsight: "Good score separation",
			shouldContain:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diag := computeDiagnostics(tt.scores)

			if tt.shouldContain {
				found := false
				for _, insight := range diag.insights {
					if contains(insight, tt.expectedInsight) {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("expected insight containing %q, got %v", tt.expectedInsight, diag.insights)
				}
			}
		})
	}
}

func TestComputeDiagnostics_Empty(t *testing.T) {
	diag := computeDiagnostics([]float32{})

	if diag.maxScore != 0 {
		t.Errorf("maxScore for empty should be 0")
	}
	if diag.minScore != 0 {
		t.Errorf("minScore for empty should be 0")
	}
	if len(diag.warnings) != 0 {
		t.Errorf("no warnings expected for empty scores")
	}
}

func TestSlicesSort(t *testing.T) {
	// Verify slices.Sort works as expected (sanity check for stdlib)
	tests := []struct {
		input    []float32
		expected []float32
	}{
		{
			input:    []float32{0.5, 0.2, 0.8, 0.1},
			expected: []float32{0.1, 0.2, 0.5, 0.8},
		},
		{
			input:    []float32{0.5},
			expected: []float32{0.5},
		},
		{
			input:    []float32{},
			expected: []float32{},
		},
	}

	for _, tt := range tests {
		// Copy input to avoid modifying test case
		input := make([]float32, len(tt.input))
		copy(input, tt.input)

		slices.Sort(input)

		if len(input) != len(tt.expected) {
			t.Errorf("length mismatch after sort")
			continue
		}
		for i := range input {
			if input[i] != tt.expected[i] {
				t.Errorf("slices.Sort: position %d = %v, want %v", i, input[i], tt.expected[i])
			}
		}
	}
}

func TestPercentileFloat32(t *testing.T) {
	sorted := []float32{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0}

	tests := []struct {
		percentile int
		expected   float32
		tolerance  float32
	}{
		{50, 0.55, 0.05}, // Median
		{25, 0.325, 0.05}, // Q1
		{75, 0.775, 0.05}, // Q3
		{0, 0.1, 0.01},    // Min
		{100, 1.0, 0.01},  // Max
	}

	for _, tt := range tests {
		result := percentileFloat32(sorted, tt.percentile)
		if math.Abs(float64(result)-float64(tt.expected)) > float64(tt.tolerance) {
			t.Errorf("percentileFloat32(sorted, %d) = %v, want ~%v", tt.percentile, result, tt.expected)
		}
	}
}

func TestClassifyScoreShape(t *testing.T) {
	tests := []struct {
		name     string
		diag     diagnostics
		expected string
	}{
		{
			name:     "tight - very low std dev",
			diag:     diagnostics{stdDev: 0.01, spread: 0.1, q1: 0.79, q3: 0.81},
			expected: "tight",
		},
		{
			name:     "spread - high spread",
			diag:     diagnostics{stdDev: 0.15, spread: 0.5, q1: 0.4, q3: 0.8},
			expected: "spread",
		},
		{
			name:     "normal - typical distribution",
			diag:     diagnostics{stdDev: 0.08, spread: 0.25, q1: 0.65, q3: 0.85},
			expected: "normal",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := classifyScoreShape(tt.diag)
			if result != tt.expected {
				t.Errorf("classifyScoreShape() = %s, want %s", result, tt.expected)
			}
		})
	}
}

// Helper function
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		(len(s) > 0 && len(substr) > 0 && findSubstring(s, substr)))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
