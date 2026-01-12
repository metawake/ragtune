// Package qdrant implements the vectorstore.Store interface for Qdrant.
package qdrant

import (
	"context"
	"fmt"

	pb "github.com/qdrant/go-client/qdrant"
	"github.com/metawake/ragtune/internal/vectorstore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Compile-time interface compliance check.
var _ vectorstore.Store = (*Client)(nil)

// Client implements vectorstore.Store for Qdrant.
type Client struct {
	conn        *grpc.ClientConn
	collections pb.CollectionsClient
	points      pb.PointsClient
}

// New creates a new Qdrant client.
func New(ctx context.Context, addr string) (*Client, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to qdrant at %s: %w", addr, err)
	}

	return &Client{
		conn:        conn,
		collections: pb.NewCollectionsClient(conn),
		points:      pb.NewPointsClient(conn),
	}, nil
}

// EnsureCollection creates a collection if it doesn't exist.
func (c *Client) EnsureCollection(ctx context.Context, name string, dim int) error {
	// Check if collection exists
	_, err := c.collections.Get(ctx, &pb.GetCollectionInfoRequest{
		CollectionName: name,
	})
	if err == nil {
		// Collection exists
		return nil
	}

	// Create collection
	_, err = c.collections.Create(ctx, &pb.CreateCollection{
		CollectionName: name,
		VectorsConfig: &pb.VectorsConfig{
			Config: &pb.VectorsConfig_Params{
				Params: &pb.VectorParams{
					Size:     uint64(dim),
					Distance: pb.Distance_Cosine,
				},
			},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to create collection %s: %w", name, err)
	}

	return nil
}

// Upsert inserts or updates points in a collection.
func (c *Client) Upsert(ctx context.Context, collection string, points []vectorstore.Point) error {
	if len(points) == 0 {
		return nil
	}

	// Convert to Qdrant points
	pbPoints := make([]*pb.PointStruct, len(points))
	for i, p := range points {
		pbPoints[i] = &pb.PointStruct{
			Id: &pb.PointId{
				PointIdOptions: &pb.PointId_Uuid{Uuid: p.ID},
			},
			Vectors: &pb.Vectors{
				VectorsOptions: &pb.Vectors_Vector{
					Vector: &pb.Vector{Data: p.Vector},
				},
			},
			Payload: toQdrantPayload(p.Payload),
		}
	}

	// Upsert in batches of 100
	batchSize := 100
	for i := 0; i < len(pbPoints); i += batchSize {
		end := i + batchSize
		if end > len(pbPoints) {
			end = len(pbPoints)
		}
		batch := pbPoints[i:end]

		_, err := c.points.Upsert(ctx, &pb.UpsertPoints{
			CollectionName: collection,
			Points:         batch,
			Wait:           boolPtr(true),
		})
		if err != nil {
			return fmt.Errorf("failed to upsert batch starting at %d: %w", i, err)
		}
	}

	return nil
}

// Search performs similarity search and returns top-k results.
func (c *Client) Search(ctx context.Context, collection string, vector []float32, topK int) ([]vectorstore.Result, error) {
	resp, err := c.points.Search(ctx, &pb.SearchPoints{
		CollectionName: collection,
		Vector:         vector,
		Limit:          uint64(topK),
		WithPayload: &pb.WithPayloadSelector{
			SelectorOptions: &pb.WithPayloadSelector_Enable{Enable: true},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("search failed: %w", err)
	}

	results := make([]vectorstore.Result, len(resp.Result))
	for i, r := range resp.Result {
		results[i] = vectorstore.Result{
			ID:      pointIDToString(r.Id),
			Score:   r.Score,
			Payload: fromQdrantPayload(r.Payload),
		}
	}

	return results, nil
}

// Count returns the number of points in a collection.
func (c *Client) Count(ctx context.Context, collection string) (int64, error) {
	resp, err := c.points.Count(ctx, &pb.CountPoints{
		CollectionName: collection,
		Exact:          boolPtr(true),
	})
	if err != nil {
		return 0, fmt.Errorf("count failed: %w", err)
	}

	return int64(resp.Result.Count), nil
}

// DeleteCollection removes a collection and all its data.
func (c *Client) DeleteCollection(ctx context.Context, name string) error {
	_, err := c.collections.Delete(ctx, &pb.DeleteCollection{
		CollectionName: name,
	})
	if err != nil {
		return fmt.Errorf("failed to delete collection %s: %w", name, err)
	}
	return nil
}

// Close releases resources.
func (c *Client) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// Helper functions

func boolPtr(b bool) *bool {
	return &b
}

func pointIDToString(id *pb.PointId) string {
	if id == nil {
		return ""
	}
	switch v := id.PointIdOptions.(type) {
	case *pb.PointId_Uuid:
		return v.Uuid
	case *pb.PointId_Num:
		return fmt.Sprintf("%d", v.Num)
	}
	return ""
}

func toQdrantPayload(payload map[string]interface{}) map[string]*pb.Value {
	if payload == nil {
		return nil
	}
	result := make(map[string]*pb.Value)
	for k, v := range payload {
		result[k] = toQdrantValue(v)
	}
	return result
}

func toQdrantValue(v interface{}) *pb.Value {
	switch val := v.(type) {
	case string:
		return &pb.Value{Kind: &pb.Value_StringValue{StringValue: val}}
	case int:
		return &pb.Value{Kind: &pb.Value_IntegerValue{IntegerValue: int64(val)}}
	case int64:
		return &pb.Value{Kind: &pb.Value_IntegerValue{IntegerValue: val}}
	case float64:
		return &pb.Value{Kind: &pb.Value_DoubleValue{DoubleValue: val}}
	case bool:
		return &pb.Value{Kind: &pb.Value_BoolValue{BoolValue: val}}
	default:
		return &pb.Value{Kind: &pb.Value_StringValue{StringValue: fmt.Sprintf("%v", val)}}
	}
}

func fromQdrantPayload(payload map[string]*pb.Value) map[string]interface{} {
	if payload == nil {
		return nil
	}
	result := make(map[string]interface{})
	for k, v := range payload {
		result[k] = fromQdrantValue(v)
	}
	return result
}

func fromQdrantValue(v *pb.Value) interface{} {
	if v == nil {
		return nil
	}
	switch val := v.Kind.(type) {
	case *pb.Value_StringValue:
		return val.StringValue
	case *pb.Value_IntegerValue:
		return val.IntegerValue
	case *pb.Value_DoubleValue:
		return val.DoubleValue
	case *pb.Value_BoolValue:
		return val.BoolValue
	default:
		return nil
	}
}

