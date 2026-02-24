// Package ctxcancel demonstrates Go's context package for cancellation,
// timeout, and propagation patterns.
package ctxcancel

import (
	"context"
	"time"
)

// WithCancelDemo creates a cancellable context, launches a goroutine that
// blocks on ctx.Done(), then cancels the context. It returns the context's
// error, which should be context.Canceled.
func WithCancelDemo() error {
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})

	go func() {
		<-ctx.Done()
		close(done)
	}()

	cancel()
	<-done
	return ctx.Err()
}

// WithTimeoutDemo creates a context with the specified timeout and waits
// for it to expire. It returns the context's error and the approximate
// elapsed duration.
func WithTimeoutDemo(timeout time.Duration) (time.Duration, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	start := time.Now()
	<-ctx.Done()
	return time.Since(start), ctx.Err()
}

// PropagateCancel creates a parent context with cancel, derives a child
// context from it, and cancels the parent. It returns both the parent's
// and child's errors, demonstrating that cancellation propagates downward.
func PropagateCancel() (parentErr, childErr error) {
	parent, cancelParent := context.WithCancel(context.Background())
	child, cancelChild := context.WithCancel(parent)
	defer cancelChild()

	cancelParent()

	// Wait for both contexts to be done.
	<-parent.Done()
	<-child.Done()

	return parent.Err(), child.Err()
}

// CleanupPattern demonstrates the idiomatic defer cancel() pattern. It
// creates a context with timeout and immediately defers cancel to ensure
// resources are released. The returned cleanup flag indicates that cancel
// was called via defer when the function returned.
func CleanupPattern() (cancelled bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer func() {
		cancel()
		cancelled = true
	}()

	// Simulate some work without waiting for the timeout.
	_ = ctx
	return false // overridden by deferred func
}
