// Package pprofbasics demonstrates how to use runtime/pprof to capture CPU
// and heap profiles for performance analysis.
//
// After collecting a profile file, analyze it with:
//
//	go tool pprof cpu.prof        # interactive CLI
//	go tool pprof -http=:8080 cpu.prof  # web UI
//
// Common pprof commands:
//
//	top       - show top functions by resource usage
//	list foo  - show annotated source for function foo
//	web       - generate SVG call graph (requires graphviz)
//
// For continuous profiling in a web server, import net/http/pprof and access
// the /debug/pprof/ endpoints.
package pprofbasics

import (
	"io"
	"runtime/pprof"
)

// Fibonacci computes the nth Fibonacci number recursively.
// This intentionally uses the naive algorithm to generate CPU load for
// profiling demonstrations.
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// CPUIntensiveWork performs CPU-bound computation by calculating Fibonacci
// numbers. It returns the sum of Fibonacci(0) through Fibonacci(limit-1).
func CPUIntensiveWork(limit int) int {
	sum := 0
	for i := range limit {
		sum += Fibonacci(i)
	}
	return sum
}

// MemoryIntensiveWork allocates memory by growing a slice in a loop.
// It returns the total number of elements allocated.
func MemoryIntensiveWork(iterations, elementsPerIteration int) int {
	total := 0
	for range iterations {
		buf := make([]byte, elementsPerIteration)
		total += len(buf)
		// Prevent the compiler from optimizing away the allocation.
		if buf[0] != 0 {
			break
		}
	}
	return total
}

// StartCPUProfile begins CPU profiling and writes the profile data to w.
// Call StopCPUProfile when the workload is complete.
func StartCPUProfile(w io.Writer) error {
	return pprof.StartCPUProfile(w)
}

// StopCPUProfile stops an active CPU profile started by StartCPUProfile.
func StopCPUProfile() {
	pprof.StopCPUProfile()
}

// WriteHeapProfile writes a snapshot of the current heap allocations to w.
// It is useful for identifying memory leaks and high-allocation code paths.
func WriteHeapProfile(w io.Writer) error {
	return pprof.WriteHeapProfile(w)
}
