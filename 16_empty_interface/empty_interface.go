// Package emptyiface demonstrates the behavior of the empty interface (any),
// type assertions with the comma-ok idiom, type switches, and boxing.
package emptyiface

import "fmt"

// TypeAssertSafe performs a type assertion using the comma-ok idiom.
// It returns the asserted string value and a boolean indicating success.
func TypeAssertSafe(i any) (string, bool) {
	val, ok := i.(string)
	return val, ok
}

// TypeSwitch classifies the dynamic type stored in an any value and returns
// a descriptive string. It handles int, string, bool, and unknown types.
func TypeSwitch(i any) string {
	switch v := i.(type) {
	case int:
		return fmt.Sprintf("int:%d", v)
	case string:
		return fmt.Sprintf("string:%s", v)
	case bool:
		return fmt.Sprintf("bool:%t", v)
	default:
		return fmt.Sprintf("unknown:%T", v)
	}
}

// AssertionPanic demonstrates that a type assertion without the comma-ok
// idiom panics when the dynamic type does not match. It asserts i.(int)
// directly, which panics if i does not hold an int.
func AssertionPanic(i any) int {
	return i.(int)
}

// StoreAny accepts any value and returns it wrapped in an any.
// When a concrete value such as an int is passed, Go boxes it into an
// interface value consisting of a type pointer and a data pointer.
func StoreAny(v any) any {
	return v
}
