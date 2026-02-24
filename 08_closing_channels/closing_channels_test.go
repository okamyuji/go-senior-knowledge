package closingchannels

import (
	"slices"
	"sort"
	"testing"
)

func TestProducerConsumer(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	got := ProducerConsumer(values)
	if !slices.Equal(got, values) {
		t.Errorf("ProducerConsumer(%v) = %v, want %v", values, got, values)
	}
}

func TestProducerConsumerEmpty(t *testing.T) {
	got := ProducerConsumer(nil)
	if len(got) != 0 {
		t.Errorf("ProducerConsumer(nil) returned %d elements, want 0", len(got))
	}
}

func TestReceiveFromClosed(t *testing.T) {
	value, ok := ReceiveFromClosed()
	if ok {
		t.Error("ReceiveFromClosed returned ok=true, want false")
	}
	if value != 0 {
		t.Errorf("ReceiveFromClosed returned value=%d, want 0", value)
	}
}

func TestFanIn(t *testing.T) {
	makeCh := func(values ...int) <-chan int {
		ch := make(chan int, len(values))
		for _, v := range values {
			ch <- v
		}
		close(ch)
		return ch
	}

	ch1 := makeCh(1, 2, 3)
	ch2 := makeCh(4, 5)
	ch3 := makeCh(6)

	got := FanIn(ch1, ch2, ch3)
	sort.Ints(got) // order is non-deterministic across goroutines
	want := []int{1, 2, 3, 4, 5, 6}
	if !slices.Equal(got, want) {
		t.Errorf("FanIn = %v, want %v", got, want)
	}
}

func TestFanInNoChannels(t *testing.T) {
	got := FanIn()
	if len(got) != 0 {
		t.Errorf("FanIn() returned %d elements, want 0", len(got))
	}
}
