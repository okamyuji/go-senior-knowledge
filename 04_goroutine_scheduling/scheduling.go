// Package scheduling demonstrates goroutine scheduling, processor
// configuration, and parallel work distribution in Go.
package scheduling

import (
	"runtime"
	"sync"
)

// GoroutineCount launches n goroutines that block on a channel, records
// the number of active goroutines via runtime.NumGoroutine, then releases
// them. It returns the goroutine count observed while all n were alive.
func GoroutineCount(n int) int {
	ready := make(chan struct{})
	done := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(n)
	for range n {
		go func() {
			wg.Done()
			<-done
		}()
	}
	// Wait until all goroutines have started.
	wg.Wait()

	count := runtime.NumGoroutine()

	// Signal goroutines to exit.
	close(done)
	_ = ready

	return count
}

// ProcessorInfo holds runtime processor configuration.
type ProcessorInfo struct {
	GOMAXPROCS int
	NumCPU     int
}

// ProcessorCount returns the current GOMAXPROCS setting and the number
// of logical CPUs available.
func ProcessorCount() ProcessorInfo {
	return ProcessorInfo{
		GOMAXPROCS: runtime.GOMAXPROCS(0),
		NumCPU:     runtime.NumCPU(),
	}
}

// ParallelWork launches n goroutines that each compute the square of their
// index. Results are collected through a channel and returned as a slice
// sorted by index.
func ParallelWork(n int) []int {
	type result struct {
		index int
		value int
	}
	ch := make(chan result, n)

	for i := range n {
		go func() {
			ch <- result{index: i, value: i * i}
		}()
	}

	results := make([]int, n)
	for range n {
		r := <-ch
		results[r.index] = r.value
	}
	return results
}
