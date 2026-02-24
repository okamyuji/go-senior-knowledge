// Package mapconcurrency demonstrates thread-safe map patterns using
// sync.RWMutex and sync.Map, and documents why concurrent access to
// regular maps causes fatal crashes.
//
// A regular Go map is not safe for concurrent use. If multiple goroutines
// read and write to the same map simultaneously without synchronization,
// the runtime detects the race and intentionally crashes the program with
// "concurrent map read and map write" or "concurrent map writes". This is
// a fatal error, not a recoverable panic.
package mapconcurrency

import "sync"

// SafeMap is a thread-safe map wrapper using sync.RWMutex.
// It provides concurrent read access via RLock and exclusive write
// access via Lock.
type SafeMap struct {
	mu sync.RWMutex
	m  map[string]int
}

// NewSafeMap creates an initialized SafeMap.
func NewSafeMap() *SafeMap {
	return &SafeMap{m: make(map[string]int)}
}

// Store sets a key-value pair in the map.
func (sm *SafeMap) Store(key string, value int) {
	sm.mu.Lock()
	sm.m[key] = value
	sm.mu.Unlock()
}

// Load retrieves a value by key. It returns the value and a boolean
// indicating whether the key was found.
func (sm *SafeMap) Load(key string) (int, bool) {
	sm.mu.RLock()
	val, ok := sm.m[key]
	sm.mu.RUnlock()
	return val, ok
}

// Delete removes a key from the map.
func (sm *SafeMap) Delete(key string) {
	sm.mu.Lock()
	delete(sm.m, key)
	sm.mu.Unlock()
}

// Len returns the number of entries in the map.
func (sm *SafeMap) Len() int {
	sm.mu.RLock()
	n := len(sm.m)
	sm.mu.RUnlock()
	return n
}

// SyncMapStore demonstrates sync.Map's Store and Load operations.
// It stores n key-value pairs and returns the count of successfully
// loaded values.
func SyncMapStore(n int) int {
	var m sync.Map
	for i := 0; i < n; i++ {
		m.Store(i, i*10)
	}

	count := 0
	for i := 0; i < n; i++ {
		if _, ok := m.Load(i); ok {
			count++
		}
	}
	return count
}

// SyncMapDelete demonstrates sync.Map's Delete operation.
// It stores a key, deletes it, and returns whether the key still exists.
func SyncMapDelete() bool {
	var m sync.Map
	m.Store("key", "value")
	m.Delete("key")
	_, ok := m.Load("key")
	return ok
}

// SyncMapRange demonstrates iterating over sync.Map using Range.
// It stores n entries and counts them via Range, returning the count.
func SyncMapRange(n int) int {
	var m sync.Map
	for i := 0; i < n; i++ {
		m.Store(i, true)
	}

	count := 0
	m.Range(func(key, value any) bool {
		count++
		return true // continue iteration
	})
	return count
}
