package channelinternals

import (
	"slices"
	"testing"
)

func TestUnbufferedChannel(t *testing.T) {
	got := UnbufferedChannel(42)
	if got != 42 {
		t.Errorf("UnbufferedChannel(42) = %d, want 42", got)
	}
}

func TestBufferedChannel(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	got := BufferedChannel(values)
	if !slices.Equal(got, values) {
		t.Errorf("BufferedChannel(%v) = %v, want %v", values, got, values)
	}
}

func TestBufferedChannelEmpty(t *testing.T) {
	got := BufferedChannel(nil)
	if len(got) != 0 {
		t.Errorf("BufferedChannel(nil) returned %d elements, want 0", len(got))
	}
}

func TestChannelDirection(t *testing.T) {
	values := []int{10, 20, 30}
	got := ChannelDirection(values)
	if !slices.Equal(got, values) {
		t.Errorf("ChannelDirection(%v) = %v, want %v", values, got, values)
	}
}
