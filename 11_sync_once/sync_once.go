// Package synconce demonstrates sync.Once for one-time initialization.
//
// sync.Once guarantees that the function passed to Do is executed exactly once,
// even when called concurrently from multiple goroutines. This is useful for
// lazy initialization of shared resources.
//
// WARNING: Calling Do recursively from inside the function passed to Do causes
// a deadlock. sync.Once uses a mutex internally, and re-entering Do from the
// same goroutine will block forever waiting to acquire the already-held lock.
//
//	var once sync.Once
//	once.Do(func() {
//	    once.Do(func() { /* deadlock! */ })
//	})
package synconce

import "sync"

// Resource represents an expensive-to-create shared resource.
type Resource struct {
	Value string
}

// ResourceLoader manages one-time initialization of a Resource.
type ResourceLoader struct {
	once     sync.Once
	resource *Resource
	initFn   func() *Resource
}

// NewResourceLoader creates a ResourceLoader with the given initialization function.
func NewResourceLoader(initFn func() *Resource) *ResourceLoader {
	return &ResourceLoader{initFn: initFn}
}

// InitOnce initializes the resource exactly once, regardless of how many
// goroutines call it concurrently. All callers block until initialization
// completes, then receive the same resource.
func (rl *ResourceLoader) InitOnce() *Resource {
	rl.once.Do(func() {
		rl.resource = rl.initFn()
	})
	return rl.resource
}

// OnceValue demonstrates a pattern where an expensive computation is performed
// once and its result is cached. This is similar to sync.OnceValue introduced
// in Go 1.21, implemented manually for illustration.
func OnceValue(f func() string) func() string {
	var (
		once  sync.Once
		value string
	)
	return func() string {
		once.Do(func() {
			value = f()
		})
		return value
	}
}
