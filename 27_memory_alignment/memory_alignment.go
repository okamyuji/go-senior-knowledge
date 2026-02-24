// Package memalign demonstrates how struct field ordering affects memory
// layout and size due to alignment padding, and how false sharing degrades
// concurrent performance.
//
// Go aligns each field to a boundary equal to its own size (e.g., int64 is
// aligned to 8 bytes). The compiler inserts padding bytes between fields to
// satisfy these constraints. By ordering fields from largest to smallest,
// you can often eliminate unnecessary padding and reduce struct size.
//
// False sharing occurs when two independently accessed variables reside on
// the same CPU cache line (typically 64 bytes). When one core writes to its
// variable, the entire cache line is invalidated on all other cores, even
// though they access a different variable. Padding between variables ensures
// they occupy separate cache lines.
package memalign

import "unsafe"

// BadLayout has fields ordered to maximize wasted padding.
// On 64-bit systems the layout is:
//
//	bool (1 byte) + 7 bytes padding + int64 (8 bytes) +
//	bool (1 byte) + 7 bytes padding + int64 (8 bytes) = 32 bytes
type BadLayout struct {
	Flag1  bool
	Value1 int64
	Flag2  bool
	Value2 int64
}

// GoodLayout reorders the same fields to minimize padding.
// On 64-bit systems the layout is:
//
//	int64 (8 bytes) + int64 (8 bytes) +
//	bool (1 byte) + bool (1 byte) + 6 bytes padding = 24 bytes
type GoodLayout struct {
	Value1 int64
	Value2 int64
	Flag1  bool
	Flag2  bool
}

// CacheLineSize is the typical cache line size on modern x86 and ARM64 CPUs.
const CacheLineSize = 64

// FalseSharingPair has two counters that likely reside on the same cache
// line. When two goroutines increment Counter1 and Counter2 concurrently,
// each write invalidates the other core's cache line, causing performance
// degradation.
type FalseSharingPair struct {
	Counter1 int64
	Counter2 int64
}

// PaddedPair separates two counters onto different cache lines using padding.
// This prevents false sharing because writes to Counter1 never invalidate
// the cache line holding Counter2.
type PaddedPair struct {
	Counter1 int64
	_        [CacheLineSize - 8]byte // push Counter2 to the next cache line
	Counter2 int64
}

// SizeOfBadLayout returns the size of BadLayout in bytes.
func SizeOfBadLayout() uintptr {
	return unsafe.Sizeof(BadLayout{})
}

// SizeOfGoodLayout returns the size of GoodLayout in bytes.
func SizeOfGoodLayout() uintptr {
	return unsafe.Sizeof(GoodLayout{})
}

// SizeOfFalseSharingPair returns the size of FalseSharingPair in bytes.
func SizeOfFalseSharingPair() uintptr {
	return unsafe.Sizeof(FalseSharingPair{})
}

// SizeOfPaddedPair returns the size of PaddedPair in bytes.
func SizeOfPaddedPair() uintptr {
	return unsafe.Sizeof(PaddedPair{})
}
