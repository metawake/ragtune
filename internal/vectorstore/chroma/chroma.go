// Package chroma implements the vectorstore.Store interface for Chroma.
package chroma

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/metawake/ragtune/internal/vectorstore"
)

// Compile-time interface compliance check.
var _ vectorstore.Store = (*Client)(nil)

// Client implements vectorstore.Store for Chroma.
type Client struct {
	baseURL    string
	httpClient *http.Client
	tenant     string
	database   string
}

// New creates a new Chroma client.
// baseURL should be the Chroma server URL, e.g., "http://localhost:8000"
func New(ctx context.Context, baseURL string) (*Client, error) {
	client := &Client{
		baseURL:  baseURL,
		tenant:   "default_tenant",
		database: "default_database",
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	// Test connection using v2 API
	_, err := client.doRequest(ctx, "GET", "/api/v2/heartbeat", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Chroma at %s: %w", baseURL, err)
	}

	return client, nil
}

// EnsureCollection creates a collection if it doesn't exist.
func (c *Client) EnsureCollection(ctx context.Context, name string, dim int) error {
	// Try to get the collection first
	_, err := c.getCollection(ctx, name)
	if err == nil {
		return nil // Collection exists
	}

	// Create collection using v2 API (tenant/database scoped)
	body := map[string]interface{}{
		"name": name,
		"metadata": map[string]interface{}{
			"hnsw:space": "cosine",
		},
	}

	path := fmt.Sprintf("/api/v2/tenants/%s/databases/%s/collections", c.tenant, c.database)
	_, err = c.doRequest(ctx, "POST", path, body)
	if err != nil {
		return fmt.Errorf("failed to create collection: %w", err)
	}

	return nil
}

// Upsert inserts or updates points in a collection.
func (c *Client) Upsert(ctx context.Context, collection string, points []vectorstore.Point) error {
	if len(points) == 0 {
		return nil
	}

	col, err := c.getCollection(ctx, collection)
	if err != nil {
		return fmt.Errorf("collection not found: %w", err)
	}

	ids := make([]string, len(points))
	embeddings := make([][]float32, len(points))
	metadatas := make([]map[string]interface{}, len(points))
	documents := make([]string, len(points))

	for i, p := range points {
		ids[i] = p.ID
		embeddings[i] = p.Vector

		meta := map[string]interface{}{}
		doc := ""
		if p.Payload != nil {
			if source, ok := p.Payload["source"].(string); ok {
				meta["source"] = source
			}
			if text, ok := p.Payload["text"].(string); ok {
				doc = text
			}
			if idx, ok := p.Payload["chunk_index"].(int); ok {
				meta["chunk_index"] = idx
			}
		}
		metadatas[i] = meta
		documents[i] = doc
	}

	body := map[string]interface{}{
		"ids":        ids,
		"embeddings": embeddings,
		"metadatas":  metadatas,
		"documents":  documents,
	}

	path := fmt.Sprintf("/api/v2/tenants/%s/databases/%s/collections/%s/upsert", c.tenant, c.database, col.ID)
	_, err = c.doRequest(ctx, "POST", path, body)
	if err != nil {
		return fmt.Errorf("upsert failed: %w", err)
	}

	return nil
}

// Search performs similarity search and returns top-k results.
func (c *Client) Search(ctx context.Context, collection string, vector []float32, topK int) ([]vectorstore.Result, error) {
	col, err := c.getCollection(ctx, collection)
	if err != nil {
		return nil, fmt.Errorf("collection not found: %w", err)
	}

	body := map[string]interface{}{
		"query_embeddings": [][]float32{vector},
		"n_results":        topK,
		"include":          []string{"metadatas", "documents", "distances"},
	}

	path := fmt.Sprintf("/api/v2/tenants/%s/databases/%s/collections/%s/query", c.tenant, c.database, col.ID)
	respBody, err := c.doRequest(ctx, "POST", path, body)
	if err != nil {
		return nil, fmt.Errorf("search failed: %w", err)
	}

	var resp queryResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if len(resp.IDs) == 0 || len(resp.IDs[0]) == 0 {
		return []vectorstore.Result{}, nil
	}

	// First query result
	ids := resp.IDs[0]
	distances := resp.Distances[0]
	metadatas := resp.Metadatas[0]
	documents := resp.Documents[0]

	results := make([]vectorstore.Result, len(ids))
	for i := range ids {
		// Convert distance to score (Chroma returns L2 or cosine distance)
		// For cosine, distance is 1 - similarity, so score = 1 - distance
		score := float32(1.0) - distances[i]

		payload := map[string]interface{}{}
		if i < len(metadatas) && metadatas[i] != nil {
			payload = metadatas[i]
		}
		if i < len(documents) && documents[i] != "" {
			payload["text"] = documents[i]
		}

		results[i] = vectorstore.Result{
			ID:      ids[i],
			Score:   score,
			Payload: payload,
		}
	}

	return results, nil
}

// Count returns the number of points in a collection.
func (c *Client) Count(ctx context.Context, collection string) (int64, error) {
	col, err := c.getCollection(ctx, collection)
	if err != nil {
		return 0, fmt.Errorf("collection not found: %w", err)
	}

	path := fmt.Sprintf("/api/v2/tenants/%s/databases/%s/collections/%s/count", c.tenant, c.database, col.ID)
	respBody, err := c.doRequest(ctx, "GET", path, nil)
	if err != nil {
		return 0, fmt.Errorf("count failed: %w", err)
	}

	var count int64
	if err := json.Unmarshal(respBody, &count); err != nil {
		return 0, fmt.Errorf("failed to parse count: %w", err)
	}

	return count, nil
}

// DeleteCollection removes a collection and all its data.
func (c *Client) DeleteCollection(ctx context.Context, name string) error {
	path := fmt.Sprintf("/api/v2/tenants/%s/databases/%s/collections/%s", c.tenant, c.database, name)
	_, err := c.doRequest(ctx, "DELETE", path, nil)
	if err != nil {
		// Ignore not found errors
		return nil
	}
	return nil
}

// Close releases resources (no-op for HTTP client).
func (c *Client) Close() error {
	return nil
}

// Helper methods

func (c *Client) getCollection(ctx context.Context, name string) (*collectionInfo, error) {
	path := fmt.Sprintf("/api/v2/tenants/%s/databases/%s/collections/%s", c.tenant, c.database, name)
	respBody, err := c.doRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}

	var col collectionInfo
	if err := json.Unmarshal(respBody, &col); err != nil {
		return nil, fmt.Errorf("failed to parse collection: %w", err)
	}

	return &col, nil
}

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

// API types

type collectionInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type queryResponse struct {
	IDs       [][]string                   `json:"ids"`
	Distances [][]float32                  `json:"distances"`
	Metadatas [][]map[string]interface{}   `json:"metadatas"`
	Documents [][]string                   `json:"documents"`
}
