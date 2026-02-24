package scheduling

import "testing"

func TestGoroutineCount(t *testing.T) {
	const n = 50
	count := GoroutineCount(n)
	// The count includes our n goroutines plus the test goroutine and
	// possibly runtime goroutines, so it must be at least n.
	if count < n {
		t.Errorf("GoroutineCount(%d) = %d, want >= %d", n, count, n)
	}
}

func TestProcessorCount(t *testing.T) {
	info := ProcessorCount()
	if info.GOMAXPROCS < 1 {
		t.Errorf("GOMAXPROCS = %d, want >= 1", info.GOMAXPROCS)
	}
	if info.NumCPU < 1 {
		t.Errorf("NumCPU = %d, want >= 1", info.NumCPU)
	}
}

func TestParallelWork(t *testing.T) {
	const n = 10
	results := ParallelWork(n)
	if len(results) != n {
		t.Fatalf("ParallelWork(%d) returned %d results, want %d", n, len(results), n)
	}
	for i, v := range results {
		want := i * i
		if v != want {
			t.Errorf("results[%d] = %d, want %d", i, v, want)
		}
	}
}
