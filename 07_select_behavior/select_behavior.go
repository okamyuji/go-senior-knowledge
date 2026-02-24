// Package selectbehavior demonstrates the behavior of Go's select statement
// including pseudo-random case selection, non-blocking operations with default,
// and timeout patterns using time.After.
package selectbehavior

import "time"

// SelectMultipleReady performs n iterations of selecting from two channels that
// both have data ready. When multiple cases are ready, Go's select picks one
// pseudo-randomly. It returns counts of how many times each channel was chosen.
func SelectMultipleReady(n int) (ch1Count, ch2Count int) {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	for range n {
		ch1 <- 1
		ch2 <- 2
		select {
		case <-ch1:
			ch1Count++
		case <-ch2:
			ch2Count++
		}
		// drain each channel individually with non-blocking receives
		select {
		case <-ch1:
		default:
		}
		select {
		case <-ch2:
		default:
		}
	}
	return ch1Count, ch2Count
}

// SelectWithDefault attempts a non-blocking receive on the given channel.
// If no value is available, it immediately returns the zero value and false.
func SelectWithDefault(ch <-chan int) (int, bool) {
	select {
	case v := <-ch:
		return v, true
	default:
		return 0, false
	}
}

// SelectWithTimeout waits for a value on the given channel up to the specified
// duration. It returns the value and true on success, or zero and false on
// timeout.
func SelectWithTimeout(ch <-chan int, timeout time.Duration) (int, bool) {
	select {
	case v := <-ch:
		return v, true
	case <-time.After(timeout):
		return 0, false
	}
}
