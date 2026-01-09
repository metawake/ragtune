# Advanced Configuration

This guide covers advanced flags and configuration options for RagTune. Most users don't need these — start with the [README](../README.md) and come back here when you need fine-grained control.

---

## Vector Store Configuration

### Qdrant (Default)

```bash
ragtune ingest ./docs --collection prod \
  --store qdrant \
  --qdrant-addr 127.0.0.1:6334
```

| Flag | Default | Description |
|------|---------|-------------|
| `--store` | `qdrant` | Vector store backend |
| `--qdrant-addr` | `127.0.0.1:6334` | Qdrant gRPC address |

### pgvector (PostgreSQL)

```bash
ragtune ingest ./docs --collection prod \
  --store pgvector \
  --pgvector-url "postgres://user:pass@localhost:5432/dbname"
```

| Flag | Default | Description |
|------|---------|-------------|
| `--pgvector-url` | *(required)* | PostgreSQL connection string |

**Connection string format:**
```
postgres://username:password@host:port/database?sslmode=disable
```

### Weaviate

```bash
ragtune ingest ./docs --collection prod \
  --store weaviate \
  --weaviate-host localhost:8080 \
  --weaviate-scheme http
```

| Flag | Default | Description |
|------|---------|-------------|
| `--weaviate-host` | `localhost:8080` | Weaviate server host |
| `--weaviate-scheme` | `http` | Protocol (http or https) |

### Chroma

```bash
ragtune ingest ./docs --collection prod \
  --store chroma \
  --chroma-url http://localhost:8000
```

| Flag | Default | Description |
|------|---------|-------------|
| `--chroma-url` | `http://localhost:8000` | Chroma server URL |

### Pinecone

```bash
ragtune ingest ./docs --collection prod \
  --store pinecone \
  --pinecone-host your-index-abc123.svc.pinecone.io \
  --pinecone-api-key $PINECONE_API_KEY
```

| Flag | Default | Description |
|------|---------|-------------|
| `--pinecone-host` | *(required)* | Index host from Pinecone console |
| `--pinecone-api-key` | *(env)* | API key (or use `PINECONE_API_KEY`) |

---

## Embedder Configuration

### OpenAI (Default)

```bash
export OPENAI_API_KEY="sk-..."
ragtune ingest ./docs --collection prod --embedder openai
```

| Flag | Default | Description |
|------|---------|-------------|
| `--embedder` | `openai` | Embedding backend |

Uses `text-embedding-3-small` (1536 dimensions).

### Ollama (Local, No API Key)

```bash
ragtune ingest ./docs --collection prod \
  --embedder ollama \
  --ollama-addr http://localhost:11434 \
  --ollama-model nomic-embed-text \
  --ollama-concurrency 8
```

| Flag | Default | Description |
|------|---------|-------------|
| `--ollama-addr` | `http://localhost:11434` | Ollama server URL |
| `--ollama-model` | `nomic-embed-text` | Embedding model name |
| `--ollama-concurrency` | `8` | Parallel requests (higher = faster, more CPU) |

**Popular Ollama models:**
- `nomic-embed-text` — 768 dimensions, fast, good quality
- `mxbai-embed-large` — 1024 dimensions, higher quality
- `all-minilm` — 384 dimensions, very fast

### Text Embeddings Inference (TEI)

```bash
# Start TEI server
docker run -p 8080:8080 ghcr.io/huggingface/text-embeddings-inference:cpu-1.2 \
  --model-id BAAI/bge-base-en-v1.5

# Use with RagTune
ragtune ingest ./docs --collection prod \
  --embedder tei \
  --tei-addr http://localhost:8080 \
  --tei-model BAAI/bge-base-en-v1.5
```

| Flag | Default | Description |
|------|---------|-------------|
| `--tei-addr` | `http://localhost:8080` | TEI server URL |
| `--tei-model` | `BAAI/bge-base-en-v1.5` | Model (for dimension inference) |

**Why TEI?** 4x faster than Ollama for batch embedding. Use for large corpora (10K+ docs).

### Cohere

```bash
export COHERE_API_KEY="..."
ragtune ingest ./docs --collection prod \
  --embedder cohere \
  --cohere-model embed-english-v3.0
```

| Flag | Default | Description |
|------|---------|-------------|
| `--cohere-model` | `embed-english-v3.0` | Cohere model |

