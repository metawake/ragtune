# RagTune

[![Go Version](https://img.shields.io/badge/go-1.22+-00ADD8?logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)
[![Release](https://img.shields.io/github/v/release/metawake/ragtune?include_prereleases)](https://github.com/metawake/ragtune/releases)

**EXPLAIN ANALYZE for RAG retrieval** — inspect, explain, benchmark, and tune your RAG retrieval layer.

<p align="center">
  <img src="assets/demo.gif" alt="RagTune demo" width="700">
</p>

<p align="center">
  <a href="#30-second-quickstart">Quickstart</a> •
  <a href="#key-concepts-60-second-primer">Concepts</a> •
  <a href="#command-guide">Commands</a> •
  <a href="#building-your-test-suite">Tutorial</a> •
  <a href="#faq">FAQ</a> •
  <a href="#troubleshooting">Troubleshooting</a>
</p>

---

| I want to... | Go to |
|--------------|-------|
| **Try it in 30 seconds** | [Quickstart](#30-second-quickstart) |
| **I don't have test queries yet** | [Bootstrap Your Test Suite](#i-dont-have-test-queries-yet) |
| **Debug a single query** | [explain command](#explain--debug-a-single-query) |
| **Run batch evaluation** | [simulate command](#simulate--batch-evaluation-with-metrics) |
| **Set up CI/CD quality gates** | [CI/CD Quality Gates](#cicd-quality-gates) |
| **Compare embedders or chunk sizes** | [compare command](#compare--side-by-side-comparison) |
| **Interpret my results** | [What Good Looks Like](#what-good-metrics-look-like) |
| **See sample output** | [examples/sample-report.md](examples/sample-report.md) |

---

## 30-Second Quickstart

```bash
# 1. Start vector store
docker run -d -p 6333:6333 -p 6334:6334 qdrant/qdrant

# 2. Ingest your documents
ragtune ingest ./docs --collection my-docs --embedder ollama

# 3. Explain what retrieval finds
ragtune explain "How do I reset my password?" --collection my-docs
```

That's it. No API keys needed with Ollama (runs locally).

### I Don't Have Test Queries Yet

No problem. Build your test suite incrementally using `explain --save`:

```bash
# 1. Ingest your docs
ragtune ingest ./docs --collection my-docs --embedder ollama

# 2. Try queries and save the good ones
ragtune explain "How do I reset my password?" --collection my-docs --save
ragtune explain "What are the rate limits?" --collection my-docs --save
ragtune explain "How to configure webhooks" --collection my-docs --save

# 3. Once you have 20+ queries, run evaluation
ragtune simulate --collection my-docs --queries golden-queries.json
```

Each `--save` adds the query to `golden-queries.json`. Start with 10-20 queries that represent real user questions.

## What You'll See

### explain — Single Query Diagnostics

```
Query: "How do I reset my password?"
Searching collection 'my-docs' (top-k=5)...

================================================================================
EXPLAIN RETRIEVAL: 5 results
================================================================================

[1] Score: 0.8934 | Source: docs/auth/password-reset.md
    Text: To reset your password: 1. Click "Forgot Password" 2. Enter your email...

[2] Score: 0.8521 | Source: docs/auth/account-security.md
    Text: Account Security ## Password Management. We recommend rotating...

--------------------------------------------------------------------------------
DIAGNOSTICS
--------------------------------------------------------------------------------
Score range:  0.7234 - 0.8934 (spread: 0.1700)
Mean score:   0.8102
Std dev:      0.0523
Quartiles:    Q1=0.7521  Median=0.8102  Q3=0.8612
Top gap:      0.0413 (distance between #1 and #2)
Distribution: normal

Insights:
  ✓ Strong top match (>0.85): likely high-quality retrieval
  ✓ Good score separation: retrieval is discriminating effectively

✓ No warnings
```

### simulate — Batch Metrics with Failure Analysis

```
Running 50 queries against collection 'prod'...

  Metrics:
    Recall@5:   0.82
    MRR:        0.76
    NDCG@5:     0.79
    Coverage:   0.94
    Redundancy: 1.20
    Latency:    p50=45ms  p95=120ms  p99=180ms  avg=52ms

  ─────────────────────────────────────────────────────────────
  FAILURES: 3 queries with Recall@5 = 0
  ─────────────────────────────────────────────────────────────

  ✗ [q-42] How do I configure SSO?
    Expected:  [sso-guide.md]
    Retrieved: [api-keys.md (0.721), auth-overview.md (0.698), ...]

  💡 Debugging hints:
     • Run `ragtune explain "<query>"` to inspect retrieval
     • Check if expected documents are in the corpus
     • Try different chunk sizes or embedders with `ragtune compare`

✓ Run saved to runs/2025-01-09T10-30-00Z.json
```

## Key Concepts (60-second primer)

New to RAG? Here's what's happening under the hood:

```mermaid
flowchart LR
    A[📄 Documents] --> B[Chunker]
    B --> C[Embedder]
    C --> D[(Vector DB)]
    
    E[❓ Query] --> F[Embedder]
    F --> G{Similarity\nSearch}
    D --> G
    G --> H[🎯 Top-K Results]
```

| Term | What It Means |
|------|---------------|
| **Embedding** | Text converted to a vector (list of numbers representing meaning). Similar texts have similar vectors. |
| **Chunk** | A piece of a document. Long docs are split into chunks before embedding. |
| **Similarity score** | 0-1 scale. Higher = query and chunk have more similar meaning. |
| **Top-K** | Return the K most similar chunks to your query. |

### Why Chunk Size Matters

```
Document: "To reset your password, go to Settings, click Security, 
           then click Reset Password. You'll receive an email..."

Chunk size 100:  ["To reset your password, go to", 
                  "Settings, click Security, then",
                  "click Reset Password. You'll..."]
                  → Query matches piece 1, but context is scattered

Chunk size 500:  ["To reset your password, go to Settings, 
                   click Security, then click Reset Password. 
                   You'll receive an email..."]
                  → Full context preserved in one chunk
```

**Too small** = context gets split, retrieval misses connections  
**Too large** = noise drowns out relevant parts

### Score Interpretation

| Score Range | Meaning | Action |
|-------------|---------|--------|
| 0.90+ | Excellent match | Exactly what user asked |
| 0.75-0.90 | Good match | Relevant, may need refinement |
| 0.60-0.75 | Weak match | Might be relevant, or noise |
| < 0.60 | Poor match | Likely not what user wants |

## Common Usage

### Ingest Documents

```bash
ragtune ingest ./your-docs --collection prod
```

Options you might want:
- `--chunk-size 512` — Characters per chunk (default: 512)
- `--embedder ollama` — Use local Ollama instead of OpenAI

### Explain a Query

```bash
ragtune explain "your question" --collection prod --top-k 10
```

See exactly what chunks are retrieved and their similarity scores.

### Run a Benchmark

```bash
ragtune simulate --collection prod --queries queries.json
```

Computes recall@k, MRR, NDCG, coverage, and latency percentiles (p50/p95/p99) across your test queries. Also shows a failure analysis for queries with Recall@K = 0.

### Compare Embedders

```bash
ragtune compare --embedders ollama,openai --docs ./docs --queries queries.json
```

Side-by-side comparison of different embedding models.

### Which Command Should I Use?

```mermaid
flowchart TD
    A[What do you want to do?] --> B{Have documents\nto load?}
    B -->|Yes| C[ragtune ingest]
    B -->|No| D{Have a collection?}
    D -->|No| E[Start with sample data]
    D -->|Yes| F{One query or many?}
    F -->|One query| G[ragtune explain]
    F -->|Many queries| H{Quick check\nor full metrics?}
    H -->|Quick| I[ragtune audit]
    H -->|Full| J[ragtune simulate]
    
    G -.->|Save for regression| K[add --save flag]
    J -.->|CI/CD| L[add --ci flag]
```

## Command Guide

### ingest — Load documents into vector store

Splits documents into chunks, generates embeddings, stores in vector DB.

```bash
ragtune ingest ./docs --collection prod --chunk-size 512 --embedder ollama
```

```
Reading documents from ./docs...
Found 42 documents
Created 187 chunks (avg 489 chars)
Using embedding dimension: 768 (auto-detected from ollama)
✓ Ingested 187 chunks into collection 'prod'
```

**What to check:** If chunk count is very high, consider larger `--chunk-size`.

---

### explain — Debug a single query

Shows exactly what chunks are retrieved for one query. Use `--save` to build your test suite incrementally.

```bash
ragtune explain "How do I reset my password?" --collection prod --save
```

**Interpreting results:**

| Signal | Meaning | Action |
|--------|---------|--------|
| Score > 0.85 | Strong match | Good retrieval |
| Score 0.60-0.85 | Moderate match | May need tuning |
| Score < 0.60 | Weak match | Check chunk size, embedder |
| Right doc missing | Retrieval failure | Increase chunk size or try different embedder |
| All scores similar | No clear winner | Query may be too vague |

---

### simulate — Batch evaluation with metrics

Runs many queries, computes aggregate metrics. Use `--ci` for automated quality gates.

```bash
ragtune simulate --collection prod --queries golden-queries.json
```

```
Running 50 queries against collection 'prod'...

METRICS:
  Recall@5:    0.82
  MRR:         0.76
  NDCG@5:      0.79
  Coverage:    0.94

LATENCY:
  p50: 45ms | p95: 120ms | p99: 180ms

FAILURES: 3 queries with Recall@5 = 0
  ✗ [q-42] "How do I configure SSO?"
    Expected: [sso-guide.md], Retrieved: [api-keys.md...]

💡 Run `ragtune explain "<query>"` for details

Run saved to: runs/2025-01-09T10-30-00Z.json
```

**CI mode:** Exit code 1 if thresholds not met:
```bash
ragtune simulate --collection prod --queries golden.json \
  --ci --min-recall 0.85 --min-coverage 0.90
```

---

### compare — Side-by-side comparison

Compares collections (different chunk sizes) or embedders.

```bash
# Compare chunk sizes
ragtune compare --collections prod-256,prod-512,prod-1024 --queries queries.json

# Compare embedders (auto-ingests)
ragtune compare --embedders ollama,openai --docs ./docs --queries queries.json
```

```
| Collection   | Recall@5 | MRR   | Coverage |
|--------------|----------|-------|----------|
| prod-256     | 0.68     | 0.62  | 0.88     |
| prod-512     | 0.78     | 0.71  | 0.92     |  ← Winner
| prod-1024    | 0.75     | 0.69  | 0.90     |
```

---

### audit — Quick health check

Pass/fail health report with recommendations. Great for daily checks or exec summaries.

```bash
ragtune audit --collection prod --queries golden-queries.json
```

Returns exit code 0 (pass) or 1 (fail). See [example output](#step-6-quick-health-check).

---

### report — Generate shareable report

Creates Markdown or JSON report from a simulation run.

```bash
ragtune report --input runs/latest.json --format markdown > report.md
```

---

### import-queries — Bulk import test queries

Import queries from CSV or JSON files.

```bash
ragtune import-queries queries.csv --output golden-queries.json
```

CSV format: `query,relevant_docs` (header required, semicolon for multiple docs)

## Building Your Test Suite

RagTune works best when you evaluate on **your own data**. Here's the journey:

### Step 1: Explore with Sample Data

Start with bundled benchmarks to understand how RAG tuning works:

```bash
ragtune ingest ./benchmarks/hotpotqa-1k/corpus --collection demo --embedder ollama
ragtune simulate --collection demo --queries ./benchmarks/hotpotqa-1k/queries.json
```

This shows you what good metrics look like and the full workflow before touching production data.

### Step 2: Debug Your Data with `explain`

Point RagTune at your actual documents:

```bash
ragtune ingest ./your-docs --collection prod --embedder ollama
ragtune explain "your real question here" --collection prod
```

### Step 3: Save Queries That Matter

When you find queries that should work (or fail), save them:

```bash
ragtune explain "How do I reset my password?" --collection prod --save
```

This builds `golden-queries.json` incrementally — like bookmarking requests in Postman.

### Step 4: Import Existing Test Cases (Optional)

If you have test cases in a spreadsheet:

```csv
query,relevant_docs
"How do I reset password?",docs/auth/password.md
"What are rate limits?",docs/api/limits.md
```

```bash
ragtune import-queries queries.csv
```

### Step 5: Run Evaluation

Once you have 20-50 queries, you get meaningful metrics:

```bash
ragtune simulate --collection prod --queries golden-queries.json
```

| Queries | Confidence Level |
|---------|-----------------|
| 10-20   | Smoke test      |
| 50-100  | CI gate         |
| 200+    | Tuning decisions |

Metrics include:
- **Recall@K**: % of relevant docs in top-K results
- **MRR**: How high the first relevant result ranks
- **NDCG@K**: Ranking quality (rewards good ordering of all results)
- **Coverage**: % of relevant docs ever retrieved
- **Latency**: p50/p95/p99 percentiles (embedding + search)

**Failure Analysis**: After computing metrics, shows queries with Recall@K = 0 (complete failures) with debugging hints.

#### What Good Metrics Look Like

| Metric | Poor | Acceptable | Good | Excellent |
|--------|------|------------|------|-----------|
| **Recall@5** | < 0.60 | 0.60–0.75 | 0.75–0.90 | > 0.90 |
| **MRR** | < 0.50 | 0.50–0.70 | 0.70–0.85 | > 0.85 |
| **NDCG@5** | < 0.50 | 0.50–0.70 | 0.70–0.85 | > 0.85 |
| **Coverage** | < 0.70 | 0.70–0.85 | 0.85–0.95 | > 0.95 |
| **Latency p95** | > 500ms | 200–500ms | 50–200ms | < 50ms |

> 💡 **Production reference:** A healthcare RAG system achieved 91.4% accuracy with proper tuning, 
> demonstrating that "Excellent" scores are achievable in production 
> ([arxiv.org/abs/2402.01733](https://arxiv.org/abs/2402.01733)).

**Context matters:**
- Legal/medical domains: aim for Excellent (accuracy critical)
- Customer support: Good is often sufficient
- Creative/exploration: Acceptable may be fine

**If metrics are low:**
1. Try larger chunks (`--chunk-size 1024`)
2. Try different embedders (`--embedder voyage` for legal)
3. Check if queries match document vocabulary
4. Add more relevant docs to your corpus

**Reality check:** Sometimes tuning doesn't produce large differences — and that's useful to know. It means you can stop optimizing and focus elsewhere.

### Step 6: Quick Health Check

Get a pass/fail health report:

```bash
ragtune audit --collection prod --queries golden-queries.json
```

Output:
```
╔════════════════════════════════════════════════════════════════╗
║  RAG HEALTH REPORT: prod                                       ║
╠════════════════════════════════════════════════════════════════╣
║  Queries:     47   ⚠ low, recommend 50+                        ║
╠════════════════════════════════════════════════════════════════╣
║  Recall@K:    0.91  ✓ PASS  (threshold: 0.85)                  ║
║  MRR:         0.82  ✓ PASS  (threshold: 0.70)                  ║
║  Coverage:    0.97  ✓ PASS  (threshold: 0.90)                  ║
╠════════════════════════════════════════════════════════════════╣
║  STATUS: ✅ HEALTHY                                            ║
╚════════════════════════════════════════════════════════════════╝
```

### Step 7: Add to CI/CD

For automation, use `simulate --ci` or `audit`:

```bash
ragtune simulate --collection prod --queries golden-queries.json \
  --ci --min-recall 0.85 --min-coverage 0.90 --max-latency-p95 500
```

Exit code 1 if thresholds not met.

## CI/CD Quality Gates

Add RAG quality checks to your CI pipeline. Here's a complete GitHub Actions example:

```yaml
name: RAG Quality Gate

on:
  push:
    paths: ['docs/**', 'tests/golden-queries.json']

jobs:
  rag-quality:
    runs-on: ubuntu-latest
    services:
      qdrant:
        image: qdrant/qdrant
        ports: ['6333:6333', '6334:6334']
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with: { go-version: '1.22' }
      
      - name: Install RagTune
        run: go install github.com/metawake/ragtune/cmd/ragtune@latest
      
      - name: Setup Ollama
        run: |
          curl -fsSL https://ollama.com/install.sh | sh
          ollama serve &
          sleep 5
          ollama pull nomic-embed-text
      
      - name: Ingest and test
        run: |
          ragtune ingest ./docs --collection ci-test --embedder ollama
          ragtune simulate --collection ci-test --queries tests/golden-queries.json \
            --embedder ollama --ci \
            --min-recall 0.85 --min-coverage 0.90 --max-latency-p95 500
```

**Key thresholds:**
- `--min-recall 0.85` — At least 85% of relevant docs in top-5
- `--min-coverage 0.90` — At least 90% of relevant docs ever retrieved  
- `--max-latency-p95 500` — 95th percentile latency under 500ms

See [examples/github-actions.yml](examples/github-actions.yml) for more options (OpenAI, different stores).

## Why RagTune?

Most teams iterate blindly on RAG retrieval. RagTune provides the diagnostics to make informed decisions.

> "By the time one method was implemented, it was often already outdated"  
> — [Industry RAG Interview Study, SCITEPRESS 2025](https://www.scitepress.org/Papers/2025/137395/137395.pdf)

> CloudFactory identified that deploying RAG "without systematic evaluation 
> can result in undetected errors" as a [top failure mode](https://www.cloudfactory.com/blog/rag-is-breaking).

### What the Data Shows

| What Matters | Impact | Evidence |
|--------------|--------|----------|
| Domain-appropriate chunking | High (7%+ recall difference) | CMU study: F1 improved 5.45% → 42.21% with hybrid search |
| Embedding model choice | Medium (5% difference) | Healthcare RAG: accuracy 80.1% → 91.4% with domain tuning |
| Retrieval evaluation | Critical | InfoObjects: 78% reduction in manual effort with proper metrics |
| Continuous monitoring | High | "Data drift" cited as failure mode in 14/18 production studies |

*See [Lessons from Production RAG](docs/articles/05-lessons-from-production.md) for detailed case study analysis.*

**Use cases:**
- Tune chunk sizes and top-k before launch
- Regression testing when corpus changes
- CI/CD quality gates
- Compare embedding models

## Embedders

| Embedder | Setup | Best For |
|----------|-------|----------|
| `ollama` | Local, no API key | Development, privacy-sensitive |
| `openai` | `OPENAI_API_KEY` | General purpose baseline |
| `tei` | Docker container | Enterprise scale (4x faster) |
| `cohere` | `COHERE_API_KEY` | Multilingual |
| `voyage` | `VOYAGE_API_KEY` | Domain-tuned (legal, code) |

Default is `openai`. Use `--embedder ollama` for zero-config local setup.

## Documentation

| Guide | When to Read |
|-------|--------------|
| [Quickstart](docs/articles/00-quickstart.md) | First-time setup with explanations |
| [Advanced Configuration](docs/advanced-configuration.md) | Vector stores, embedders, performance tuning |
| [Benchmarking Guide](docs/articles/03-benchmarking-guide.md) | Scale testing, expected runtimes |
| [Deployment Patterns](docs/articles/04-deployment-patterns.md) | CI/CD, monitoring, production |
| [Sample Report](examples/sample-report.md) | See example output before running |

## Installation

### Option A: Download Binary (Recommended)

Download the latest release for your platform:

```bash
# macOS (Apple Silicon)
curl -L https://github.com/metawake/ragtune/releases/latest/download/ragtune_darwin_arm64.tar.gz | tar xz
sudo mv ragtune /usr/local/bin/

# macOS (Intel)
curl -L https://github.com/metawake/ragtune/releases/latest/download/ragtune_darwin_amd64.tar.gz | tar xz
sudo mv ragtune /usr/local/bin/

# Linux
curl -L https://github.com/metawake/ragtune/releases/latest/download/ragtune_linux_amd64.tar.gz | tar xz
sudo mv ragtune /usr/local/bin/
```

Or download from [GitHub Releases](https://github.com/metawake/ragtune/releases).

### Option B: Go Install

```bash
go install github.com/metawake/ragtune/cmd/ragtune@latest
```

### Option C: Build from Source

```bash
git clone https://github.com/metawake/ragtune.git
cd ragtune && go build -o ragtune ./cmd/ragtune
```

### Prerequisites

- Go 1.22+
- Docker (for Qdrant)
- Ollama (for local embeddings) or API key for cloud embedders

## CLI Reference

### Commands

| Command | Purpose |
|---------|---------|
| `ingest` | Load documents into vector store |
| `explain` | Debug retrieval for a single query with score distribution analysis |
| `simulate` | Batch benchmark with metrics (Recall, MRR, NDCG, Coverage) + failure analysis |
| `compare` | Compare embedders or configs |
| `report` | Generate markdown reports |
| `import-queries` | Import queries from CSV or JSON |
| `audit` | Quick health check with pass/fail |

### Metrics Computed

| Metric | Description |
|--------|-------------|
| **Recall@K** | Fraction of relevant docs found in top-K results |
| **MRR** | Mean Reciprocal Rank — how high the first relevant result ranks |
| **NDCG@K** | Normalized Discounted Cumulative Gain — rewards good ranking of ALL results |
| **Coverage** | Fraction of relevant docs ever retrieved across all queries |
| **Redundancy** | Average times a doc is retrieved (detects over-representation) |
| **Diversity@K** | Fraction of unique docs in top-K (detects wasted slots from duplicates) |

### Common Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--collection` | *required* | Collection name |
| `--embedder` | `openai` | Embedding backend |
| `--top-k` | `5` | Results to retrieve |

### Ingest Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--chunk-size` | `512` | Characters per chunk |
| `--chunk-overlap` | `64` | Overlap between chunks |

### CI Mode Flags (simulate)

| Flag | Default | Description |
|------|---------|-------------|
| `--ci` | `false` | Enable CI mode (exit 1 if thresholds fail) |
| `--min-recall` | `0` | Minimum Recall@K threshold |
| `--min-mrr` | `0` | Minimum MRR threshold |
| `--min-coverage` | `0` | Minimum Coverage threshold |
| `--max-latency-p95` | `0` | Maximum p95 latency in ms (0 = no limit) |

### Explain Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--save` | `false` | Save query to golden queries file |
| `--golden-file` | `golden-queries.json` | Path to golden queries file |
| `--relevant` | *(inferred)* | Explicit relevant doc path |

**Diagnostics Output:**
- Score statistics (range, mean, std dev)
- Quartiles (Q1, median, Q3) and distribution shape
- Top gap analysis (distance between #1 and #2)
- Automatic insights and warnings

### Audit Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--min-recall` | `0.85` | Minimum Recall@K threshold |
| `--min-mrr` | `0.70` | Minimum MRR threshold |
| `--min-coverage` | `0.90` | Minimum Coverage threshold |
| `--max-latency-p95` | `0` | Maximum p95 latency in ms (0 = no limit) |

For vector store configuration, embedder options, and performance tuning, see **[Advanced Configuration](docs/advanced-configuration.md)**.

## Included Benchmarks

| Dataset | Documents | Purpose |
|---------|-----------|---------|
| `data/` | 9 | Quick testing |
| `benchmarks/hotpotqa-1k/` | 398 | General knowledge |
| `benchmarks/casehold-500/` | 500 | Legal domain |
| `benchmarks/synthetic-50k/` | 50,000 | Scale testing |

### Experimenting with Sample Data

Before using your own data, learn RagTune with these experiments:

#### Quick Sanity Check (2 min)

```bash
ragtune ingest ./data/docs --collection quick-test --embedder ollama
ragtune simulate --collection quick-test --queries ./data/queries.json
```

Expected: Recall@5 > 0.80. If lower, check that Ollama is running.

#### Chunk Size Experiment (10 min)

See how chunk size affects retrieval quality:

```bash
# Ingest with three different chunk sizes
ragtune ingest ./benchmarks/hotpotqa-1k/corpus --collection hotpot-256 \
  --chunk-size 256 --embedder ollama

ragtune ingest ./benchmarks/hotpotqa-1k/corpus --collection hotpot-512 \
  --chunk-size 512 --embedder ollama

ragtune ingest ./benchmarks/hotpotqa-1k/corpus --collection hotpot-1024 \
  --chunk-size 1024 --embedder ollama

# Compare results
ragtune compare --collections hotpot-256,hotpot-512,hotpot-1024 \
  --queries ./benchmarks/hotpotqa-1k/queries.json
```

**What you'll learn:** Chunk size typically affects recall by 5-10%.

#### Domain-Specific Test: Legal (15 min)

Legal text has different vocabulary patterns:

```bash
ragtune ingest ./benchmarks/casehold-500/corpus --collection legal --embedder ollama
ragtune simulate --collection legal --queries ./benchmarks/casehold-500/queries.json
```

**What you'll learn:** General embedders struggle with legal text. 
Try `--embedder voyage` with `--voyage-model voyage-law-2` for better results.

#### Expected Baselines

Use these as reference points (your results may vary ±5%):

| Benchmark | Embedder | Chunk Size | Recall@5 | MRR |
|-----------|----------|------------|----------|-----|
| hotpotqa-1k | ollama | 512 | ~0.72 | ~0.65 |
| hotpotqa-1k | openai | 512 | ~0.78 | ~0.71 |
| casehold-500 | ollama | 512 | ~0.58 | ~0.52 |
| casehold-500 | voyage-law-2 | 512 | ~0.71 | ~0.66 |

## Expected Runtimes

| Corpus | Ingestion (TEI) | Query Benchmark |
|--------|-----------------|-----------------|
| 1K docs | 4 min | 5 sec |
| 10K docs | 40 min | 15 sec |
| 50K docs | 3 hours | 25 sec |

With GPU: divide ingestion time by 5-10x.

## Vector Stores

| Store | Setup | Flag |
|-------|-------|------|
| Qdrant | `docker run -p 6333:6333 -p 6334:6334 qdrant/qdrant` | `--store qdrant` (default) |
| pgvector | PostgreSQL + pgvector extension | `--store pgvector --pgvector-url postgres://...` |
| Weaviate | `docker run -p 8080:8080 semitechnologies/weaviate` | `--store weaviate --weaviate-host localhost:8080` |
| Chroma | `docker run -p 8000:8000 chromadb/chroma` | `--store chroma --chroma-url http://localhost:8000` |
| Pinecone | Cloud (requires API key) | `--store pinecone --pinecone-host HOST --pinecone-api-key KEY` |

### Using pgvector

```bash
# 1. Set up PostgreSQL with pgvector (Docker example)
docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=password pgvector/pgvector:pg16

# 2. Use with RagTune
ragtune ingest ./docs --collection demo --embedder ollama \
  --store pgvector --pgvector-url "postgres://postgres:password@localhost:5432/postgres"
```

pgvector is ideal if you already use PostgreSQL and want vector search without a separate database.

## Architecture

RagTune is **vector-store agnostic**. Supported: Qdrant, pgvector, Weaviate, Chroma, Pinecone.

```
ragtune/
├── cmd/ragtune/        # CLI
├── internal/
│   ├── cli/            # Commands
│   ├── vectorstore/    # Store adapters
│   ├── embedder/       # Embedding backends
│   ├── chunker/        # Text splitting
│   └── metrics/        # Recall, MRR, etc.
├── benchmarks/         # Test datasets
└── docs/               # Guides
```

## FAQ

<details>
<summary><strong>Do I need API keys?</strong></summary>

No, if you use Ollama (runs locally). Yes, for cloud embedders:
- OpenAI: `OPENAI_API_KEY`
- Cohere: `COHERE_API_KEY`
- Voyage: `VOYAGE_API_KEY`
</details>

<details>
<summary><strong>How many test queries do I need?</strong></summary>

| Queries | Use Case |
|---------|----------|
| 10-20 | Smoke test |
| 50-100 | CI quality gate |
| 200+ | Tuning decisions |

Start with 20, add more as you discover edge cases.
</details>

<details>
<summary><strong>What's a good Recall@5 score?</strong></summary>

- **< 0.60**: Poor — needs investigation
- **0.60-0.75**: Acceptable for exploration use cases
- **0.75-0.90**: Good for most production uses
- **> 0.90**: Excellent — typical for well-tuned systems

Context matters: legal/medical needs higher scores than casual search.
</details>

<details>
<summary><strong>Can I use my existing vector database?</strong></summary>

Yes! RagTune supports:
- Qdrant (default)
- pgvector (PostgreSQL)
- Weaviate
- Chroma
- Pinecone

Use `--store <name>` to switch. See [Vector Stores](#vector-stores).
</details>

<details>
<summary><strong>How do I create queries.json?</strong></summary>

Three options:
1. **Incrementally**: Use `ragtune explain "query" --save` to build up golden queries
2. **Import CSV**: `ragtune import-queries queries.csv`
3. **Manual**: Create JSON with `{"queries": [{"text": "...", "relevant_docs": [...]}]}`

See [Building Your Test Suite](#building-your-test-suite).
</details>

<details>
<summary><strong>Which embedder should I use?</strong></summary>

| Embedder | Best For |
|----------|----------|
| `ollama` | Development, privacy, no API costs |
| `openai` | General purpose, good baseline |
| `voyage` | Domain-specific (legal, code) |
| `cohere` | Multilingual |
| `tei` | High throughput, self-hosted |
</details>

## Troubleshooting

### "Connection refused" to Qdrant

Qdrant isn't running. Start it:

```bash
docker run -d -p 6333:6333 -p 6334:6334 --name qdrant qdrant/qdrant
```

Or restart if it exists:
```bash
docker start qdrant
```

### "Is Ollama running?"

Start the Ollama server and pull an embedding model:

```bash
ollama serve  # In one terminal
ollama pull nomic-embed-text  # In another
```

### "OPENAI_API_KEY not set"

Either set the key:
```bash
export OPENAI_API_KEY="sk-..."
```

Or use Ollama instead (no key needed):
```bash
ragtune ingest ./docs --collection test --embedder ollama
```

### Slow ingestion

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

### Low recall scores

1. **Check chunk size**: Try larger chunks (`--chunk-size 1024`)
2. **Try different embedder**: `voyage-law-2` for legal, `voyage-code-2` for code
3. **Verify relevant_docs paths**: Must match exactly (case-sensitive)
4. **Inspect with explain**: `ragtune explain "your query" --collection test`

### Embedding dimension mismatch

If you get dimension errors, the collection was created with a different embedder.
Delete and recreate:

```bash
# Delete collection (via vector DB admin or recreate)
ragtune ingest ./docs --collection new-name --embedder ollama
```

## Contributing

Contributions welcome. Please open an issue first to discuss significant changes.

## Sponsor

If RagTune saves you time, consider supporting development:

- **Star this repo** — It helps others discover the project
- **ETH:** `0x0a542565b3615e8fc934cc3cc4921a0c22e5dc5e`

## License

MIT
