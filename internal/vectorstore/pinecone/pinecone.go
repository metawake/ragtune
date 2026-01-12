// Package pinecone implements the vectorstore.Store interface for Pinecone.
package pinecone

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/metawake/ragtune/internal/vectorstore"
)

// Compile-time interface compliance check.
var _ vectorstore.Store = (*Client)(nil)

// Client implements vectorstore.Store for Pinecone.
type Client struct {
	apiKey     string
	host       string // e.g., "index-name-project.svc.environment.pinecone.io"
	httpClient *http.Client
}

// New creates a new Pinecone client.
// host should be the full Pinecone index host (from the Pinecone console).
// API key is read from PINECONE_API_KEY environment variable if not provided.
func New(ctx context.Context, host string, apiKey string) (*Client, error) {
	if apiKey == "" {
		apiKey = os.Getenv("PINECONE_API_KEY")
	}
	if apiKey == "" {
		return nil, fmt.Errorf("PINECONE_API_KEY not set")
	}

	if host == "" {
		return nil, fmt.Errorf("Pinecone host is required")
	}

	client := &Client{
		apiKey: apiKey,
		host:   host,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	// Test connection by describing index stats
	_, err := client.describeIndexStats(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Pinecone: %w", err)
	}

	return client, nil
}

// EnsureCollection is a no-op for Pinecone (index must be created via console/API).
// Pinecone uses namespaces within an index, which are created automatically on upsert.
func (c *Client) EnsureCollection(ctx context.Context, name string, dim int) error {
	// Namespaces are created automatically on first upsert
	return nil
}

// Upsert inserts or updates vectors in the index.
func (c *Client) Upsert(ctx context.Context, collection string, points []vectorstore.Point) error {
	if len(points) == 0 {
		return nil
	}

	// Pinecone has a limit of 100 vectors per upsert
	batchSize := 100
	for i := 0; i < len(points); i += batchSize {
		end := i + batchSize
		if end > len(points) {
			end = len(points)
		}
		batch := points[i:end]

		vectors := make([]pineconeVector, len(batch))
		for j, p := range batch {
			vectors[j] = pineconeVector{
				ID:       p.ID,
				Values:   p.Vector,
				Metadata: p.Payload,
			}
		}

		body := upsertRequest{
			Vectors:   vectors,
			Namespace: collection,
		}

		_, err := c.doRequest(ctx, "POST", "/vectors/upsert", body)
		if err != nil {
			return fmt.Errorf("upsert failed: %w", err)
		}
	}

	return nil
}

// Search performs similarity search and returns top-k results.
func (c *Client) Search(ctx context.Context, collection string, vector []float32, topK int) ([]vectorstore.Result, error) {
	body := queryRequest{
		Vector:          vector,
		TopK:            topK,
		Namespace:       collection,
		IncludeMetadata: true,
	}

	respBody, err := c.doRequest(ctx, "POST", "/query", body)
	if err != nil {
		return nil, fmt.Errorf("search failed: %w", err)
	}

	var resp queryResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("failed to parse search response: %w", err)
	}

	results := make([]vectorstore.Result, len(resp.Matches))
	for i, m := range resp.Matches {
		results[i] = vectorstore.Result{
			ID:      m.ID,
			Score:   m.Score,
			Payload: m.Metadata,
		}
	}

	return results, nil
}

// Count returns the number of vectors in a namespace.
func (c *Client) Count(ctx context.Context, collection string) (int64, error) {
	stats, err := c.describeIndexStats(ctx)
	if err != nil {
		return 0, err
	}

	if ns, ok := stats.Namespaces[collection]; ok {
		return int64(ns.VectorCount), nil
	}

	return 0, nil
}

// DeleteCollection deletes all vectors in a namespace.
func (c *Client) DeleteCollection(ctx context.Context, name string) error {
	body := deleteRequest{
		DeleteAll: true,
		Namespace: name,
	}

	_, err := c.doRequest(ctx, "POST", "/vectors/delete", body)
	if err != nil {
		return fmt.Errorf("delete failed: %w", err)
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

	url := fmt.Sprintf("https://%s%s", c.host, path)
	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Api-Key", c.apiKey)
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

func (c *Client) describeIndexStats(ctx context.Context) (*indexStats, error) {
	respBody, err := c.doRequest(ctx, "POST", "/describe_index_stats", struct{}{})
	if err != nil {
		return nil, err
	}

	var stats indexStats
	if err := json.Unmarshal(respBody, &stats); err != nil {
		return nil, fmt.Errorf("failed to parse stats: %w", err)
	}

	return &stats, nil
}

// API types

type pineconeVector struct {
	ID       string                 `json:"id"`
	Values   []float32              `json:"values"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type upsertRequest struct {
	Vectors   []pineconeVector `json:"vectors"`
	Namespace string           `json:"namespace,omitempty"`
}

type queryRequest struct {
	Vector          []float32 `json:"vector"`
	TopK            int       `json:"topK"`
	Namespace       string    `json:"namespace,omitempty"`
	IncludeMetadata bool      `json:"includeMetadata"`
}

type queryResponse struct {
	Matches []struct {
		ID       string                 `json:"id"`
		Score    float32                `json:"score"`
		Metadata map[string]interface{} `json:"metadata,omitempty"`
	} `json:"matches"`
}

type deleteRequest struct {
	DeleteAll bool   `json:"deleteAll"`
	Namespace string `json:"namespace,omitempty"`
}

type indexStats struct {
	Namespaces map[string]struct {
		VectorCount int `json:"vectorCount"`
	} `json:"namespaces"`
	TotalVectorCount int `json:"totalVectorCount"`
}
