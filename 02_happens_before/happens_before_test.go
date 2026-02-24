package happensbefore

import "testing"

func TestMutexOrder(t *testing.T) {
	const n = 500
	got := MutexOrder(n)
	if got != n {
		t.Errorf("MutexOrder(%d) = %d, want %d", n, got, n)
	}
}

func TestChannelOrder(t *testing.T) {
	got := ChannelOrder(99)
	if got != 99 {
		t.Errorf("ChannelOrder(99) = %d, want 99", got)
	}
}

func TestAtomicOrder(t *testing.T) {
	const n = 500
	got := AtomicOrder(n)
	if got != int64(n) {
		t.Errorf("AtomicOrder(%d) = %d, want %d", n, got, n)
	}
}

func TestWaitGroupOrder(t *testing.T) {
	const n = 500
	got := WaitGroupOrder(n)
	if got != n {
		t.Errorf("WaitGroupOrder(%d) = %d, want %d", n, got, n)
	}
}