**Available models:**
- `embed-english-v3.0` — English, 1024 dimensions
- `embed-multilingual-v3.0` — 100+ languages

### Voyage

```bash
export VOYAGE_API_KEY="..."
ragtune ingest ./docs --collection prod \
  --embedder voyage \
  --voyage-model voyage-2
```

| Flag | Default | Description |
|------|---------|-------------|
| `--voyage-model` | `voyage-2` | Voyage model |

**Domain-specific models:**
- `voyage-2` — General purpose
- `voyage-law-2` — Legal documents (contracts, cases)
- `voyage-code-2` — Source code

---

## Embedding Dimension

Most embedders auto-detect dimensions. Override if needed:

```bash
ragtune ingest ./docs --collection prod --embedding-dim 768
```

| Flag | Default | Description |
|------|---------|-------------|
| `--embedding-dim` | *(auto)* | Force embedding dimension |

**Common dimensions:**
- OpenAI: 1536
- Ollama (nomic): 768
- TEI (bge-base): 768
- Cohere: 1024
- Voyage: 1024

---

## Chunking Options

```bash
ragtune ingest ./docs --collection prod \
  --chunk-size 512 \
  --chunk-overlap 64
```

| Flag | Default | Description |
|------|---------|-------------|
| `--chunk-size` | `512` | Characters per chunk |
| `--chunk-overlap` | `64` | Overlap between adjacent chunks |

**Guidelines:**
- **Small chunks (256-384):** Better for precise Q&A, more chunks to search
- **Medium chunks (512-768):** Good balance for most use cases
- **Large chunks (1024+):** Better for summarization, fewer chunks

---

## CI/CD Thresholds

### simulate --ci

```bash
ragtune simulate --collection prod --queries golden.json \
  --ci \
  --min-recall 0.85 \
  --min-mrr 0.70 \
  --min-coverage 0.90 \
  --max-latency-p95 500
```

| Flag | Default | Description |
|------|---------|-------------|
| `--ci` | `false` | Enable CI mode (exit 1 on failure) |
| `--min-recall` | `0` | Minimum Recall@K |
| `--min-mrr` | `0` | Minimum MRR |
| `--min-coverage` | `0` | Minimum Coverage |
| `--max-latency-p95` | `0` | Maximum p95 latency in ms (0 = no limit) |

**Example with all thresholds:**
```bash
ragtune simulate --collection prod --queries golden.json \
  --ci --min-recall 0.85 --min-mrr 0.70 --min-coverage 0.90 --max-latency-p95 500
```

### audit (Built-in Thresholds)

```bash
ragtune audit --collection prod --queries golden.json \
  --min-recall 0.85 \
  --min-mrr 0.70 \
  --min-coverage 0.90 \
  --max-latency-p95 500
```

`audit` has sensible defaults. Override only when needed. The `--max-latency-p95` flag defaults to 0 (no limit) — set it explicitly for production checks.

---

## Environment Variables

| Variable | Purpose |
|----------|---------|
| `OPENAI_API_KEY` | OpenAI embeddings |
| `COHERE_API_KEY` | Cohere embeddings |
| `VOYAGE_API_KEY` | Voyage embeddings |
| `PINECONE_API_KEY` | Pinecone vector store |

---

## Performance Tuning

### Large Corpus Ingestion (10K+ docs)

1. **Use TEI instead of Ollama:**
   ```bash
   docker run -p 8080:8080 ghcr.io/huggingface/text-embeddings-inference:cpu-1.2 \
     --model-id BAAI/bge-base-en-v1.5
   ragtune ingest ./large-corpus --collection prod --embedder tei
   ```

2. **Increase Ollama concurrency:**
   ```bash
   ragtune ingest ./docs --collection prod \
     --embedder ollama \
     --ollama-concurrency 16
   ```

3. **Use larger chunk sizes (fewer embedding calls):**
   ```bash
   ragtune ingest ./docs --collection prod --chunk-size 1024
   ```

### GPU Acceleration

TEI with GPU is 10x faster:

```bash
docker run --gpus all -p 8080:8080 \
  ghcr.io/huggingface/text-embeddings-inference:1.2 \
  --model-id BAAI/bge-base-en-v1.5
```

---

## All Flags Reference

### Global Flags (All Commands)

| Flag | Default | Description |
|------|---------|-------------|
| `--collection` | *(required)* | Collection name |
| `--store` | `qdrant` | Vector store backend |
| `--embedder` | `openai` | Embedding backend |
| `--top-k` | `5` | Results to retrieve |

