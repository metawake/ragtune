# RagTune Content Plan

> Publishing schedule, feature releases, and distribution strategy.

---

## Timing Strategy

| Period | Engagement | Action |
|--------|------------|--------|
| Dec 24-26 | 🔴 Dead | Don't publish |
| Dec 27-31 | 🟡 Low | Prep, polish, schedule |
| Jan 2-3 | 🟡 Ramp-up | Soft launch (tweet, not full push) |
| **Jan 7-9** | 🟢 Peak | **First major article** |
| Jan onwards | 🟢 Normal | Weekly cadence |

**Best days:** Tuesday, Wednesday, Thursday (10am-12pm local timezone)

---

## Master Schedule: Content + Features

| Week | Date | Content | Feature Release |
|------|------|---------|-----------------|
| 0 | Dec 24-31 | Prep & polish | — |
| 1 | **Jan 7** | 📝 Article 1: Chunk Size | — |
| 2 | Jan 8-13 | Engage, promote | 🔧 pgvector backend |
| 2 | **Jan 14** | 📝 Article 2: Embeddings | — |
| 3 | Jan 15-20 | Engage, promote | 🔧 CI/CD mode (--ci, thresholds) |
| 3 | Jan 15-20 | — | 📊 Dataset upgrade (HotpotQA 10k, SciFact, FiQA) |
| 4 | Jan 21-27 | Run pgvector benchmarks | 🔧 Latency tracking |
| 4 | **Jan 28** | 📝 Article 3: Qdrant vs pgvector | — |
| 5 | **Feb 4** | 📝 Article 4: Optimization Hierarchy | 🔧 More embedders (Cohere) |
| 6 | **Feb 11** | 📝 Article 5: Query Types | 🔧 Query type analysis |
| 7 | Feb 12-17 | Build Weaviate backend | 🔧 Weaviate backend |
| 8 | **Feb 25** | 📝 Article 6: 3 DBs Compared | — |

---

## Detailed Schedule

### Week 0: Dec 24-31 (Prep)

**Content tasks:**
- [x] Articles 1 & 2 written
- [ ] Polish Article 1 (chunk size discovery)
- [ ] Create GitHub repo banner/social image
- [ ] Write tweet/LinkedIn drafts
- [ ] Ensure repo README is clean

**Feature tasks:**
- [ ] None — don't add features, polish only
- [ ] Fix any obvious bugs
- [ ] Test all commands work cleanly

---

### Week 1: Jan 6-10 — LAUNCH WEEK

| Day | Content | Features |
|-----|---------|----------|
| Mon Jan 6 | Final prep | — |
| **Tue Jan 7** | 🚀 PUBLISH Article 1 | — |
| Wed Jan 8 | Engage comments | — |
| Thu Jan 9 | Reddit cross-post | Start pgvector backend |
| Fri Jan 10 | Monitor, respond | Continue pgvector |

---

### Week 2: Jan 13-17

| Day | Content | Features |
|-----|---------|----------|
| Mon Jan 13 | Prep Article 2 | Finish pgvector backend |
| **Tue Jan 14** | 🚀 PUBLISH Article 2 | pgvector released |
| Wed Jan 15 | Engage comments | — |
| Thu Jan 16 | Reddit cross-post | Start CI/CD mode |
| Fri Jan 17 | Monitor, respond | Continue CI/CD |

---

### Week 3: Jan 20-24

| Day | Content | Features |
|-----|---------|----------|
| Mon Jan 20 | — | Finish CI/CD mode |
| Tue Jan 21 | — | CI/CD released, start latency tracking |
| Wed Jan 22 | Run pgvector benchmarks | Continue latency |
| Thu Jan 23 | Analyze results | Finish latency tracking |
| Fri Jan 24 | Write Article 3 | Latency released |

---

### Week 4: Jan 27-31

| Day | Content | Features |
|-----|---------|----------|
| Mon Jan 27 | Polish Article 3 | — |
| **Tue Jan 28** | 🚀 PUBLISH Article 3 | — |
| Wed Jan 29 | Engage comments | Start Cohere embedder |
| Thu Jan 30 | Reddit cross-post | Continue embedders |
| Fri Jan 31 | Monitor | Finish Cohere embedder |

---

### February Schedule

| Date | Content | Features |
|------|---------|----------|
| **Feb 4** | 🚀 Article 4: Optimization Hierarchy | Cohere released |
| Feb 5-10 | Promote | Query type analysis feature |
| **Feb 11** | 🚀 Article 5: Query Types | Query types released |
| Feb 12-17 | Promote | Weaviate backend |
| Feb 18-24 | Run Weaviate benchmarks | Weaviate released |
| **Feb 25** | 🚀 Article 6: 3 DBs Compared | — |

---

## Feature Release Checklist

### pgvector Backend (Release: Jan 14)
- [ ] Implement Store interface
- [ ] Test with HotpotQA dataset
- [ ] Test with CaseHOLD dataset
- [ ] Update README
- [ ] Commit & tag v0.2.0

