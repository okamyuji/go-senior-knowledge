package syncprimitives

import "testing"

func TestAtomicCounter(t *testing.T) {
	const total = 10000
	const goroutines = 10
	got := AtomicCounter(total, goroutines)
	if got != int64(total) {
		t.Errorf("AtomicCounter(%d, %d) = %d, want %d", total, goroutines, got, total)
	}
}

func TestMutexCounter(t *testing.T) {
	const total = 10000
	const goroutines = 10
	got := MutexCounter(total, goroutines)
	if got != int64(total) {
		t.Errorf("MutexCounter(%d, %d) = %d, want %d", total, goroutines, got, total)
	}
}

func TestChannelCounter(t *testing.T) {
	const total = 10000
	const goroutines = 10
	got := ChannelCounter(total, goroutines)
	if got != int64(total) {
		t.Errorf("ChannelCounter(%d, %d) = %d, want %d", total, goroutines, got, total)
	}
}

func TestCountersAgree(t *testing.T) {
	const total = 5000
	const goroutines = 5
	a := AtomicCounter(total, goroutines)
	m := MutexCounter(total, goroutines)
	c := ChannelCounter(total, goroutines)
	if a != m || m != c {
		t.Errorf("counters disagree: atomic=%d, mutex=%d, channel=%d", a, m, c)
	}
}
