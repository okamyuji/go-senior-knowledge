package dataraces

import (
	"strings"
	"testing"
)

func TestSafeCounter(t *testing.T) {
	const n = 1000
	got := SafeCounter(n)
	if got != n {
		t.Errorf("SafeCounter(%d) = %d, want %d", n, got, n)
	}
}

func TestSafeAtomicCounter(t *testing.T) {
	const n = 1000
	got := SafeAtomicCounter(n)
	if got != int64(n) {
		t.Errorf("SafeAtomicCounter(%d) = %d, want %d", n, got, n)
	}
}

func TestUnsafePatternDemo(t *testing.T) {
	desc := UnsafePatternDemo()
	if !strings.Contains(desc, "data race") {
		t.Error("UnsafePatternDemo should mention data race")
	}
	if !strings.Contains(desc, "sync.Mutex") {
		t.Error("UnsafePatternDemo should mention sync.Mutex as a solution")
	}
}
