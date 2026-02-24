package deferinternals

import (
	"reflect"
	"testing"
)

func TestDeferOrder(t *testing.T) {
	// Defers execute in LIFO order: third registered runs first, then
	// second, then first. But DeferOrder appends to the slice during
	// execution, so the slice records the execution order.
	got := DeferOrder()
	want := []string{"third", "second", "first"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("DeferOrder() = %v, want %v", got, want)
	}
}

func TestDeferWithClosureBug(t *testing.T) {
	got := DeferWithClosureBug(3)
	// All closures capture the same shared variable, which equals n after the loop.
	want := []int{3, 3, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("DeferWithClosureBug(3) = %v, want %v", got, want)
	}
}

func TestDeferWithClosureFix(t *testing.T) {
	got := DeferWithClosureFix(3)
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("DeferWithClosureFix(3) = %v, want %v", got, want)
	}
}

func TestDeferInLoop(t *testing.T) {
	const n = 100
	got := DeferInLoop(n)
	if got != n {
		t.Errorf("DeferInLoop(%d) = %d, want %d", n, got, n)
	}
}

func TestDeferModifyReturn(t *testing.T) {
	got := DeferModifyReturn()
	if got != 11 {
		t.Errorf("DeferModifyReturn() = %d, want 11", got)
	}
}
