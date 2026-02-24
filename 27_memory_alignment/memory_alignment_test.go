package memalign

import (
	"testing"
	"unsafe"
)

func TestGoodLayoutSmallerOrEqual(t *testing.T) {
	bad := SizeOfBadLayout()
	good := SizeOfGoodLayout()
	if good > bad {
		t.Errorf("GoodLayout (%d bytes) should be <= BadLayout (%d bytes)", good, bad)
	}
	t.Logf("BadLayout: %d bytes, GoodLayout: %d bytes, saved: %d bytes", bad, good, bad-good)
}

func TestBadLayoutSize(t *testing.T) {
	// On 64-bit platforms: bool(1)+pad(7)+int64(8)+bool(1)+pad(7)+int64(8)=32
	size := SizeOfBadLayout()
	if size < 32 {
		t.Errorf("BadLayout size = %d, expected at least 32 on 64-bit", size)
	}
}

func TestGoodLayoutSize(t *testing.T) {
	// On 64-bit platforms: int64(8)+int64(8)+bool(1)+bool(1)+pad(6)=24
	size := SizeOfGoodLayout()
	if size < 24 {
		t.Errorf("GoodLayout size = %d, expected at least 24 on 64-bit", size)
	}
}

func TestPaddedPairLargerThanFalseSharing(t *testing.T) {
	fs := SizeOfFalseSharingPair()
	pp := SizeOfPaddedPair()
	if pp <= fs {
		t.Errorf("PaddedPair (%d bytes) should be larger than FalseSharingPair (%d bytes)", pp, fs)
	}
	t.Logf("FalseSharingPair: %d bytes, PaddedPair: %d bytes", fs, pp)
}

func TestPaddedPairCacheLineSeparation(t *testing.T) {
	var p PaddedPair
	addr1 := uintptr(unsafe.Pointer(&p.Counter1))
	addr2 := uintptr(unsafe.Pointer(&p.Counter2))
	distance := addr2 - addr1
	if distance < CacheLineSize {
		t.Errorf("Counter2 is %d bytes from Counter1, want at least %d", distance, CacheLineSize)
	}
	t.Logf("distance between Counter1 and Counter2: %d bytes", distance)
}

func TestFieldsAccessible(t *testing.T) {
	bad := BadLayout{Flag1: true, Value1: 100, Flag2: false, Value2: 200}
	if !bad.Flag1 || bad.Value1 != 100 || bad.Flag2 || bad.Value2 != 200 {
		t.Error("BadLayout fields not accessible correctly")
	}

	good := GoodLayout{Value1: 100, Value2: 200, Flag1: true, Flag2: false}
	if good.Value1 != 100 || good.Value2 != 200 || !good.Flag1 || good.Flag2 {
		t.Error("GoodLayout fields not accessible correctly")
	}
}
