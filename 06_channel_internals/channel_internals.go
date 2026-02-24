// Package channelinternals demonstrates Go channel fundamentals: unbuffered
// synchronous handoffs, buffered asynchronous sends, and directional channel types.
package channelinternals

// UnbufferedChannel sends a value through an unbuffered channel between two
// goroutines and returns the received value. The send blocks until the
// receiver is ready, guaranteeing a synchronous handoff.
func UnbufferedChannel(v int) int {
	ch := make(chan int)
	go func() {
		ch <- v
	}()
	return <-ch
}

// BufferedChannel sends all values from the slice into a buffered channel
// whose capacity equals the slice length, then receives and returns them
// in FIFO order. Because the buffer is large enough, all sends complete
// without a concurrent receiver.
func BufferedChannel(values []int) []int {
	ch := make(chan int, len(values))
	for _, v := range values {
		ch <- v
	}
	result := make([]int, 0, len(values))
	for range len(values) {
		result = append(result, <-ch)
	}
	return result
}

// ChannelDirection demonstrates directional channel types. The producer
// function writes values to a send-only channel, while the consumer reads
// from a receive-only channel. This enforces correct usage at compile time.
func ChannelDirection(values []int) []int {
	ch := make(chan int, len(values))
	producer(ch, values)
	return consumer(ch, len(values))
}

func producer(out chan<- int, values []int) {
	for _, v := range values {
		out <- v
	}
}

func consumer(in <-chan int, n int) []int {
	result := make([]int, 0, n)
	for range n {
		result = append(result, <-in)
	}
	return result
}
