// Package cli implements the command-line interface for RagTune.
package cli

import "errors"

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
