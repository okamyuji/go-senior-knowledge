package racedetector

import "testing"

func TestSafeIncrement(t *testing.T) {
	const n = 1000
	got := SafeIncrement(n)
	if got != n {
		t.Errorf("SafeIncrement(%d) = %d, want %d", n, got, n)
	}
}

func TestSafeIncrementZero(t *testing.T) {
	got := SafeIncrement(0)
	if got != 0 {
		t.Errorf("SafeIncrement(0) = %d, want 0", got)
	}
}

func TestSafeAtomicIncrement(t *testing.T) {
	const n = 1000
	got := SafeAtomicIncrement(n)
	if got != int64(n) {
		t.Errorf("SafeAtomicIncrement(%d) = %d, want %d", n, got, n)
	}
}

func TestSafeAtomicIncrementZero(t *testing.T) {
	got := SafeAtomicIncrement(0)
	if got != 0 {
		t.Errorf("SafeAtomicIncrement(0) = %d, want 0", got)
	}
}

func TestSafeChannelCommunication(t *testing.T) {
	const n = 100
	got := SafeChannelCommunication(n)
	// sum of 1..n = n*(n+1)/2
	want := n * (n + 1) / 2
	if got != want {
		t.Errorf("SafeChannelCommunication(%d) = %d, want %d", n, got, want)
	}
}

func TestSafeChannelCommunicationZero(t *testing.T) {
	got := SafeChannelCommunication(0)
	if got != 0 {
		t.Errorf("SafeChannelCommunication(0) = %d, want 0", got)
	}
}
