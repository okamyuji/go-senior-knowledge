package selectbehavior

import (
	"testing"
	"time"
)

func TestSelectMultipleReady(t *testing.T) {
	const iterations = 1000
	ch1Count, ch2Count := SelectMultipleReady(iterations)

	if ch1Count+ch2Count != iterations {
		t.Fatalf("total selections = %d, want %d", ch1Count+ch2Count, iterations)
	}
	// With pseudo-random selection over 1000 iterations, each channel
	// should be picked at least once. A threshold of 1% is conservative.
	if ch1Count == 0 || ch2Count == 0 {
		t.Errorf("expected both channels to be selected at least once: ch1=%d, ch2=%d", ch1Count, ch2Count)
	}
}

func TestSelectWithDefaultNoData(t *testing.T) {
	ch := make(chan int)
	start := time.Now()
	v, ok := SelectWithDefault(ch)
	elapsed := time.Since(start)

	if ok {
		t.Error("SelectWithDefault returned ok=true on empty channel")
	}
	if v != 0 {
		t.Errorf("SelectWithDefault returned value %d, want 0", v)
	}
	if elapsed > 10*time.Millisecond {
		t.Errorf("SelectWithDefault took %v, expected near-instant return", elapsed)
	}
}

func TestSelectWithDefaultHasData(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 99
	v, ok := SelectWithDefault(ch)
	if !ok {
		t.Error("SelectWithDefault returned ok=false when data was available")
	}
	if v != 99 {
		t.Errorf("SelectWithDefault = %d, want 99", v)
	}
}

func TestSelectWithTimeoutSuccess(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 7
	v, ok := SelectWithTimeout(ch, time.Second)
	if !ok {
		t.Error("SelectWithTimeout returned ok=false when data was available")
	}
	if v != 7 {
		t.Errorf("SelectWithTimeout = %d, want 7", v)
	}
}

func TestSelectWithTimeoutExpires(t *testing.T) {
	ch := make(chan int)
	start := time.Now()
	v, ok := SelectWithTimeout(ch, 50*time.Millisecond)
	elapsed := time.Since(start)

	if ok {
		t.Error("SelectWithTimeout returned ok=true on timeout")
	}
	if v != 0 {
		t.Errorf("SelectWithTimeout returned value %d on timeout, want 0", v)
	}
	if elapsed < 40*time.Millisecond {
		t.Errorf("SelectWithTimeout returned too quickly: %v", elapsed)
	}
}
