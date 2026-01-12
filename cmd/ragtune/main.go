// RagTune — EXPLAIN ANALYZE for RAG retrieval
// A CLI tool to inspect, explain, benchmark, and tune RAG retrieval layers.
package main

import (
	"errors"
	"os"

	"github.com/metawake/ragtune/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		// Exit code 1 for audit/CI failures (expected failures)
		// Exit code 2 for unexpected errors
		if errors.Is(err, cli.ErrAuditFailed) || errors.Is(err, cli.ErrCICheckFailed) {
			os.Exit(1)
		}
		os.Exit(2)
	}
}



