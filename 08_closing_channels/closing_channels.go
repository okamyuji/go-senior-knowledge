// Package closingchannels demonstrates channel closing patterns: producer-consumer
// with range, receiving from closed channels, and fan-in merging of multiple channels.
package closingchannels

import "sync"

// ProducerConsumer launches a producer goroutine that sends all values into a
// channel and then closes it. The consumer ranges over the channel, collecting
// results in FIFO order.
func ProducerConsumer(values []int) []int {
	ch := make(chan int)
	go func() {
		for _, v := range values {
			ch <- v
		}
		close(ch)
	}()

	var result []int
	for v := range ch {
		result = append(result, v)
	}
	return result
}

// ReceiveFromClosed demonstrates the behavior of receiving from a closed
// channel. It returns the zero value and ok=false, indicating the channel
// is closed and empty.
func ReceiveFromClosed() (value int, ok bool) {
	ch := make(chan int, 1)
	ch <- 10
	close(ch)

	// first receive gets the buffered value
	<-ch

	// second receive from closed, empty channel
	value, ok = <-ch
	return value, ok
}

// FanIn merges values from all input channels into a single output slice.
// It launches one goroutine per input channel, and a coordinator goroutine
// that waits for all of them to finish before closing the output channel.
func FanIn(channels ...<-chan int) []int {
	out := make(chan int)
	var wg sync.WaitGroup

	wg.Add(len(channels))
	for _, ch := range channels {
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	var result []int
	for v := range out {
		result = append(result, v)
	}
	return result
}
