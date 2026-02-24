// Package preemption demonstrates cooperative yielding, asynchronous
// preemption (Go 1.14+), and concurrent task completion.
package preemption

import (
	"runtime"
	"sync"
)

// CooperativeYield runs iterations loops, calling runtime.Gosched on each
// iteration to voluntarily yield the processor to other goroutines. It
// returns the number of completed iterations.
func CooperativeYield(iterations int) int {
	count := 0
	for range iterations {
		runtime.Gosched()
		count++
	}
	return count
}

// PreemptibleWork performs a CPU-bound computation (summing 1 to n) that
// can be preempted by the Go runtime's asynchronous preemption mechanism.
// It returns the computed sum.
func PreemptibleWork(n int) int64 {
	var sum int64
	for i := 1; i <= n; i++ {
		sum += int64(i)
	}
	return sum
}

// ConcurrentTasks launches n goroutines, each computing the sum of 1 to
// (index+1)*1000. All goroutines run concurrently and are subject to the
// scheduler's preemption. Results are collected and returned as a slice.
func ConcurrentTasks(n int) []int64 {
	type result struct {
		index int
		value int64
	}

	ch := make(chan result, n)
	var wg sync.WaitGroup
	wg.Add(n)

	for i := range n {
		go func(i int) {
			defer wg.Done()
			limit := (i + 1) * 1000
			ch <- result{index: i, value: PreemptibleWork(limit)}
		}(i)
	}

	wg.Wait()
	close(ch)

	results := make([]int64, n)
	for r := range ch {
		results[r.index] = r.value
	}
	return results
}
