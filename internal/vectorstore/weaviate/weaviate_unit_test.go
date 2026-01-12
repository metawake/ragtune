package weaviate

import "testing"

func TestClassName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple", "demo", "Ragtune_Demo"},
		{"with hyphen", "my-collection", "Ragtune_My_collection"},
		{"with underscore", "test_123", "Ragtune_Test_123"},
		{"with dots", "my.collection.name", "Ragtune_My_collection_name"},
		{"with spaces", "my collection", "Ragtune_My_collection"},
		{"all special", "---", "Ragtune____"},
		{"empty", "", "Ragtune"},
		{"single char", "a", "Ragtune_A"},
		{"uppercase", "DEMO", "Ragtune_DEMO"},
		{"mixed case", "MyCollection", "Ragtune_MyCollection"},
		{"numbers only", "123", "Ragtune_123"},
		{"starts with number", "1test", "Ragtune_1test"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := className(tt.input)
			if got != tt.expected {
				t.Errorf("className(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestVectorToJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    []float32
		expected string
	}{
		{"simple", []float32{0.1, 0.2, 0.3}, "[0.100000,0.200000,0.300000]"},
		{"empty", []float32{}, "[]"},
		{"single", []float32{1.0}, "[1.000000]"},
		{"negative", []float32{-0.5, 0.5}, "[-0.500000,0.500000]"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := vectorToJSON(tt.input)
			if got != tt.expected {
				t.Errorf("vectorToJSON(%v) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}
