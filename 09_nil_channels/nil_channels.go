// Package nilchannels demonstrates how nil channels behave in select statements
// and how to use them to dynamically enable or disable select cases at runtime.
package nilchannels

// NilChannelInSelect demonstrates that a nil channel case in select is never
// selected. It sets up a select with one nil channel and one ready channel,
// returning the value from the ready channel.
func NilChannelInSelect(readyValue int) int {
	var nilCh chan int
	readyCh := make(chan int, 1)
	readyCh <- readyValue

	select {
	case v := <-nilCh:
		return v // never reached
	case v := <-readyCh:
		return v
	}
}

// DisableChannel reads values from two channels until both are closed.
// When a channel is closed, it is set to nil so the corresponding select
// case is disabled, preventing busy-loop zero-value reads.
func DisableChannel(ch1, ch2 <-chan int) []int {
	var result []int
	a, b := ch1, ch2

	for a != nil || b != nil {
		select {
		case v, ok := <-a:
			if !ok {
				a = nil
				continue
			}
			result = append(result, v)
		case v, ok := <-b:
			if !ok {
				b = nil
				continue
			}
			result = append(result, v)
		}
	}
	return result
}

// DynamicSelect reads from ch and doubles each value. When the doubled value
// exceeds the threshold, it stops processing and returns. It returns all
// processed values.
func DynamicSelect(ch <-chan int, threshold int) []int {
	var result []int
	for v := range ch {
		doubled := v * 2
		result = append(result, doubled)
		if doubled > threshold {
			return result
		}
	}
	return result
}
