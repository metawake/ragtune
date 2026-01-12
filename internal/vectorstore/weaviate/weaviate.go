// Package weaviate implements the vectorstore.Store interface for Weaviate.
package weaviate

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/metawake/ragtune/internal/vectorstore"
)

// Compile-time interface compliance check.
var _ vectorstore.Store = (*Client)(nil)

// Client implements vectorstore.Store for Weaviate using REST API.
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// New creates a new Weaviate client.
// host should be the Weaviate server address, e.g., "localhost:8080"
func New(ctx context.Context, host string, scheme string) (*Client, error) {
	baseURL := fmt.Sprintf("%s://%s", scheme, host)

	client := &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	// Test connection
	_, err := client.doRequest(ctx, "GET", "/v1/.well-known/ready", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Weaviate at %s: %w", baseURL, err)
	}

	return client, nil
}

// className converts a collection name to a valid Weaviate class name.
func className(collection string) string {
	clean := strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			return r
		}
		return '_'
	}, collection)

	if len(clean) == 0 {
		return "Ragtune"
	}

	// Capitalize first letter (Weaviate requires PascalCase)
	first := strings.ToUpper(string(clean[0]))
	if len(clean) == 1 {
		return "Ragtune_" + first
	}
	return "Ragtune_" + first + clean[1:]
}

// EnsureCollection creates a class for the collection if it doesn't exist.
func (c *Client) EnsureCollection(ctx context.Context, name string, dim int) error {
	class := className(name)

	// Check if class exists
	_, err := c.doRequest(ctx, "GET", "/v1/schema/"+class, nil)
	if err == nil {
		return nil // Class exists
	}

	// Create class
	classObj := map[string]interface{}{
		"class":       class,
		"description": "RagTune collection: " + name,
		"vectorIndexConfig": map[string]interface{}{
			"distance": "cosine",
		},
		"properties": []map[string]interface{}{
			{"name": "source", "dataType": []string{"text"}},
			{"name": "text", "dataType": []string{"text"}},
			{"name": "chunk_index", "dataType": []string{"int"}},
		},
	}

	_, err = c.doRequest(ctx, "POST", "/v1/schema", classObj)
	if err != nil {
		return fmt.Errorf("failed to create class: %w", err)
	}

	return nil
}

// Upsert inserts or updates points in a collection.
func (c *Client) Upsert(ctx context.Context, collection string, points []vectorstore.Point) error {
	if len(points) == 0 {
		return nil
	}

	class := className(collection)

	// Batch insert
	objects := make([]map[string]interface{}, len(points))
	for i, p := range points {
		props := map[string]interface{}{}
		if p.Payload != nil {
			if source, ok := p.Payload["source"].(string); ok {
				props["source"] = source
			}
			if text, ok := p.Payload["text"].(string); ok {
				props["text"] = text
			}
			if idx, ok := p.Payload["chunk_index"].(int); ok {
				props["chunk_index"] = idx
			}
		}

		objects[i] = map[string]interface{}{
			"class":      class,
			"id":         p.ID,
			"properties": props,
			"vector":     p.Vector,
		}
	}

	body := map[string]interface{}{
		"objects": objects,
	}

	_, err := c.doRequest(ctx, "POST", "/v1/batch/objects", body)
	if err != nil {
		return fmt.Errorf("batch upsert failed: %w", err)
	}

	return nil
}

// Search performs similarity search and returns top-k results.
func (c *Client) Search(ctx context.Context, collection string, vector []float32, topK int) ([]vectorstore.Result, error) {
	class := className(collection)

	// GraphQL query for nearVector search
	query := fmt.Sprintf(`{
		Get {
			%s(nearVector: {vector: %s}, limit: %d) {
				_additional { id distance }
				source
				text
				chunk_index
			}
		}
	}`, class, vectorToJSON(vector), topK)

	body := map[string]interface{}{
		"query": query,
	}

	respBody, err := c.doRequest(ctx, "POST", "/v1/graphql", body)
	if err != nil {
		return nil, fmt.Errorf("search failed: %w", err)
	}

	var resp graphQLResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if len(resp.Errors) > 0 {
		return nil, fmt.Errorf("graphql errors: %s", resp.Errors[0].Message)
	}

	items, ok := resp.Data.Get[class]
	if !ok || len(items) == 0 {
		return []vectorstore.Result{}, nil
	}

	results := make([]vectorstore.Result, len(items))
	for i, item := range items {
		distance := float32(0)
		id := ""
		if item.Additional != nil {
			if d, ok := item.Additional["distance"].(float64); ok {
				distance = float32(d)
			}
			if idVal, ok := item.Additional["id"].(string); ok {
				id = idVal
			}
		}

		payload := map[string]interface{}{}
		if item.Source != "" {
			payload["source"] = item.Source
		}
		if item.Text != "" {
			payload["text"] = item.Text
		}
		if item.ChunkIndex != 0 {
			payload["chunk_index"] = item.ChunkIndex
		}

		results[i] = vectorstore.Result{
			ID:      id,
			Score:   1 - distance, // Convert distance to similarity
			Payload: payload,
		}
	}

	return results, nil
}

// Count returns the number of objects in a collection.
func (c *Client) Count(ctx context.Context, collection string) (int64, error) {
	class := className(collection)

	query := fmt.Sprintf(`{
		Aggregate {
			%s {
				meta { count }
			}
		}
	}`, class)

	body := map[string]interface{}{
		"query": query,
	}

	respBody, err := c.doRequest(ctx, "POST", "/v1/graphql", body)
	if err != nil {
		return 0, fmt.Errorf("count failed: %w", err)
	}

	var resp struct {
		Data struct {
			Aggregate map[string][]struct {
				Meta struct {
					Count int64 `json:"count"`
				} `json:"meta"`
			} `json:"Aggregate"`
		} `json:"data"`
	}

	if err := json.Unmarshal(respBody, &resp); err != nil {
		return 0, fmt.Errorf("failed to parse response: %w", err)
	}

	items := resp.Data.Aggregate[class]
	if len(items) == 0 {
		return 0, nil
	}

	return items[0].Meta.Count, nil
}

// DeleteCollection removes a class and all its data.
func (c *Client) DeleteCollection(ctx context.Context, name string) error {
	class := className(name)
	_, err := c.doRequest(ctx, "DELETE", "/v1/schema/"+class, nil)
	if err != nil {
		// Ignore not found errors
		if strings.Contains(err.Error(), "404") {
			return nil
		}
		return fmt.Errorf("failed to delete class: %w", err)
	}
	return nil
}

// Close releases resources (no-op for HTTP client).
func (c *Client) Close() error {
	return nil
}

// Helper methods

func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}) ([]byte, error) {
	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request: %w", err)
		}
		reqBody = bytes.NewReader(jsonBody)
	}

	url := c.baseURL + path
	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API error %d: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

func vectorToJSON(vec []float32) string {
	parts := make([]string, len(vec))
	for i, v := range vec {
		parts[i] = fmt.Sprintf("%f", v)
	}
	return "[" + strings.Join(parts, ",") + "]"
}

// API types

type graphQLResponse struct {
	Data struct {
		Get map[string][]searchResult `json:"Get"`
	} `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

type searchResult struct {
	Additional map[string]interface{} `json:"_additional"`
	Source     string                 `json:"source"`
	Text       string                 `json:"text"`
	ChunkIndex int                    `json:"chunk_index"`
}
