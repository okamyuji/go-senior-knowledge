// Package racedetector demonstrates race-safe patterns that pass the Go
// race detector. All functions in this package are safe for concurrent use.
//
// # Running the Race Detector
//
// The race detector is enabled with the -race flag:
//
//	go test -race ./38_race_detector/...
//	go run -race main.go
//	go build -race -o myapp
//
// # What the Race Detector Catches
//
// The race detector instruments memory accesses at compile time and monitors
// them at runtime. It detects actual data races that occur during execution:
// unsynchronized concurrent reads and writes to the same memory location.
// It reports the goroutines involved and the stack traces of the conflicting
// accesses.
//
// # What the Race Detector Cannot Prove
//
// The race detector can only find races that actually happen during a test run.
// It cannot prove the absence of races. A program may have latent races that
// are not triggered by a particular execution path or timing. Comprehensive
// tests with good coverage improve the chance of detecting races, but absence
// of reports does not guarantee race-freedom.
//
// # Performance Overhead
//
// The race detector is sampling-based and adds approximately 5-10x CPU
// overhead and 5-10x memory overhead. It is intended for testing and
// development, not for production use.
package racedetector

import (
	"sync"
	"sync/atomic"
)

// SafeIncrement launches n goroutines that each increment a shared counter
// once, using sync.Mutex for synchronization. It returns the final value.
func SafeIncrement(n int) int {
	var (
		mu      sync.Mutex
		wg      sync.WaitGroup
		counter int
	)
	wg.Add(n)
	for range n {
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	return counter
}

// SafeAtomicIncrement launches n goroutines that each atomically increment
// a shared counter once using sync/atomic. It returns the final value.
func SafeAtomicIncrement(n int) int64 {
	var (
		wg      sync.WaitGroup
		counter atomic.Int64
	)
	wg.Add(n)
	for range n {
		go func() {
			defer wg.Done()
			counter.Add(1)
		}()
	}
	wg.Wait()
	return counter.Load()
}

// SafeChannelCommunication launches n producer goroutines that each send
// a value through a channel. A single consumer goroutine receives all values
// and computes their sum. No shared memory is accessed by multiple goroutines.
// This follows Go's principle: "Do not communicate by sharing memory;
// instead, share memory by communicating."
func SafeChannelCommunication(n int) int {
	ch := make(chan int, n)
	var wg sync.WaitGroup

	wg.Add(n)
	for i := range n {
		go func(i int) {
			defer wg.Done()
			ch <- i + 1
		}(i)
	}

	// Close the channel after all producers finish.
	go func() {
		wg.Wait()
		close(ch)
	}()

	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}
