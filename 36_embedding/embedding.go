// Package embedding demonstrates Go's struct embedding mechanism including
// method promotion, ambiguity resolution, zero-value usability, the
// distinction from class-based inheritance, and implicit interface satisfaction.
//
// Go embedding is composition, not inheritance. When a struct embeds another
// struct, the embedded type's methods are promoted to the outer type. However,
// the receiver of those methods remains the embedded type, not the outer type.
// There is no polymorphic dispatch: the embedded type has no knowledge of
// the outer type. This is fundamentally different from class-based inheritance
// where a base class method can be overridden and the derived class is passed
// as the receiver.
package embedding

import "fmt"

// Logger provides a simple logging capability.
type Logger struct {
	Prefix string
}

// Log returns a formatted log message with the prefix.
func (l Logger) Log(msg string) string {
	if l.Prefix == "" {
		return msg
	}
	return fmt.Sprintf("[%s] %s", l.Prefix, msg)
}

// Counter tracks a count value.
type Counter struct {
	N int
}

// Increment adds 1 to the counter and returns the new value.
func (c *Counter) Increment() int {
	c.N++
	return c.N
}

// Service embeds both Logger and Counter. Their methods are promoted
// to Service, so callers can invoke s.Log(msg) and s.Increment() directly.
type Service struct {
	Logger
	Counter
}

// Promotion demonstrates that an embedded struct's method is callable
// on the outer struct directly. It returns the result of Logger.Log
// called through a Service value.
func Promotion() string {
	s := Service{
		Logger: Logger{Prefix: "SVC"},
	}
	return s.Log("started")
}

// ----- Ambiguity Resolution -----

// Reader provides a Describe method.
type Reader struct{}

// Describe returns a description for Reader.
func (Reader) Describe() string { return "reader" }

// Writer provides a Describe method with the same name as Reader.
type Writer struct{}

// Describe returns a description for Writer.
func (Writer) Describe() string { return "writer" }

// ReadWriter embeds both Reader and Writer. Because both have a Describe
// method, calling rw.Describe() would be ambiguous. The outer struct must
// define its own Describe to resolve the ambiguity.
type ReadWriter struct {
	Reader
	Writer
}

// Describe resolves the ambiguity by providing an explicit method on
// ReadWriter. Without this method, rw.Describe() would not compile.
func (rw ReadWriter) Describe() string {
	return rw.Reader.Describe() + "+" + rw.Writer.Describe()
}

// AmbiguityResolution shows that the outer struct's Describe method
// is called, resolving the conflict between the two embedded types.
func AmbiguityResolution() string {
	rw := ReadWriter{}
	return rw.Describe()
}

// ----- Zero-Value Embedding -----

// ZeroValueEmbedding demonstrates that an embedded struct is usable
// even when the outer struct is initialized as a zero value. Logger's
// zero value has Prefix="" which produces a valid log message.
func ZeroValueEmbedding() string {
	var s Service
	return s.Log("hello")
}

// ----- Interface Satisfaction via Embedding -----

// Stringer is an interface requiring a String method.
type Stringer interface {
	String() string
}

// Name holds a name string.
type Name struct {
	Value string
}

// String returns the name value, satisfying the Stringer interface.
func (n Name) String() string {
	return n.Value
}

// Person embeds Name. Because Name has a String method, Person also
// satisfies the Stringer interface without declaring its own String.
type Person struct {
	Name
	Age int
}

// InterfaceSatisfaction demonstrates that embedding Name makes Person
// satisfy the Stringer interface. The function accepts a Stringer and
// returns its String() result.
func InterfaceSatisfaction() string {
	p := Person{
		Name: Name{Value: "Alice"},
		Age:  30,
	}
	var s Stringer = p
	return s.String()
}
