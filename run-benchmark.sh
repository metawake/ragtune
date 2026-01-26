#!/bin/bash
set -e

# Benchmark: Qdrant vs pgvector
# Runs 3 iterations each for variance measurement

PGVECTOR_URL="postgres://postgres:test@localhost:5432/postgres"

mkdir -p runs/qdrant-casehold runs/pgvector-casehold runs/qdrant-hotpot runs/pgvector-hotpot

echo "============================================"
echo "Starting benchmark: $(date)"
echo "============================================"

# --- HotpotQA 5K: Qdrant (3 runs) ---
for i in 1 2 3; do
  echo ""
  echo "=== Qdrant HotpotQA run $i/3 ==="
  ./ragtune simulate \
    --collection hotpot5k-qdrant \
    --store qdrant \
    --embedder ollama \
    --queries ./benchmarks/hotpotqa-5k/queries.json \
    --output runs/qdrant-hotpot
  sleep 2
done

# --- HotpotQA 5K: pgvector (3 runs) ---
for i in 1 2 3; do
  echo ""
  echo "=== pgvector HotpotQA run $i/3 ==="
  ./ragtune simulate \
    --collection hotpot5k-pgvector \
    --store pgvector \
    --pgvector-url "$PGVECTOR_URL" \
    --embedder ollama \
    --queries ./benchmarks/hotpotqa-5k/queries.json \
    --output runs/pgvector-hotpot
  sleep 2
done

# --- CaseHOLD 500: Qdrant (3 runs) ---
for i in 1 2 3; do
  echo ""
  echo "=== Qdrant CaseHOLD run $i/3 ==="
  ./ragtune simulate \
    --collection casehold-qdrant \
    --store qdrant \
    --embedder ollama \
    --queries ./benchmarks/casehold-500/queries.json \
    --output runs/qdrant-casehold
  sleep 2
done

# --- CaseHOLD 500: pgvector (3 runs) ---
for i in 1 2 3; do
  echo ""
  echo "=== pgvector CaseHOLD run $i/3 ==="
  ./ragtune simulate \
    --collection casehold-pgvector \
    --store pgvector \
    --pgvector-url "$PGVECTOR_URL" \
    --embedder ollama \
    --queries ./benchmarks/casehold-500/queries.json \
    --output runs/pgvector-casehold
  sleep 2
done

echo ""
echo "============================================"
echo "Benchmark complete: $(date)"
echo "============================================"
echo ""
echo "Results saved to:"
ls -la runs/*/
