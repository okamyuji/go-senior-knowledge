package memorymodel

import "testing"

func TestSynchronizedIncrement(t *testing.T) {
	const n = 1000
	got := SynchronizedIncrement(n)
	if got != n {
		t.Errorf("SynchronizedIncrement(%d) = %d, want %d", n, got, n)
	}
}

func TestChannelVisibility(t *testing.T) {
	got := ChannelVisibility()
	if got != 42 {
		t.Errorf("ChannelVisibility() = %d, want 42", got)
	}
}
