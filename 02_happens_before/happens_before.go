// Package happensbefore demonstrates the four primary happens-before
// relationships in Go: mutex, channel, atomic, and WaitGroup ordering.
package happensbefore

import (
	"sync"
	"sync/atomic"
)

// MutexOrder launches n goroutines that increment a shared counter protected
// by a mutex. The mutex lock/unlock pairs create happens-before edges,
// guaranteeing the final value equals n.
func MutexOrder(n int) int {
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

// ChannelOrder sends a value on a channel from a goroutine and receives it
// in the caller. The channel send happens-before the corresponding receive,
// so the sent value is always visible to the receiver.
func ChannelOrder(value int) int {
	ch := make(chan int)
	go func() {
		ch <- value
	}()
	return <-ch
}

// AtomicOrder launches n goroutines that each atomically add 1 to a counter.
// Atomic operations provide happens-before guarantees for the affected
// memory location, ensuring the final value equals n.
func AtomicOrder(n int) int64 {
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

// WaitGroupOrder launches n goroutines that each append a value to a shared
// slice. The call to wg.Done happens-before the corresponding wg.Wait
// returns, so all values are visible when Wait completes.
func WaitGroupOrder(n int) int {
	var (
		mu      sync.Mutex
		wg      sync.WaitGroup
		results []int
	)
	wg.Add(n)
	for i := range n {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			results = append(results, i)
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	return len(results)
}
