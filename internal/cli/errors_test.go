package cli

import (
	"bytes"
	"errors"
	"io"
	"testing"
)

func TestAuditError_Error(t *testing.T) {
	err := &AuditError{FailCount: 3}
	if err.Error() != "audit failed" {
		t.Errorf("Error() = %q, want %q", err.Error(), "audit failed")
	}
}

func TestAuditError_Unwrap(t *testing.T) {
	err := &AuditError{FailCount: 2}
	if !errors.Is(err, ErrAuditFailed) {
		t.Error("expected AuditError to unwrap to ErrAuditFailed")
	}
}

func TestAuditError_FailCount(t *testing.T) {
	err := &AuditError{FailCount: 5}
	if err.FailCount != 5 {
		t.Errorf("FailCount = %d, want 5", err.FailCount)
	}
}

func TestCICheckError_Error(t *testing.T) {
	err := &CICheckError{FailedChecks: []string{"recall", "mrr"}}
	if err.Error() != "CI check failed" {
		t.Errorf("Error() = %q, want %q", err.Error(), "CI check failed")
	}
}

func TestCICheckError_Unwrap(t *testing.T) {
	err := &CICheckError{FailedChecks: []string{"coverage"}}
	if !errors.Is(err, ErrCICheckFailed) {
		t.Error("expected CICheckError to unwrap to ErrCICheckFailed")
	}
}

func TestCICheckError_FailedChecks(t *testing.T) {
	checks := []string{"recall", "mrr", "coverage"}
	err := &CICheckError{FailedChecks: checks}
	if len(err.FailedChecks) != 3 {
		t.Errorf("FailedChecks count = %d, want 3", len(err.FailedChecks))
	}
}

func TestSentinelErrors(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		sentinel error
	}{
		{"audit error", &AuditError{FailCount: 1}, ErrAuditFailed},
		{"ci check error", &CICheckError{FailedChecks: []string{"x"}}, ErrCICheckFailed},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !errors.Is(tt.err, tt.sentinel) {
				t.Errorf("errors.Is(%T, %v) = false, want true", tt.err, tt.sentinel)
			}
		})
	}
}

func TestErrorsAs(t *testing.T) {
	t.Run("AuditError", func(t *testing.T) {
		var wrapped error = &AuditError{FailCount: 7}
		var auditErr *AuditError
		if !errors.As(wrapped, &auditErr) {
			t.Error("errors.As failed for AuditError")
		}
		if auditErr.FailCount != 7 {
			t.Errorf("FailCount = %d, want 7", auditErr.FailCount)
		}
	})

	t.Run("CICheckError", func(t *testing.T) {
		var wrapped error = &CICheckError{FailedChecks: []string{"a", "b"}}
		var ciErr *CICheckError
		if !errors.As(wrapped, &ciErr) {
			t.Error("errors.As failed for CICheckError")
		}
		if len(ciErr.FailedChecks) != 2 {
			t.Errorf("FailedChecks count = %d, want 2", len(ciErr.FailedChecks))
		}
	})
}

// mockCloser implements io.Closer for testing closeWithLog
type mockCloser struct {
	err error
}

func (m *mockCloser) Close() error {
	return m.err
}

func TestCloseWithLog_Success(t *testing.T) {
	mc := &mockCloser{err: nil}
	// Should not panic or produce output
	closeWithLog(mc, "test resource")
}

func TestCloseWithLog_Error(t *testing.T) {
	mc := &mockCloser{err: io.ErrClosedPipe}
	// Just verify it doesn't panic - we can't easily capture stderr without more setup
	closeWithLog(mc, "test resource")
}

// captureStderr helper for testing stderr output
func captureStderr(t *testing.T, fn func()) string {
	t.Helper()
	// Note: This is a simplified version; full stderr capture would require
	// redirecting os.Stderr which can be fragile in tests
	var buf bytes.Buffer
	// For a complete test, you'd capture os.Stderr here
	fn()
	return buf.String()
}
