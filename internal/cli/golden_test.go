package cli

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestAppendGoldenQuery_NewFile(t *testing.T) {
	// Create temp directory
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "golden.json")

	// Append first query
	err := appendGoldenQuery(filePath, "How do I reset my password?", "docs/auth/password.md")
	if err != nil {
		t.Fatalf("appendGoldenQuery failed: %v", err)
	}

	// Verify file was created and has correct content
	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("failed to read golden file: %v", err)
	}

	var gq GoldenQueries
	if err := json.Unmarshal(data, &gq); err != nil {
		t.Fatalf("failed to parse golden file: %v", err)
	}

	if len(gq.Queries) != 1 {
		t.Errorf("expected 1 query, got %d", len(gq.Queries))
	}

	q := gq.Queries[0]
	if q.Text != "How do I reset my password?" {
		t.Errorf("expected query text 'How do I reset my password?', got %q", q.Text)
	}
	if len(q.RelevantDocs) != 1 || q.RelevantDocs[0] != "password.md" {
		t.Errorf("expected relevant docs ['password.md'], got %v", q.RelevantDocs)
	}
	if !strings.HasPrefix(q.ID, "q-") {
		t.Errorf("expected ID to start with 'q-', got %q", q.ID)
	}
}

func TestAppendGoldenQuery_AppendToExisting(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "golden.json")

	// Create initial file with one query
	initial := GoldenQueries{
		Queries: []GoldenQuery{
			{ID: "q-1", Text: "First query", RelevantDocs: []string{"first.md"}},
		},
	}
	data, _ := json.MarshalIndent(initial, "", "  ")
	os.WriteFile(filePath, data, 0644)

	// Append second query
	err := appendGoldenQuery(filePath, "Second query", "second.md")
	if err != nil {
		t.Fatalf("appendGoldenQuery failed: %v", err)
	}

	// Verify both queries exist
	data, _ = os.ReadFile(filePath)
	var gq GoldenQueries
	json.Unmarshal(data, &gq)

	if len(gq.Queries) != 2 {
		t.Errorf("expected 2 queries, got %d", len(gq.Queries))
	}

	// First query should be preserved
	if gq.Queries[0].Text != "First query" {
		t.Errorf("first query was modified")
	}

	// Second query should be added
	if gq.Queries[1].Text != "Second query" {
		t.Errorf("second query text incorrect: %q", gq.Queries[1].Text)
	}
}

func TestAppendGoldenQuery_DuplicateRejected(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "golden.json")

	// Add first query
	err := appendGoldenQuery(filePath, "Duplicate query", "doc.md")
	if err != nil {
		t.Fatalf("first append failed: %v", err)
	}

	// Try to add same query again
	err = appendGoldenQuery(filePath, "Duplicate query", "other.md")
	if err == nil {
		t.Error("expected error for duplicate query, got nil")
	}
	if !strings.Contains(err.Error(), "already exists") {
		t.Errorf("expected 'already exists' error, got: %v", err)
	}

	// Verify only one query exists
	data, _ := os.ReadFile(filePath)
	var gq GoldenQueries
	json.Unmarshal(data, &gq)

	if len(gq.Queries) != 1 {
		t.Errorf("expected 1 query (duplicate rejected), got %d", len(gq.Queries))
	}
}

func TestAppendGoldenQuery_PathNormalization(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "golden.json")

	// Append with full path
	err := appendGoldenQuery(filePath, "Test query", "/Users/user/docs/nested/dir/file.md")
	if err != nil {
		t.Fatalf("appendGoldenQuery failed: %v", err)
	}

	data, _ := os.ReadFile(filePath)
	var gq GoldenQueries
	json.Unmarshal(data, &gq)

	// Should only contain filename, not full path
	if gq.Queries[0].RelevantDocs[0] != "file.md" {
		t.Errorf("expected 'file.md', got %q", gq.Queries[0].RelevantDocs[0])
	}
}

func TestAppendGoldenQuery_UniqueIDs(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "golden.json")

	// Add multiple queries rapidly
	queries := []string{"Query 1", "Query 2", "Query 3"}
	for _, q := range queries {
		err := appendGoldenQuery(filePath, q, "doc.md")
		if err != nil {
			t.Fatalf("appendGoldenQuery failed for %q: %v", q, err)
		}
	}

	data, _ := os.ReadFile(filePath)
	var gq GoldenQueries
	json.Unmarshal(data, &gq)

	// Verify all IDs are unique
	ids := make(map[string]bool)
	for _, q := range gq.Queries {
		if ids[q.ID] {
			t.Errorf("duplicate ID found: %q", q.ID)
		}
		ids[q.ID] = true
	}
}

func TestAppendGoldenQuery_InvalidExistingFile(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "golden.json")

	// Write invalid JSON
	os.WriteFile(filePath, []byte("not valid json"), 0644)

	err := appendGoldenQuery(filePath, "Test query", "doc.md")
	if err == nil {
		t.Error("expected error for invalid JSON, got nil")
	}
	if !strings.Contains(err.Error(), "failed to parse") {
		t.Errorf("expected parse error, got: %v", err)
	}
}
