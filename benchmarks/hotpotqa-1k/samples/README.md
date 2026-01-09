# Sample Dataset

A small subset (~30 docs) for quick "does it run?" testing.

```bash
ragtune ingest ./benchmarks/hotpotqa-1k/samples --collection demo --embedder ollama
ragtune simulate --collection demo --queries ./benchmarks/hotpotqa-1k/samples/queries.json --embedder ollama
```

> ⚠️ **This sample is for testing only.** Config effects (chunk size, embedder differences) only emerge at full scale (500+ docs). To reproduce article results, run the full benchmark:

```bash
cd benchmarks/hotpotqa-1k
pip install datasets
python prepare.py
ragtune ingest ./corpus --collection hotpotqa --embedder ollama
ragtune simulate --collection hotpotqa --queries queries.json --embedder ollama
```
