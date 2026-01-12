// Package cli implements the command-line interface for RagTune.
// It provides commands for ingesting documents, explaining retrieval results,
// running simulations, auditing RAG quality, and comparing configurations.
package cli

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// Sentinel errors for CLI operations.
// These allow callers to check specific error conditions with errors.Is().
var (
	// ErrAuditFailed indicates one or more audit thresholds were not met.
	ErrAuditFailed = errors.New("audit failed")

	// ErrCICheckFailed indicates CI threshold checks did not pass.
	ErrCICheckFailed = errors.New("CI check failed")

	// ErrValidation indicates invalid input parameters.
	ErrValidation = errors.New("validation error")
)

// AuditError provides details about audit failures.
type AuditError struct {
	FailCount int
}

func (e *AuditError) Error() string {
	return "audit failed"
}

func (e *AuditError) Unwrap() error {
	return ErrAuditFailed
}

// CICheckError provides details about CI check failures.
type CICheckError struct {
	FailedChecks []string
}

func (e *CICheckError) Error() string {
	return "CI check failed"
}

func (e *CICheckError) Unwrap() error {
	return ErrCICheckFailed
}

// closeWithLog closes a resource and logs any error to stderr.
// Use this with defer for resources where Close() errors should be noted but not fatal.
func closeWithLog(c io.Closer, name string) {
	if err := c.Close(); err != nil {
		fmt.Fprintf(os.Stderr, "warning: failed to close %s: %v\n", name, err)
	}
}
