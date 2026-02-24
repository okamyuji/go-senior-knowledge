package nilchannels

import (
	"slices"
	"sort"
	"testing"
)

func TestNilChannelInSelect(t *testing.T) {
	got := NilChannelInSelect(42)
	if got != 42 {
		t.Errorf("NilChannelInSelect(42) = %d, want 42", got)
	}
}

func TestDisableChannel(t *testing.T) {
	ch1 := make(chan int, 3)
	ch2 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	close(ch1)
	ch2 <- 4
	ch2 <- 5
	close(ch2)

	got := DisableChannel(ch1, ch2)
	sort.Ints(got) // order may vary
	want := []int{1, 2, 3, 4, 5}
	if !slices.Equal(got, want) {
		t.Errorf("DisableChannel = %v, want %v", got, want)
	}
}

func TestDisableChannelOneEmpty(t *testing.T) {
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 20
	close(ch1)

	ch2 := make(chan int)
	close(ch2)

	got := DisableChannel(ch1, ch2)
	sort.Ints(got)
	want := []int{10, 20}
	if !slices.Equal(got, want) {
		t.Errorf("DisableChannel = %v, want %v", got, want)
	}
}

func TestDynamicSelect(t *testing.T) {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 3
	ch <- 10
	ch <- 100
	ch <- 200
	close(ch)

	got := DynamicSelect(ch, 15)
	// Values: 1*2=2, 3*2=6, 10*2=20 (exceeds 15, stops)
	want := []int{2, 6, 20}
	if !slices.Equal(got, want) {
		t.Errorf("DynamicSelect = %v, want %v", got, want)
	}
}

func TestDynamicSelectAllBelowThreshold(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	got := DynamicSelect(ch, 100)
	want := []int{2, 4, 6}
	if !slices.Equal(got, want) {
		t.Errorf("DynamicSelect = %v, want %v", got, want)
	}
}
