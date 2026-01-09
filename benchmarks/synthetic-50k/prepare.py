#!/usr/bin/env python3
"""
Prepare synthetic 50K benchmark dataset for RagTune enterprise-scale testing.

Generates:
- 50,000 synthetic documents (realistic enterprise content)
- 500 benchmark queries with known relevant documents
- Timing/throughput metrics

This benchmark tests:
- Embedding throughput at scale (target: <5 min for 50k docs)
- Batching efficiency across concurrency levels
- Memory stability under load

Usage:
    python prepare.py [--docs 50000] [--queries 500]
    
Example:
    python prepare.py                    # Default 50k docs, 500 queries
    python prepare.py --docs 10000       # Quick test with 10k docs
    python prepare.py --docs 100000      # Stress test with 100k docs
"""

import json
import os
import random
import hashlib
import argparse
from pathlib import Path
from datetime import datetime, timedelta
import time

# Seed for reproducibility
random.seed(42)

# =============================================================================
# Content Templates - Realistic enterprise document types
# =============================================================================

DEPARTMENTS = [
    "Engineering", "Product", "Sales", "Marketing", "Finance",
    "Legal", "HR", "Operations", "Customer Success", "Security",
    "Data Science", "DevOps", "QA", "Design", "Research"
]

DOCUMENT_TYPES = [
    "policy", "procedure", "guide", "specification", "report",
    "memo", "proposal", "review", "analysis", "summary"
]

TOPICS = [
    "authentication", "authorization", "data processing", "API integration",
    "performance optimization", "security compliance", "user onboarding",
    "billing procedures", "incident response", "change management",
    "deployment pipeline", "code review", "testing standards", "monitoring",
    "disaster recovery", "data retention", "access control", "encryption",
    "audit logging", "rate limiting", "caching strategy", "database migration",
    "service mesh", "container orchestration", "CI/CD automation",
    "load balancing", "fault tolerance", "backup procedures", "SLA management",
    "vendor evaluation", "cost optimization", "capacity planning", "scaling",
    "microservices", "API versioning", "schema evolution", "event sourcing"
]

TECH_TERMS = [
    "Kubernetes", "Docker", "PostgreSQL", "Redis", "Kafka", "gRPC",
    "REST API", "GraphQL", "OAuth 2.0", "JWT", "TLS", "mTLS",
    "Prometheus", "Grafana", "ELK Stack", "Terraform", "Ansible",
    "GitHub Actions", "Jenkins", "ArgoCD", "Istio", "Envoy",
    "S3", "CloudFront", "Lambda", "DynamoDB", "RDS", "EC2",
    "VPC", "IAM", "KMS", "CloudWatch", "SNS", "SQS"
]

ACTIONS = [
    "configure", "implement", "deploy", "monitor", "optimize",
    "troubleshoot", "validate", "migrate", "upgrade", "secure",
    "audit", "document", "review", "test", "benchmark"
]

REQUIREMENTS = [
    "must be completed within 24 hours",
    "requires manager approval",
    "should follow security guidelines",
    "needs to be documented in Confluence",
    "must pass QA review",
    "requires load testing",
    "should be backwards compatible",
    "needs stakeholder sign-off",
    "must meet SLA requirements",
    "requires security review"
]


def generate_document_id() -> str:
    """Generate a unique document ID."""
    return hashlib.md5(str(random.random()).encode()).hexdigest()[:12]


