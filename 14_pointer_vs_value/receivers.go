// Package receivers demonstrates the difference between pointer and value
// receivers in Go, and how they affect method sets and interface satisfaction.
//
// Key rules:
//   - A value of type T has a method set containing only value receiver methods.
//   - A value of type *T has a method set containing both pointer and value
//     receiver methods.
//   - Therefore, *T satisfies any interface that T satisfies, but not vice versa.
//   - A value receiver method gets a copy of the receiver; modifications do not
//     affect the original.
//   - A pointer receiver method operates on the original value.
package receivers

import "fmt"

// Counter holds a simple integer count.
type Counter struct {
	N int
}

// Increment adds 1 to the counter. It uses a pointer receiver so the
// modification is visible to the caller.
func (c *Counter) Increment() {
	c.N++
}

// Value returns the current count. It uses a value receiver because it
// only reads the state without modifying it.
func (c Counter) Value() int {
	return c.N
}

// String returns a string representation. It uses a value receiver.
func (c Counter) String() string {
	return fmt.Sprintf("Counter(%d)", c.N)
}

// Incrementer is an interface with a pointer receiver method.
// Only *Counter satisfies this interface, not Counter.
type Incrementer interface {
	Increment()
}

// Valuer is an interface with a value receiver method.
// Both Counter and *Counter satisfy this interface.
type Valuer interface {
	Value() int
}

// InterfaceSatisfaction demonstrates which types satisfy which interfaces.
// It returns true if the type relationships are as expected.
func InterfaceSatisfaction() bool {
	var c Counter

	// Both Counter and *Counter satisfy Valuer.
	var _ Valuer = c
	var _ Valuer = &c

	// Only *Counter satisfies Incrementer.
	// var _ Incrementer = c  // This would NOT compile.
	var _ Incrementer = &c

	return true
}

// PointerVsValueSemantics demonstrates the difference between pointer and
// value receivers. It returns the counter values after calling methods on
// each receiver type.
func PointerVsValueSemantics() (pointerResult, valueResult int) {
	c := Counter{N: 0}

	// Pointer receiver: modifies the original.
	c.Increment()
	c.Increment()
	c.Increment()
	pointerResult = c.Value()

	// Value receiver gets a copy. Calling a hypothetical mutating
	// value-receiver method would not affect the original.
	copy := c
	valueReceiverDemo(copy)
	valueResult = copy.Value()

	return pointerResult, valueResult
}

// valueReceiverDemo simulates modifying a Counter received by value.
// Changes are lost because the function receives a copy.
func valueReceiverDemo(c Counter) {
	c.N = 999
}
