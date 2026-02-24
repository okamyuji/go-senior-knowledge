// Package dataraces demonstrates safe concurrent access patterns and explains
// why unsynchronized access leads to data races.
package dataraces

import (
	"sync"
	"sync/atomic"
)

// SafeCounter launches n goroutines that each increment a shared counter
// once, protected by sync.Mutex. It returns the final counter value.
func SafeCounter(n int) int {
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

// SafeAtomicCounter launches n goroutines that each atomically add 1
// to a shared counter using atomic.AddInt64. It returns the final value.
func SafeAtomicCounter(n int) int64 {
	var (
		wg      sync.WaitGroup
		counter int64
	)
	wg.Add(n)
	for range n {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}
	wg.Wait()
	return counter
}

// UnsafePatternDemo returns a description of why unsynchronized concurrent
// access causes a data race. It does not actually perform unsafe operations.
func UnsafePatternDemo() string {
	return "When multiple goroutines read and write a shared variable " +
		"without synchronization, the program has a data race. " +
		"The Go memory model does not guarantee the order or visibility " +
		"of such operations, so the result is undefined. " +
		"Use sync.Mutex, sync/atomic, or channels to prevent data races."
}
