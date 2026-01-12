# CLI Reference

Complete reference for all RagTune commands and flags.

## Commands Overview

| Command | Purpose |
|---------|---------|
| `ingest` | Load documents into vector store |
| `explain` | Debug retrieval for a single query with score distribution analysis |
| `simulate` | Batch benchmark with metrics (Recall, MRR, NDCG, Coverage) + failure analysis |
| `compare` | Compare embedders or configs |
| `report` | Generate markdown reports |
| `import-queries` | Import queries from CSV or JSON |
| `audit` | Quick health check with pass/fail |

---

## ingest

Splits documents into chunks, generates embeddings, stores in vector DB.

```bash
ragtune ingest ./docs --collection prod --chunk-size 512 --embedder ollama
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--collection` | *required* | Collection name |
| `--embedder` | `openai` | Embedding backend |
| `--chunk-size` | `512` | Characters per chunk |
| `--chunk-overlap` | `64` | Overlap between chunks |
| `--store` | `qdrant` | Vector store backend |

### Example Output

```
Reading documents from ./docs...
Found 42 documents
Created 187 chunks (avg 489 chars)
Using embedding dimension: 768 (auto-detected from ollama)
✓ Ingested 187 chunks into collection 'prod'
```

---

## explain

Shows exactly what chunks are retrieved for one query. Use `--save` to build your test suite incrementally.

```bash
ragtune explain "How do I reset my password?" --collection prod --save
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--collection` | *required* | Collection name |
| `--embedder` | `openai` | Embedding backend |
| `--top-k` | `5` | Results to retrieve |
| `--save` | `false` | Save query to golden queries file |
| `--golden-file` | `golden-queries.json` | Path to golden queries file |
| `--relevant` | *(inferred)* | Explicit relevant doc path |

### Diagnostics Output

- Score statistics (range, mean, std dev)
- Quartiles (Q1, median, Q3) and distribution shape
- Top gap analysis (distance between #1 and #2)
- Automatic insights and warnings

### Interpreting Results

| Signal | Meaning | Action |
|--------|---------|--------|
| Score > 0.85 | Strong match | Good retrieval |
| Score 0.60-0.85 | Moderate match | May need tuning |
| Score < 0.60 | Weak match | Check chunk size, embedder |
| Right doc missing | Retrieval failure | Increase chunk size or try different embedder |
| All scores similar | No clear winner | Query may be too vague |

---

## simulate

Runs many queries, computes aggregate metrics. Use `--ci` for automated quality gates.

```bash
ragtune simulate --collection prod --queries golden-queries.json
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--collection` | *required* | Collection name |
| `--queries` | *required* | Path to queries JSON file |
| `--embedder` | `openai` | Embedding backend |
| `--top-k` | `5` | Results to retrieve |

### CI Mode Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--ci` | `false` | Enable CI mode (exit 1 if thresholds fail) |
| `--min-recall` | `0` | Minimum Recall@K threshold |
| `--min-mrr` | `0` | Minimum MRR threshold |
| `--min-coverage` | `0` | Minimum Coverage threshold |
| `--max-latency-p95` | `0` | Maximum p95 latency in ms (0 = no limit) |

### Example

```bash
ragtune simulate --collection prod --queries golden.json \
  --ci --min-recall 0.85 --min-coverage 0.90
```

Exit code 1 if thresholds not met.

---

## compare

Compares collections (different chunk sizes) or embedders.

```bash
# Compare chunk sizes
ragtune compare --collections prod-256,prod-512,prod-1024 --queries queries.json

# Compare embedders (auto-ingests)
ragtune compare --embedders ollama,openai --docs ./docs --queries queries.json
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--collections` | | Comma-separated collection names |
| `--embedders` | | Comma-separated embedder names |
| `--docs` | | Path to documents (required with `--embedders`) |
| `--queries` | *required* | Path to queries JSON file |

---

## audit

Pass/fail health report with recommendations. Great for daily checks or exec summaries.

```bash
ragtune audit --collection prod --queries golden-queries.json
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--collection` | *required* | Collection name |
| `--queries` | *required* | Path to queries JSON file |
| `--min-recall` | `0.85` | Minimum Recall@K threshold |
| `--min-mrr` | `0.70` | Minimum MRR threshold |
| `--min-coverage` | `0.90` | Minimum Coverage threshold |
| `--max-latency-p95` | `0` | Maximum p95 latency in ms (0 = no limit) |

Returns exit code 0 (pass) or 1 (fail).

---

## report

Creates Markdown or JSON report from a simulation run.

```bash
ragtune report --input runs/latest.json --format markdown > report.md
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--input` | *required* | Path to simulation run JSON |
| `--format` | `markdown` | Output format (`markdown` or `json`) |

---

## import-queries

Import queries from CSV or JSON files.

```bash
ragtune import-queries queries.csv --output golden-queries.json
```

### CSV Format

```csv
query,relevant_docs
"How do I reset password?",docs/auth/password.md
"What are rate limits?",docs/api/limits.md;docs/api/quotas.md
```

Header required. Use semicolon for multiple relevant docs.

---

## Common Flags

These flags work across most commands:

| Flag | Default | Description |
|------|---------|-------------|
| `--collection` | *required* | Collection name |
| `--embedder` | `openai` | Embedding backend (`ollama`, `openai`, `tei`, `cohere`, `voyage`) |
| `--top-k` | `5` | Results to retrieve |
| `--store` | `qdrant` | Vector store backend |

## Vector Store Flags

| Flag | Description |
|------|-------------|
| `--store qdrant` | Use Qdrant (default) |
| `--store pgvector --pgvector-url URL` | Use PostgreSQL with pgvector |
| `--store weaviate --weaviate-host HOST` | Use Weaviate |
| `--store chroma --chroma-url URL` | Use ChromaDB |
| `--store pinecone --pinecone-host HOST --pinecone-api-key KEY` | Use Pinecone |

## Embedder Flags

| Embedder | Required Flags / Environment |
|----------|------------------------------|
| `ollama` | Ollama must be running locally |
| `openai` | `OPENAI_API_KEY` environment variable |
| `tei` | `--tei-addr http://localhost:8080` |
| `cohere` | `COHERE_API_KEY` environment variable |
| `voyage` | `VOYAGE_API_KEY` environment variable, optional `--voyage-model` |
