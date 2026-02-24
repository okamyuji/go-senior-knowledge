// Package gomaxprocs demonstrates GOMAXPROCS, which controls the number of
// OS threads that can execute user-level Go code simultaneously.
//
// GOMAXPROCS determines how many goroutines can run in parallel (on separate
// CPU cores). By default, it is set to the number of logical CPUs available.
// Setting GOMAXPROCS to 1 means only one goroutine runs at a time (though
// many can be scheduled cooperatively). Increasing it allows true parallelism
// on multi-core machines.
//
// Note that GOMAXPROCS limits user Go code execution threads. The runtime may
// use additional threads for system calls, GC, and other background tasks.
package gomaxprocs

import (
	"runtime"
	"sync"
	"time"
)

// GetMaxProcs returns the current GOMAXPROCS value.
func GetMaxProcs() int {
	return runtime.GOMAXPROCS(0)
}

// SetMaxProcs sets GOMAXPROCS to n, executes the given function, and then
// restores the original value. It returns the original GOMAXPROCS value.
func SetMaxProcs(n int, fn func()) int {
	original := runtime.GOMAXPROCS(n)
	defer runtime.GOMAXPROCS(original)
	fn()
	return original
}

// ParallelismDemo launches numGoroutines goroutines that each perform a
// CPU-bound task with the given GOMAXPROCS value. It returns the elapsed
// wall-clock time. By comparing results with different GOMAXPROCS values,
// you can observe the effect of parallelism on execution time.
func ParallelismDemo(maxProcs, numGoroutines int) time.Duration {
	original := runtime.GOMAXPROCS(maxProcs)
	defer runtime.GOMAXPROCS(original)

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	start := time.Now()
	for range numGoroutines {
		go func() {
			defer wg.Done()
			cpuWork()
		}()
	}
	wg.Wait()
	return time.Since(start)
}

// cpuWork performs a simple CPU-bound computation to simulate workload.
func cpuWork() {
	sum := 0
	for i := range 1_000_000 {
		sum += i
	}
	runtime.KeepAlive(sum)
}
