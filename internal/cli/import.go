package cli

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	importOutput string
	importFormat string
)

var importCmd = &cobra.Command{
	Use:   "import-queries <file>",
	Short: "Import queries from CSV or JSON file",
	Long: `Import queries from a CSV or JSON file into golden queries format.

CSV Format:
  query,relevant_docs
  "How do I reset my password?",docs/auth/password.md
  "What are rate limits?",docs/api/limits.md;docs/api/quotas.md

  - Header row required
  - Multiple relevant docs separated by semicolon

JSON Format:
  {
    "queries": [
      {"id": "q1", "text": "query text", "relevant_docs": ["doc1.md"]}
    ]
  }

Examples:
  ragtune import-queries queries.csv
  ragtune import-queries queries.json --output golden-queries.json`,
	Args: cobra.ExactArgs(1),
	RunE: runImport,
}

func init() {
	importCmd.Flags().StringVar(&importOutput, "output", "golden-queries.json", "Output file path")
	importCmd.Flags().StringVar(&importFormat, "format", "", "Input format (csv, json) - auto-detected from extension if not specified")

	rootCmd.AddCommand(importCmd)
}

func runImport(cmd *cobra.Command, args []string) error {
	inputPath := args[0]

	// Determine format
	format := importFormat
	if format == "" {
		ext := strings.ToLower(filepath.Ext(inputPath))
		switch ext {
		case ".csv":
			format = "csv"
		case ".json":
			format = "json"
		default:
			return fmt.Errorf("cannot determine format from extension %q (use --format csv or --format json)", ext)
		}
	}

	var queries []GoldenQuery
	var err error

	switch format {
	case "csv":
		queries, err = parseCSVQueries(inputPath)
	case "json":
		queries, err = parseJSONQueries(inputPath)
	default:
		return fmt.Errorf("unsupported format: %s (use csv or json)", format)
	}

	if err != nil {
		return fmt.Errorf("failed to parse %s: %w", inputPath, err)
	}

	if len(queries) == 0 {
		return fmt.Errorf("no queries found in %s", inputPath)
	}

	// Load existing golden queries if output file exists
	var existing GoldenQueries
	if data, err := os.ReadFile(importOutput); err == nil {
		if err := json.Unmarshal(data, &existing); err != nil {
			return fmt.Errorf("failed to parse existing %s: %w", importOutput, err)
		}
	}

	// Build set of existing query texts for deduplication
	existingTexts := make(map[string]bool)
	for _, q := range existing.Queries {
		existingTexts[q.Text] = true
	}

	// Add new queries, skip duplicates
	added := 0
	skipped := 0
	for _, q := range queries {
		if existingTexts[q.Text] {
			skipped++
			continue
		}
		existing.Queries = append(existing.Queries, q)
		existingTexts[q.Text] = true
		added++
	}

	// Write output
	data, err := json.MarshalIndent(existing, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal output: %w", err)
	}

	if err := os.WriteFile(importOutput, data, 0644); err != nil {
		return fmt.Errorf("failed to write %s: %w", importOutput, err)
	}

	fmt.Printf("✓ Imported %d queries to %s", added, importOutput)
	if skipped > 0 {
		fmt.Printf(" (%d duplicates skipped)", skipped)
	}
	fmt.Println()

	return nil
}

// parseCSVQueries parses a CSV file with query,relevant_docs columns.
func parseCSVQueries(path string) ([]GoldenQuery, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV: %w", err)
	}

	if len(records) < 2 {
		return nil, fmt.Errorf("CSV must have header row and at least one data row")
	}

	// Find column indices from header
	header := records[0]
	queryCol := -1
	relevantCol := -1
	idCol := -1

	for i, col := range header {
		col = strings.ToLower(strings.TrimSpace(col))
		switch col {
		case "query", "text", "question":
			queryCol = i
		case "relevant_docs", "relevant", "docs", "answer_docs":
			relevantCol = i
		case "id":
			idCol = i
		}
	}

	if queryCol == -1 {
		return nil, fmt.Errorf("CSV must have 'query' or 'text' column")
	}
	if relevantCol == -1 {
		return nil, fmt.Errorf("CSV must have 'relevant_docs' or 'relevant' column")
	}

	var queries []GoldenQuery
	for i, row := range records[1:] {
		if len(row) <= queryCol || len(row) <= relevantCol {
			continue // Skip malformed rows
		}

		queryText := strings.TrimSpace(row[queryCol])
		if queryText == "" {
			continue
		}

		// Parse relevant docs (semicolon-separated)
		relevantStr := strings.TrimSpace(row[relevantCol])
		var relevantDocs []string
		for _, doc := range strings.Split(relevantStr, ";") {
			doc = strings.TrimSpace(doc)
			if doc != "" {
				relevantDocs = append(relevantDocs, doc)
			}
		}

		// Get or generate ID
		id := fmt.Sprintf("imported-%d", i+1)
		if idCol >= 0 && len(row) > idCol && row[idCol] != "" {
			id = strings.TrimSpace(row[idCol])
		}

		queries = append(queries, GoldenQuery{
			ID:           id,
			Text:         queryText,
			RelevantDocs: relevantDocs,
		})
	}

	return queries, nil
}

// parseJSONQueries parses a JSON file in golden queries format.
func parseJSONQueries(path string) ([]GoldenQuery, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var gq GoldenQueries
	if err := json.Unmarshal(data, &gq); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Ensure all queries have IDs
	for i := range gq.Queries {
		if gq.Queries[i].ID == "" {
			gq.Queries[i].ID = fmt.Sprintf("imported-%d", i+1)
		}
	}

	return gq.Queries, nil
}
