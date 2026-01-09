// RagTune — EXPLAIN ANALYZE for RAG retrieval
// A CLI tool to inspect, explain, benchmark, and tune RAG retrieval layers.
package main

import (
	"os"

	"github.com/metawake/ragtune/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}



