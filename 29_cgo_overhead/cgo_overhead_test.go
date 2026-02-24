package cgooverhead

import "testing"

func TestPureGoAdd(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-5, 5, 0},
		{100, 200, 300},
	}
	for _, tt := range tests {
		got := PureGoAdd(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("PureGoAdd(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestSimulateCgoBoundary(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-5, 5, 0},
		{100, 200, 300},
	}
	for _, tt := range tests {
		got := SimulateCgoBoundary(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("SimulateCgoBoundary(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestPureGoAndSimulatedAgree(t *testing.T) {
	for a := -10; a <= 10; a++ {
		for b := -10; b <= 10; b++ {
			pure := PureGoAdd(a, b)
			sim := SimulateCgoBoundary(a, b)
			if pure != sim {
				t.Fatalf("PureGoAdd(%d,%d)=%d != SimulateCgoBoundary(%d,%d)=%d",
					a, b, pure, a, b, sim)
			}
		}
	}
}

func TestBatchPureGoAdd(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{10, 20, 30, 40, 50}
	want := []int{11, 22, 33, 44, 55}
	got := BatchPureGoAdd(a, b)
	if len(got) != len(want) {
		t.Fatalf("BatchPureGoAdd: len=%d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("BatchPureGoAdd[%d] = %d, want %d", i, got[i], want[i])
		}
	}
}

func TestBatchSimulateCgoBoundary(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{10, 20, 30, 40, 50}
	want := []int{11, 22, 33, 44, 55}
	got := BatchSimulateCgoBoundary(a, b)
	if len(got) != len(want) {
		t.Fatalf("BatchSimulateCgoBoundary: len=%d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("BatchSimulateCgoBoundary[%d] = %d, want %d", i, got[i], want[i])
		}
	}
}

func TestBatchUnequalLengths(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{10, 20}
	got := BatchPureGoAdd(a, b)
	if len(got) != 2 {
		t.Errorf("BatchPureGoAdd with unequal slices: len=%d, want 2", len(got))
	}
}

func BenchmarkPureGoAdd(b *testing.B) {
	for range b.N {
		PureGoAdd(42, 58)
	}
}

func BenchmarkSimulateCgoBoundary(b *testing.B) {
	for range b.N {
		SimulateCgoBoundary(42, 58)
	}
}
