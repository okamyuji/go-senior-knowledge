package errorhandling

import (
	"errors"
	"testing"
)

func TestSentinelErrorNotFound(t *testing.T) {
	err := SentinelError(0)
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("SentinelError(0) = %v, want ErrNotFound", err)
	}
}

func TestSentinelErrorPermission(t *testing.T) {
	err := SentinelError(-1)
	if !errors.Is(err, ErrPermission) {
		t.Errorf("SentinelError(-1) = %v, want ErrPermission", err)
	}
}

func TestSentinelErrorNil(t *testing.T) {
	err := SentinelError(1)
	if err != nil {
		t.Errorf("SentinelError(1) = %v, want nil", err)
	}
}

func TestWrapErrorPreservesChain(t *testing.T) {
	wrapped := WrapError(ErrNotFound, "fetch user")
	if !errors.Is(wrapped, ErrNotFound) {
		t.Error("wrapped error does not contain ErrNotFound in chain")
	}
	want := "fetch user: not found"
	if got := wrapped.Error(); got != want {
		t.Errorf("wrapped.Error() = %q, want %q", got, want)
	}
}

func TestWrapErrorDoubleWrap(t *testing.T) {
	first := WrapError(ErrPermission, "read file")
	second := WrapError(first, "open config")
	if !errors.Is(second, ErrPermission) {
		t.Error("double-wrapped error does not contain ErrPermission")
	}
	want := "open config: read file: permission denied"
	if got := second.Error(); got != want {
		t.Errorf("second.Error() = %q, want %q", got, want)
	}
}

func TestUnwrapErrorMatch(t *testing.T) {
	wrapped := WrapError(ErrNotFound, "lookup")
	if !UnwrapError(wrapped, ErrNotFound) {
		t.Error("UnwrapError should find ErrNotFound in chain")
	}
}

func TestUnwrapErrorNoMatch(t *testing.T) {
	wrapped := WrapError(ErrNotFound, "lookup")
	if UnwrapError(wrapped, ErrPermission) {
		t.Error("UnwrapError should not find ErrPermission in ErrNotFound chain")
	}
}

func TestErrorAsExtractsValidationError(t *testing.T) {
	err := WrapValidationError("email", "invalid format", "create user")
	ve, ok := ErrorAs(err)
	if !ok {
		t.Fatal("ErrorAs should extract ValidationError from chain")
	}
	if ve.Field != "email" {
		t.Errorf("ve.Field = %q, want %q", ve.Field, "email")
	}
	if ve.Message != "invalid format" {
		t.Errorf("ve.Message = %q, want %q", ve.Message, "invalid format")
	}
}

func TestErrorAsReturnsFalseForNonMatch(t *testing.T) {
	err := WrapError(ErrNotFound, "lookup")
	ve, ok := ErrorAs(err)
	if ok {
		t.Error("ErrorAs should return false for non-ValidationError chain")
	}
	if ve != nil {
		t.Error("ErrorAs should return nil for non-ValidationError chain")
	}
}

func TestCustomErrorMessage(t *testing.T) {
	err := NewValidationError("age", "must be positive")
	want := "validation failed on age: must be positive"
	if got := err.Error(); got != want {
		t.Errorf("err.Error() = %q, want %q", got, want)
	}
}
