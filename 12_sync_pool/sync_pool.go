// Package syncpool demonstrates sync.Pool for reusing temporary objects.
//
// sync.Pool is a set of temporary objects that can be saved and retrieved
// individually. It is safe for concurrent use.
//
// IMPORTANT: sync.Pool is NOT a connection pool or a general-purpose object
// cache. Objects in the pool may be removed at any time without notification,
// typically during garbage collection. Do not store objects whose lifecycle
// must be managed explicitly (e.g., database connections, file handles).
//
// sync.Pool is best suited for reducing allocation pressure by reusing
// short-lived temporary objects such as byte buffers or scratch structs.
package syncpool

import (
	"bytes"
	"sync"
)

// BufferPool is a pool of reusable bytes.Buffer instances.
// It reduces GC pressure by recycling buffers instead of allocating new ones.
var BufferPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

// GetBuffer retrieves a buffer from the pool and resets it for reuse.
func GetBuffer() *bytes.Buffer {
	buf := BufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	return buf
}

// PutBuffer returns a buffer to the pool for reuse.
// The buffer should not be used after calling PutBuffer.
func PutBuffer(buf *bytes.Buffer) {
	BufferPool.Put(buf)
}

// GetAndPut demonstrates the typical pool usage pattern: get a buffer,
// use it, then return it. The function writes data to a pooled buffer
// and returns the resulting string.
func GetAndPut(data string) string {
	buf := GetBuffer()
	defer PutBuffer(buf)

	buf.WriteString("processed: ")
	buf.WriteString(data)
	return buf.String()
}

// PoolWithNew demonstrates that the New function is called only when the
// pool has no available objects. It returns a pool and a counter that
// tracks how many times New was called.
func PoolWithNew() (*sync.Pool, *int) {
	newCalls := new(int)
	pool := &sync.Pool{
		New: func() any {
			*newCalls++
			return make([]byte, 0, 1024)
		},
	}
	return pool, newCalls
}
