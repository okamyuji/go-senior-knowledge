// Package deferinternals demonstrates defer execution order, closure capture
// gotchas, resource accumulation in loops, and named return modification.
package deferinternals

// DeferOrder registers multiple defers and returns the order in which they
// executed. Defers run in LIFO (last-in, first-out) order.
func DeferOrder() (order []string) {
	defer func() { order = append(order, "first") }()
	defer func() { order = append(order, "second") }()
	defer func() { order = append(order, "third") }()
	return
}

// DeferWithClosureBug demonstrates a closure capture gotcha. A single
// variable (shared) is captured by reference by multiple closures. All
// closures see the last value written to the variable.
//
// Note: Go 1.22 changed for-loop variable scoping so each iteration gets
// its own copy. This example uses a shared variable outside the loop to
// illustrate the underlying concept that still applies to non-loop variables.
func DeferWithClosureBug(n int) []int {
	var result []int
	var defers []func()
	shared := 0
	for range n {
		shared++
		defers = append(defers, func() {
			result = append(result, shared)
		})
	}
	for _, fn := range defers {
		fn()
	}
	return result
}

// DeferWithClosureFix demonstrates the correct pattern: creating a local
// copy per iteration so each closure captures its own independent value.
func DeferWithClosureFix(n int) []int {
	var result []int
	var defers []func()
	shared := 0
	for range n {
		shared++
		val := shared // local copy for this iteration
		defers = append(defers, func() {
			result = append(result, val)
		})
	}
	for _, fn := range defers {
		fn()
	}
	return result
}

// DeferInLoop shows why placing defer inside a hot loop is problematic.
// Each iteration pushes a deferred call onto the stack; none execute until
// the enclosing function returns. This returns the count of accumulated
// deferred calls that executed after the loop completed.
func DeferInLoop(n int) int {
	count := 0
	func() {
		for range n {
			defer func() { count++ }()
		}
		// All n deferred calls are stacked up and execute here when
		// the anonymous function returns.
	}()
	return count
}

// DeferModifyReturn uses a named return value that is modified inside a
// deferred function. The deferred function runs after the return statement
// evaluates but before the function actually returns to the caller,
// allowing it to change the returned value.
func DeferModifyReturn() (result int) {
	defer func() {
		result += 10
	}()
	return 1 // result is set to 1, then defer adds 10 -> returns 11
}