def generate_paragraph(topic: str, min_sentences: int = 3, max_sentences: int = 8) -> str:
    """Generate a realistic paragraph about a topic."""
    templates = [
        f"The {topic} system is designed to handle enterprise-scale workloads efficiently. "
        f"It integrates with {random.choice(TECH_TERMS)} and {random.choice(TECH_TERMS)} for optimal performance. "
        f"Teams should {random.choice(ACTIONS)} the configuration according to their specific requirements.",
        
        f"When working with {topic}, it's important to consider scalability and reliability. "
        f"The recommended approach involves using {random.choice(TECH_TERMS)} in conjunction with {random.choice(TECH_TERMS)}. "
        f"This {random.choice(REQUIREMENTS)}.",
        
        f"Best practices for {topic} include regular monitoring and proactive maintenance. "
        f"The {random.choice(DEPARTMENTS)} team is responsible for ensuring compliance with internal standards. "
        f"All changes {random.choice(REQUIREMENTS)}.",
        
        f"The implementation of {topic} follows our standard {random.choice(DOCUMENT_TYPES)} guidelines. "
        f"Key technologies include {random.choice(TECH_TERMS)}, {random.choice(TECH_TERMS)}, and {random.choice(TECH_TERMS)}. "
        f"Performance benchmarks show consistent throughput under load.",
        
        f"For {topic} operations, the system leverages {random.choice(TECH_TERMS)} as the primary infrastructure. "
        f"The {random.choice(DEPARTMENTS)} team maintains documentation and provides support. "
        f"Emergency procedures are outlined in the incident response playbook.",
    ]
    
    paragraphs = []
    num_paragraphs = random.randint(min_sentences, max_sentences) // 3 + 1
    
    for _ in range(num_paragraphs):
        paragraphs.append(random.choice(templates))
    
    return "\n\n".join(paragraphs)


def generate_document(doc_id: str, topic: str, doc_type: str, department: str) -> dict:
    """Generate a complete synthetic document."""
    
    title = f"{department} {doc_type.title()}: {topic.replace('_', ' ').title()}"
    
    # Header section
    header = f"""# {title}

**Document ID:** DOC-{doc_id.upper()}
**Department:** {department}
**Type:** {doc_type.title()}
**Last Updated:** {datetime.now() - timedelta(days=random.randint(1, 365))}
**Owner:** {department} Team

---

## Overview

This document describes the {topic} procedures and guidelines for the {department} department.
It covers implementation details, best practices, and operational requirements.
"""
    
    # Body sections
    sections = []
    
    # Purpose section
    sections.append(f"""## Purpose

{generate_paragraph(topic, 2, 4)}
""")
    
    # Scope section
    sections.append(f"""## Scope

This {doc_type} applies to all {department} personnel and systems involved in {topic} operations.
The guidelines outlined here {random.choice(REQUIREMENTS)}.

Key stakeholders include:
- {random.choice(DEPARTMENTS)} Team
- {random.choice(DEPARTMENTS)} Team  
- {random.choice(DEPARTMENTS)} Team
""")
    
    # Technical details
    sections.append(f"""## Technical Requirements

{generate_paragraph(topic, 4, 8)}

### Infrastructure Components

The following components are required for {topic}:

1. **{random.choice(TECH_TERMS)}** - Primary processing layer
2. **{random.choice(TECH_TERMS)}** - Data persistence
3. **{random.choice(TECH_TERMS)}** - Monitoring and observability
4. **{random.choice(TECH_TERMS)}** - Security and access control
""")
    
    # Procedures
    sections.append(f"""## Procedures

### Standard Operations

To {random.choice(ACTIONS)} the {topic} system:

1. Verify prerequisites and access permissions
2. Review current configuration in {random.choice(TECH_TERMS)}
3. Execute the {random.choice(ACTIONS)} procedure
4. Validate results using monitoring dashboards
5. Document changes according to policy

{generate_paragraph(topic, 2, 4)}
""")
    
    # Compliance
    sections.append(f"""## Compliance and Security

All {topic} operations must comply with:

- SOC 2 Type II requirements
- Internal security policies
- Data retention guidelines
- Access control standards

{generate_paragraph(topic, 2, 3)}
""")
    
    content = header + "\n".join(sections)
    
    return {
        "id": doc_id,
        "title": title,
        "topic": topic,
        "doc_type": doc_type,
        "department": department,
        "content": content
    }


