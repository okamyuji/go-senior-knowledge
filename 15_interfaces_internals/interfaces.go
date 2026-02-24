// Package interfaces demonstrates the internal behavior of Go interfaces,
// including nil interface semantics, dynamic dispatch, and type information.
//
// An interface value in Go is internally represented as two words:
//   - A pointer to type information (the "type" or "itab" pointer)
//   - A pointer to the underlying data
//
// A nil interface has both pointers set to nil. An interface holding a nil
// pointer has a non-nil type pointer but a nil data pointer. These two are
// NOT equal.
package interfaces

import (
	"fmt"
	"reflect"
)

// Stringer is a simple interface for demonstration.
type Stringer interface {
	String() string
}

// Dog implements Stringer.
type Dog struct {
	Name string
}

func (d *Dog) String() string {
	if d == nil {
		return "<nil dog>"
	}
	return "Dog: " + d.Name
}

// Cat implements Stringer.
type Cat struct {
	Name string
}

func (c Cat) String() string {
	return "Cat: " + c.Name
}

// NilInterface demonstrates the difference between a nil interface and
// an interface that holds a nil pointer.
//
// A nil interface has no type and no value:
//
//	var s Stringer       // s == nil is true
//
// An interface holding a nil pointer has a type but no value:
//
//	var d *Dog = nil
//	var s Stringer = d   // s == nil is false!
func NilInterface() (nilIface, nilPointerIface bool) {
	var s Stringer
	nilIface = (s == nil)

	var d *Dog
	s = d
	nilPointerIface = isNilInterface(s)

	return nilIface, nilPointerIface
}

// isNilInterface reports whether an interface value is nil. Because the
// concrete type information is lost at the call boundary, the comparison
// correctly reflects the runtime semantics: an interface holding a typed
// nil pointer is NOT nil.
func isNilInterface(s Stringer) bool {
	return s == nil
}

// InterfaceComparison compares interface values. Two interface values are
// equal if they have the same dynamic type and equal dynamic values.
func InterfaceComparison() (sameType, diffType, bothNil bool) {
	var a Stringer = Cat{Name: "Mimi"}
	var b Stringer = Cat{Name: "Mimi"}
	sameType = (a == b)

	var c Stringer = Cat{Name: "Mimi"}
	var d Stringer = &Dog{Name: "Mimi"}
	diffType = (c == d)

	var e, f Stringer
	bothNil = (e == f)

	return sameType, diffType, bothNil
}

// DynamicDispatch calls the String method through the Stringer interface.
// The actual method called depends on the dynamic type stored in the
// interface at runtime.
func DynamicDispatch(s Stringer) string {
	return s.String()
}

// TypeInfo returns the dynamic type name stored in an interface value
// using reflect and fmt.Sprintf.
func TypeInfo(v any) (typeName, reflectType string) {
	typeName = fmt.Sprintf("%T", v)
	if v != nil {
		reflectType = reflect.TypeOf(v).String()
	}
	return typeName, reflectType
}
