// Package sliceinternals demonstrates the internal structure of Go slices,
// including the slice header (len/cap), append growth behavior, shared backing
// arrays, and proper copying techniques.
//
// A slice header consists of three fields: a pointer to the underlying array,
// the length (number of elements), and the capacity (maximum elements before
// reallocation). Understanding this structure is essential for writing
// efficient and correct Go code.
package sliceinternals

// SliceHeader returns the length and capacity of the given slice.
// Internally, a slice is represented as a struct with a pointer to the
// backing array, a length, and a capacity. This function exposes len and cap.
func SliceHeader(s []int) (length, capacity int) {
	return len(s), cap(s)
}

// AppendGrowth demonstrates how append doubles capacity when the backing array
// is full. It starts with a nil slice and appends n elements, recording the
// capacity at each step. The returned slice contains the capacity values.
//
// Go's runtime grows slices roughly by doubling for small slices and by ~1.25x
// for larger ones (threshold around 256 elements as of Go 1.21+).
func AppendGrowth(n int) []int {
	var s []int
	caps := make([]int, 0, n)
	for i := range n {
		s = append(s, i)
		caps = append(caps, cap(s))
	}
	return caps
}

// SharedBackingArray creates a slice of length 5 and derives a sub-slice from
// it. It then mutates the sub-slice at the given index and returns both
// slices. Because both slices share the same backing array, the mutation
// is visible in both.
func SharedBackingArray(index, value int) (original, sub []int) {
	original = []int{0, 1, 2, 3, 4}
	sub = original[1:4] // shares backing array, elements [1,2,3]
	sub[index] = value
	return original, sub
}

// CopySlice creates an independent copy of src using the built-in copy
// function. Mutations to the returned slice do not affect src.
func CopySlice(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}
