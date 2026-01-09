# Sample Dataset

A small subset (~30 legal docs) for quick "does it run?" testing.

```bash
ragtune ingest ./benchmarks/casehold-500/samples --collection demo-legal --embedder ollama
ragtune simulate --collection demo-legal --queries ./benchmarks/casehold-500/samples/queries.json --embedder ollama
```

> ⚠️ **This sample is for testing only.** Config effects (chunk size, embedder differences) only emerge at full scale (500+ docs). To reproduce article results, run the full benchmark:

```bash
cd benchmarks/casehold-500
pip install datasets
python prepare.py
ragtune ingest ./corpus --collection legal --embedder ollama
ragtune simulate --collection legal --queries queries.json --embedder ollama
```
