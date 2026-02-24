// Package gcbasics demonstrates Go's garbage collector behavior, including
// manual GC invocation, finalizers, and GC statistics.
//
// Go uses a concurrent, tri-color mark-and-sweep garbage collector.
//
// Tri-color marking works as follows:
//   - White: objects not yet visited (candidates for collection)
//   - Gray: objects visited but whose references have not been fully scanned
//   - Black: objects fully scanned (reachable, will not be collected)
//
// The GC starts by marking root objects (globals, stacks) as gray, then
// iteratively scans gray objects, marking their referents gray and the
// scanned object black. When no gray objects remain, all white objects
// are unreachable and can be collected.
//
// Stop-the-world (STW) phases:
// The GC has two brief STW phases. The first STW phase enables the write
// barrier and prepares for concurrent marking. The second STW phase occurs
// after marking is complete to disable the write barrier and perform
// final cleanup. Between these two phases, marking runs concurrently
// with application goroutines.
package gcbasics

import (
	"runtime"
	"sync/atomic"
)

// AllocateAndCollect allocates n objects on the heap, triggers a GC cycle,
// and returns the number of completed GC cycles (NumGC) from runtime.MemStats.
func AllocateAndCollect(n int) uint32 {
	// Allocate objects that escape to the heap.
	holders := make([]*[1024]byte, n)
	for i := range n {
		holders[i] = new([1024]byte)
	}
	// Prevent the compiler from optimizing away the allocations.
	runtime.KeepAlive(holders)

	runtime.GC()

	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	return stats.NumGC
}

// finalizerCount is used by FinalizerDemo to track finalizer invocations.
var finalizerCount atomic.Int64

// FinalizerDemo sets a finalizer on a heap-allocated object and triggers GC.
// It returns the number of finalizer invocations observed. Note that finalizer
// execution is not guaranteed to happen immediately, so the returned count
// may be 0 in some cases.
func FinalizerDemo() int64 {
	finalizerCount.Store(0)

	obj := new(int)
	runtime.SetFinalizer(obj, func(_ *int) {
		finalizerCount.Add(1)
	})
	// Drop the reference so the object becomes eligible for collection.
	obj = nil
	runtime.KeepAlive(obj)

	runtime.GC()
	// Finalizers run in a separate goroutine. Give it a moment.
	runtime.Gosched()

	return finalizerCount.Load()
}

// GCStatSnapshot holds a subset of GC statistics.
type GCStatSnapshot struct {
	NumGC      uint32 // number of completed GC cycles
	PauseTotal uint64 // total STW pause time in nanoseconds
	HeapAlloc  uint64 // bytes of allocated heap objects
}

// GCStats reads runtime.MemStats and returns a snapshot of GC-related
// statistics. NumGC is the number of completed GC cycles, PauseTotal is the
// cumulative STW pause time, and HeapAlloc is the current heap allocation.
func GCStats() GCStatSnapshot {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	return GCStatSnapshot{
		NumGC:      stats.NumGC,
		PauseTotal: stats.PauseTotalNs,
		HeapAlloc:  stats.HeapAlloc,
	}
}