### CI/CD Mode (Release: Jan 21)
- [ ] Add --ci flag
- [ ] Add --min-recall threshold
- [ ] Add --min-coverage threshold
- [ ] Exit code based on thresholds
- [ ] Update README
- [ ] Commit & tag v0.3.0

### Latency Tracking (Release: Jan 24)
- [ ] Add timing to search operations
- [ ] Report p50, p95, p99
- [ ] Add --benchmark-latency flag
- [ ] Update README
- [ ] Commit & tag v0.3.1

### More Embedders - Cohere (Release: Feb 4)
- [ ] Implement Cohere embedder
- [ ] Add --embedder cohere flag
- [ ] Test with datasets
- [ ] Update README
- [ ] Commit & tag v0.4.0

### Query Type Analysis (Release: Feb 11)
- [ ] Add query type field support
- [ ] Group metrics by type
- [ ] Add analyze-queries command
- [ ] Update README
- [ ] Commit & tag v0.5.0

### Weaviate Backend (Release: Feb 18)
- [ ] Implement Store interface
- [ ] Test with all datasets
- [ ] Update README
- [ ] Commit & tag v0.6.0

### Dataset Upgrade (Release: Jan 20)
Upgrade from "toy" to "convincing" scale:

| Dataset | Current | Target | Domain |
|---------|---------|--------|--------|
| HotpotQA | 398 docs | 10k docs | General knowledge |
| CaseHOLD | 500 docs | Keep (legal) | Legal |
| SciFact | — | 5k docs | Scientific |
| FiQA | — | 10k subset | Financial |
| MS MARCO | — | 10k subset | Web search |

**Result:** ~35k docs, ~5k queries, 5 domains → "Convincing" level

Tasks:
- [ ] Modify HotpotQA prepare.py for 10k docs
- [ ] Add SciFact prepare.py
- [ ] Add FiQA prepare.py  
- [ ] Add MS MARCO prepare.py
- [ ] Run benchmarks on upgraded datasets
- [ ] Update article results if needed

---

## Content Backlog

| # | Article | Status | Target Date |
|---|---------|--------|-------------|
| 1 | I Built a RAG Tuning Tool and Discovered Intuition Fails on Legal Text | ✅ Ready | Jan 7 |
| 2 | OpenAI vs Ollama: I Benchmarked Both on Real Legal Data | ✅ Ready | Jan 14 |
| 3 | Qdrant vs pgvector: The Benchmark Nobody Asked For | 📝 Need backend | Jan 28 |
| 4 | The RAG Optimization Hierarchy: What Consultants Won't Tell You | 📝 Outline exists | Feb 4 |
| 5 | Why Query Types Matter More Than Your Vector Database | 💡 Idea | Feb 11 |
| 6 | Three Vector DBs, One Answer: What Actually Matters | 💡 Idea | Feb 25 |
| 7 | What Happens to RAG at 100K Documents? I Tested It | 💡 Idea | Mar |

---

## Methodology Progression

Gradually increase rigor as audience grows:

| Phase | Articles | Rigor Level | What to Include |
|-------|----------|-------------|-----------------|
| **Launch** | 1-2 | Basic | Results + brief limitations note |
| **Establish** | 3-4 | Basic+ | Full limitations section |
| **Credibility** | 5-6 | Intermediate | Multiple runs, mean ± std |
| **Authority** | 7+ | Advanced | BEIR benchmark, BM25 baseline, stats |

### Article 1-2: Basic Rigor
- Single run results (acceptable for blog)
- Clear methodology description
- Brief limitations note (1-2 sentences)
- Reproducible (code available)

### Article 3-4: Add Limitations Section
```markdown
## Limitations
- Tested on 2 datasets; results may differ for other domains
- Single embedding model per test; other models may vary
- Fixed chunk overlap (12.5%); not systematically explored
- No comparison to keyword search (BM25)
```

### Article 5-6: Add Statistical Confidence
- Run each experiment 3x
- Report mean ± standard deviation
- Note variance in results
- Example: "Recall@5 = 0.664 ± 0.008 (n=3)"

### Article 7+: Full Rigor
- Use BEIR benchmark (industry standard)
- Include BM25 baseline comparison
- Multiple embedding models in same study
- Statistical significance tests if applicable
- Comprehensive limitations section

---

## Distribution Checklist (Per Article)

```
□ Medium (primary)
□ dev.to (cross-post, reaches different audience)
□ LinkedIn (personal post + article link)
□ Twitter/X (thread summarizing key points)
□ Hacker News (Show HN or regular post)
□ Reddit: r/MachineLearning, r/LocalLLaMA, r/LangChain
□ Relevant Discord servers (LangChain, AI/ML communities)
□ Lobsters (if appropriate)
```

---

## Platform-Specific Notes

