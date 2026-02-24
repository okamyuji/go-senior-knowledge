// Package cgooverhead explains cgo overhead conceptually and simulates the
// cost of crossing the Go/C boundary, without actually using cgo.
//
// Why cgo calls are expensive:
//
//  1. Thread pinning: A goroutine calling C must be locked to an OS thread
//     (runtime.LockOSThread) because C code is not aware of Go's cooperative
//     scheduling. This prevents the scheduler from multiplexing other
//     goroutines onto that thread.
//
//  2. Stack switch: Go goroutines use segmented/growable stacks starting at a
//     few KB. C code expects a large, contiguous stack. The runtime must switch
//     from the goroutine stack to a system stack (typically 8 MB) before
//     entering C.
//
//  3. No GC during C calls: The garbage collector cannot inspect or move
//     memory that C code is using. All Go pointers passed to C must be pinned,
//     and the GC must wait for C calls to complete before it can fully collect.
//
//  4. No preemption: The Go scheduler cannot preempt a goroutine while it is
//     executing C code. Long-running C functions block the OS thread entirely,
//     reducing the effective GOMAXPROCS.
//
//  5. Function call overhead: Each cgo call involves multiple runtime
//     transitions (Go -> C ABI adapter -> C -> C ABI adapter -> Go) with
//     associated bookkeeping.
//
// As a rule of thumb, a single cgo call costs roughly 50-100 ns of overhead
// compared to a pure Go function call (~1-2 ns). Avoid cgo in hot loops; batch
// work on the C side instead.
package cgooverhead

import "runtime"

// PureGoAdd performs a simple addition entirely in Go. This represents the
// baseline cost of a function call with no cgo boundary crossing.
func PureGoAdd(a, b int) int {
	return a + b
}

// SimulateCgoBoundary simulates the overhead of a cgo call by performing the
// same addition but with the extra steps the runtime takes for a real C call:
// locking the goroutine to the current OS thread and unlocking afterward. In
// a real cgo call the runtime also switches stacks, disables preemption, and
// pins Go pointers, but those operations are not accessible from pure Go.
func SimulateCgoBoundary(a, b int) int {
	runtime.LockOSThread()
	result := a + b
	runtime.UnlockOSThread()
	return result
}

// BatchPureGoAdd adds pairs of numbers from two slices element-wise and
// returns the results. This demonstrates the idiomatic approach: perform
// all work in Go without per-element boundary crossings.
func BatchPureGoAdd(a, b []int) []int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	result := make([]int, n)
	for i := range n {
		result[i] = a[i] + b[i]
	}
	return result
}

// BatchSimulateCgoBoundary adds pairs of numbers but locks the OS thread
// once for the entire batch, simulating the recommended pattern of batching
// work on the C side to amortize the cgo call overhead.
func BatchSimulateCgoBoundary(a, b []int) []int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	runtime.LockOSThread()
	result := make([]int, n)
	for i := range n {
		result[i] = a[i] + b[i]
	}
	runtime.UnlockOSThread()
	return result
}