### Store-Specific Flags

| Flag | Default | For Store |
|------|---------|-----------|
| `--qdrant-addr` | `127.0.0.1:6334` | qdrant |
| `--pgvector-url` | | pgvector |
| `--weaviate-host` | `localhost:8080` | weaviate |
| `--weaviate-scheme` | `http` | weaviate |
| `--chroma-url` | `http://localhost:8000` | chroma |
| `--pinecone-host` | | pinecone |
| `--pinecone-api-key` | | pinecone |

### Embedder-Specific Flags

| Flag | Default | For Embedder |
|------|---------|--------------|
| `--ollama-addr` | `http://localhost:11434` | ollama |
| `--ollama-model` | `nomic-embed-text` | ollama |
| `--ollama-concurrency` | `8` | ollama |
| `--tei-addr` | `http://localhost:8080` | tei |
| `--tei-model` | `BAAI/bge-base-en-v1.5` | tei |
| `--cohere-model` | `embed-english-v3.0` | cohere |
| `--voyage-model` | `voyage-2` | voyage |

### Ingest Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--chunk-size` | `512` | Characters per chunk |
| `--chunk-overlap` | `64` | Overlap between chunks |
| `--embedding-dim` | *(auto)* | Force embedding dimension |

### Explain Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--save` | `false` | Save to golden queries file |
| `--golden-file` | `golden-queries.json` | Golden queries path |
| `--relevant` | *(inferred)* | Explicit relevant doc path |

**Diagnostics Output:**

`explain` provides comprehensive score distribution analysis:

- **Score statistics:** Range, mean, standard deviation
- **Quartiles:** Q1 (25th percentile), median, Q3 (75th percentile)
- **Top gap:** Distance between #1 and #2 scores — large gap indicates confident retrieval
- **Distribution type:** `tight` (< 0.05 spread), `spread` (> 0.30), `bimodal` (Q2-Q1 > Q3-Q2 * 1.5), or `normal`
- **Insights:** Positive indicators (✓ Strong top match, ✓ Good score separation)
- **Warnings:** Issues needing attention (⚠ Low similarity, ⚠ High variance)

### Simulate Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--queries` | *(required)* | Path to queries JSON |
| `--ci` | `false` | CI mode (exit 1 on failure) |
| `--min-recall` | `0` | Minimum Recall@K |
| `--min-mrr` | `0` | Minimum MRR |
| `--min-coverage` | `0` | Minimum Coverage |
| `--max-latency-p95` | `0` | Maximum p95 latency (ms) |

**Metrics Computed:**

| Metric | Description |
|--------|-------------|
| **Recall@K** | Fraction of relevant docs found in top-K results |
| **MRR** | Mean Reciprocal Rank — how high the first relevant result ranks |
| **NDCG@K** | Normalized Discounted Cumulative Gain — rewards good ranking with logarithmic discount |
| **Coverage** | Fraction of relevant docs ever retrieved across all queries |
| **Redundancy** | Average times a doc is retrieved (detects over-representation) |
| **Latency** | p50/p95/p99 percentiles for embedding + search |

**Failure Analysis:**

After computing metrics, `simulate` shows queries with Recall@K = 0 (complete retrieval failures):

```
FAILURES: 3 queries with Recall@5 = 0
  ✗ [q-42] "How do I configure SSO?"
    Expected: [sso-guide.md], Retrieved: [api-keys.md (0.721), ...]

💡 Debugging hints:
   • Run `ragtune explain "<query>"` to inspect retrieval
   • Check if expected documents are in the corpus
   • Try different chunk sizes or embedders with `ragtune compare`
```

This helps identify specific queries needing attention rather than just aggregate metrics.

### Audit Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--queries` | *(required)* | Path to queries JSON |
| `--min-recall` | `0.85` | Minimum Recall@K |
| `--min-mrr` | `0.70` | Minimum MRR |
| `--min-coverage` | `0.90` | Minimum Coverage |
| `--max-latency-p95` | `0` | Maximum p95 latency (ms) |

---

## See Also

- [README](../README.md) — Quick start and common usage
- [Benchmarking Guide](articles/03-benchmarking-guide.md) — Scale testing
- [Deployment Patterns](articles/04-deployment-patterns.md) — CI/CD integration
