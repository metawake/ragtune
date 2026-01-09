# RagTune Research Ideas

> Future experiments and research directions for the platform.

---

## 1. 🏆 Embedding Model Showdown

**Question:** Which embedding model wins for different domains?

**Models to test:**
| Model | Type | Cost | Dimension |
|-------|------|------|-----------|
| OpenAI text-embedding-3-small | Cloud | $$ | 1536 |
| OpenAI text-embedding-3-large | Cloud | $$$ | 3072 |
| Cohere embed-v3 | Cloud | $$ | 1024 |
| Voyage-2 | Cloud | $$ | 1024 |
| nomic-embed-text (Ollama) | Local | Free | 768 |
| mxbai-embed-large (Ollama) | Local | Free | 1024 |
| bge-large-en-v1.5 | Local | Free | 1024 |

**Command:**
```bash
ragtune compare --embedders openai,cohere,voyage,nomic,bge \
    --docs ./corpus --queries ./queries.json
```

**Why it matters:** Everyone asks "which embedder?" — no comprehensive cross-domain benchmarks exist.

---

## 2. 📊 Query Type Taxonomy

**Question:** Do different query types need different RAG configurations?

**Categories:**
| Type | Example | Expected Behavior |
|------|---------|-------------------|
| Factual/KB | "What is the refund policy?" | Single chunk, high precision |
| How-to/Procedural | "How do I integrate webhooks?" | Multiple chunks, sequence |
| Conceptual/Broad | "Explain the auth architecture" | Many chunks, context breadth |
| Comparison | "Difference between plans?" | Multiple related chunks |
| Troubleshooting | "Why is API returning 403?" | Specific + contextual |

**Experiment:**
- Tag queries by type in queries.json
- Run same corpus with different top-k per type
- Measure: Does optimal k vary by query type?

**Potential finding:** Adaptive retrieval — different k for different query types.

---

## 3. 📈 Scale Degradation Study

**Question:** How does recall degrade as corpus grows?

| Corpus Size | Expected Recall | Why? |
|-------------|-----------------|------|
| 100 docs | 95%+ | Easy |
| 1,000 docs | 90%+ | Manageable |
| 10,000 docs | 80%? | Semantic overlap |
| 100,000 docs | ??? | Embedding space crowded |

**Datasets:**
- Wikipedia dumps (scalable)
- ArXiv abstracts (100k+ papers)
- StackOverflow posts

**Why it matters:** Most tutorials test tiny datasets. Production has millions.

---

## 4. 🧬 Domain-Specific Embedding Battle

**Question:** Do specialized embeddings beat general-purpose?

| Domain | Generic | Specialized |
|--------|---------|-------------|
| Legal | text-embedding-3 | legal-bert, saul-bert |
| Medical | nomic-embed | pubmed-bert, biobert |
| Code | bge-large | codebert, code-search-net |
| Scientific | voyage-2 | specter, scibert |

**Hypothesis:** Specialized models win by 10%+ on domain queries.

**Why it matters:** Answers "should I fine-tune or use off-the-shelf?"

---

## 5. 🔀 Chunking Strategy Deep Dive

**Question:** Does *how* you chunk matter as much as size?

| Strategy | Description |
|----------|-------------|
| Fixed-size | Current default (512 chars) |
| Sentence-based | Split on sentence boundaries |
| Paragraph-based | Respect document structure |
| Semantic | Split on topic shifts (NLP) |
| Overlap variations | 0%, 10%, 25%, 50% |

**Why it matters:** "Chunk by sentences or paragraphs?" is unanswered.

---

## 6. 🌍 Multilingual Retrieval

**Question:** Cross-lingual embedding performance?

| Setup | Query Lang | Doc Lang |
|-------|------------|----------|
| Monolingual | English | English |
| Cross-lingual | English | French |
| Mixed corpus | English | EN + FR + DE |

**Models:**
- multilingual-e5-large
- cohere-multilingual-v3
- OpenAI (claims multilingual)

**Why it matters:** Global companies need this. Few benchmarks exist.

---

