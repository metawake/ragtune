# Synthetic 50K Benchmark

Enterprise-scale benchmark for testing RagTune embedding throughput and batching efficiency.

## Dataset Statistics

| Metric | Value |
|--------|-------|
| Documents | 50,000 |
| Queries | 500 |
| Total Size | 142.8 MB |
| Avg Doc Size | 2996 chars |
| Generation Time | 93.7s |

## Document Types

The corpus contains synthetic enterprise documents across 15 departments:
Engineering, Product, Sales, Marketing, Finance, Legal, HR, Operations, Customer Success, Security, Data Science, DevOps, QA, Design, Research

Document types include: policy, procedure, guide, specification, report, memo, proposal, review, analysis, summary

## Performance Targets

For enterprise viability, embedding 50k documents should complete in:

| Embedder | Target Time | Docs/Second |
|----------|-------------|-------------|
| TEI (GPU) | < 2 min | ~400/s |
| TEI (CPU) | < 10 min | ~80/s |
| Ollama (8 concurrent) | < 5 min | ~160/s |
| Ollama (sequential) | < 30 min | ~28/s |

## Usage

```bash
# Generate the benchmark dataset
python prepare.py

# Quick test with fewer documents
python prepare.py --docs 1000 --queries 50

# Run the benchmark
ragtune ingest ./benchmarks/synthetic-50k/corpus --collection synthetic-50k
ragtune simulate --collection synthetic-50k --queries ./benchmarks/synthetic-50k/queries.json
```

## Benchmark Results

### 1K Document Subset (7,087 chunks)

| Embedder | Time | Chunks/sec | 50K Projected |
|----------|------|------------|---------------|
| **TEI (CPU)** | 3m 45s | 31.5/sec | ~31 min ✓ |
| Ollama (concurrency=8) | 15m 20s | 7.7/sec | ~2 hours ⚠️ |

**Key Finding:** TEI is **4x faster** than Ollama due to native batching.

### Recommendations

| Scale | Recommended Embedder | Notes |
|-------|---------------------|-------|
| < 1K docs | Ollama | Simple setup, local |
| 1K-10K docs | TEI (CPU) | 10-30 min, no GPU needed |
| 10K-100K docs | TEI (GPU) | 5-15 min with GPU |
| 100K+ docs | TEI (GPU) + batching | Production deployment |

### Full Results Log

| Date | Embedder | Model | Time | Chunks/Sec | Notes |
|------|----------|-------|------|------------|-------|
| 2025-12-28 | TEI CPU | bge-small-en-v1.5 | 3m45s | 31.5 | 1K docs, 7087 chunks |
| 2025-12-28 | Ollama | nomic-embed-text | 15m20s | 7.7 | 1K docs, 7087 chunks |

---

*Generated on 2025-12-27T18:03:26.014519*
