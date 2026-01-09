# RagTune Documentation

Welcome to RagTune docs. Find what you need based on where you are in your journey.

## Getting Started

| Guide | Time | For |
|-------|------|-----|
| [Quickstart](articles/00-quickstart.md) | 5 min | First-time setup, see it working |

## Guides by Use Case

### I want to...

| Goal | Guide |
|------|-------|
| **Understand expected runtimes** | [Benchmarking Guide](articles/03-benchmarking-guide.md) |
| **Set up CI/CD quality gates** | [Deployment Patterns](articles/04-deployment-patterns.md#stage-2-cicd-regression-testing) |
| **Monitor production quality** | [Deployment Patterns](articles/04-deployment-patterns.md#stage-3-monitoring-production-drift-detection) |
| **Tune chunk sizes and top-k** | [Deployment Patterns](articles/04-deployment-patterns.md#stage-1-tuning-configuration-exploration) |
| **Compare embedding models** | [Embeddings Comparison](articles/02-embeddings-comparison.md) |
| **Optimize for large scale** | [Benchmarking Guide](articles/03-benchmarking-guide.md#embedder-comparison) |

## Guides by Complexity

### Beginner

- [Quickstart](articles/00-quickstart.md) — Get RagTune running in 5 minutes

### Intermediate

- [Benchmarking Guide](articles/03-benchmarking-guide.md) — Expected runtimes, query counts, scale planning
- [Deployment Patterns](articles/04-deployment-patterns.md) — Tuning vs CI/CD vs Monitoring

### Advanced

- [Chunk Size Discovery](articles/01-chunk-size-discovery.md) — Finding optimal chunking for your domain
- [Embeddings Comparison](articles/02-embeddings-comparison.md) — Deep dive into model selection

## Quick Reference

### Commands

```bash
ragtune ingest ./docs --collection name    # Load documents
ragtune explain "query" --collection name  # Debug single query
ragtune simulate --collection name --queries file.json  # Batch benchmark
ragtune compare --embedders a,b --docs ./docs --queries file.json  # Compare
ragtune report --run runs/latest.json      # Generate report
```

### Embedders

| Name | Command | Notes |
|------|---------|-------|
| Ollama (local) | `--embedder ollama` | Free, no API key |
| OpenAI | `--embedder openai` | Default |
| TEI | `--embedder tei` | 4x faster for scale |
| Cohere | `--embedder cohere` | Multilingual |
| Voyage | `--embedder voyage` | Domain-tuned |

### Recommended Query Counts

```
Queries = max(100, corpus_size × 1%, cap 5000)

1K docs   → 100 queries
10K docs  → 300 queries
50K docs  → 500 queries
100K docs → 1000 queries
```

### Expected Runtimes (TEI CPU)

| Corpus | Ingestion | Query Benchmark |
|--------|-----------|-----------------|
| 1K docs | 4 min | 5 sec |
| 10K docs | 40 min | 15 sec |
| 50K docs | 3 hours | 25 sec |

GPU: 5-10x faster for ingestion.

## All Articles

| Article | Description |
|---------|-------------|
| [00-quickstart.md](articles/00-quickstart.md) | First-time setup |
| [01-chunk-size-discovery.md](articles/01-chunk-size-discovery.md) | Chunking optimization |
| [02-embeddings-comparison.md](articles/02-embeddings-comparison.md) | Model selection |
| [03-benchmarking-guide.md](articles/03-benchmarking-guide.md) | Scale and runtimes |
| [04-deployment-patterns.md](articles/04-deployment-patterns.md) | CI/CD and monitoring |

## Need Help?

- Check [Troubleshooting](articles/00-quickstart.md#troubleshooting) for common issues
- Open an issue on GitHub for bugs or feature requests


