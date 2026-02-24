package ctxcancel

import (
	"context"
	"testing"
	"time"
)

func TestWithCancelDemo(t *testing.T) {
	err := WithCancelDemo()
	if err != context.Canceled {
		t.Errorf("WithCancelDemo() = %v, want context.Canceled", err)
	}
}

func TestWithTimeoutDemo(t *testing.T) {
	timeout := 50 * time.Millisecond
	elapsed, err := WithTimeoutDemo(timeout)

	if err != context.DeadlineExceeded {
		t.Errorf("WithTimeoutDemo error = %v, want context.DeadlineExceeded", err)
	}
	if elapsed < 40*time.Millisecond {
		t.Errorf("WithTimeoutDemo elapsed = %v, expected >= 40ms", elapsed)
	}
	if elapsed > 200*time.Millisecond {
		t.Errorf("WithTimeoutDemo elapsed = %v, expected < 200ms", elapsed)
	}
}

func TestPropagateCancel(t *testing.T) {
	parentErr, childErr := PropagateCancel()

	if parentErr != context.Canceled {
		t.Errorf("parent error = %v, want context.Canceled", parentErr)
	}
	if childErr != context.Canceled {
		t.Errorf("child error = %v, want context.Canceled", childErr)
	}
}

func TestCleanupPattern(t *testing.T) {
	cancelled := CleanupPattern()
	if !cancelled {
		t.Error("CleanupPattern() returned false, want true (defer cancel should run)")
	}
}
