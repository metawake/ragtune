# CaseHOLD Legal Benchmark

A legal domain benchmark derived from [CaseHOLD](https://github.com/reglab/casehold) (Harvard Law School).

## Why This Dataset?

Legal text is **notoriously difficult** for RAG systems:

- **Dense, precise language** — small word changes have big legal implications
- **Domain-specific terminology** — generic embeddings may miss nuance
- **Long, complex sentences** — chunk boundaries can split key concepts
- **Semantic similarity traps** — holdings that sound alike mean different things

This makes CaseHOLD an excellent stress test for chunking strategies.

## Dataset Structure

- **500 legal holdings** (court decisions) as documents
- **500 citing contexts** as queries (each should retrieve its corresponding holding)
- Ground truth: 1 relevant document per query

## Preparation

```bash
pip install datasets
python prepare.py
```

## Running the Benchmark

```bash
# Ingest with different chunk sizes
ragtune ingest ./docs --collection casehold-256 --chunk-size 256 --embedder ollama
ragtune ingest ./docs --collection casehold-512 --chunk-size 512 --embedder ollama
ragtune ingest ./docs --collection casehold-1024 --chunk-size 1024 --embedder ollama

# Compare performance
ragtune compare --collections casehold-256,casehold-512,casehold-1024 \
    --queries ./queries.json --embedder ollama --top-k 5
```

## Expected Insights

Unlike general-knowledge datasets (like HotpotQA), legal text often shows:

1. **Larger chunk sizes may hurt** — legal precision requires exact phrase matching
2. **Smaller chunks may help** — holdings are self-contained statements
3. **Domain embeddings may outperform** — legal-specific models vs generic

## Source

- Dataset: [CaseHOLD on HuggingFace](https://huggingface.co/datasets/casehold/casehold)
- Paper: [When Does Pretraining Help? (Zheng et al., 2021)](https://arxiv.org/abs/2104.08671)
- License: Research use




