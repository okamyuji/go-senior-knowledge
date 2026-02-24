package preemption

import "testing"

func TestCooperativeYield(t *testing.T) {
	const n = 100
	got := CooperativeYield(n)
	if got != n {
		t.Errorf("CooperativeYield(%d) = %d, want %d", n, got, n)
	}
}

func TestPreemptibleWork(t *testing.T) {
	// Sum of 1..1000 = 1000*1001/2 = 500500
	got := PreemptibleWork(1000)
	if got != 500500 {
		t.Errorf("PreemptibleWork(1000) = %d, want 500500", got)
	}
}

func TestConcurrentTasks(t *testing.T) {
	const n = 5
	results := ConcurrentTasks(n)
	if len(results) != n {
		t.Fatalf("ConcurrentTasks(%d) returned %d results, want %d", n, len(results), n)
	}
	for i, got := range results {
		limit := int64((i + 1) * 1000)
		want := limit * (limit + 1) / 2
		if got != want {
			t.Errorf("results[%d] = %d, want %d", i, got, want)
		}
	}
}
