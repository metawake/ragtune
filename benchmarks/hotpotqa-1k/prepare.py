#!/usr/bin/env python3
"""
Prepare HotpotQA benchmark dataset for RagTune.

Downloads HotpotQA from HuggingFace and extracts:
- 1,000 unique Wikipedia paragraphs (docs)
- 200 multi-hop questions with ground truth

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


def sanitize_filename(title: str) -> str:
    """Convert a title to a safe filename."""
    safe = "".join(c if c.isalnum() or c in "._- " else "_" for c in title)
    safe = safe.strip().replace(" ", "_")[:50]
    hash_suffix = hashlib.md5(title.encode()).hexdigest()[:6]
    return f"{safe}_{hash_suffix}"


def sanitize_text(text: str) -> str:
    """Remove invalid UTF-8 characters and normalize text."""
    # Encode to UTF-8 and back, replacing invalid chars
    text = text.encode('utf-8', errors='replace').decode('utf-8')
    # Remove null bytes and other control characters
    text = "".join(c for c in text if c.isprintable() or c in '\n\t ')
    return text


def main():
    print("Loading HotpotQA dataset from HuggingFace...")
    dataset = load_dataset("hotpot_qa", "fullwiki", split="validation")
    
    target_queries = 200
    
    # First pass: collect good multi-hop examples with their supporting docs
    print(f"Finding {target_queries} multi-hop questions...")
    
    examples = []
    for item in dataset:
        supporting_titles = list(set(item["supporting_facts"]["title"]))
        
        # Must have exactly 2 supporting docs (clean multi-hop)
        if len(supporting_titles) != 2:
            continue
        
        # Get the paragraphs for supporting docs
        context_titles = item["context"]["title"]
        context_sentences = item["context"]["sentences"]
        
        supporting_docs = {}
        for i, title in enumerate(context_titles):
            if title in supporting_titles:
                paragraph = " ".join(context_sentences[i])
                paragraph = sanitize_text(paragraph)
                if len(paragraph) > 100:
                    supporting_docs[title] = paragraph
        
        # Only keep if we found both supporting docs
        if len(supporting_docs) == 2:
            examples.append({
                "question": item["question"],
                "answer": item["answer"],
                "supporting_docs": supporting_docs,
                "type": item.get("type", "unknown")
            })
        
        if len(examples) >= target_queries:
            break
    
    print(f"Found {len(examples)} valid multi-hop examples")
    
    # Collect all unique documents
    docs = {}
    for ex in examples:
        for title, text in ex["supporting_docs"].items():
            if title not in docs:
                docs[title] = {"title": title, "text": text}
    
    print(f"Collected {len(docs)} unique supporting documents")
    
    # Build queries
    queries = []
    for i, ex in enumerate(examples):
        titles = list(ex["supporting_docs"].keys())
        queries.append({
            "id": f"q{i+1}",
            "text": ex["question"],
            "answer": ex["answer"],
            "relevant_docs": [f"{sanitize_filename(t)}.txt" for t in titles],
            "type": ex["type"]
        })
    
    # Create output directories
    docs_dir = Path(__file__).parent / "corpus"
    docs_dir.mkdir(exist_ok=True)
    
    # Clear existing docs
    for f in docs_dir.glob("*.txt"):
        f.unlink()
    
    # Write documents
    print("Writing documents...")
    for title, doc in docs.items():
        filename = sanitize_filename(title) + ".txt"
        filepath = docs_dir / filename
        content = f"# {doc['title']}\n\n{doc['text']}"
        filepath.write_text(content)
    
    # Write queries
    print("Writing queries...")
    queries_file = Path(__file__).parent / "queries.json"
    with open(queries_file, "w") as f:
        json.dump({"queries": queries}, f, indent=2)
    
    # Write configs
    print("Writing configs...")
    configs_file = Path(__file__).parent / "configs.yaml"
    configs_file.write_text("""configs:
  - name: top3
    top_k: 3

  - name: top5
    top_k: 5

  - name: top10
    top_k: 10

  - name: top20
    top_k: 20
""")
    
    # Stats
    print(f"""
Done! Created:
  - {len(docs)} documents in {docs_dir}/
  - {len(queries)} queries in {queries_file}
  - Config file at {configs_file}
  
Query type distribution:
""")
    types = {}
    for q in queries:
        t = q.get("type", "unknown")
        types[t] = types.get(t, 0) + 1
    for t, count in sorted(types.items()):
        print(f"  - {t}: {count}")

    print(f"""
To run the benchmark:
  ragtune ingest ./benchmarks/hotpotqa-1k/docs --collection hotpotqa
  ragtune simulate --collection hotpotqa --queries ./benchmarks/hotpotqa-1k/queries.json
""")


if __name__ == "__main__":
    main()
