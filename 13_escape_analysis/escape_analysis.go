// Package escapeanalysis demonstrates how the Go compiler decides whether to
// allocate variables on the stack or the heap.
//
// To see escape analysis output, run:
//
//	go build -gcflags="-m" ./13_escape_analysis/
//
// Adding -m multiple times increases verbosity:
//
//	go build -gcflags="-m -m" ./13_escape_analysis/
//
// Stack allocation is preferred because it is cheaper: the memory is freed
// automatically when the function returns. Heap allocation requires garbage
// collection. A variable "escapes to heap" when the compiler cannot prove
// that the variable's lifetime does not exceed the stack frame.
package escapeanalysis

// StackAlloc creates a variable that stays on the stack.
// The integer is used only within the function and never shared outside,
// so the compiler can safely allocate it on the stack.
func StackAlloc(x, y int) int {
	sum := x + y
	return sum
}

// HeapAlloc returns a pointer to a local variable.
// Because the pointer outlives the function's stack frame, the compiler
// must allocate the variable on the heap.
func HeapAlloc(value int) *int {
	v := value
	return &v
}

// SliceEscape creates a slice and returns it.
// The backing array escapes to the heap because the caller holds a
// reference after this function returns.
func SliceEscape(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = i * i
	}
	return s
}

// NoEscape takes a pointer and reads from it without letting it escape.
// The compiler can determine that the pointer is not stored or returned,
// so the pointed-to value does not need to be heap-allocated by the caller.
func NoEscape(p *int) int {
	return *p + 1
}

// sumWithPointer is used internally to show that passing a pointer to a
// function that does not let it escape allows stack allocation.
func sumWithPointer(a, b int) int {
	total := a + b
	return NoEscape(&total)
}

// SumViaPointer exposes sumWithPointer for testing.
func SumViaPointer(a, b int) int {
	return sumWithPointer(a, b)
}
