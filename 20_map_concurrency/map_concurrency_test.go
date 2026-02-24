package mapconcurrency

import (
	"sync"
	"testing"
)

func TestSafeMap_StoreAndLoad(t *testing.T) {
	sm := NewSafeMap()
	sm.Store("a", 1)
	sm.Store("b", 2)

	if v, ok := sm.Load("a"); !ok || v != 1 {
		t.Errorf("Load(\"a\") = (%d, %v), want (1, true)", v, ok)
	}
	if v, ok := sm.Load("b"); !ok || v != 2 {
		t.Errorf("Load(\"b\") = (%d, %v), want (2, true)", v, ok)
	}
	if _, ok := sm.Load("missing"); ok {
		t.Error("Load(\"missing\") should return ok=false")
	}
}

func TestSafeMap_Delete(t *testing.T) {
	sm := NewSafeMap()
	sm.Store("x", 99)
	sm.Delete("x")
	if _, ok := sm.Load("x"); ok {
		t.Error("Load(\"x\") should return ok=false after Delete")
	}
}

func TestSafeMap_ConcurrentAccess(t *testing.T) {
	sm := NewSafeMap()
	const goroutines = 100
	const opsPerGoroutine = 100

	var wg sync.WaitGroup
	wg.Add(goroutines * 2)

	// Concurrent writers
	for g := 0; g < goroutines; g++ {
		go func(id int) {
			defer wg.Done()
			for i := 0; i < opsPerGoroutine; i++ {
				sm.Store("shared", id*opsPerGoroutine+i)
			}
		}(g)
	}

	// Concurrent readers
	for g := 0; g < goroutines; g++ {
		go func() {
			defer wg.Done()
			for i := 0; i < opsPerGoroutine; i++ {
				sm.Load("shared")
			}
		}()
	}

	wg.Wait()

	// The map should still be functional after concurrent access.
	if _, ok := sm.Load("shared"); !ok {
		t.Error("\"shared\" key should exist after concurrent writes")
	}
}

func TestSafeMap_Len(t *testing.T) {
	sm := NewSafeMap()
	sm.Store("a", 1)
	sm.Store("b", 2)
	sm.Store("c", 3)
	if n := sm.Len(); n != 3 {
		t.Errorf("Len() = %d, want 3", n)
	}
}

func TestSyncMapStore(t *testing.T) {
	const n = 50
	got := SyncMapStore(n)
	if got != n {
		t.Errorf("SyncMapStore(%d) = %d, want %d", n, got, n)
	}
}

func TestSyncMapDelete(t *testing.T) {
	if SyncMapDelete() {
		t.Error("SyncMapDelete() = true, want false after deletion")
	}
}

func TestSyncMapRange(t *testing.T) {
	const n = 25
	got := SyncMapRange(n)
	if got != n {
		t.Errorf("SyncMapRange(%d) = %d, want %d", n, got, n)
	}
}

func TestSyncMap_ConcurrentAccess(t *testing.T) {
	var m sync.Map
	const goroutines = 100
	const opsPerGoroutine = 100

	var wg sync.WaitGroup
	wg.Add(goroutines * 2)

	// Concurrent writers
	for g := 0; g < goroutines; g++ {
		go func(id int) {
			defer wg.Done()
			for i := 0; i < opsPerGoroutine; i++ {
				m.Store(i, id)
			}
		}(g)
	}

	// Concurrent readers
	for g := 0; g < goroutines; g++ {
		go func() {
			defer wg.Done()
			for i := 0; i < opsPerGoroutine; i++ {
				m.Load(i)
			}
		}()
	}

	wg.Wait()

	// Verify all keys are present.
	count := 0
	m.Range(func(_, _ any) bool {
		count++
		return true
	})
	if count != opsPerGoroutine {
		t.Errorf("sync.Map has %d entries, want %d", count, opsPerGoroutine)
	}
}
