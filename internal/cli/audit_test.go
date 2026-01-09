package cli

import (
	"testing"

	"github.com/metawake/ragtune/internal/metrics"
)

func TestBuildCheck_Pass(t *testing.T) {
	// 0.95 is well above 0.85 (more than 10% above)
	check := buildCheck("Recall@K", 0.95, 0.85, false, "issue", "recommendation")

	if !check.passed {
		t.Error("expected check to pass when value > threshold")
	}
	if check.warning {
		t.Error("expected no warning when well above threshold")
	}
}

func TestBuildCheck_Fail(t *testing.T) {
	check := buildCheck("Recall@K", 0.70, 0.85, false, "issue text", "recommendation text")

	if check.passed {
		t.Error("expected check to fail when value < threshold")
	}
	// Recommendation should include issue when failing
	if check.recommendation != "issue text. recommendation text" {
		t.Errorf("expected combined recommendation, got: %s", check.recommendation)
	}
}

func TestBuildCheck_Warning(t *testing.T) {
	// Value is above threshold but within 10% - should warn
	check := buildCheck("Recall@K", 0.86, 0.85, false, "issue", "recommendation")

	if !check.passed {
		t.Error("expected check to pass")
	}
	if !check.warning {
		t.Error("expected warning when value is within 10% of threshold")
	}
}

func TestBuildCheck_NoWarningWellAbove(t *testing.T) {
	// Value is well above threshold (>10%) - no warning
	check := buildCheck("Recall@K", 0.95, 0.85, false, "issue", "recommendation")

	if !check.passed {
		t.Error("expected check to pass")
	}
	if check.warning {
		t.Error("expected no warning when well above threshold")
	}
}

func TestTruncateStr(t *testing.T) {
	tests := []struct {
		input    string
		maxLen   int
		expected string
	}{
		{"short", 10, "short"},
		{"exactly10!", 10, "exactly10!"},
		{"this is a longer string", 10, "this is..."},
		{"", 10, ""},
	}

	for _, tt := range tests {
		got := truncateStr(tt.input, tt.maxLen)
		if got != tt.expected {
			t.Errorf("truncateStr(%q, %d) = %q, want %q", tt.input, tt.maxLen, got, tt.expected)
		}
	}
}

func TestAuditCheck_AllScenarios(t *testing.T) {
	tests := []struct {
		name       string
		value      float64
		threshold  float64
		wantPass   bool
		wantWarn   bool
	}{
		{"well above", 0.95, 0.85, true, false},
		{"exactly at", 0.85, 0.85, true, true},  // At threshold = warning
		{"just above", 0.86, 0.85, true, true},  // Within 10%
		{"10% above", 0.935, 0.85, true, false}, // Exactly 10% above = no warning
		{"just below", 0.84, 0.85, false, false},
		{"well below", 0.50, 0.85, false, false},
		{"zero", 0.0, 0.85, false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			check := buildCheck("Test", tt.value, tt.threshold, false, "issue", "rec")
			if check.passed != tt.wantPass {
				t.Errorf("passed = %v, want %v", check.passed, tt.wantPass)
			}
			if check.warning != tt.wantWarn {
				t.Errorf("warning = %v, want %v", check.warning, tt.wantWarn)
			}
		})
	}
}

func TestAuditCheck_LatencyScenarios(t *testing.T) {
	// For latency, lower is better (value should be <= threshold)
	tests := []struct {
		name       string
		value      float64
		threshold  float64
		wantPass   bool
		wantWarn   bool
	}{
		{"well below", 50.0, 500.0, true, false},
		{"exactly at", 500.0, 500.0, true, true},    // At threshold = warning
		{"just below", 400.0, 500.0, true, false},   // Below 90% of threshold = no warning
		{"within 10%", 460.0, 500.0, true, true},    // Within 10% of threshold = warning
		{"just above", 501.0, 500.0, false, false},  // Fails
		{"well above", 1000.0, 500.0, false, false}, // Fails
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			check := buildCheck("Latency p95", tt.value, tt.threshold, true, "issue", "rec")
			if check.passed != tt.wantPass {
				t.Errorf("passed = %v, want %v (value=%.1f, threshold=%.1f)", check.passed, tt.wantPass, tt.value, tt.threshold)
			}
			if check.warning != tt.wantWarn {
				t.Errorf("warning = %v, want %v (value=%.1f, threshold=%.1f)", check.warning, tt.wantWarn, tt.value, tt.threshold)
			}
		})
	}
}

func TestAuditMetricsIntegration(t *testing.T) {
	// Test that audit correctly interprets metrics.Result
	result := metrics.Result{
		RecallAtK:  0.88,
		MRR:        0.72,
		Coverage:   0.95,
		Redundancy: 1.5,
		LatencyP95: 100.0,
	}

	// Build checks with default thresholds
	checks := []auditCheck{
		buildCheck("Recall@K", result.RecallAtK, 0.85, false, "", ""),
		buildCheck("MRR", result.MRR, 0.70, false, "", ""),
		buildCheck("Coverage", result.Coverage, 0.90, false, "", ""),
		buildCheck("Latency p95", result.LatencyP95, 500.0, true, "", ""),
	}

	// All should pass with these values
	for _, c := range checks {
		if !c.passed {
			t.Errorf("%s should pass: value=%.3f, threshold=%.3f", c.name, c.value, c.threshold)
		}
	}
}

func TestAuditMetricsFailure(t *testing.T) {
	// Test with failing metrics
	result := metrics.Result{
		RecallAtK:  0.70,   // Below 0.85
		MRR:        0.50,   // Below 0.70
		Coverage:   0.80,   // Below 0.90
		Redundancy: 2.0,
		LatencyP95: 1000.0, // Above 500ms
	}

	checks := []auditCheck{
		buildCheck("Recall@K", result.RecallAtK, 0.85, false, "", ""),
		buildCheck("MRR", result.MRR, 0.70, false, "", ""),
		buildCheck("Coverage", result.Coverage, 0.90, false, "", ""),
		buildCheck("Latency p95", result.LatencyP95, 500.0, true, "", ""),
	}

	// All should fail
	for _, c := range checks {
		if c.passed {
			t.Errorf("%s should fail: value=%.3f, threshold=%.3f", c.name, c.value, c.threshold)
		}
	}
}

func TestAuditQueryCountGuidance(t *testing.T) {
	// Test the logic for query count guidance
	tests := []struct {
		count    int
		category string
	}{
		{10, "very low"},
		{19, "very low"},
		{20, "low"},
		{49, "low"},
		{50, "okay"},
		{99, "okay"},
		{100, "good"},
		{500, "good"},
	}

	for _, tt := range tests {
		var guidance string
		if tt.count < 20 {
			guidance = "very low"
		} else if tt.count < 50 {
			guidance = "low"
		} else if tt.count < 100 {
			guidance = "okay"
		} else {
			guidance = "good"
		}

		if guidance != tt.category {
			t.Errorf("count=%d: got %s, want %s", tt.count, guidance, tt.category)
		}
	}
}
