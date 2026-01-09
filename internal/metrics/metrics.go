// Package metrics provides RAG retrieval quality metrics.
package metrics

import (
	"math"
	"sort"
)

// Result holds computed metrics for a simulation run.
type Result struct {
	RecallAtK    float64 `json:"recall_at_k"`
	MRR          float64 `json:"mrr"`
	NDCGAtK      float64 `json:"ndcg_at_k"`      // Normalized Discounted Cumulative Gain
	Coverage     float64 `json:"coverage"`
	Redundancy   float64 `json:"redundancy"`
	DiversityAtK float64 `json:"diversity_at_k"` // Fraction of unique docs in top-K results
	// Latency stats in milliseconds
	LatencyP50 float64 `json:"latency_p50_ms,omitempty"`
	LatencyP95 float64 `json:"latency_p95_ms,omitempty"`
	LatencyP99 float64 `json:"latency_p99_ms,omitempty"`
	LatencyAvg float64 `json:"latency_avg_ms,omitempty"`
}

// QueryResult represents retrieval results for a single query.
type QueryResult struct {
	QueryID      string    `json:"query_id"`
	Query        string    `json:"query"`
	RetrievedIDs []string  `json:"retrieved_ids"`
	RelevantIDs  []string  `json:"relevant_ids"`
	Scores       []float32 `json:"scores"`
	LatencyMs    float64   `json:"latency_ms,omitempty"` // Query latency in milliseconds
}

// Compute calculates all metrics from query results.
func Compute(results []QueryResult, k int) Result {
	if len(results) == 0 {
		return Result{}
	}

	var totalRecall, totalRR, totalNDCG, totalDiversity float64
	allRetrieved := make(map[string]int)   // doc -> count
	allRelevant := make(map[string]struct{})

	for _, r := range results {
		// Track all relevant docs
		for _, id := range r.RelevantIDs {
			allRelevant[id] = struct{}{}
		}

		// Recall@k: fraction of relevant docs in top-k
		totalRecall += RecallAtK(r.RetrievedIDs, r.RelevantIDs, k)

		// MRR: reciprocal rank of first relevant doc
		totalRR += ReciprocalRank(r.RetrievedIDs, r.RelevantIDs)

		// NDCG@k: normalized discounted cumulative gain
		totalNDCG += NDCGAtK(r.RetrievedIDs, r.RelevantIDs, k)

		// Diversity@k: fraction of unique docs in top-k
		totalDiversity += DiversityAtK(r.RetrievedIDs, k)

		// Track retrieved docs for coverage/redundancy
		topK := r.RetrievedIDs
		if len(topK) > k {
			topK = topK[:k]
		}
		for _, id := range topK {
			allRetrieved[id]++
		}
	}

	n := float64(len(results))

	// Coverage: fraction of relevant docs ever retrieved
	coverage := Coverage(allRetrieved, allRelevant)

	// Redundancy: average times a doc is retrieved across queries
	redundancy := Redundancy(allRetrieved)

	// Compute latency stats
	latencyP50, latencyP95, latencyP99, latencyAvg := ComputeLatencyStats(results)

	return Result{
		RecallAtK:    totalRecall / n,
		MRR:          totalRR / n,
		NDCGAtK:      totalNDCG / n,
		Coverage:     coverage,
		Redundancy:   redundancy,
		DiversityAtK: totalDiversity / n,
		LatencyP50:   latencyP50,
		LatencyP95:   latencyP95,
		LatencyP99:   latencyP99,
		LatencyAvg:   latencyAvg,
	}
}

// RecallAtK computes recall for a single query.
// Returns the fraction of relevant documents found in the top-k retrieved.
// Deduplicates retrieved IDs (multiple chunks from same doc count once).
func RecallAtK(retrieved, relevant []string, k int) float64 {
	if len(relevant) == 0 {
		return 1.0 // No relevant docs = perfect recall
	}

	topK := retrieved
	if len(topK) > k {
		topK = topK[:k]
	}

	relevantSet := make(map[string]struct{})
	for _, id := range relevant {
		relevantSet[id] = struct{}{}
	}

	// Track which relevant docs were found (deduplicated)
	foundSet := make(map[string]struct{})
	for _, id := range topK {
		if _, ok := relevantSet[id]; ok {
			foundSet[id] = struct{}{}
		}
	}

	return float64(len(foundSet)) / float64(len(relevant))
}

// ReciprocalRank computes the reciprocal rank for a single query.
// Returns 1/rank of the first relevant document, or 0 if none found.
func ReciprocalRank(retrieved, relevant []string) float64 {
	relevantSet := make(map[string]struct{})
	for _, id := range relevant {
		relevantSet[id] = struct{}{}
	}

	for i, id := range retrieved {
		if _, ok := relevantSet[id]; ok {
			return 1.0 / float64(i+1)
		}
	}
	return 0
}

