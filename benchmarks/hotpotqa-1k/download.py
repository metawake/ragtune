#!/usr/bin/env python3
"""
Download and prepare HotpotQA-1K benchmark dataset.

Usage:
    pip install datasets
    python download.py
"""

import json
import os
from pathlib import Path

def main():
    try:
        from datasets import load_dataset
    except ImportError:
        print("Please install datasets: pip install datasets")
        return

    print("Loading HotpotQA dataset from HuggingFace...")
    dataset = load_dataset("hotpot_qa", "distractor", split="validation")
    
    # Take first 1000 examples
    subset = dataset.select(range(min(1000, len(dataset))))
    
    docs_dir = Path(__file__).parent / "corpus"
    docs_dir.mkdir(exist_ok=True)
    
    queries = []
    doc_id_counter = 0
    doc_id_map = {}  # title -> doc_id
    
    print(f"Processing {len(subset)} examples...")
    
    for idx, example in enumerate(subset):
        question = example["question"]
        answer = example["answer"]
        supporting_facts = example["supporting_facts"]
        context = example["context"]
        
        # Extract supporting fact titles (ground truth docs)
        supporting_titles = set(supporting_facts["title"])
        
        # Process context paragraphs
        relevant_doc_ids = []
        
        for title, sentences in zip(context["title"], context["sentences"]):
            # Create doc ID from title
            if title not in doc_id_map:
                doc_id = f"doc_{doc_id_counter:04d}"
                doc_id_counter += 1
                doc_id_map[title] = doc_id
                
                # Write document
                content = f"# {title}\n\n" + " ".join(sentences)
                doc_path = docs_dir / f"{doc_id}.txt"
                doc_path.write_text(content)
            
            # Track if this is a supporting document
            if title in supporting_titles:
                relevant_doc_ids.append(doc_id_map[title] + ".txt")
        
        # Create query entry
        queries.append({
            "id": f"q{idx:04d}",
            "text": question,
            "relevant_docs": relevant_doc_ids,
            "answer": answer,
            "type": example.get("type", "unknown")
        })
        
        if (idx + 1) % 100 == 0:
            print(f"  Processed {idx + 1}/{len(subset)}")
    
    # Write queries.json
    queries_path = Path(__file__).parent / "queries.json"
    with open(queries_path, "w") as f:
        json.dump({"queries": queries}, f, indent=2)
    
    print(f"\nDone!")
    print(f"  Documents: {doc_id_counter} files in {docs_dir}")
    print(f"  Queries: {len(queries)} in {queries_path}")

if __name__ == "__main__":
    main()

