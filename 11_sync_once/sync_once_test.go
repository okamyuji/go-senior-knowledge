package synconce

import (
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
)

func TestInitOnce_CalledOnce(t *testing.T) {
	var callCount atomic.Int64

	loader := NewResourceLoader(func() *Resource {
		callCount.Add(1)
		return &Resource{Value: "initialized"}
	})

	const goroutines = 100
	var wg sync.WaitGroup
	wg.Add(goroutines)

	results := make([]*Resource, goroutines)
	for i := range goroutines {
		go func(idx int) {
			defer wg.Done()
			results[idx] = loader.InitOnce()
		}(i)
	}
	wg.Wait()

	if got := callCount.Load(); got != 1 {
		t.Errorf("init function called %d times, want 1", got)
	}

	for i, r := range results {
		if r != results[0] {
			t.Errorf("goroutine %d got different resource pointer", i)
		}
		if r.Value != "initialized" {
			t.Errorf("goroutine %d got Value=%q, want %q", i, r.Value, "initialized")
		}
	}
}

func TestOnceValue_CachesResult(t *testing.T) {
	var callCount atomic.Int64

	getter := OnceValue(func() string {
		callCount.Add(1)
		return "expensive-result"
	})

	const goroutines = 50
	var wg sync.WaitGroup
	wg.Add(goroutines)

	results := make([]string, goroutines)
	for i := range goroutines {
		go func(idx int) {
			defer wg.Done()
			results[idx] = getter()
		}(i)
	}
	wg.Wait()

	if got := callCount.Load(); got != 1 {
		t.Errorf("computation ran %d times, want 1", got)
	}

	for i, v := range results {
		if v != "expensive-result" {
			t.Errorf("goroutine %d got %q, want %q", i, v, "expensive-result")
		}
	}
}

func TestOnceValue_MultipleIndependent(t *testing.T) {
	g1 := OnceValue(func() string { return "first" })
	g2 := OnceValue(func() string { return "second" })

	if got := g1(); got != "first" {
		t.Errorf("g1() = %q, want %q", got, "first")
	}
	if got := g2(); got != "second" {
		t.Errorf("g2() = %q, want %q", got, "second")
	}
}

func TestInitOnce_ResourceValue(t *testing.T) {
	loader := NewResourceLoader(func() *Resource {
		return &Resource{Value: "db-connection"}
	})

	r1 := loader.InitOnce()
	r2 := loader.InitOnce()

	if r1 != r2 {
		t.Error("subsequent calls returned different pointers")
	}
	if r1.Value != "db-connection" {
		t.Errorf("Value = %q, want %q", r1.Value, "db-connection")
	}
}

func BenchmarkOnceValue(b *testing.B) {
	getter := OnceValue(func() string {
		return "result-" + strconv.Itoa(42)
	})
	// After the first call, subsequent calls have near-zero cost.
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = getter()
		}
	})
}