// Coverage computes the fraction of relevant docs that were ever retrieved.
func Coverage(retrieved map[string]int, relevant map[string]struct{}) float64 {
	if len(relevant) == 0 {
		return 1.0
	}

	found := 0
	for id := range relevant {
		if _, ok := retrieved[id]; ok {
			found++
		}
	}
	return float64(found) / float64(len(relevant))
}

// Redundancy computes how often docs are retrieved on average.
// Higher values indicate the same docs appear across many queries.
func Redundancy(retrieved map[string]int) float64 {
	if len(retrieved) == 0 {
		return 0
	}

	total := 0
	for _, count := range retrieved {
		total += count
	}
	return float64(total) / float64(len(retrieved))
}

// DiversityAtK computes the fraction of unique documents in the top-K results.
// Returns 1.0 if all K results are from different documents.
// Returns < 1.0 if multiple chunks from the same document appear in top-K.
// This metric surfaces the "wasted slots" problem when chunk overlap is too high.
func DiversityAtK(retrieved []string, k int) float64 {
	if len(retrieved) == 0 || k == 0 {
		return 0
	}

	topK := retrieved
	if len(topK) > k {
		topK = topK[:k]
	}

	// Count unique documents
	unique := make(map[string]struct{})
	for _, id := range topK {
		unique[id] = struct{}{}
	}

	// Diversity = unique docs / actual results returned
	// Use actual length, not k, in case fewer than k results were returned
	return float64(len(unique)) / float64(len(topK))
}

// NDCGAtK computes Normalized Discounted Cumulative Gain for a single query.
// NDCG rewards having relevant documents ranked higher in the result list.
// Unlike MRR which only considers the first relevant result, NDCG considers
// ALL relevant results and penalizes those ranked lower with a logarithmic discount.
//
// Formula: NDCG = DCG / IDCG
// where DCG = Σ (relevance_i / log2(rank_i + 1))
// and IDCG is the ideal DCG (all relevant docs ranked first)
//
// Returns 1.0 for perfect ranking, 0.0 if no relevant docs retrieved.
// This metric is an industry standard (used in search engines, academic papers).
func NDCGAtK(retrieved, relevant []string, k int) float64 {
	if len(relevant) == 0 {
		return 1.0 // No relevant docs = perfect NDCG
	}

	topK := retrieved
	if len(topK) > k {
		topK = topK[:k]
	}

	// Build set of relevant docs for O(1) lookup
	relevantSet := make(map[string]struct{})
	for _, id := range relevant {
		relevantSet[id] = struct{}{}
	}

	// Compute DCG with binary relevance (1 if relevant, 0 otherwise)
	var dcg float64
	for i, id := range topK {
		if _, ok := relevantSet[id]; ok {
			// log2(rank+1) where rank is 1-indexed, so log2(i+2)
			dcg += 1.0 / math.Log2(float64(i+2))
		}
	}

	// Compute IDCG: ideal DCG if all relevant docs ranked first
	// Number of relevant docs that could fit in top-K
	idealK := k
	if len(relevant) < idealK {
		idealK = len(relevant)
	}

	var idcg float64
	for i := 0; i < idealK; i++ {
		idcg += 1.0 / math.Log2(float64(i+2))
	}

	if idcg == 0 {
		return 1.0
	}

	return dcg / idcg
}

// ComputeLatencyStats calculates p50, p95, p99, and average latency from query results.
// Returns zeros if no latency data is available.
func ComputeLatencyStats(results []QueryResult) (p50, p95, p99, avg float64) {
	if len(results) == 0 {
		return 0, 0, 0, 0
	}

	// Collect all latencies (skip zeros)
	var latencies []float64
	var sum float64
	for _, r := range results {
		if r.LatencyMs > 0 {
			latencies = append(latencies, r.LatencyMs)
			sum += r.LatencyMs
		}
	}

	if len(latencies) == 0 {
		return 0, 0, 0, 0
	}

	// Sort for percentile calculation
	sort.Float64s(latencies)

	avg = sum / float64(len(latencies))
	p50 = percentile(latencies, 50)
	p95 = percentile(latencies, 95)
	p99 = percentile(latencies, 99)

	return p50, p95, p99, avg
}

// percentile calculates the p-th percentile of a sorted slice.
func percentile(sorted []float64, p int) float64 {
	if len(sorted) == 0 {
		return 0
	}
	if len(sorted) == 1 {
		return sorted[0]
	}

	// Use nearest-rank method
	rank := float64(p) / 100.0 * float64(len(sorted)-1)
	lower := int(rank)
	upper := lower + 1
	if upper >= len(sorted) {
		return sorted[len(sorted)-1]
	}

	// Linear interpolation
	weight := rank - float64(lower)
	return sorted[lower]*(1-weight) + sorted[upper]*weight
}
