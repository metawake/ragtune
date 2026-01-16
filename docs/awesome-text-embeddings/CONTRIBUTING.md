# Contributing to Awesome Text Embeddings

Thanks for your interest in contributing! This document explains what belongs in this list and how to add it.

## Scope

This list focuses on **text embedding models** and **tools for evaluating/comparing them**.

### ✅ In Scope

- **Embedding models** — open source or API-based models that convert text to vectors
- **Benchmarking tools** — tools that help evaluate or compare embedding quality
- **Fine-tuning frameworks** — tools specifically for training embedding models
- **Embedding visualization** — tools for exploring embedding spaces
- **Relevant papers** — foundational or influential research on text embeddings

### ❌ Out of Scope

These are great topics, but belong in other lists:

| Topic | Why not here | Better list |
|-------|--------------|-------------|
| Vector databases | Storage, not embeddings themselves | [awesome-vector-databases](https://github.com/dangkhoasdc/awesome-vector-database) |
| RAG frameworks | Orchestration, embeddings are one piece | [awesome-rag](https://github.com/coree/awesome-rag) |
| LLM inference | Text generation, not embeddings | [awesome-local-ai](https://github.com/janhq/awesome-local-ai) |
| General ML frameworks | Too broad | [awesome-machine-learning](https://github.com/josephmisiti/awesome-machine-learning) |

**Not sure?** Open an issue and ask — we're happy to discuss edge cases.

---

## Quality Guidelines

We aim for a **curated** list, not a comprehensive dump. Models should be:

### For embedding models:

- [ ] **Notable** — widely used, top MTEB performer, or fills a unique niche
- [ ] **Accessible** — available on Hugging Face, via API, or documented source
- [ ] **Documented** — has usage instructions or paper
- [ ] **Not a minor variant** — we don't need every fine-tuned version of the same base

### For tools:

- [ ] **Actively maintained** — commits within last 6 months
- [ ] **Documented** — has README with usage instructions
- [ ] **Focused on embeddings** — primary purpose relates to embedding models

---

## How to Add an Entry

### 1. Fork and create a branch

```bash
git checkout -b add-model-name
```

### 2. Add your entry in the appropriate section

Follow the existing table format:

**For open source models:**
```markdown
| [model-name](https://huggingface.co/org/model-name) | Provider | 1024 | 512 | 64.5 | Apache 2.0 | Brief note |
```

**For API services:**
```markdown
| [model-name](https://docs.provider.com/embeddings) | Provider | 1024 | 8192 | $0.10 | Brief note |
```

**For tools:**
```markdown
| **tool-name** | One-line description | [GitHub](https://github.com/org/repo) |
```

### 3. Keep entries sorted

- Models: alphabetically by name within each section
- Tools: alphabetically by name

### 4. Submit a pull request

Include:
- What you're adding
- Why it belongs (e.g., "Top 5 on MTEB retrieval", "Only open-source legal embedding")

---

## Entry Format Reference

### Required fields for models

| Field | Description | Example |
|-------|-------------|---------|
| Name | Linked to source | `[bge-large-en-v1.5](https://...)` |
| Provider | Organization | `BAAI` |
| Dims | Output dimensions | `1024` |
| Max Tokens | Context length | `512` |
| Quality metric | MTEB score or pricing | `64.2` or `$0.10` |
| License/Notes | License or key info | `MIT` |

### Optional but helpful

- Specific MTEB retrieval score (more relevant for RAG use cases)
- Quantization availability (GGUF, ONNX)
- Special requirements (e.g., "requires query: prefix")

---

## Updating Existing Entries

Found outdated info? PRs welcome for:

- Updated benchmark scores
- New model versions (replace old, don't add as separate entry)
- Corrected links or pricing
- Better descriptions

---

## Questions?

- **Not sure if something fits?** Open an issue
- **Want to suggest a new section?** Open an issue
- **Found a bug or broken link?** Open a PR

Thanks for helping make this list useful!
