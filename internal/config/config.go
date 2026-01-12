// Package config handles loading and parsing of simulation configurations
// and query files. It supports both YAML and JSON formats for flexibility.
package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// SimConfig represents a simulation configuration variant.
type SimConfig struct {
	Name       string `json:"name" yaml:"name"`
	TopK       int    `json:"top_k" yaml:"top_k"`
	ChunkSize  int    `json:"chunk_size,omitempty" yaml:"chunk_size,omitempty"`
	Overlap    int    `json:"overlap,omitempty" yaml:"overlap,omitempty"`
}

// ConfigFile represents the configs file structure.
type ConfigFile struct {
	Configs []SimConfig `json:"configs" yaml:"configs"`
}

// Query represents a query with ground truth.
type Query struct {
	ID          string   `json:"id" yaml:"id"`
	Text        string   `json:"text" yaml:"text"`
	RelevantDocs []string `json:"relevant_docs" yaml:"relevant_docs"`
	Notes       string   `json:"notes,omitempty" yaml:"notes,omitempty"`
}

// QueriesFile represents the queries file structure.
type QueriesFile struct {
	Queries []Query `json:"queries" yaml:"queries"`
}

// LoadConfigs loads simulation configs from a YAML or JSON file.
func LoadConfigs(path string) ([]SimConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cf ConfigFile
	ext := filepath.Ext(path)
	switch ext {
	case ".yaml", ".yml":
		if err := yaml.Unmarshal(data, &cf); err != nil {
			return nil, fmt.Errorf("failed to parse YAML: %w", err)
		}
	case ".json":
		if err := json.Unmarshal(data, &cf); err != nil {
			return nil, fmt.Errorf("failed to parse JSON: %w", err)
		}
	default:
		return nil, fmt.Errorf("unsupported config format: %s (use .yaml or .json)", ext)
	}

	// Set defaults
	for i := range cf.Configs {
		if cf.Configs[i].TopK == 0 {
			cf.Configs[i].TopK = 5
		}
	}

	return cf.Configs, nil
}

// LoadQueries loads queries from a JSON file.
func LoadQueries(path string) ([]Query, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read queries file: %w", err)
	}

	var qf QueriesFile
	if err := json.Unmarshal(data, &qf); err != nil {
		return nil, fmt.Errorf("failed to parse queries JSON: %w", err)
	}

	return qf.Queries, nil
}




