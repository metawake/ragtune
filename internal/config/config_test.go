package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfigs_YAML(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "configs.yaml")

	content := `configs:
  - name: default
    top_k: 5
  - name: large
    top_k: 10
    chunk_size: 512
    overlap: 50
`
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	configs, err := LoadConfigs(path)
	if err != nil {
		t.Fatalf("LoadConfigs failed: %v", err)
	}

	if len(configs) != 2 {
		t.Errorf("expected 2 configs, got %d", len(configs))
	}

	// Verify first config
	if configs[0].Name != "default" {
		t.Errorf("configs[0].Name = %q, want %q", configs[0].Name, "default")
	}
	if configs[0].TopK != 5 {
		t.Errorf("configs[0].TopK = %d, want 5", configs[0].TopK)
	}

	// Verify second config with optional fields
	if configs[1].Name != "large" {
		t.Errorf("configs[1].Name = %q, want %q", configs[1].Name, "large")
	}
	if configs[1].ChunkSize != 512 {
		t.Errorf("configs[1].ChunkSize = %d, want 512", configs[1].ChunkSize)
	}
	if configs[1].Overlap != 50 {
		t.Errorf("configs[1].Overlap = %d, want 50", configs[1].Overlap)
	}
}

func TestLoadConfigs_YML(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "configs.yml") // .yml extension

	content := `configs:
  - name: test
    top_k: 3
`
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	configs, err := LoadConfigs(path)
	if err != nil {
		t.Fatalf("LoadConfigs failed for .yml: %v", err)
	}

	if len(configs) != 1 || configs[0].Name != "test" {
		t.Errorf("unexpected configs: %+v", configs)
	}
}

func TestLoadConfigs_JSON(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "configs.json")

	content := `{
  "configs": [
    {"name": "json-config", "top_k": 7},
    {"name": "another", "top_k": 3, "chunk_size": 256}
  ]
}`
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	configs, err := LoadConfigs(path)
	if err != nil {
		t.Fatalf("LoadConfigs failed: %v", err)
	}

	if len(configs) != 2 {
		t.Errorf("expected 2 configs, got %d", len(configs))
	}
	if configs[0].Name != "json-config" {
		t.Errorf("configs[0].Name = %q, want %q", configs[0].Name, "json-config")
	}
	if configs[0].TopK != 7 {
		t.Errorf("configs[0].TopK = %d, want 7", configs[0].TopK)
	}
}

func TestLoadConfigs_Defaults(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "configs.yaml")

	// Config without top_k - should default to 5
	content := `configs:
  - name: no-topk
`
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	configs, err := LoadConfigs(path)
	if err != nil {
		t.Fatalf("LoadConfigs failed: %v", err)
	}

	if configs[0].TopK != 5 {
		t.Errorf("expected default TopK=5, got %d", configs[0].TopK)
	}
}

func TestLoadConfigs_InvalidExtension(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "configs.txt")

	if err := os.WriteFile(path, []byte("some content"), 0644); err != nil {
		t.Fatal(err)
	}

	_, err := LoadConfigs(path)
	if err == nil {
		t.Error("expected error for unsupported extension")
	}
}

func TestLoadConfigs_MissingFile(t *testing.T) {
	_, err := LoadConfigs("/nonexistent/path/configs.yaml")
	if err == nil {
		t.Error("expected error for missing file")
	}
}

func TestLoadConfigs_MalformedYAML(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "bad.yaml")

	content := `configs:
  - name: test
    top_k: [invalid yaml structure
`
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	_, err := LoadConfigs(path)
	if err == nil {
		t.Error("expected error for malformed YAML")
	}
}

func TestLoadConfigs_MalformedJSON(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "bad.json")

	content := `{"configs": [{"name": "test", broken json`
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	_, err := LoadConfigs(path)
	if err == nil {
		t.Error("expected error for malformed JSON")
	}
}

func TestLoadConfigs_EmptyFile(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "empty.yaml")

	if err := os.WriteFile(path, []byte(""), 0644); err != nil {
		t.Fatal(err)
	}

	configs, err := LoadConfigs(path)
	if err != nil {
		t.Fatalf("LoadConfigs failed: %v", err)
	}

	// Empty configs should return nil/empty slice
	if len(configs) != 0 {
		t.Errorf("expected 0 configs for empty file, got %d", len(configs))
	}
}

// --- LoadQueries Tests ---

func TestLoadQueries_Valid(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "queries.json")

	content := `{
  "queries": [
    {
      "id": "q1",
      "text": "How do I reset my password?",
      "relevant_docs": ["auth/password.md"]
    },
    {
      "id": "q2",
      "text": "What are the API limits?",
      "relevant_docs": ["api/limits.md", "api/quotas.md"],
      "notes": "Common question"
    }
  ]
}`
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	queries, err := LoadQueries(path)
	if err != nil {
		t.Fatalf("LoadQueries failed: %v", err)
	}

	if len(queries) != 2 {
		t.Errorf("expected 2 queries, got %d", len(queries))
	}

	// Verify first query
	if queries[0].ID != "q1" {
		t.Errorf("queries[0].ID = %q, want %q", queries[0].ID, "q1")
	}
	if queries[0].Text != "How do I reset my password?" {
		t.Errorf("queries[0].Text mismatch")
	}
	if len(queries[0].RelevantDocs) != 1 {
		t.Errorf("queries[0].RelevantDocs count = %d, want 1", len(queries[0].RelevantDocs))
	}

	// Verify second query with multiple docs and notes
	if len(queries[1].RelevantDocs) != 2 {
		t.Errorf("queries[1].RelevantDocs count = %d, want 2", len(queries[1].RelevantDocs))
	}
	if queries[1].Notes != "Common question" {
		t.Errorf("queries[1].Notes = %q, want %q", queries[1].Notes, "Common question")
	}
}

func TestLoadQueries_MissingFile(t *testing.T) {
	_, err := LoadQueries("/nonexistent/queries.json")
	if err == nil {
		t.Error("expected error for missing file")
	}
}

func TestLoadQueries_MalformedJSON(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "bad.json")

	content := `{"queries": [{"id": "q1", broken`
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	_, err := LoadQueries(path)
	if err == nil {
		t.Error("expected error for malformed JSON")
	}
}

func TestLoadQueries_EmptyQueries(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "empty.json")

	content := `{"queries": []}`
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	queries, err := LoadQueries(path)
	if err != nil {
		t.Fatalf("LoadQueries failed: %v", err)
	}

	if len(queries) != 0 {
		t.Errorf("expected 0 queries, got %d", len(queries))
	}
}

func TestLoadQueries_MinimalFields(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "minimal.json")

	// Only required fields
	content := `{
  "queries": [
    {"id": "q1", "text": "test", "relevant_docs": ["doc.md"]}
  ]
}`
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	queries, err := LoadQueries(path)
	if err != nil {
		t.Fatalf("LoadQueries failed: %v", err)
	}

	if len(queries) != 1 {
		t.Fatalf("expected 1 query, got %d", len(queries))
	}

	q := queries[0]
	if q.ID != "q1" || q.Text != "test" || len(q.RelevantDocs) != 1 {
		t.Errorf("unexpected query: %+v", q)
	}
	if q.Notes != "" {
		t.Errorf("expected empty notes, got %q", q.Notes)
	}
}
