// Package memorymodel demonstrates Go's memory model through synchronized
// counter increments and channel-based visibility guarantees.
package memorymodel

import "sync"

// SynchronizedIncrement launches n goroutines that each increment a shared
// counter once, protected by sync.Mutex. It returns the final counter value.
func SynchronizedIncrement(n int) int {
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

// ChannelVisibility demonstrates that a write to a variable before a channel
// send is visible to the goroutine that receives from that channel.
// It returns the value observed after the signal is received.
func ChannelVisibility() int {
	var value int
	done := make(chan struct{})

	go func() {
		value = 42
		done <- struct{}{}
	}()

	<-done
	return value
}
