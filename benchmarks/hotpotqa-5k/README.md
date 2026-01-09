# HotpotQA 5K Benchmark

Benchmark with **5,000 documents** to test retrieval at scale.

## Dataset

- **Source:** HotpotQA (HuggingFace, train split)
- **Documents:** 5,001 Wikipedia paragraphs
- **Queries:** 2,581 multi-hop questions
- **Ground truth:** Each query maps to exactly 2 supporting documents

## Results (December 2024)

Tested with OpenAI text-embedding-3-small:

| Chunk Size | Recall@5 | MRR | Coverage | Chunks |
|------------|----------|-----|----------|--------|
| 256 | 0.874 | 0.983 | 0.943 | 12,617 |
| 512 | 0.883 | 0.980 | 0.957 | 6,785 |
| **1024** | **0.893** | 0.982 | **0.971** | 5,094 |

**Key finding:** Larger chunks (1024) perform best at this scale — opposite of the 400-doc benchmark where smaller chunks (256) won.

## Comparison with 1K

| Metric | 1K (400 docs) | 5K (5,000 docs) | Change |
|--------|---------------|-----------------|--------|
| Documents | 398 | 5,001 | 12.5x |
| Queries | 200 | 2,581 | 13x |
| Best chunk | 256 | 1024 | **Flipped** |
| Best Recall@5 | 0.995 | 0.893 | -10% |

## Usage

```bash
# Prepare dataset (requires: pip install datasets)
python prepare.py

# Run benchmark
ragtune ingest ./docs --collection hotpotqa-5k-512 --chunk-size 512
ragtune simulate --collection hotpotqa-5k-512 --queries ./queries.json
```
