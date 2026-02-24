package gcbasics

import (
	"runtime"
	"testing"
)

func TestAllocateAndCollect(t *testing.T) {
	numGC := AllocateAndCollect(100)
	if numGC == 0 {
		t.Error("AllocateAndCollect returned NumGC = 0, want > 0")
	}
}

func TestFinalizerDemo(t *testing.T) {
	// FinalizerDemo should not panic. The finalizer count may or may not be
	// incremented depending on GC timing, so we only verify no errors occur.
	count := FinalizerDemo()
	if count < 0 {
		t.Errorf("FinalizerDemo returned negative count: %d", count)
	}
}

func TestGCStats(t *testing.T) {
	// Force at least one GC cycle.
	runtime.GC()

	snap := GCStats()
	if snap.NumGC == 0 {
		t.Error("GCStats.NumGC = 0 after runtime.GC(), want > 0")
	}
	if snap.HeapAlloc == 0 {
		t.Error("GCStats.HeapAlloc = 0, want > 0")
	}
}

func TestGCStatsAfterMultipleCycles(t *testing.T) {
	snap1 := GCStats()
	runtime.GC()
	snap2 := GCStats()

	if snap2.NumGC <= snap1.NumGC {
		t.Errorf("NumGC did not increase: before=%d, after=%d", snap1.NumGC, snap2.NumGC)
	}
}
