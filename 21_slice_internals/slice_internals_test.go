package sliceinternals

import "testing"

func TestSliceHeader(t *testing.T) {
	s := make([]int, 3, 10)
	length, capacity := SliceHeader(s)
	if length != 3 {
		t.Errorf("SliceHeader length = %d, want 3", length)
	}
	if capacity != 10 {
		t.Errorf("SliceHeader capacity = %d, want 10", capacity)
	}
}

func TestSliceHeaderNil(t *testing.T) {
	var s []int
	length, capacity := SliceHeader(s)
	if length != 0 || capacity != 0 {
		t.Errorf("SliceHeader(nil) = (%d, %d), want (0, 0)", length, capacity)
	}
}

func TestAppendGrowth(t *testing.T) {
	caps := AppendGrowth(10)
	if len(caps) != 10 {
		t.Fatalf("AppendGrowth(10) returned %d entries, want 10", len(caps))
	}
	// First append to a nil slice should yield capacity >= 1.
	if caps[0] < 1 {
		t.Errorf("initial capacity = %d, want >= 1", caps[0])
	}
	// Capacity must be non-decreasing.
	for i := 1; i < len(caps); i++ {
		if caps[i] < caps[i-1] {
			t.Errorf("capacity decreased at index %d: %d < %d", i, caps[i], caps[i-1])
		}
	}
	// After 10 appends starting from nil, capacity should exceed 10's half
	// due to doubling behavior.
	if caps[9] < 10 {
		t.Errorf("final capacity = %d, want >= 10", caps[9])
	}
}

func TestSharedBackingArray(t *testing.T) {
	original, sub := SharedBackingArray(0, 99)
	// sub[0] corresponds to original[1], so original[1] should be 99.
	if original[1] != 99 {
		t.Errorf("original[1] = %d, want 99 (shared backing array)", original[1])
	}
	if sub[0] != 99 {
		t.Errorf("sub[0] = %d, want 99", sub[0])
	}
}

func TestCopySlice(t *testing.T) {
	src := []int{10, 20, 30}
	dst := CopySlice(src)

	// Verify the copy has the same values.
	for i, v := range src {
		if dst[i] != v {
			t.Errorf("dst[%d] = %d, want %d", i, dst[i], v)
		}
	}

	// Mutate the copy and verify the original is unchanged.
	dst[0] = 999
	if src[0] != 10 {
		t.Errorf("src[0] = %d after mutating copy, want 10", src[0])
	}
}
