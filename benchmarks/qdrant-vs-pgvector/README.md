# Qdrant vs pgvector Benchmark

Raw data from [docs/articles/03-qdrant-vs-pgvector.md](../../docs/articles/03-qdrant-vs-pgvector.md).

## Key Finding

**pgvector with default `ef_search=40` doesn't use the HNSW index on small tables.** The query planner falls back to sequential scan. Once you set `ef_search ≥ 87` (on this dataset), pgvector and Qdrant have nearly identical latency.

## Results Summary

### HotpotQA (5K docs, 2581 queries)

| Store | Config | p50 | p95 | Recall@5 |
|-------|--------|-----|-----|----------|
| Qdrant | HNSW | 48ms ± 1 | 114ms ± 14 | **0.911** |
| pgvector | HNSW (ef=100) | 49ms ± 1 | 119ms ± 6 | 0.900 |
| pgvector | Seq scan (ef=40) | 104ms ± 2 | 223ms ± 14 | 0.911 |

### CaseHOLD (500 docs, 500 queries)

| Store | Config | p50 | p95 | Recall@5 |
|-------|--------|-----|-----|----------|
| Qdrant | HNSW | 334ms ± 35 | 564ms ± 114 | 0.658 |
| pgvector | HNSW (ef=100) | 319ms ± 5 | 489ms ± 15 | 0.658 |
| pgvector | Seq scan (ef=40) | 330ms ± 31 | 524ms ± 74 | 0.658 |

## Environment

| Component | Version |
|-----------|---------|
| Machine | MacBook Pro M1, 16GB RAM |
| Qdrant | 1.16.3 (Docker) |
| PostgreSQL | 16.11 |
| pgvector | 0.8.1, HNSW (m=16, ef_construction=64) |
| Ollama | 0.6.8 |
| Embedding model | nomic-embed-text (768 dim) |
| RagTune | dev (commit 9f4b45e) |

## Raw Data Files

```
# Qdrant HotpotQA (3 runs)
2026-01-21T17-14-35Z.json
2026-01-21T17-16-57Z.json
2026-01-21T17-19-26Z.json

# pgvector HotpotQA - Seq Scan ef_search=40 (3 runs)
2026-01-21T17-22-00Z.json
2026-01-21T17-26-59Z.json
2026-01-21T17-32-28Z.json

# pgvector HotpotQA - HNSW ef_search=100 (3 runs)
2026-01-21T18-32-14Z.json
2026-01-21T18-34-44Z.json
2026-01-21T18-37-21Z.json

# Qdrant CaseHOLD (3 runs)
2026-01-21T17-37-52Z.json
2026-01-21T17-40-39Z.json
2026-01-21T17-43-35Z.json

# pgvector CaseHOLD - Seq Scan ef_search=40 (3 runs)
2026-01-21T17-47-06Z.json
2026-01-21T17-49-58Z.json
2026-01-21T17-53-17Z.json

# pgvector CaseHOLD - HNSW ef_search=100 (3 runs)
2026-01-21T18-39-57Z.json
2026-01-21T18-42-45Z.json
2026-01-21T18-45-38Z.json
```

## The ef_search Gotcha

pgvector's query planner uses `ef_search` to estimate HNSW index cost. With low values, it may choose sequential scan instead.

```sql
-- Check if your queries use the index
EXPLAIN ANALYZE SELECT id, embedding <=> $1 AS dist
FROM my_table ORDER BY dist LIMIT 5;

-- If you see "Seq Scan", increase ef_search:
ALTER DATABASE mydb SET hnsw.ef_search = 100;
```

On this dataset (6,785 chunks, 768 dimensions), the threshold was `ef_search ≥ 87`.

## Reproduce

```bash
# Start services
docker run -p 6333:6333 -p 6334:6334 qdrant/qdrant:v1.16.3

# Set ef_search BEFORE ingesting/benchmarking
psql -c "ALTER DATABASE postgres SET hnsw.ef_search = 100;"

# Ingest
ragtune ingest ./benchmarks/hotpotqa-5k/corpus \
  --collection hotpot5k-qdrant --store qdrant \
  --chunk-size 512 --embedder ollama

ragtune ingest ./benchmarks/hotpotqa-5k/corpus \
  --collection hotpot5k-pgvector --store pgvector \
  --pgvector-url "postgres://..." --chunk-size 512 --embedder ollama

# Benchmark (3x each)
for i in 1 2 3; do
  ragtune simulate --collection hotpot5k-qdrant --store qdrant \
    --queries ./benchmarks/hotpotqa-5k/queries.json --embedder ollama
  ragtune simulate --collection hotpot5k-pgvector --store pgvector \
    --pgvector-url "postgres://..." --queries ./benchmarks/hotpotqa-5k/queries.json --embedder ollama
done
```

## Extract Metrics

```bash
for f in *.json; do
  echo "$f"
  jq '{recall: .configs[0].metrics.recall_at_k, p50: .configs[0].metrics.latency_p50_ms, p95: .configs[0].metrics.latency_p95_ms}' "$f"
done
```
