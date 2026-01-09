package embedder

import (
	"context"
	"os"
	"testing"
)

func TestNewCohereEmbedder(t *testing.T) {
	// Save and clear env var for testing
	original := os.Getenv("COHERE_API_KEY")
	defer os.Setenv("COHERE_API_KEY", original)
	os.Setenv("COHERE_API_KEY", "test-key")

	t.Run("default configuration", func(t *testing.T) {
		e := NewCohereEmbedder()

		if e.model != "embed-english-v3.0" {
			t.Errorf("model = %v, want embed-english-v3.0", e.model)
		}
		if e.dim != 1024 {
			t.Errorf("dim = %v, want 1024", e.dim)
		}
		if e.inputType != "search_document" {
			t.Errorf("inputType = %v, want search_document", e.inputType)
		}
	})

	t.Run("with custom model", func(t *testing.T) {
		e := NewCohereEmbedder(WithCohereModel("embed-multilingual-v3.0"))

		if e.model != "embed-multilingual-v3.0" {
			t.Errorf("model = %v, want embed-multilingual-v3.0", e.model)
		}
		if e.dim != 1024 {
			t.Errorf("dim = %v, want 1024", e.dim)
		}
	})

	t.Run("light model has smaller dimension", func(t *testing.T) {
		e := NewCohereEmbedder(WithCohereModel("embed-english-light-v3.0"))

		if e.dim != 384 {
			t.Errorf("dim = %v, want 384 for light model", e.dim)
		}
	})

	t.Run("with custom input type", func(t *testing.T) {
		e := NewCohereEmbedder(WithCohereInputType("search_query"))

		if e.inputType != "search_query" {
			t.Errorf("inputType = %v, want search_query", e.inputType)
		}
	})
}

func TestCohereEmbedder_Dim(t *testing.T) {
	tests := []struct {
		model    string
		expected int
	}{
		{"embed-english-v3.0", 1024},
		{"embed-multilingual-v3.0", 1024},
		{"embed-english-light-v3.0", 384},
		{"embed-multilingual-light-v3.0", 384},
		{"unknown-model", 1024}, // defaults to 1024
	}

	for _, tt := range tests {
		t.Run(tt.model, func(t *testing.T) {
			e := NewCohereEmbedder(WithCohereModel(tt.model))
			if e.Dim() != tt.expected {
				t.Errorf("Dim() = %v, want %v", e.Dim(), tt.expected)
			}
		})
	}
}

func TestCohereEmbedder_MissingAPIKey(t *testing.T) {
	// Clear env var
	original := os.Getenv("COHERE_API_KEY")
	defer os.Setenv("COHERE_API_KEY", original)
	os.Unsetenv("COHERE_API_KEY")

	e := NewCohereEmbedder()
	_, err := e.Embed(context.Background(), "test text")

	if err == nil {
		t.Error("expected error for missing API key")
	}
	if err.Error() != "COHERE_API_KEY environment variable not set" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCohereEmbedder_EmbedBatchMissingAPIKey(t *testing.T) {
	// Clear env var
	original := os.Getenv("COHERE_API_KEY")
	defer os.Setenv("COHERE_API_KEY", original)
	os.Unsetenv("COHERE_API_KEY")

	e := NewCohereEmbedder()
	_, err := e.EmbedBatch(context.Background(), []string{"text1", "text2"})

	if err == nil {
		t.Error("expected error for missing API key")
	}
}



