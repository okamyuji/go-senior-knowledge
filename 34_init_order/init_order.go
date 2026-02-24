// Package initorder demonstrates Go's package initialization order.
//
// Package-level variables are initialized in the order they are declared
// in the source file. After all variables are initialized, init() functions
// are executed in the order they appear. A single file can contain multiple
// init() functions, and they all run in declaration order.
//
// Key rules:
//   - init() runs before main() in the main package.
//   - Packages are initialized in dependency order: if package A imports
//     package B, then B's init runs before A's init.
//   - init() functions cannot be called or referenced from user code.
//
// Common traps:
//   - Side effects in init() (e.g., opening files, connecting to databases)
//     make testing and reuse difficult.
//   - Circular import dependencies cause a compile error.
//   - Doing too much work in init() slows down program startup and makes
//     failures harder to diagnose because init errors often manifest as
//     panics without clear stack traces.
//   - The order of init() across multiple files in the same package depends
//     on the lexicographic order of file names, which can be surprising.
package initorder

// initLog records the order in which variables and init() functions
// execute. Each step appends a string to this slice.
var initLog []string

// first is a package-level variable initialized before the init() calls.
// Variable declarations are evaluated in source order.
var first = recordVar("var_first")

// second is another package-level variable, initialized after first.
var second = recordVar("var_second")

// recordVar appends the label to initLog and returns it. This is used
// as a side effect during package-level variable initialization.
func recordVar(label string) string {
	initLog = append(initLog, label)
	return label
}

// First init() function in this file. Multiple init() functions in the
// same file are executed in the order they appear.
func init() {
	initLog = append(initLog, "init_1")
}

// Second init() function in this file. Runs after the first init().
func init() {
	initLog = append(initLog, "init_2")
}

// Third init() function. Runs after the second init().
func init() {
	initLog = append(initLog, "init_3")
}

// InitOrder returns the recorded initialization order as a slice of
// strings. The expected order is: var_first, var_second, init_1, init_2,
// init_3.
func InitOrder() []string {
	return initLog
}

// These variables exist to confirm that first and second are used.
var _ = first
var _ = second