### Hacker News
- Title matters enormously
- Format: "Show HN: RagTune – I benchmarked RAG chunk sizes on legal text"
- Post in evening (US time) for visibility
- Don't ask for upvotes (against rules)
- Engage genuinely in comments

### LinkedIn
- Personal story angle works best
- "I built this because..." resonates
- Tag relevant people/companies
- Use 3-5 hashtags max

### Twitter/X
- Thread format (5-8 tweets)
- Lead with the surprising finding
- Include screenshot of results table
- End with CTA to GitHub

### Reddit
- Each subreddit has different rules
- r/MachineLearning: more academic
- r/LocalLLaMA: Ollama angle
- r/LangChain: practical RAG users

---

## Quick Wins Before Jan 7

| Task | Time | Impact |
|------|------|--------|
| Add social preview image to repo | 30 min | Better link sharing |
| Polish README intro paragraph | 15 min | First impression |
| Create Twitter thread draft for Article 1 | 30 min | Ready to post |
| Set up Medium account if needed | 15 min | Publishing ready |
| Prepare HN title | 15 min | Critical for clicks |

---

## Success Metrics

### First Month Targets

| Metric | Target | Stretch |
|--------|--------|---------|
| GitHub stars | 100+ | 500+ |
| Article views (total) | 5,000+ | 20,000+ |
| LinkedIn impressions | 10,000+ | 50,000+ |
| Hacker News frontpage | 1 article | 2 articles |
| Inbound messages/emails | 5+ | 20+ |
| Newsletter subscribers | 50+ | 200+ |

### Tracking
- GitHub: Stars, forks, issues
- Medium: Views, reads, claps
- LinkedIn: Impressions, engagement
- Twitter: Impressions, followers gained
- HN: Points, comments

---

## Article Templates

### Title Formulas That Work
- "I [did X] and discovered [surprising Y]"
- "[Tool A] vs [Tool B]: I benchmarked both on real data"
- "Why [common belief] is wrong for [specific domain]"
- "The [X] Optimization Hierarchy: What Actually Matters"

### Structure
1. Hook (surprising finding)
2. Problem (why this matters)
3. What I built (brief)
4. Experiment (methodology)
5. Results (table/data)
6. Analysis (why)
7. Takeaway (actionable)
8. CTA (try it, star repo)

---

## Long-Term Vision

| Phase | Timeline | Goal |
|-------|----------|------|
| Launch | Jan 2025 | Establish presence, 500 stars |
| Growth | Feb-Mar 2025 | 2,000 stars, recognized in community |
| Authority | Q2 2025 | Invited to podcasts/conferences |
| Standard | H2 2025 | Vector DB companies submit benchmarks |
| Multimodal | Q3-Q4 2025 | Expand to audio/image/video RAG |

---

## Market Differentiation Strategy

Based on RAG consulting market research (Dec 2024):

### Key Insight

Most RAG consultants (Vstorm, Prismetric, Valprovia, etc.) offer:
- Full-cycle services (consulting → data prep → development → support)
- Business-value language ("improve decision-making", "reduce hallucinations")
- Case studies with ROI claims

**None offer:** Transparent, reproducible benchmarking methodology.

### Content Positioning

| Competitor Angle | Our Counter-Angle |
|-----------------|-------------------|
| "Trust our expertise" | "Run the benchmark yourself" |
| "We optimize RAG" | "Here's exactly what improved and by how much" |
| "Domain-specific solutions" | "Here's how your domain actually behaves vs. generic text" |

### Articles Should:

1. **Challenge conventional wisdom** — "Why 512 chunk size is wrong for legal text"
2. **Show measurable results** — Always include Recall@K, MRR tables
3. **Invite reproduction** — Link to code, datasets, and exact commands
4. **Position against consultants** — Subtly contrast with "vibes-based" optimization

---

## Multimodal RAG Expansion (Q3-Q4 2025)

*Only pursue after text RAG is established (1,000+ stars, community traction).*

### Audio RAG
- **Partner:** AssemblyAI for transcription
- **Features:** Speaker-aware chunking, timestamp mapping
- **Use cases:** Meeting search, podcast RAG, call analytics
- **Article:** "Benchmarking Audio RAG: From Meetings to Answers"

### Image RAG
- **Embeddings:** CLIP, SigLIP
- **Features:** Text-to-image retrieval, image-to-image search
- **Use cases:** Product catalogs, visual documentation
- **Article:** "Text RAG vs Image RAG: When Pictures Beat Words"

### Video RAG
- **Approach:** Frame sampling + transcript + audio
- **Complexity:** 🔴 High
- **Use cases:** Training video search, content moderation
- **Article:** "Video RAG: The Next Frontier"

### Signals to Start Multimodal
- [ ] 1,000+ GitHub stars on text version
- [ ] 10+ requests for multimodal support
- [ ] Potential sponsor/partner interested
- [ ] Text RAG roadmap mostly complete

---

*Last updated: December 2024 (market research integration)*

