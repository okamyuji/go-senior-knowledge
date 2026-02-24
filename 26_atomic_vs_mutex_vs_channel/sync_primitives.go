// Package syncprimitives compares three synchronization approaches for
// concurrent counter increments: atomic operations, sync.Mutex, and channels.
//
// Use atomic for simple numeric counters where the operation is a single
// read-modify-write. Use sync.Mutex when a critical section involves multiple
// steps or complex invariants. Use channels when goroutines need to
// communicate values or coordinate workflows, not just protect shared state.
package syncprimitives

import (
	"sync"
	"sync/atomic"
)

// AtomicCounter increments a counter n times across numGoroutines goroutines
// using sync/atomic. Atomic operations are lock-free and have the lowest
// overhead for simple numeric updates, but they only support individual
// read-modify-write operations on a single variable.
func AtomicCounter(n, numGoroutines int) int64 {
	var counter int64
	perGoroutine := n / numGoroutines
	var wg sync.WaitGroup
	wg.Add(numGoroutines)
	for range numGoroutines {
		go func() {
			defer wg.Done()
			for range perGoroutine {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}
	wg.Wait()
	return counter
}

// MutexCounter increments a counter n times across numGoroutines goroutines
// using sync.Mutex. Mutex is appropriate when the critical section contains
// multiple operations that must be performed atomically together, such as
// updating several related fields or checking a condition before writing.
func MutexCounter(n, numGoroutines int) int64 {
	var (
		mu      sync.Mutex
		counter int64
	)
	perGoroutine := n / numGoroutines
	var wg sync.WaitGroup
	wg.Add(numGoroutines)
	for range numGoroutines {
		go func() {
			defer wg.Done()
			for range perGoroutine {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	return counter
}

// ChannelCounter increments a counter n times across numGoroutines goroutines
// using a channel-based approach. A single manager goroutine owns the counter
// state and processes increment requests from a channel. This pattern is
// idiomatic Go ("share memory by communicating") and works well when the
// state owner also needs to perform side effects, fan-out results, or
// coordinate multiple pieces of state.
func ChannelCounter(n, numGoroutines int) int64 {
	inc := make(chan struct{}, numGoroutines)
	done := make(chan int64)

	// Manager goroutine: the sole owner of the counter.
	go func() {
		var counter int64
		for range n {
			<-inc
			counter++
		}
		done <- counter
	}()

	perGoroutine := n / numGoroutines
	var wg sync.WaitGroup
	wg.Add(numGoroutines)
	for range numGoroutines {
		go func() {
			defer wg.Done()
			for range perGoroutine {
				inc <- struct{}{}
			}
		}()
	}
	wg.Wait()

	return <-done
}