## 7. 🎯 Edge Case Stress Tests

**Question:** How does RAG handle tricky cases?

| Scenario | Challenge |
|----------|-----------|
| Near-duplicates | 10 docs with 90% identical text |
| Contradictory info | Doc A says X, Doc B says not-X |
| Outdated vs current | Old policy vs new policy |
| Adversarial | "Ignore previous instructions..." |
| Empty/minimal docs | Single-sentence documents |

**Why it matters:** Production edge cases rarely tested.

---

## 8. 📉 Embedding Dimension vs Quality

**Question:** Is bigger always better?

| Model | Dimensions | Speed | Quality? |
|-------|------------|-------|----------|
| text-embedding-3-small | 1536 | Fast | Baseline |
| text-embedding-3-large | 3072 | Slower | Better? |
| Matryoshka (truncated) | 256/512/1024 | Variable | Test each |

**OpenAI Matryoshka:** Truncate embeddings to smaller dims. Does 512-dim match 1536?

**Why it matters:** Smaller = faster search, less storage, lower cost.

---

## 9. 🔄 Query Paraphrase Robustness

**Question:** How sensitive is retrieval to query phrasing?

| Original Query | Paraphrases |
|----------------|-------------|
| "How to reset password?" | "password reset steps", "forgot my password", "change password process" |

**Experiment:**
- Generate 5-10 paraphrases per query (AI or manual)
- Measure variance in retrieval results
- Identify fragile vs robust queries

**Why it matters:** Users don't phrase queries consistently.

---

## 10. ⚡ Latency vs Quality Tradeoffs

**Question:** What's the speed/quality curve?

| Configuration | Latency | Quality |
|---------------|---------|---------|
| top-k=3, small embeddings | Fast | ? |
| top-k=10, large embeddings | Slow | ? |
| Approximate NN vs exact | ? | ? |

**Why it matters:** Production needs sub-100ms responses.

---

## 11. 🔍 RAG Audit Methodology

**Question:** Can we create a standardized "RAG Health Check"?

**Components:**

| Check | Metric | Threshold |
|-------|--------|-----------|
| Retrieval quality | Recall@5 | > 0.85 |
| Ranking quality | MRR | > 0.70 |
| Coverage | Coverage | > 0.95 |
| Scale readiness | Recall degradation at 10x docs | < 5% drop |

**Output:** "RAG Audit Report" — shareable artifact for consulting engagements or internal reviews.

**Why it matters:** Creates a productized consulting offering. No competitor has a standardized assessment framework.

**Implementation:**
```bash
ragtune audit --docs ./corpus --queries ./queries.json --output audit-report.md
```

Generates:
- Current metrics vs. thresholds
- Pass/fail status per check
- Recommendations for improvement
- Comparison to baseline benchmarks

---

## Priority Ranking

| Priority | Research | Impact | Effort | Market Signal |
|----------|----------|--------|--------|---------------|
| **1** | **Scale Degradation Study** | Enterprise credibility | Medium | Big cos care about 100K+ docs |
| 2 | Query Type Taxonomy | Novel, publishable | Medium | Differentiation |
| 3 | Embedding Showdown | High demand | Low | Commodity but expected |
| 4 | Chunking Strategies | Practical | Medium | Already proven valuable |
| 5 | Domain Embeddings | Specialized value | High | Legal/medical verticals |

### Why Scale is Priority 1

Market research shows enterprise buyers (Deloitte, Accenture clients) don't trust benchmarks on tiny datasets. Most RAG tutorials test 100-1000 docs. Production has millions.

Article: "What Happens to RAG at 100K Documents?" would:
- Signal production-level understanding
- Differentiate from tutorial-level competitors
- Attract enterprise attention

---

## Features Needed

| Feature | For Research |
|---------|--------------|
| More embedders (Cohere, Voyage, BGE) | Embedding showdown |
| Query type field + grouping | Query taxonomy |
| Per-type metrics in reports | Query taxonomy |
| Semantic chunker | Chunking study |
| Latency tracking | Performance study |

---

*Last updated: December 2024 (market research integration)*


