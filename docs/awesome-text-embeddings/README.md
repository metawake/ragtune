# Awesome Text Embeddings [![Awesome](https://awesome.re/badge.svg)](https://awesome.re)

> A curated list of text embedding models and tools for evaluating them.

Text embeddings convert text into dense vectors for semantic search, retrieval, clustering, and classification. This list helps you choose the right embedding model for your use case.

## Quick Picks

Just want a recommendation? Start here:

| Use Case | Model | Why |
|----------|-------|-----|
| **Best overall (API)** | [text-embedding-3-large](https://platform.openai.com/docs/guides/embeddings) | Highest quality, 8k context, adjustable dims |
| **Best overall (open)** | [gte-large-en-v1.5](https://huggingface.co/Alibaba-NLP/gte-large-en-v1.5) | Top MTEB, 8k context, Apache 2.0 |
| **Best budget** | [text-embedding-3-small](https://platform.openai.com/docs/guides/embeddings) | $0.02/1M tokens, still good quality |
| **Best local/private** | [nomic-embed-text-v1.5](https://huggingface.co/nomic-ai/nomic-embed-text-v1.5) | GGUF available, runs on CPU, 8k context |
| **Best multilingual** | [multilingual-e5-large](https://huggingface.co/intfloat/multilingual-e5-large) | 100+ languages, MIT license |
| **Best for code** | [voyage-code-2](https://docs.voyageai.com/docs/embeddings) | Purpose-built, 16k context |

---

## Contents

- [Quick Picks](#quick-picks)
- [How to Choose](#how-to-choose)
- [General Purpose](#general-purpose)
  - [Open Source](#open-source)
  - [API Services](#api-services)
- [Specialized](#specialized)
  - [Multilingual](#multilingual)
  - [Code Embeddings](#code-embeddings)
  - [Long-Context](#long-context)
  - [Domain-Specific](#domain-specific)
- [Rerankers](#rerankers)
- [Benchmarks & Leaderboards](#benchmarks--leaderboards)
- [Tools & Evaluation](#tools--evaluation)
- [Resources](#resources)
  - [Papers](#papers)
  - [Tutorials](#tutorials)
- [Related Lists](#related-lists)

---

## How to Choose

| Question | Recommendation |
|----------|----------------|
| Need best quality, don't mind API costs? | OpenAI `text-embedding-3-large` or Cohere `embed-v3` |
| Want open source, good quality? | `gte-large-en-v1.5` or `bge-large-en-v1.5` |
| Need multilingual? | `multilingual-e5-large` or Cohere `embed-multilingual-v3` |
| Working with code? | `voyage-code-2` |
| Have very long documents? | `jina-embeddings-v2-base-en` (8k) or `nomic-embed-text-v1.5` (8k) |
| Running locally/edge? | `nomic-embed-text-v1.5` (GGUF available) |
| Need on-prem / data privacy? | Open source models only — see [Open Source](#open-source) section |

**Key tradeoffs:**

- **Dimensions**: Higher = more expressive but more storage/compute. 768-1024 is the sweet spot for most use cases.
- **Context length**: Most models max at 512 tokens; some go to 8k+. Longer = fewer chunks needed.
- **Open vs API**: Open = privacy, cost control, on-prem; API = simplicity, no infrastructure.
- **Quality vs speed**: Larger models score higher on benchmarks but have higher latency.

---

## General Purpose

### Open Source

| Model | Provider | Dims | Max Tokens | MTEB Avg | License | Notes |
|-------|----------|------|------------|----------|---------|-------|
| [stella-en-1.5B-v5](https://huggingface.co/NovaSearch/stella-en-1.5B-v5) | NovaSearch | 1024 | 512 | 66.9 | MIT | Largest, highest quality |
| [gte-large-en-v1.5](https://huggingface.co/Alibaba-NLP/gte-large-en-v1.5) | Alibaba | 1024 | 8192 | 65.4 | Apache 2.0 | Long context, top tier |
| [mxbai-embed-large-v1](https://huggingface.co/mixedbread-ai/mxbai-embed-large-v1) | Mixedbread | 1024 | 512 | 64.7 | Apache 2.0 | Strong MTEB performer |
| [snowflake-arctic-embed-l](https://huggingface.co/Snowflake/snowflake-arctic-embed-l) | Snowflake | 1024 | 512 | 64.5 | Apache 2.0 | Strong retrieval |
| [bge-large-en-v1.5](https://huggingface.co/BAAI/bge-large-en-v1.5) | BAAI | 1024 | 512 | 64.2 | MIT | Widely adopted, battle-tested |
| [gte-base-en-v1.5](https://huggingface.co/Alibaba-NLP/gte-base-en-v1.5) | Alibaba | 768 | 8192 | 64.1 | Apache 2.0 | Smaller + long context |
| [bge-base-en-v1.5](https://huggingface.co/BAAI/bge-base-en-v1.5) | BAAI | 768 | 512 | 63.5 | MIT | Good speed/quality balance |
| [nomic-embed-text-v1.5](https://huggingface.co/nomic-ai/nomic-embed-text-v1.5) | Nomic | 768 | 8192 | 62.3 | Apache 2.0 | Long context, GGUF for local |
| [e5-large-v2](https://huggingface.co/intfloat/e5-large-v2) | Microsoft | 1024 | 512 | 62.2 | MIT | Requires "query:" prefix |
| [e5-base-v2](https://huggingface.co/intfloat/e5-base-v2) | Microsoft | 768 | 512 | 61.5 | MIT | Smaller variant |

### API Services

| Model | Provider | Dims | Max Tokens | Pricing (per 1M tokens) | Notes |
|-------|----------|------|------------|-------------------------|-------|
| [text-embedding-3-large](https://platform.openai.com/docs/guides/embeddings) | OpenAI | 3072 | 8191 | $0.13 | Best quality, adjustable dims |
| [voyage-large-2](https://docs.voyageai.com/docs/embeddings) | Voyage AI | 1536 | 16000 | $0.12 | Longest context |
| [embed-english-v3.0](https://docs.cohere.com/docs/embeddings) | Cohere | 1024 | 512 | $0.10 | Strong retrieval |
| [embed-large-v1](https://www.mixedbread.ai/docs/embeddings) | Mixedbread | 1024 | 512 | $0.05 | Good quality/price |
| [embedding-001](https://cloud.google.com/vertex-ai/docs/generative-ai/embeddings/get-text-embeddings) | Google | 768 | 2048 | $0.025 | Vertex AI |
| [text-embedding-3-small](https://platform.openai.com/docs/guides/embeddings) | OpenAI | 1536 | 8191 | $0.02 | Best budget option |
| [jina-embeddings-v2-base-en](https://jina.ai/embeddings/) | Jina AI | 768 | 8192 | $0.02 | Open weights also available |

---

## Specialized

### Multilingual

| Model | Provider | Dims | Languages | Max Tokens | Notes |
|-------|----------|------|-----------|------------|-------|
| [bge-m3](https://huggingface.co/BAAI/bge-m3) | BAAI | 1024 | 100+ | 8192 | Hybrid dense+sparse, long context |
| [multilingual-e5-large](https://huggingface.co/intfloat/multilingual-e5-large) | Microsoft | 1024 | 100+ | 512 | Best open multilingual |
| [multilingual-e5-base](https://huggingface.co/intfloat/multilingual-e5-base) | Microsoft | 768 | 100+ | 512 | Smaller variant |
| [embed-multilingual-v3.0](https://docs.cohere.com/docs/embeddings) | Cohere | 1024 | 100+ | 512 | API, strong quality |
| [paraphrase-multilingual-mpnet-base-v2](https://huggingface.co/sentence-transformers/paraphrase-multilingual-mpnet-base-v2) | SBERT | 768 | 50+ | 512 | Sentence-transformers |

### Code Embeddings

| Model | Provider | Dims | Languages | Notes |
|-------|----------|------|-----------|-------|
| [voyage-code-2](https://docs.voyageai.com/docs/embeddings) | Voyage AI | 1536 | 20+ | Best code retrieval, 16k context |
| [StarEncoder](https://huggingface.co/bigcode/starencoder) | BigCode | 768 | 80+ | StarCoder-based, open source |
| [codebert-base](https://huggingface.co/microsoft/codebert-base) | Microsoft | 768 | 6 | Open source, smaller |
| [code-search-ada-002](https://platform.openai.com/docs/guides/embeddings) | OpenAI | 1536 | Multiple | Legacy but still used |

### Long-Context

Models supporting 4k+ tokens — useful for embedding full documents without chunking.

| Model | Provider | Dims | Max Tokens | Notes |
|-------|----------|------|------------|-------|
| [voyage-large-2](https://docs.voyageai.com/docs/embeddings) | Voyage AI | 1536 | 16000 | Longest context (API) |
| [gte-large-en-v1.5](https://huggingface.co/Alibaba-NLP/gte-large-en-v1.5) | Alibaba | 1024 | 8192 | Top quality (open) |
| [jina-embeddings-v2-base-en](https://huggingface.co/jinaai/jina-embeddings-v2-base-en) | Jina AI | 768 | 8192 | Open + API available |
| [nomic-embed-text-v1.5](https://huggingface.co/nomic-ai/nomic-embed-text-v1.5) | Nomic | 768 | 8192 | GGUF for local inference |
| [text-embedding-3-large](https://platform.openai.com/docs/guides/embeddings) | OpenAI | 3072 | 8191 | Adjustable dimensions |
| [bge-m3](https://huggingface.co/BAAI/bge-m3) | BAAI | 1024 | 8192 | Also multilingual |

### Domain-Specific

| Model | Provider | Domain | Dims | Notes |
|-------|----------|--------|------|-------|
| [legal-bert-base-uncased](https://huggingface.co/nlpaueb/legal-bert-base-uncased) | NLP@AUEb | Legal | 768 | Trained on legal corpora |
| [PubMedBERT](https://huggingface.co/microsoft/BiomedNLP-PubMedBERT-base-uncased-abstract) | Microsoft | Biomedical | 768 | PubMed abstracts |
| [SciBERT](https://huggingface.co/allenai/scibert_scivocab_uncased) | Allen AI | Scientific | 768 | Scientific papers |
| [finbert](https://huggingface.co/yiyanghkust/finbert-tone) | FinBERT | Finance | 768 | Financial sentiment |

---

## Rerankers

Rerankers improve retrieval quality by rescoring initial results. Use after embedding-based retrieval.

| Model | Provider | Type | Notes |
|-------|----------|------|-------|
| [rerank-english-v3.0](https://docs.cohere.com/docs/rerank) | Cohere | API | Production-ready, easy to integrate |
| [rerank-multilingual-v3.0](https://docs.cohere.com/docs/rerank) | Cohere | API | 100+ languages |
| [bge-reranker-v2-m3](https://huggingface.co/BAAI/bge-reranker-v2-m3) | BAAI | Open | Multilingual, pairs with BGE embeddings |
| [bge-reranker-large](https://huggingface.co/BAAI/bge-reranker-large) | BAAI | Open | English-focused, strong quality |
| [ms-marco-MiniLM-L-12-v2](https://huggingface.co/cross-encoder/ms-marco-MiniLM-L-12-v2) | SBERT | Open | Lightweight, fast |
| [jina-reranker-v2-base-multilingual](https://huggingface.co/jinaai/jina-reranker-v2-base-multilingual) | Jina AI | Open | 100+ languages, 1k context |
| [mxbai-rerank-large-v1](https://huggingface.co/mixedbread-ai/mxbai-rerank-large-v1) | Mixedbread | Open | Strong quality |

**When to use a reranker:**
- You have more than ~20 candidates from initial retrieval
- Quality matters more than latency
- Your embedding model's ranking isn't precise enough

---

## Benchmarks & Leaderboards

| Benchmark | What it measures | Best for | Link |
|-----------|------------------|----------|------|
| **MTEB** | 8 task types (retrieval, classification, clustering, etc.) across 58 datasets, 112 languages | Overall embedding quality comparison | [Leaderboard](https://huggingface.co/spaces/mteb/leaderboard) |
| **BEIR** | Zero-shot retrieval across 18 diverse datasets | Retrieval-focused evaluation | [GitHub](https://github.com/beir-cellar/beir) |
| **MIRACL** | Multilingual retrieval across 18 languages | Non-English retrieval | [GitHub](https://github.com/project-miracl/miracl) |
| **C-MTEB** | Chinese-specific embedding tasks | Chinese language models | [Leaderboard](https://huggingface.co/spaces/mteb/leaderboard) |

**Note:** MTEB scores are useful for comparison but don't always predict real-world performance. Test on your own data with tools like [ragtune](#tools--evaluation).

---

## Tools & Evaluation

### Benchmarking & Comparison

| Tool | Description | Link |
|------|-------------|------|
| **ragtune** | CLI for benchmarking RAG retrieval quality. Compare embedding models on your queries and documents. | [GitHub](https://github.com/user/ragtune) |
| **MTEB** | Official benchmark toolkit for evaluating embeddings on standard tasks | [GitHub](https://github.com/embeddings-benchmark/mteb) |
| **sentence-transformers** | Framework for using, comparing, and training embeddings | [GitHub](https://github.com/UKPLab/sentence-transformers) |
| **Embeddings Projector** | Visualize high-dimensional embeddings in 2D/3D | [TensorFlow](https://projector.tensorflow.org/) |

### Fine-tuning

| Tool | Description | Link |
|------|-------------|------|
| **sentence-transformers** | Training custom embedding models with contrastive learning | [Docs](https://www.sbert.net/docs/training/overview.html) |
| **FlagEmbedding** | BAAI's toolkit for fine-tuning BGE models | [GitHub](https://github.com/FlagOpen/FlagEmbedding) |
| **uniem** | Unified embedding model training framework | [GitHub](https://github.com/wangyuxinwhy/uniem) |

### Local Inference

| Tool | Description | Link |
|------|-------------|------|
| **FastEmbed** | Fast, lightweight embedding inference by Qdrant | [GitHub](https://github.com/qdrant/fastembed) |
| **Ollama** | Run embedding models locally (GGUF format) | [Ollama](https://ollama.ai/) |
| **llama.cpp** | C++ inference for quantized models | [GitHub](https://github.com/ggerganov/llama.cpp) |
| **TEI** | Hugging Face's Text Embeddings Inference server | [GitHub](https://github.com/huggingface/text-embeddings-inference) |

---

## Resources

### Papers

**Foundational:**
- [Sentence-BERT: Sentence Embeddings using Siamese BERT-Networks](https://arxiv.org/abs/1908.10084) (2019) — Started modern sentence embeddings
- [MTEB: Massive Text Embedding Benchmark](https://arxiv.org/abs/2210.07316) (2022) — The standard benchmark
- [Text and Code Embeddings by Contrastive Pre-Training](https://arxiv.org/abs/2201.10005) (2022) — OpenAI's approach

**Recent advances:**
- [Improving Text Embeddings with Large Language Models](https://arxiv.org/abs/2401.00368) (2024) — LLM-based embedding training (E5-mistral)
- [BGE M3-Embedding](https://arxiv.org/abs/2402.03216) (2024) — Multi-lingual, multi-functionality, multi-granularity
- [Matryoshka Representation Learning](https://arxiv.org/abs/2205.13147) (2022) — Flexible dimension embeddings
- [NV-Embed: Improved Techniques for Training LLMs as Generalist Embedding Models](https://arxiv.org/abs/2405.17428) (2024) — NVIDIA's approach
- [Jina Embeddings 2: 8192-Token General-Purpose Text Embeddings](https://arxiv.org/abs/2310.19923) (2023) — Long-context embeddings

**Understanding embeddings:**
- [Text Embeddings Reveal (Almost) As Much As Text](https://arxiv.org/abs/2310.06816) (2023) — Privacy implications
- [BEIR: A Heterogeneous Benchmark for Zero-shot Evaluation](https://arxiv.org/abs/2104.08663) (2021) — Retrieval benchmark

### Tutorials

- [Sentence-Transformers Documentation](https://www.sbert.net/) — Comprehensive embedding guide
- [Hugging Face NLP Course](https://huggingface.co/learn/nlp-course) — Includes embedding fundamentals
- [Choosing an Embedding Model](https://www.pinecone.io/learn/series/rag/embedding-models/) — Pinecone's practical guide
- [Cohere Embed Guide](https://docs.cohere.com/docs/embeddings) — Good API-focused tutorial

---

## Related Lists

For adjacent topics, see these curated lists:

- [awesome-vector-databases](https://github.com/dangkhoasdc/awesome-vector-database) — Vector storage and retrieval
- [awesome-rag](https://github.com/coree/awesome-rag) — Retrieval-augmented generation
- [awesome-semantic-search](https://github.com/Agrover112/awesome-semantic-search) — Semantic search resources
- [awesome-local-ai](https://github.com/janhq/awesome-local-ai) — Local AI inference

---

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on adding new models or tools.

---

## License

[![CC0](https://licensebuttons.net/p/zero/1.0/88x31.png)](https://creativecommons.org/publicdomain/zero/1.0/)

To the extent possible under law, the authors have waived all copyright and related rights to this work.
