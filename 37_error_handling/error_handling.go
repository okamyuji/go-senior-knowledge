// Package errorhandling demonstrates Go's error handling patterns including
// sentinel errors, error wrapping with %w, errors.Is, errors.As, and
// custom error types implementing the error interface.
package errorhandling

import (
	"errors"
	"fmt"
)

// Sentinel errors are package-level values compared by identity.
// Callers use errors.Is to check if a sentinel exists in an error chain.
var (
	ErrNotFound   = errors.New("not found")
	ErrPermission = errors.New("permission denied")
)

// SentinelError returns ErrNotFound when id is 0, ErrPermission when id
// is negative, and nil otherwise.
func SentinelError(id int) error {
	switch {
	case id == 0:
		return ErrNotFound
	case id < 0:
		return ErrPermission
	default:
		return nil
	}
}

// WrapError wraps the given error with additional context using fmt.Errorf
// and the %w verb. The original error remains accessible through the chain.
func WrapError(err error, context string) error {
	return fmt.Errorf("%s: %w", context, err)
}

// UnwrapError checks whether target exists anywhere in err's chain using
// errors.Is. This works even when the error has been wrapped multiple times.
func UnwrapError(err, target error) bool {
	return errors.Is(err, target)
}

// ValidationError is a custom error type that carries a field name and
// a description of the problem.
type ValidationError struct {
	Field   string
	Message string
}

// Error satisfies the error interface.
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed on %s: %s", e.Field, e.Message)
}

// NewValidationError creates a ValidationError for the given field.
func NewValidationError(field, message string) error {
	return &ValidationError{Field: field, Message: message}
}

// WrapValidationError creates a wrapped ValidationError so that callers
// can use errors.As to extract the underlying ValidationError from the chain.
func WrapValidationError(field, message, context string) error {
	ve := &ValidationError{Field: field, Message: message}
	return fmt.Errorf("%s: %w", context, ve)
}

// ErrorAs attempts to extract a *ValidationError from err's chain.
// It returns the extracted error and true on success, or nil and false.
func ErrorAs(err error) (*ValidationError, bool) {
	var ve *ValidationError
	if errors.As(err, &ve) {
		return ve, true
	}
	return nil, false
}
