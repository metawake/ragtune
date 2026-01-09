# HotpotQA Benchmark (1K subset)

A real-world benchmark for RagTune using the HotpotQA dataset.

## About HotpotQA

[HotpotQA](https://hotpotqa.github.io/) is a question answering dataset featuring natural, multi-hop questions that require reasoning over multiple Wikipedia paragraphs.

**Why it's perfect for RagTune:**
- Questions require finding 2+ documents to answer correctly
- Has ground truth "supporting facts" (which paragraphs are relevant)
- Real Wikipedia content, not synthetic
- Tests multi-document retrieval quality

## Dataset Statistics

- **Documents:** 1,000 Wikipedia paragraphs
- **Queries:** 200 multi-hop questions
- **Avg relevant docs per query:** 2

## Usage

```bash
# Ingest the benchmark documents
ragtune ingest ./benchmarks/hotpotqa-1k/docs --collection hotpotqa

# Run simulation
ragtune simulate --collection hotpotqa \
  --queries ./benchmarks/hotpotqa-1k/queries.json \
  --configs ./benchmarks/hotpotqa-1k/configs.yaml

# Generate report
ragtune report --run runs/latest.json --out hotpotqa-report.md
```

## Expected Results

With default settings, you should see recall improve significantly with higher top-k:

| Config | Recall@K | Notes |
|--------|----------|-------|
| top-3  | ~0.50    | Misses second supporting doc |
| top-5  | ~0.70    | Better coverage |
| top-10 | ~0.85    | Good for multi-hop |

## License

HotpotQA is released under [CC BY-SA 4.0](https://creativecommons.org/licenses/by-sa/4.0/).

## Citation

```bibtex
@inproceedings{yang2018hotpotqa,
  title={{HotpotQA}: A Dataset for Diverse, Explainable Multi-hop Question Answering},
  author={Yang, Zhiping and others},
  booktitle={EMNLP},
  year={2018}
}
```

## Regenerating the Dataset

To regenerate from scratch:

```bash
python benchmarks/hotpotqa-1k/prepare.py
```

Requires: `pip install datasets`