def generate_query(doc: dict, related_docs: list) -> dict:
    """Generate a query that targets specific documents."""
    
    query_templates = [
        f"How do I {random.choice(ACTIONS)} {doc['topic']} in the {doc['department']} department?",
        f"What are the requirements for {doc['topic']} according to {doc['department']} guidelines?",
        f"Where can I find the {doc['doc_type']} for {doc['topic']}?",
        f"What is the procedure for {doc['topic']} operations?",
        f"Who is responsible for {doc['topic']} in {doc['department']}?",
        f"What technologies are used for {doc['topic']}?",
        f"How should I handle {doc['topic']} compliance requirements?",
        f"What are the best practices for {doc['topic']}?",
    ]
    
    return {
        "id": f"q{doc['id']}",
        "text": random.choice(query_templates),
        "answer": f"See {doc['title']}",
        "relevant_docs": [f"doc_{doc['id']}.txt"] + [f"doc_{d['id']}.txt" for d in related_docs[:2]],
        "primary_topic": doc['topic'],
        "department": doc['department']
    }


def main():
    parser = argparse.ArgumentParser(description="Generate synthetic benchmark dataset")
    parser.add_argument("--docs", type=int, default=50000, help="Number of documents to generate")
    parser.add_argument("--queries", type=int, default=500, help="Number of queries to generate")
    args = parser.parse_args()
    
    num_docs = args.docs
    num_queries = args.queries
    
    print(f"""
╔══════════════════════════════════════════════════════════════╗
║  RagTune Synthetic Benchmark Generator                       ║
║  Enterprise-Scale Performance Testing                        ║
╠══════════════════════════════════════════════════════════════╣
║  Target Documents: {num_docs:>6,}                                    ║
║  Target Queries:   {num_queries:>6,}                                    ║
╚══════════════════════════════════════════════════════════════╝
""")
    
    start_time = time.time()
    
    # Create output directory
    base_dir = Path(__file__).parent
    corpus_dir = base_dir / "corpus"
    corpus_dir.mkdir(exist_ok=True)
    
    # Clear existing documents
    print("Clearing existing corpus...")
    for f in corpus_dir.glob("*.txt"):
        f.unlink()
    
    # Generate documents
    print(f"Generating {num_docs:,} documents...")
    documents = []
    docs_by_topic = {}
    
    for i in range(num_docs):
        doc_id = generate_document_id()
        topic = random.choice(TOPICS)
        doc_type = random.choice(DOCUMENT_TYPES)
        department = random.choice(DEPARTMENTS)
        
        doc = generate_document(doc_id, topic, doc_type, department)
        documents.append(doc)
        
        # Index by topic for query generation
        if topic not in docs_by_topic:
            docs_by_topic[topic] = []
        docs_by_topic[topic].append(doc)
        
        # Progress update
        if (i + 1) % 5000 == 0:
            elapsed = time.time() - start_time
            rate = (i + 1) / elapsed
            eta = (num_docs - i - 1) / rate
            print(f"  Generated {i + 1:,} documents ({rate:.0f} docs/sec, ETA: {eta:.0f}s)")
    
    # Write documents
    print("Writing documents to corpus/...")
    write_start = time.time()
    
    for doc in documents:
        filepath = corpus_dir / f"doc_{doc['id']}.txt"
        filepath.write_text(doc["content"])
    
    write_time = time.time() - write_start
    print(f"  Wrote {num_docs:,} files in {write_time:.1f}s ({num_docs/write_time:.0f} files/sec)")
    
    # Generate queries
    print(f"Generating {num_queries:,} queries...")
    queries = []
    
    # Select documents for queries, preferring topics with multiple docs
    query_docs = []
    for topic, topic_docs in docs_by_topic.items():
        if len(topic_docs) >= 3:
            query_docs.extend(topic_docs[:3])
    
    random.shuffle(query_docs)
    query_docs = query_docs[:num_queries]
    
    # If not enough, fill with random docs
    if len(query_docs) < num_queries:
        remaining = random.sample(documents, min(num_queries - len(query_docs), len(documents)))
        query_docs.extend(remaining)
    
    for doc in query_docs[:num_queries]:
        # Find related docs (same topic)
        related = [d for d in docs_by_topic.get(doc["topic"], []) if d["id"] != doc["id"]]
        query = generate_query(doc, related)
        queries.append(query)
    
    # Write queries
    queries_file = base_dir / "queries.json"
    with open(queries_file, "w") as f:
        json.dump({"queries": queries}, f, indent=2)
    
    # Write configs
    configs_file = base_dir / "configs.yaml"
    configs_file.write_text("""# Benchmark configurations for enterprise-scale testing
# Tests various top_k values to measure recall at different retrieval depths

configs:
  - name: top3
    top_k: 3
    description: "Minimal retrieval - tests precision"

  - name: top5
    top_k: 5
    description: "Standard retrieval depth"

  - name: top10
    top_k: 10
    description: "Extended retrieval for complex queries"

  - name: top20
    top_k: 20
    description: "Deep retrieval for multi-hop reasoning"

  - name: top50
    top_k: 50
    description: "Maximum retrieval for exhaustive search"
""")
    
    # Calculate stats
    total_time = time.time() - start_time
    total_chars = sum(len(doc["content"]) for doc in documents)
    avg_doc_size = total_chars / num_docs
    
    # Write README
    readme_file = base_dir / "README.md"
    readme_file.write_text(f"""# Synthetic 50K Benchmark

Enterprise-scale benchmark for testing RagTune embedding throughput and batching efficiency.

## Dataset Statistics

| Metric | Value |
|--------|-------|
| Documents | {num_docs:,} |
| Queries | {num_queries:,} |
| Total Size | {total_chars / 1024 / 1024:.1f} MB |
| Avg Doc Size | {avg_doc_size:.0f} chars |
| Generation Time | {total_time:.1f}s |

## Document Types

The corpus contains synthetic enterprise documents across {len(DEPARTMENTS)} departments:
{', '.join(DEPARTMENTS)}

Document types include: {', '.join(DOCUMENT_TYPES)}

## Performance Targets

For enterprise viability, embedding 50k documents should complete in:

| Embedder | Target Time | Docs/Second |
|----------|-------------|-------------|
| TEI (GPU) | < 2 min | ~400/s |
| TEI (CPU) | < 10 min | ~80/s |
| Ollama (8 concurrent) | < 5 min | ~160/s |
| Ollama (sequential) | < 30 min | ~28/s |

## Usage

```bash
# Generate the benchmark dataset
python prepare.py

# Quick test with fewer documents
python prepare.py --docs 1000 --queries 50

# Run the benchmark
ragtune ingest ./benchmarks/synthetic-50k/corpus --collection synthetic-50k
ragtune simulate --collection synthetic-50k --queries ./benchmarks/synthetic-50k/queries.json
```

## Benchmark Results

Record your results here:

| Date | Embedder | Concurrency | Time | Docs/Sec | Notes |
|------|----------|-------------|------|----------|-------|
| | | | | | |

---

*Generated on {datetime.now().isoformat()}*
""")
    
    # Final summary
    print(f"""
╔══════════════════════════════════════════════════════════════╗
║  ✓ Benchmark Generation Complete                             ║
╠══════════════════════════════════════════════════════════════╣
║  Documents:     {num_docs:>8,}                                      ║
║  Queries:       {num_queries:>8,}                                      ║
║  Total Size:    {total_chars / 1024 / 1024:>8.1f} MB                                   ║
║  Avg Doc Size:  {avg_doc_size:>8.0f} chars                                ║
║  Generation:    {total_time:>8.1f} seconds                              ║
╚══════════════════════════════════════════════════════════════╝

Files created:
  - {corpus_dir}/ ({num_docs:,} documents)
  - {queries_file}
  - {configs_file}
  - {readme_file}

Next steps:
  1. Run the embedding benchmark:
     INTEGRATION_TEST=1 go test -v -run TestScale ./internal/embedder/
  
  2. Or ingest with ragtune:
     ragtune ingest ./benchmarks/synthetic-50k/corpus --collection synthetic-50k
""")


if __name__ == "__main__":
    main()



