# RagTune Product Hypotheses

> Internal document capturing value proposition, target users, and learnings.

## Core Question

**Is RagTune useful, or can an experienced RAG developer intuit optimal parameters?**

## What Intuition Already Provides (~80%)

An experienced RAG developer can reasonably guess:

| Parameter | Rule of Thumb |
|-----------|---------------|
| Chunk size | 256-512 for precise Q&A, 1024+ for summarization |
| Top-k | 3-5 for simple queries, 10-20 for multi-hop |
| Overlap | 10-20% of chunk size |
| Embedding model | "Use a good one" (OpenAI, Cohere, etc.) |

**Implication:** RagTune is NOT about discovering unknown truths for standard use cases.

---

## Where RagTune Adds Real Value

### 1. Domain-Specific Tuning
Intuition fails on unusual data: legal contracts, source code, medical records, multilingual content. Your dataset might behave differently than "typical" text.

### 2. Quantified Proof
- "0.995 recall" beats "I think it works"
- Useful for stakeholders, production sign-off, SLAs
- Enables data-driven decisions vs. vibes

### 3. Regression Detection
After adding 50k new docs, did recall drop from 0.98 to 0.91? You can't intuit that. RagTune catches regressions before users do.

### 4. Finding Specific Failures
Which queries fail? The `explain` command shows:
- What chunks were retrieved
- Their similarity scores
- Why a query missed its ground truth

### 5. CI/CD Integration
Automated quality gates on every deployment. No human intuition scales to this. Example: "Block deploy if Recall@5 drops below 0.95."

### 6. A/B Decisions with Data
"Is `nomic-embed-text` actually better than `text-embedding-3-small` for *my* data?" Run both, compare metrics, decide with evidence.

---

## Market Positioning

### The RAG Consulting Landscape (2025)

Based on market research, RAG consulting is a growing but underserved market:

| What Consultants Offer | What They Lack |
|------------------------|----------------|
| "We optimize RAG" | Reproducible benchmarks |
| "Domain-specific solutions" | Public methodology |
| Full-cycle services | Open-source tooling |
| Case studies with ROI claims | Verifiable metrics |

**RagTune's Edge:** The first tool that makes RAG optimization transparent and reproducible. Consultants sell vibes + experience. RagTune sells measurable proof.

### Three Strategic Paths

| Path | Advantage | Success Signal |
|------|-----------|----------------|
| **Consulting** | Tool = unique pitch ("Here's your Recall@5") | Inbound leads from articles |
| **Startup** | First-mover in benchmarking space | 500+ stars, requests for GUI |
| **Senior Role** | Demonstrable depth vs. typical candidates | Strong portfolio regardless |

All three paths benefit from the same work (tool + articles). Decision point: March 2025.

---

## Target Users

| User Type | Why They Need RagTune |
|-----------|----------------------|
| **Junior/Mid RAG devs** | Learn what good looks like, build intuition |
| **Senior RAG devs** | Validate assumptions, optimize edge cases |
| **ML Platform teams** | CI/CD integration, quality gates |
| **Consulting/Agency** | Prove quality to clients with metrics |
| **Regulated industries** | Audit trail, documented performance |

---

## Key Use Cases

### 1. Initial Tuning
Compare chunk sizes, embedding models, top-k values on your actual data before production.

### 2. Regression Testing
Run `simulate` in CI after:
- Adding new documents
- Changing embedding models
- Updating chunking logic

### 3. Debugging Failures
When users report "bad answers," use `explain` to see exactly what was retrieved and why.

### 4. Chunk Size Experimentation
Use `compare` to test multiple configurations side-by-side with real metrics.

---

## Learnings from Testing

### HotpotQA 1k Benchmark (200 multi-hop queries, 398 docs)

| Chunk Size | Recall@5 | MRR | Coverage | Chunks |
|------------|----------|-----|----------|--------|
| 256 | 0.995 | 0.993 | 1.000 | 1098 |
| 512 | 0.990 | 0.998 | 0.997 | 584 |
| 1024 | 0.993 | 0.998 | 0.997 | 409 |

**Observations:**
- All chunk sizes perform very well (>99% recall)
- Smaller chunks (256) = best recall & coverage
- Larger chunks (512, 1024) = better MRR (relevant doc ranks #1 more often)
- At k=10, all sizes achieve perfect recall

**Conclusion:** For this dataset, 512 with k=10 is optimal: perfect recall, best ranking, fewer chunks.

**Meta-insight:** An experienced dev would guess "512 is probably fine" — and they'd be right. But they wouldn't know the precise trade-offs without measurement.

---

## Positioning

### What RagTune IS
- A measurement and debugging tool for RAG retrieval
- CI/CD-friendly quality assurance
- Framework-agnostic (not tied to LangChain, LlamaIndex)
- Production-minded, not a toy

### What RagTune is NOT
- A RAG framework
- An LLM wrapper
- A "magic optimizer" that finds perfect params automatically
- Essential for trivial/one-time RAG setups

---

## Open Questions

1. **Auto-tuning:** Should RagTune recommend optimal params, or just measure? (Current: measure only)
2. **LLM-as-judge:** Add evaluation of *answer quality*, not just retrieval? (Scope creep risk)
3. **Visualization:** CLI-only or add a web dashboard? (Phase 4+)
4. **Hosted version:** SaaS offering for teams without local setup? (Future)
5. **Audit mode:** Add `ragtune audit` that wraps simulate + recommendations for consulting use cases?
6. **Scale benchmarks:** Prioritize 100K+ document tests for enterprise credibility?
7. **LLM-as-judge:** Add optional `--eval-answers` to close the retrieval→answer loop?

---

## Success Metrics (for the project itself)

- GitHub stars as social proof
- Cited in RAG blog posts / tutorials
- Used in production by at least 3 companies
- Contributor PRs from community

---

*Last updated: December 2024 (market research integration)*


