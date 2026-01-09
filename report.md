# RagTune Retrieval Report

**Timestamp:** 2025-12-23T15:02:54Z  
**Collection:** demo2  
**Store:** qdrant  

---

## Summary

| Config | Top-K | Recall@K | MRR | Coverage | Redundancy |
|--------|-------|----------|-----|----------|------------|
| top3 | 3 | 0.689 | 0.967 | 0.778 | 6.43 |
| top5 | 5 | 0.856 | 0.967 | 1.000 | 8.33 |
| top10 | 10 | 0.978 | 0.967 | 1.000 | 16.67 |

---

## Detailed Results

### Config: top3 (top_k=3)

| Query | Retrieved | Relevant | Hit? |
|-------|-----------|----------|------|
| q1 | security.md, security.md, security.md | security.md, api-keys.md, authentication.md | ✅ |
| q2 | troubleshooting.md, authentication.md, authentication.md | troubleshooting.md, authentication.md | ✅ |
| q3 | webhooks.md, webhooks.md, integrations.md | webhooks.md, getting-started.md, integrations.md | ✅ |
| q4 | authentication.md, authentication.md, integrations.md | rate-limits.md, authentication.md | ✅ |
| q5 | integrations.md, integrations.md, authentication.md | troubleshooting.md, integrations.md, rate-limits.md | ✅ |
| q6 | authentication.md, authentication.md, security.md | api-keys.md, security.md, authentication.md | ✅ |
| q7 | billing.md, webhooks.md, troubleshooting.md | webhooks.md, integrations.md | ✅ |
| q8 | troubleshooting.md, authentication.md, rate-limits.md | rate-limits.md, troubleshooting.md | ✅ |
| q9 | authentication.md, security.md, security.md | security.md, authentication.md, api-keys.md | ✅ |
| q10 | integrations.md, integrations.md, integrations.md | integrations.md | ✅ |
| q11 | billing.md, troubleshooting.md, troubleshooting.md | billing.md | ✅ |
| q12 | security.md, webhooks.md, security.md | webhooks.md, security.md, integrations.md | ✅ |
| q13 | troubleshooting.md, troubleshooting.md, billing.md | billing.md, troubleshooting.md | ✅ |
| q14 | billing.md, billing.md, troubleshooting.md | billing.md, rate-limits.md | ✅ |
| q15 | security.md, security.md, troubleshooting.md | security.md, authentication.md | ✅ |

### Config: top5 (top_k=5)

| Query | Retrieved | Relevant | Hit? |
|-------|-----------|----------|------|
| q1 | security.md, security.md, security.md (+2) | security.md, api-keys.md, authentication.md | ✅ |
| q2 | troubleshooting.md, authentication.md, authentication.md (+2) | troubleshooting.md, authentication.md | ✅ |
| q3 | webhooks.md, webhooks.md, integrations.md (+2) | webhooks.md, getting-started.md, integrations.md | ✅ |
| q4 | authentication.md, authentication.md, integrations.md (+2) | rate-limits.md, authentication.md | ✅ |
| q5 | integrations.md, integrations.md, authentication.md (+2) | troubleshooting.md, integrations.md, rate-limits.md | ✅ |
| q6 | authentication.md, authentication.md, security.md (+2) | api-keys.md, security.md, authentication.md | ✅ |
| q7 | billing.md, webhooks.md, troubleshooting.md (+2) | webhooks.md, integrations.md | ✅ |
| q8 | troubleshooting.md, authentication.md, rate-limits.md (+2) | rate-limits.md, troubleshooting.md | ✅ |
| q9 | authentication.md, security.md, security.md (+2) | security.md, authentication.md, api-keys.md | ✅ |
| q10 | integrations.md, integrations.md, integrations.md (+2) | integrations.md | ✅ |
| q11 | billing.md, troubleshooting.md, troubleshooting.md (+2) | billing.md | ✅ |
| q12 | security.md, webhooks.md, security.md (+2) | webhooks.md, security.md, integrations.md | ✅ |
| q13 | troubleshooting.md, troubleshooting.md, billing.md (+2) | billing.md, troubleshooting.md | ✅ |
| q14 | billing.md, billing.md, troubleshooting.md (+2) | billing.md, rate-limits.md | ✅ |
| q15 | security.md, security.md, troubleshooting.md (+2) | security.md, authentication.md | ✅ |

### Config: top10 (top_k=10)

| Query | Retrieved | Relevant | Hit? |
|-------|-----------|----------|------|
| q1 | security.md, security.md, security.md (+7) | security.md, api-keys.md, authentication.md | ✅ |
| q2 | troubleshooting.md, authentication.md, authentication.md (+7) | troubleshooting.md, authentication.md | ✅ |
| q3 | webhooks.md, webhooks.md, integrations.md (+7) | webhooks.md, getting-started.md, integrations.md | ✅ |
| q4 | authentication.md, authentication.md, integrations.md (+7) | rate-limits.md, authentication.md | ✅ |
| q5 | integrations.md, integrations.md, authentication.md (+7) | troubleshooting.md, integrations.md, rate-limits.md | ✅ |
| q6 | authentication.md, authentication.md, security.md (+7) | api-keys.md, security.md, authentication.md | ✅ |
| q7 | billing.md, webhooks.md, troubleshooting.md (+7) | webhooks.md, integrations.md | ✅ |
| q8 | troubleshooting.md, authentication.md, rate-limits.md (+7) | rate-limits.md, troubleshooting.md | ✅ |
| q9 | authentication.md, security.md, security.md (+7) | security.md, authentication.md, api-keys.md | ✅ |
| q10 | integrations.md, integrations.md, integrations.md (+7) | integrations.md | ✅ |
| q11 | billing.md, troubleshooting.md, troubleshooting.md (+7) | billing.md | ✅ |
| q12 | security.md, webhooks.md, security.md (+7) | webhooks.md, security.md, integrations.md | ✅ |
| q13 | troubleshooting.md, troubleshooting.md, billing.md (+7) | billing.md, troubleshooting.md | ✅ |
| q14 | billing.md, billing.md, troubleshooting.md (+7) | billing.md, rate-limits.md | ✅ |
| q15 | security.md, security.md, troubleshooting.md (+7) | security.md, authentication.md | ✅ |

---

## Metrics Guide

- **Recall@K**: Fraction of relevant docs found in top-K results (higher is better)
- **MRR**: Mean Reciprocal Rank — how high the first relevant result ranks (higher is better)
- **Coverage**: Fraction of all relevant docs ever retrieved across queries (higher is better)
- **Redundancy**: Average times each doc is retrieved (lower may indicate diverse results)

---

*Generated by [RagTune](https://github.com/metawake/ragtune)*
