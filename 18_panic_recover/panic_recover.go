// Package panicrecover demonstrates panic recovery patterns including
// SafeExecute, recover placement rules, and panic value conversion.
//
// Important: recover only works when called directly inside a deferred
// function. Calling recover outside a defer or in a nested function
// within a defer does not catch panics. Additionally, recover cannot
// catch panics originating from a different goroutine.
package panicrecover

import "fmt"

// SafeExecute runs fn and recovers from any panic, returning it as an error.
// If fn completes without panicking, it returns nil.
func SafeExecute(fn func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = PanicToError(r)
		}
	}()
	fn()
	return nil
}

// RecoverInDefer demonstrates that recover must be called inside a deferred
// function to work. It returns the recovered value if a panic occurred,
// or nil otherwise.
func RecoverInDefer(fn func()) (recovered any) {
	defer func() {
		recovered = recover()
	}()
	fn()
	return nil
}

// RecoverOutsideDefer shows that calling recover outside a deferred function
// always returns nil, even when a panic is active. This function panics
// because the recover call has no effect.
func RecoverOutsideDefer() {
	// This recover is not inside a deferred function, so it does nothing.
	_ = recover()
	panic("this panic is not caught")
}

// PanicToError converts a recovered panic value to an error.
// It handles error, string, and arbitrary types.
func PanicToError(r any) error {
	switch v := r.(type) {
	case error:
		return v
	case string:
		return fmt.Errorf("panic: %s", v)
	default:
		return fmt.Errorf("panic: %v", v)
	}
}
