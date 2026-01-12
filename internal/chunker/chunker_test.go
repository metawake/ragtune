package chunker

import (
	"strings"
	"testing"
)

func TestChunker_BasicChunking(t *testing.T) {
	c, err := New(100, 20)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	text := "This is a test document. It has multiple sentences. We want to see how it gets chunked into smaller pieces for embedding."
	chunks := c.Chunk(text, "test.md")

	if len(chunks) == 0 {
		t.Fatal("expected at least one chunk")
	}

	// Verify all chunks have required fields
	for i, chunk := range chunks {
		if chunk.ID == "" {
			t.Errorf("chunk %d: missing ID", i)
		}
		if chunk.Text == "" {
			t.Errorf("chunk %d: missing Text", i)
		}
		if chunk.Source != "test.md" {
			t.Errorf("chunk %d: expected source 'test.md', got %q", i, chunk.Source)
		}
		if chunk.Index != i {
			t.Errorf("chunk %d: expected index %d, got %d", i, i, chunk.Index)
		}
	}
}

func TestChunker_EmptyText(t *testing.T) {
	c := MustNew(100, 20)

	chunks := c.Chunk("", "empty.md")
	if len(chunks) != 0 {
		t.Errorf("expected 0 chunks for empty text, got %d", len(chunks))
	}

	chunks = c.Chunk("   \n\t  ", "whitespace.md")
	if len(chunks) != 0 {
		t.Errorf("expected 0 chunks for whitespace-only text, got %d", len(chunks))
	}
}

func TestChunker_SmallText(t *testing.T) {
	c := MustNew(1000, 100)

	text := "Short text."
	chunks := c.Chunk(text, "short.md")

	if len(chunks) != 1 {
		t.Fatalf("expected 1 chunk for short text, got %d", len(chunks))
	}

	if chunks[0].Text != text {
		t.Errorf("expected chunk text %q, got %q", text, chunks[0].Text)
	}
}

func TestChunker_Overlap(t *testing.T) {
	c := MustNew(50, 10)

	// Create text with distinct words to verify overlap behavior
	text := "The quick brown fox jumps over the lazy dog. A wonderful serenity has taken possession of my entire soul."
	chunks := c.Chunk(text, "overlap.md")

	if len(chunks) < 2 {
		t.Fatalf("expected multiple chunks, got %d", len(chunks))
	}

	// Verify chunks are not identical (forward progress is made)
	for i := 0; i < len(chunks)-1; i++ {
		if chunks[i].Text == chunks[i+1].Text {
			t.Errorf("chunks %d and %d are identical: %q", i, i+1, chunks[i].Text)
		}
	}

	// Verify each chunk has content
	for i, chunk := range chunks {
		if len(chunk.Text) == 0 {
			t.Errorf("chunk %d is empty", i)
		}
	}
}

func TestChunker_DeterministicIDs(t *testing.T) {
	c := MustNew(100, 20)

	text := "Deterministic ID test content."
	chunks1 := c.Chunk(text, "test.md")
	chunks2 := c.Chunk(text, "test.md")

	if len(chunks1) != len(chunks2) {
		t.Fatalf("expected same number of chunks, got %d vs %d", len(chunks1), len(chunks2))
	}

	for i := range chunks1 {
		if chunks1[i].ID != chunks2[i].ID {
			t.Errorf("chunk %d: IDs not deterministic: %q vs %q", i, chunks1[i].ID, chunks2[i].ID)
		}
	}
}

func TestChunker_DifferentSourcesDifferentIDs(t *testing.T) {
	c := MustNew(100, 20)

	text := "Same content different source."
	chunks1 := c.Chunk(text, "file1.md")
	chunks2 := c.Chunk(text, "file2.md")

	if chunks1[0].ID == chunks2[0].ID {
		t.Error("expected different IDs for different sources with same content")
	}
}

func TestChunker_UUIDFormat(t *testing.T) {
	c := MustNew(100, 20)

	chunks := c.Chunk("Test content for UUID validation.", "test.md")
	if len(chunks) == 0 {
		t.Fatal("expected at least one chunk")
	}

	// UUID format: 8-4-4-4-12 hex chars with dashes
	id := chunks[0].ID
	parts := strings.Split(id, "-")
	if len(parts) != 5 {
		t.Errorf("expected UUID with 5 parts, got %d: %s", len(parts), id)
	}

	expectedLengths := []int{8, 4, 4, 4, 12}
	for i, part := range parts {
		if len(part) != expectedLengths[i] {
			t.Errorf("UUID part %d: expected length %d, got %d", i, expectedLengths[i], len(part))
		}
	}
}

func TestChunker_OverlapClampedToSize(t *testing.T) {
	// Overlap >= size should be clamped to size/4
	c, err := New(100, 200)
	if err != nil {
		t.Fatalf("New() should clamp overlap, got error: %v", err)
	}

	// Should not panic and should create valid chunks
	chunks := c.Chunk("Test content that is reasonably long enough to chunk.", "test.md")
	if len(chunks) == 0 {
		t.Error("expected at least one chunk even with clamped overlap")
	}
}

func TestNew_ValidationErrors(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		overlap int
		wantErr error
	}{
		{"zero size", 0, 10, ErrInvalidChunkSize},
		{"negative size", -100, 10, ErrInvalidChunkSize},
		{"negative overlap", 100, -10, ErrNegativeOverlap},
		{"valid params", 100, 20, nil},
		{"overlap equals size (clamped)", 100, 100, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.size, tt.overlap)
			if tt.wantErr != nil {
				if err != tt.wantErr {
					t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else if err != nil {
				t.Errorf("New() unexpected error = %v", err)
			}
		})
	}
}

func TestNew_Defaults(t *testing.T) {
	c, err := New(512, 64)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	if c.size != 512 {
		t.Errorf("expected size 512, got %d", c.size)
	}
	if c.overlap != 64 {
		t.Errorf("expected overlap 64, got %d", c.overlap)
	}
}

