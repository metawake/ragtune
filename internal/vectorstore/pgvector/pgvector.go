// Package pgvector implements the vectorstore.Store interface for PostgreSQL with pgvector.
package pgvector

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/metawake/ragtune/internal/vectorstore"
	pgvec "github.com/pgvector/pgvector-go"
	pgvecpgx "github.com/pgvector/pgvector-go/pgx"
)

// Client implements vectorstore.Store for PostgreSQL with pgvector extension.
type Client struct {
	pool *pgxpool.Pool
}

// New creates a new pgvector client.
// connStr should be a PostgreSQL connection string, e.g.:
// "postgres://user:password@localhost:5432/dbname"
func New(ctx context.Context, connStr string) (*Client, error) {
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse connection string: %w", err)
	}

	// First, create a simple connection to ensure pgvector extension exists
	// (must happen before pool creation because AfterConnect registers types)
	initConn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	_, err = initConn.Exec(ctx, "CREATE EXTENSION IF NOT EXISTS vector")
	initConn.Close(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to enable pgvector extension: %w", err)
	}

	// Register pgvector types on each connection
	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		return pgvecpgx.RegisterTypes(ctx, conn)
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	return &Client{pool: pool}, nil
}

// tableName converts a collection name to a valid table name.
// Prefixes with "ragtune_" and sanitizes the name.
func tableName(collection string) string {
	// Basic sanitization: replace non-alphanumeric chars with underscore
	sanitized := strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_' {
			return r
		}
		return '_'
	}, collection)
	return "ragtune_" + sanitized
}

// EnsureCollection creates a table for the collection if it doesn't exist.
func (c *Client) EnsureCollection(ctx context.Context, name string, dim int) error {
	table := tableName(name)

	// Create table with vector column
	query := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id TEXT PRIMARY KEY,
			embedding vector(%d) NOT NULL,
			payload JSONB
		)
	`, table, dim)

	_, err := c.pool.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to create collection table: %w", err)
	}

	// Create HNSW index for fast similarity search
	indexQuery := fmt.Sprintf(`
		CREATE INDEX IF NOT EXISTS %s_embedding_idx 
		ON %s USING hnsw (embedding vector_cosine_ops)
	`, table, table)

	_, err = c.pool.Exec(ctx, indexQuery)
	if err != nil {
		return fmt.Errorf("failed to create embedding index: %w", err)
	}

	return nil
}

// Upsert inserts or updates points in a collection.
func (c *Client) Upsert(ctx context.Context, collection string, points []vectorstore.Point) error {
	if len(points) == 0 {
		return nil
	}

	table := tableName(collection)

	// Use batch for efficiency
	batch := &pgx.Batch{}

	for _, p := range points {
		payloadJSON, err := json.Marshal(p.Payload)
		if err != nil {
			return fmt.Errorf("failed to marshal payload for %s: %w", p.ID, err)
		}

		query := fmt.Sprintf(`
			INSERT INTO %s (id, embedding, payload) 
			VALUES ($1, $2, $3)
			ON CONFLICT (id) DO UPDATE SET 
				embedding = EXCLUDED.embedding,
				payload = EXCLUDED.payload
		`, table)

		batch.Queue(query, p.ID, pgvec.NewVector(p.Vector), payloadJSON)
	}

	results := c.pool.SendBatch(ctx, batch)
	defer results.Close()

	// Check all results
	for range points {
		_, err := results.Exec()
		if err != nil {
			return fmt.Errorf("failed to upsert point: %w", err)
		}
	}

	return nil
}

// Search performs similarity search and returns top-k results.
// Uses cosine distance (1 - cosine_similarity) for ranking.
func (c *Client) Search(ctx context.Context, collection string, vector []float32, topK int) ([]vectorstore.Result, error) {
	table := tableName(collection)

	// Query using cosine distance (lower is better, so we order ASC)
	// We return 1 - distance as score (higher is better, like cosine similarity)
	query := fmt.Sprintf(`
		SELECT id, 1 - (embedding <=> $1) as score, payload
		FROM %s
		ORDER BY embedding <=> $1
		LIMIT $2
	`, table)

	rows, err := c.pool.Query(ctx, query, pgvec.NewVector(vector), topK)
	if err != nil {
		return nil, fmt.Errorf("search failed: %w", err)
	}
	defer rows.Close()

	var results []vectorstore.Result
	for rows.Next() {
		var id string
		var score float32
		var payloadJSON []byte

		if err := rows.Scan(&id, &score, &payloadJSON); err != nil {
			return nil, fmt.Errorf("failed to scan result: %w", err)
		}

		var payload map[string]interface{}
		if len(payloadJSON) > 0 {
			if err := json.Unmarshal(payloadJSON, &payload); err != nil {
				return nil, fmt.Errorf("failed to unmarshal payload: %w", err)
			}
		}

		results = append(results, vectorstore.Result{
			ID:      id,
			Score:   score,
			Payload: payload,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating results: %w", err)
	}

	return results, nil
}

// Count returns the number of points in a collection.
func (c *Client) Count(ctx context.Context, collection string) (int64, error) {
	table := tableName(collection)

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", table)

	var count int64
	err := c.pool.QueryRow(ctx, query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("count failed: %w", err)
	}

	return count, nil
}

// DeleteCollection removes a collection table and all its data.
func (c *Client) DeleteCollection(ctx context.Context, name string) error {
	table := tableName(name)

	query := fmt.Sprintf("DROP TABLE IF EXISTS %s", table)
	_, err := c.pool.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to delete collection: %w", err)
	}

	return nil
}

// Close releases resources.
func (c *Client) Close() error {
	if c.pool != nil {
		c.pool.Close()
	}
	return nil
}
