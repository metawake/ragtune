// Package chunker provides simple text chunking for document ingestion.
package chunker

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/google/uuid"
)

// Validation errors for chunker configuration.
var (
	// ErrInvalidChunkSize is returned when chunk size is not positive.
	ErrInvalidChunkSize = errors.New("chunk size must be positive")

	// ErrNegativeOverlap is returned when overlap is negative.
	ErrNegativeOverlap = errors.New("overlap cannot be negative")
)

// Chunk represents a text chunk with metadata.
type Chunk struct {
	ID     string // Unique identifier (hash-based)
	Text   string // The chunk text
	Source string // Source document path
	Index  int    // Chunk index within the document
}

// Chunker splits text into overlapping chunks.
type Chunker struct {
	size    int
	overlap int
}

// New creates a new Chunker with the specified size and overlap.
// Returns an error if size <= 0 or overlap < 0.
// If overlap >= size, it is automatically clamped to size/4.
func New(size, overlap int) (*Chunker, error) {
	if size <= 0 {
		return nil, ErrInvalidChunkSize
	}
	if overlap < 0 {
		return nil, ErrNegativeOverlap
	}
	if overlap >= size {
		overlap = size / 4
	}
	return &Chunker{
		size:    size,
		overlap: overlap,
	}, nil
}

// MustNew creates a new Chunker, panicking on invalid input.
// Use this only in tests or when inputs are known-valid constants.
func MustNew(size, overlap int) *Chunker {
	c, err := New(size, overlap)
	if err != nil {
		panic(err)
	}
	return c
}

// Chunk splits text into chunks and returns them with metadata.
func (c *Chunker) Chunk(text, source string) []Chunk {
	// Sanitize and normalize
	text = sanitizeUTF8(text)
	text = strings.TrimSpace(text)
	if len(text) == 0 {
		return nil
	}

	var chunks []Chunk
	start := 0
	index := 0

	for start < len(text) {
		end := start + c.size
		if end > len(text) {
			end = len(text)
		}

		// Try to break at word boundary
		if end < len(text) {
			lastSpace := strings.LastIndex(text[start:end], " ")
			if lastSpace > c.size/2 {
				end = start + lastSpace
			}
		}

		chunkText := strings.TrimSpace(text[start:end])
		if len(chunkText) > 0 {
			chunks = append(chunks, Chunk{
				ID:     generateChunkID(source, index, chunkText),
				Text:   chunkText,
				Source: source,
				Index:  index,
			})
			index++
		}

		// Break if we've reached the end
		if end >= len(text) {
			break
		}

		// Move start position forward, ensuring progress
		// Step forward by (chunk length - overlap), minimum 1 char
		step := (end - start) - c.overlap
		if step < 1 {
			step = 1
		}
		start = start + step
	}

	return chunks
}

// generateChunkID creates a deterministic UUID for a chunk based on source and content.
// Uses UUID v5 (SHA-1 namespace) for reproducibility.
func generateChunkID(source string, index int, text string) string {
	// Create a deterministic UUID based on content hash
	h := sha256.Sum256([]byte(fmt.Sprintf("%s:%d:%s", source, index, text)))
	// Use first 16 bytes of hash to create a valid UUID
	u, _ := uuid.FromBytes(h[:16])
	return u.String()
}

// sanitizeUTF8 removes invalid UTF-8 sequences and control characters.
func sanitizeUTF8(s string) string {
	if utf8.ValidString(s) {
		return removeControlChars(s)
	}
	// Replace invalid UTF-8 with replacement character
	var b strings.Builder
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		if r == utf8.RuneError && size == 1 {
			b.WriteRune('?')
			i++
		} else {
			b.WriteRune(r)
			i += size
		}
	}
	return removeControlChars(b.String())
}

// removeControlChars removes non-printable control characters except newline/tab.
func removeControlChars(s string) string {
	var b strings.Builder
	for _, r := range s {
		if r == '\n' || r == '\t' || r == '\r' || r >= 32 {
			b.WriteRune(r)
		}
	}
	return b.String()
}

