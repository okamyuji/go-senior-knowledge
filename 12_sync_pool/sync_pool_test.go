package syncpool

import (
	"bytes"
	"sync"
	"testing"
)

func TestBufferPool_ReturnsBuffer(t *testing.T) {
	buf := GetBuffer()
	if buf == nil {
		t.Fatal("GetBuffer returned nil")
	}
	buf.WriteString("hello")
	if got := buf.String(); got != "hello" {
		t.Errorf("buf.String() = %q, want %q", got, "hello")
	}
	PutBuffer(buf)
}

func TestBufferPool_ResetsOnGet(t *testing.T) {
	buf := GetBuffer()
	buf.WriteString("dirty data")
	PutBuffer(buf)

	// The next Get may or may not return the same buffer, but if it
	// does, it should be reset.
	buf2 := GetBuffer()
	if buf2.Len() != 0 {
		t.Errorf("buffer not reset: Len() = %d, want 0", buf2.Len())
	}
	PutBuffer(buf2)
}

func TestGetAndPut(t *testing.T) {
	got := GetAndPut("hello")
	want := "processed: hello"
	if got != want {
		t.Errorf("GetAndPut(%q) = %q, want %q", "hello", got, want)
	}
}

func TestGetAndPut_Concurrent(t *testing.T) {
	const goroutines = 100
	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := range goroutines {
		go func(n int) {
			defer wg.Done()
			result := GetAndPut("data")
			if result != "processed: data" {
				t.Errorf("unexpected result: %q", result)
			}
		}(i)
	}
	wg.Wait()
}

func TestPoolWithNew_CallsNewWhenEmpty(t *testing.T) {
	pool, newCalls := PoolWithNew()

	// First Get should trigger New because the pool is empty.
	obj := pool.Get()
	if obj == nil {
		t.Fatal("pool.Get() returned nil")
	}
	if *newCalls != 1 {
		t.Errorf("New called %d times, want 1", *newCalls)
	}

	// Put the object back and Get again; New should not be called.
	pool.Put(obj)
	before := *newCalls
	obj2 := pool.Get()
	if obj2 == nil {
		t.Fatal("pool.Get() returned nil after Put")
	}
	if *newCalls != before {
		t.Errorf("New called after Put: calls=%d, want %d", *newCalls, before)
	}
}

func TestPoolWithNew_ReturnsCorrectType(t *testing.T) {
	pool, _ := PoolWithNew()
	obj := pool.Get()
	buf, ok := obj.([]byte)
	if !ok {
		t.Fatalf("pool returned %T, want []byte", obj)
	}
	if cap(buf) != 1024 {
		t.Errorf("cap = %d, want 1024", cap(buf))
	}
}

func BenchmarkBufferPool(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			buf := GetBuffer()
			buf.WriteString("benchmark data")
			_ = buf.String()
			PutBuffer(buf)
		}
	})
}

func BenchmarkNewBuffer(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			buf := new(bytes.Buffer)
			buf.WriteString("benchmark data")
			_ = buf.String()
		}
	})
}
