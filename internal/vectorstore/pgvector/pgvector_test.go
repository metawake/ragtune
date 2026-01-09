package pgvector

import (
	"testing"
)

func TestTableName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"demo", "ragtune_demo"},
		{"my-collection", "ragtune_my_collection"},
		{"test_123", "ragtune_test_123"},
		{"My Collection!", "ragtune_My_Collection_"},
		{"prod", "ragtune_prod"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := tableName(tt.input)
			if got != tt.expected {
				t.Errorf("tableName(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}
