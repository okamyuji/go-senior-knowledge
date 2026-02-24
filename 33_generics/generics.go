// Package generics demonstrates Go's type parameter features introduced in
// Go 1.18. It shows how to write generic functions (Min, Map, Filter) and
// generic data structures (Stack) using custom constraint interfaces.
//
// Generics reduce code duplication by letting a single function operate on
// multiple types while retaining compile-time type safety.
package generics

// Ordered is a constraint that permits any type whose underlying type is
// int, float64, or string. These types support the < operator, which is
// required for comparison-based algorithms like Min.
type Ordered interface {
	~int | ~float64 | ~string
}

// Min returns the smaller of two values. The type parameter T is
// constrained to Ordered, ensuring that the < operator is available.
func Min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Map transforms each element of the input slice by applying the given
// function fn. It returns a new slice of type []U with the same length
// as the input.
func Map[T any, U any](input []T, fn func(T) U) []U {
	result := make([]U, len(input))
	for i, v := range input {
		result[i] = fn(v)
	}
	return result
}

// Filter returns a new slice containing only the elements of the input
// for which the predicate function returns true.
func Filter[T any](input []T, pred func(T) bool) []T {
	var result []T
	for _, v := range input {
		if pred(v) {
			result = append(result, v)
		}
	}
	return result
}

// Stack is a generic LIFO data structure. The zero value is an empty
// stack ready to use.
type Stack[T any] struct {
	items []T
}

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

// Pop removes and returns the top element. It returns the value and true
// if the stack is non-empty, or the zero value and false if empty.
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

// Peek returns the top element without removing it. It returns the value
// and true if the stack is non-empty, or the zero value and false if empty.
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

// Len returns the number of elements in the stack.
func (s *Stack[T]) Len() int {
	return len(s.items)
}
