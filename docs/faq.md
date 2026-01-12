# Frequently Asked Questions

## Getting Started

### Do I need API keys?

No, if you use Ollama (runs locally). Yes, for cloud embedders:

| Embedder | Environment Variable |
|----------|---------------------|
| OpenAI | `OPENAI_API_KEY` |
| Cohere | `COHERE_API_KEY` |
| Voyage | `VOYAGE_API_KEY` |

### How many test queries do I need?

| Queries | Use Case |
|---------|----------|
| 10-20 | Smoke test |
| 50-100 | CI quality gate |
| 200+ | Tuning decisions |

Start with 20, add more as you discover edge cases.

### How do I create queries.json?

Three options:

1. **Incrementally**: Use `ragtune explain "query" --save` to build up golden queries
2. **Import CSV**: `ragtune import-queries queries.csv`
3. **Manual**: Create JSON with `{"queries": [{"text": "...", "relevant_docs": [...]}]}`

Example JSON:

```json
{
  "queries": [
    {
      "id": "q1",
      "text": "How do I reset my password?",
      "relevant_docs": ["docs/auth/password.md"]
    },
    {
      "id": "q2", 
      "text": "What are the rate limits?",
      "relevant_docs": ["docs/api/limits.md"]
    }
  ]
}
```

---

## Metrics & Interpretation

### What's a good Recall@5 score?

| Score | Interpretation |
|-------|----------------|
| < 0.60 | Poor — needs investigation |
| 0.60-0.75 | Acceptable for exploration use cases |
| 0.75-0.90 | Good for most production uses |
| > 0.90 | Excellent — typical for well-tuned systems |

Context matters: legal/medical needs higher scores than casual search.

### What do the different metrics mean?

| Metric | What It Measures |
|--------|------------------|
| **Recall@K** | % of relevant docs found in top-K results |
| **MRR** | How high the first relevant result ranks (1.0 = always first) |
| **NDCG@K** | Ranking quality — rewards good ordering of all results |
| **Coverage** | % of relevant docs ever retrieved across all queries |
| **Latency** | Time to embed query + search (p50/p95/p99 percentiles) |

### Why is my recall low?

Common causes:

1. **Chunk size too small**: Context gets scattered across chunks
2. **Wrong embedder for domain**: General embedders struggle with legal/medical/code
3. **Query vocabulary mismatch**: Queries use different terms than documents
4. **Missing documents**: Expected docs aren't in the corpus

Debugging steps:

```bash
# 1. Inspect what's being retrieved
ragtune explain "your failing query" --collection prod

# 2. Try larger chunks
ragtune ingest ./docs --collection prod-1024 --chunk-size 1024

# 3. Try domain-specific embedder
ragtune compare --embedders ollama,voyage --docs ./docs --queries queries.json
```

---

## Embedders

### Which embedder should I use?

| Embedder | Best For |
|----------|----------|
| `ollama` | Development, privacy, no API costs |
| `openai` | General purpose, good baseline |
| `voyage` | Domain-specific (legal, code) |
| `cohere` | Multilingual |
| `tei` | High throughput, self-hosted |

### How do I use a domain-specific Voyage model?

```bash
export VOYAGE_API_KEY="your-key"
ragtune ingest ./docs --collection legal --embedder voyage --voyage-model voyage-law-2
```

Available models:
- `voyage-law-2` — Legal documents
- `voyage-code-2` — Source code
- `voyage-2` — General purpose

---

## Vector Stores

### Can I use my existing vector database?

Yes! RagTune supports:

| Store | Flag |
|-------|------|
| Qdrant (default) | `--store qdrant` |
| pgvector (PostgreSQL) | `--store pgvector --pgvector-url postgres://...` |
| Weaviate | `--store weaviate --weaviate-host localhost:8080` |
| Chroma | `--store chroma --chroma-url http://localhost:8000` |
| Pinecone | `--store pinecone --pinecone-host HOST --pinecone-api-key KEY` |

### How do I switch from Qdrant to pgvector?

```bash
# 1. Set up PostgreSQL with pgvector
docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=password pgvector/pgvector:pg16

# 2. Re-ingest with pgvector
ragtune ingest ./docs --collection demo --embedder ollama \
  --store pgvector --pgvector-url "postgres://postgres:password@localhost:5432/postgres"
```

---

## CI/CD Integration

### How do I add RagTune to my CI pipeline?

Use `simulate --ci` with thresholds:

```bash
ragtune simulate --collection prod --queries golden-queries.json \
  --ci --min-recall 0.85 --min-coverage 0.90 --max-latency-p95 500
```

Exit code 1 if any threshold fails.

See [examples/github-actions.yml](../examples/github-actions.yml) for a complete GitHub Actions example.

### What thresholds should I set?

| Use Case | Recall | Coverage | Latency p95 |
|----------|--------|----------|-------------|
| Development | 0.70 | 0.80 | 1000ms |
| Staging | 0.80 | 0.85 | 500ms |
| Production | 0.85 | 0.90 | 200ms |
| Critical (legal/medical) | 0.90 | 0.95 | 100ms |

---

## Performance

### Ingestion is slow. How can I speed it up?

1. **Use TEI** instead of Ollama (4x faster):
   ```bash
   docker run -p 8080:8080 ghcr.io/huggingface/text-embeddings-inference:cpu-1.2 \
     --model-id BAAI/bge-base-en-v1.5
   ragtune ingest ./docs --collection test --embedder tei
   ```

2. **Increase chunk size** (fewer chunks = fewer embedding calls):
   ```bash
   ragtune ingest ./docs --collection test --chunk-size 1024
   ```

3. **Use GPU** if available (10x faster for TEI)

### How long will ingestion take?

| Corpus | TEI CPU | TEI GPU | Ollama |
|--------|---------|---------|--------|
| 1K docs | 4 min | 30 sec | 15 min |
| 10K docs | 40 min | 5 min | 2.5 hours |
| 50K docs | 3 hours | 20 min | 12+ hours |

See [Benchmarking Guide](articles/03-benchmarking-guide.md) for detailed runtimes.
