// Package mapinternals demonstrates fundamental map operations, iteration
// order non-determinism, dynamic growth, and nil map behavior.
package mapinternals

// BasicMapOps demonstrates create, insert, lookup, delete, and existence
// check on a map. It returns the final state of the map.
func BasicMapOps() map[string]int {
	m := make(map[string]int)

	// Insert
	m["alpha"] = 1
	m["beta"] = 2
	m["gamma"] = 3

	// Delete
	delete(m, "gamma")

	// Update
	m["alpha"] = 10

	return m
}

// MapLookup checks whether a key exists using the comma-ok idiom.
// It returns the value and a boolean indicating whether the key was found.
func MapLookup(m map[string]int, key string) (int, bool) {
	val, ok := m[key]
	return val, ok
}

// IterationOrder iterates over the same map multiple times and returns
// the collected key orders. Go randomizes map iteration order, so
// different iterations may yield different key sequences.
func IterationOrder(m map[string]int, iterations int) [][]string {
	results := make([][]string, iterations)
	for i := 0; i < iterations; i++ {
		var keys []string
		for k := range m {
			keys = append(keys, k)
		}
		results[i] = keys
	}
	return results
}

// MapGrowth inserts n entries into a map, demonstrating that maps grow
// dynamically without requiring explicit capacity management.
// It returns the final length of the map.
func MapGrowth(n int) int {
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		m[i] = i * i
	}
	return len(m)
}

// NilMapRead reads from a nil map, which is safe and returns the zero
// value for the value type.
func NilMapRead(key string) int {
	var m map[string]int // nil map
	return m[key]
}

// NilMapWrite attempts to write to a nil map, which causes a panic.
// The nil map is obtained via nilStringIntMap so that the write target
// cannot be statically proven nil.
func NilMapWrite(key string, value int) {
	m := nilStringIntMap()
	m[key] = value // panics: assignment to entry in nil map
}

// nilStringIntMap returns a nil map. Extracting this into a function
// prevents static analyzers from proving the map is nil at the call site,
// while the runtime behavior (panic on write) remains the same.
func nilStringIntMap() map[string]int {
	return nil
}
