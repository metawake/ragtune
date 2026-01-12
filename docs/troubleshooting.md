# Troubleshooting

Common issues and solutions when using RagTune.

## Connection Issues

### "Connection refused" to Qdrant

Qdrant isn't running. Start it:

```bash
docker run -d -p 6333:6333 -p 6334:6334 --name qdrant qdrant/qdrant
```

Or restart if it exists:

```bash
docker start qdrant
```

Verify it's running:

```bash
curl http://localhost:6333/collections
# Should return: {"result":{"collections":[]},"status":"ok",...}
```

### "Is Ollama running?"

Start the Ollama server and pull an embedding model:

```bash
# Terminal 1: Start Ollama
ollama serve

# Terminal 2: Pull embedding model
ollama pull nomic-embed-text
```

Verify it's running:

```bash
curl http://localhost:11434/api/tags
```

### "OPENAI_API_KEY not set"

Either set the key:

```bash
export OPENAI_API_KEY="sk-..."
```

Or use Ollama instead (no key needed):

```bash
ragtune ingest ./docs --collection test --embedder ollama
```

---

## Performance Issues

### Slow ingestion

1. **Use TEI** instead of Ollama (4x faster):
   ```bash
   docker run -p 8080:8080 ghcr.io/huggingface/text-embeddings-inference:cpu-1.2 \
     --model-id BAAI/bge-base-en-v1.5
   ragtune ingest ./docs --collection test --embedder tei
   ```

2. **Increase chunk size** (fewer chunks = fewer embedding calls):
   ```bash
   ragtune ingest ./docs --collection test --chunk-size 1024
   ```

3. **Use GPU** if available (10x faster for TEI):
   ```bash
   docker run --gpus all -p 8080:8080 \
     ghcr.io/huggingface/text-embeddings-inference:latest \
     --model-id BAAI/bge-base-en-v1.5
   ```

### High latency in queries

1. Check vector store is running locally (network latency for remote stores)
2. Use smaller `--top-k` value
3. Consider using a faster embedder for queries

---

## Retrieval Quality Issues

### Low recall scores

1. **Check chunk size**: Try larger chunks
   ```bash
   ragtune ingest ./docs --collection prod-1024 --chunk-size 1024
   ```

2. **Try different embedder**: Domain-specific embedders help
   ```bash
   # For legal text
   ragtune ingest ./docs --collection legal --embedder voyage --voyage-model voyage-law-2
   
   # For code
   ragtune ingest ./docs --collection code --embedder voyage --voyage-model voyage-code-2
   ```

3. **Verify relevant_docs paths**: Must match exactly (case-sensitive)
   ```bash
   # Check what paths are in your collection
   ragtune explain "test query" --collection prod
   # Compare Source: paths with your queries.json relevant_docs
   ```

4. **Inspect with explain**: See exactly what's being retrieved
   ```bash
   ragtune explain "your failing query" --collection test
   ```

### All scores are similar (no discrimination)

This usually means:
- Query is too vague or generic
- Documents are too similar to each other
- Chunk size is too large (everything becomes similar)

Try:
1. More specific queries
2. Smaller chunk sizes
3. Different embedder

### Right document not retrieved at all

1. **Verify document was ingested**:
   ```bash
   ragtune explain "keyword from document" --collection prod
   ```

2. **Check document format**: Ensure it's readable (not binary, not corrupted)

3. **Re-ingest with verbose output**:
   ```bash
   ragtune ingest ./docs --collection prod --embedder ollama
   # Check "Found X documents" count
   ```

---

## Embedding Issues

### Embedding dimension mismatch

If you get dimension errors, the collection was created with a different embedder.

**Solution**: Delete and recreate with the correct embedder:

```bash
# The collection must be recreated with matching embedder
ragtune ingest ./docs --collection new-name --embedder ollama
```

To delete a collection in Qdrant:
```bash
curl -X DELETE http://localhost:6333/collections/old-collection
```

### "Model not found" with Ollama

Pull the embedding model:

```bash
ollama pull nomic-embed-text
```

For other models:
```bash
ollama pull mxbai-embed-large
```

---

## Query File Issues

### "Invalid queries file"

Ensure your JSON is valid:

```json
{
  "queries": [
    {
      "id": "q1",
      "text": "How do I reset my password?",
      "relevant_docs": ["docs/auth/password.md"]
    }
  ]
}
```

Common mistakes:
- Missing `"queries"` wrapper array
- Trailing commas
- Unquoted strings
- Wrong field names (`query` instead of `text`)

### "No relevant documents found"

Your `relevant_docs` paths don't match any ingested documents.

1. Run `explain` to see actual paths:
   ```bash
   ragtune explain "test" --collection prod
   # Note the Source: field paths
   ```

2. Update your queries.json to use matching paths

---

## CI/CD Issues

### Exit code 1 but metrics look fine

Check the exact threshold that failed:

```bash
ragtune simulate --collection prod --queries queries.json \
  --ci --min-recall 0.85 --min-coverage 0.90 --max-latency-p95 500
```

The output will show which threshold failed:
- `FAIL: Recall@5 = 0.82 < 0.85`

### Ingestion fails in GitHub Actions

Common causes:

1. **Qdrant not ready**: Add a health check wait
   ```yaml
   - name: Wait for Qdrant
     run: |
       for i in {1..30}; do
         curl -s http://localhost:6333/collections && break
         sleep 1
       done
   ```

2. **Ollama model not pulled**: Pull before ingest
   ```yaml
   - name: Setup Ollama
     run: |
       curl -fsSL https://ollama.com/install.sh | sh
       ollama serve &
       sleep 5
       ollama pull nomic-embed-text
   ```

---

## Getting Help

If your issue isn't covered here:

1. Run with verbose output and capture the error
2. Check [GitHub Issues](https://github.com/metawake/ragtune/issues)
3. Open a new issue with:
   - RagTune version (`ragtune --version`)
   - Command you ran
   - Full error output
   - Your OS and Docker version
