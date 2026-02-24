package escapeanalysis

import "testing"

func TestStackAlloc(t *testing.T) {
	got := StackAlloc(3, 4)
	if got != 7 {
		t.Errorf("StackAlloc(3, 4) = %d, want 7", got)
	}
}

func TestHeapAlloc(t *testing.T) {
	p := HeapAlloc(42)
	if p == nil {
		t.Fatal("HeapAlloc returned nil")
	}
	if *p != 42 {
		t.Errorf("*HeapAlloc(42) = %d, want 42", *p)
	}

	// Each call should return a distinct pointer.
	p2 := HeapAlloc(42)
	if p == p2 {
		t.Error("two HeapAlloc calls returned the same pointer")
	}
}

func TestSliceEscape(t *testing.T) {
	s := SliceEscape(5)
	if len(s) != 5 {
		t.Fatalf("len = %d, want 5", len(s))
	}
	want := []int{0, 1, 4, 9, 16}
	for i, v := range s {
		if v != want[i] {
			t.Errorf("s[%d] = %d, want %d", i, v, want[i])
		}
	}
}

func TestSliceEscape_Empty(t *testing.T) {
	s := SliceEscape(0)
	if len(s) != 0 {
		t.Errorf("len = %d, want 0", len(s))
	}
}

func TestNoEscape(t *testing.T) {
	v := 10
	got := NoEscape(&v)
	if got != 11 {
		t.Errorf("NoEscape(&10) = %d, want 11", got)
	}
	// The original value should be unchanged.
	if v != 10 {
		t.Errorf("v changed to %d, want 10", v)
	}
}

func TestSumViaPointer(t *testing.T) {
	got := SumViaPointer(5, 3)
	if got != 9 {
		t.Errorf("SumViaPointer(5, 3) = %d, want 9", got)
	}
}

func BenchmarkStackAlloc(b *testing.B) {
	for b.Loop() {
		_ = StackAlloc(1, 2)
	}
}

func BenchmarkHeapAlloc(b *testing.B) {
	for b.Loop() {
		_ = HeapAlloc(42)
	}
}
