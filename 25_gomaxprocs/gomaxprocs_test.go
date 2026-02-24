package gomaxprocs

import (
	"runtime"
	"testing"
)

func TestGetMaxProcs(t *testing.T) {
	got := GetMaxProcs()
	if got <= 0 {
		t.Errorf("GetMaxProcs() = %d, want > 0", got)
	}
}

func TestGetMaxProcsMatchesRuntime(t *testing.T) {
	want := runtime.GOMAXPROCS(0)
	got := GetMaxProcs()
	if got != want {
		t.Errorf("GetMaxProcs() = %d, want %d", got, want)
	}
}

func TestSetMaxProcs(t *testing.T) {
	original := runtime.GOMAXPROCS(0)
	defer runtime.GOMAXPROCS(original)

	called := false
	returned := SetMaxProcs(2, func() {
		called = true
		current := runtime.GOMAXPROCS(0)
		if current != 2 {
			t.Errorf("inside SetMaxProcs: GOMAXPROCS = %d, want 2", current)
		}
	})

	if !called {
		t.Error("SetMaxProcs did not call the provided function")
	}
	if returned != original {
		t.Errorf("SetMaxProcs returned %d, want original %d", returned, original)
	}

	// Verify restoration.
	restored := runtime.GOMAXPROCS(0)
	if restored != original {
		t.Errorf("GOMAXPROCS after SetMaxProcs = %d, want original %d", restored, original)
	}
}

func TestParallelismDemo(t *testing.T) {
	// Just verify that ParallelismDemo runs without error and returns
	// a positive duration.
	d := ParallelismDemo(1, 2)
	if d <= 0 {
		t.Errorf("ParallelismDemo(1, 2) = %v, want > 0", d)
	}
}

func TestParallelismDemoRestoresMaxProcs(t *testing.T) {
	original := runtime.GOMAXPROCS(0)
	ParallelismDemo(1, 1)
	after := runtime.GOMAXPROCS(0)
	if after != original {
		t.Errorf("GOMAXPROCS after ParallelismDemo = %d, want %d", after, original)
	}
}
