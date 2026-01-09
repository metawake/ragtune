#!/usr/bin/env python3
"""
Prepare CaseHOLD benchmark dataset for RagTune.

CaseHOLD contains legal case holdings from U.S. court decisions.
Each example has a citing context and 5 possible holdings (1 correct).

We reframe it as a retrieval task:
- Documents: All correct holdings (unique legal statements)
- Queries: The citing context that should retrieve the correct holding

This tests RAG on dense legal language where:
- Similar-sounding clauses mean different things
- Chunk boundaries matter for legal precision
- Domain terminology challenges generic embeddings

Usage:
    pip install datasets
    python prepare.py
"""

import json
import os
import hashlib
from pathlib import Path

try:
    from datasets import load_dataset
except ImportError:
    print("Please install datasets: pip install datasets")
    exit(1)


def sanitize_text(text: str) -> str:
    """Remove invalid UTF-8 characters and normalize text."""
    text = text.encode('utf-8', errors='replace').decode('utf-8')
    text = "".join(c for c in text if c.isprintable() or c in '\n\t ')
    return text.strip()


def make_doc_id(text: str) -> str:
    """Create a deterministic short ID for a document."""
    return hashlib.md5(text.encode()).hexdigest()[:12]


def main():
    print("Loading CaseHOLD dataset from LexGLUE on HuggingFace...")
    # CaseHOLD is available via lex_glue dataset
    dataset = load_dataset("lex_glue", "case_hold", split="train")
    
    target_queries = 500
    min_context_len = 200  # Filter out very short contexts
    min_holding_len = 50   # Filter out very short holdings
    
    print(f"Processing dataset ({len(dataset)} examples)...")
    
    # Collect examples with valid holdings
    examples = []
    seen_holdings = set()  # Avoid duplicate holdings
    
    for item in dataset:
        # Context has <HOLDING> placeholder - remove it for cleaner query
        citing_context = sanitize_text(item["context"].replace("<HOLDING>", ""))
        label = item["label"]
        
        # Get the correct holding from endings list
        correct_holding = sanitize_text(item["endings"][label])
        
        # Filter criteria
        if len(citing_context) < min_context_len:
            continue
        if len(correct_holding) < min_holding_len:
            continue
        
        # Skip if we've seen this exact holding
        holding_hash = hashlib.md5(correct_holding.encode()).hexdigest()
        if holding_hash in seen_holdings:
            continue
        seen_holdings.add(holding_hash)
        
        examples.append({
            "context": citing_context,
            "holding": correct_holding,
            "holding_id": make_doc_id(correct_holding)
        })
        
        if len(examples) >= target_queries:
            break
    
    print(f"Collected {len(examples)} valid examples")
    
    # Create output directories
    docs_dir = Path(__file__).parent / "corpus"
    docs_dir.mkdir(exist_ok=True)
    
    # Clear existing docs
    for f in docs_dir.glob("*.txt"):
        f.unlink()
    
    # Write documents (holdings)
    print("Writing holdings as documents...")
    for ex in examples:
        filename = f"holding_{ex['holding_id']}.txt"
        filepath = docs_dir / filename
        # Include a header to give context
        content = f"# Legal Holding\n\n{ex['holding']}"
        filepath.write_text(content)
    
    # Build queries (citing contexts)
    queries = []
    for i, ex in enumerate(examples):
        queries.append({
            "id": f"case{i+1}",
            "text": ex["context"],
            "relevant_docs": [f"holding_{ex['holding_id']}.txt"]
        })
    
    # Write queries
    print("Writing queries...")
    queries_file = Path(__file__).parent / "queries.json"
    with open(queries_file, "w") as f:
        json.dump({"queries": queries}, f, indent=2)
    
    # Write configs
    print("Writing configs...")
    configs_file = Path(__file__).parent / "configs.yaml"
    configs_file.write_text("""configs:
  - name: top1
    top_k: 1

  - name: top3
    top_k: 3

  - name: top5
    top_k: 5

  - name: top10
    top_k: 10
""")
    
    # Sample document lengths for stats
    doc_lengths = [len(ex["holding"]) for ex in examples]
    query_lengths = [len(ex["context"]) for ex in examples]
    
    print(f"""
Done! Created:
  - {len(examples)} holding documents in {docs_dir}/
  - {len(queries)} queries in {queries_file}
  - Config file at {configs_file}

Document stats (holdings):
  - Min length: {min(doc_lengths)} chars
  - Max length: {max(doc_lengths)} chars
  - Avg length: {sum(doc_lengths)//len(doc_lengths)} chars

Query stats (citing contexts):
  - Min length: {min(query_lengths)} chars
  - Max length: {max(query_lengths)} chars
  - Avg length: {sum(query_lengths)//len(query_lengths)} chars

Why this dataset is challenging:
  - Dense legal language with domain-specific terminology
  - Holdings are semantically similar but legally distinct
  - Citing contexts are long and complex
  - Standard chunking may split legal concepts incorrectly

To run the benchmark:
  ragtune ingest ./benchmarks/casehold-500/docs --collection casehold
  ragtune simulate --collection casehold --queries ./benchmarks/casehold-500/queries.json
""")


if __name__ == "__main__":
    main()

