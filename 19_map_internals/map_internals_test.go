package mapinternals

import (
	"sort"
	"testing"
)

func TestBasicMapOps(t *testing.T) {
	m := BasicMapOps()

	// alpha was updated to 10
	if v, ok := m["alpha"]; !ok || v != 10 {
		t.Errorf("m[\"alpha\"] = %d, ok=%v; want 10, true", v, ok)
	}
	// beta remains 2
	if v, ok := m["beta"]; !ok || v != 2 {
		t.Errorf("m[\"beta\"] = %d, ok=%v; want 2, true", v, ok)
	}
	// gamma was deleted
	if _, ok := m["gamma"]; ok {
		t.Error("m[\"gamma\"] should not exist after delete")
	}
	if len(m) != 2 {
		t.Errorf("len(m) = %d, want 2", len(m))
	}
}

func TestMapLookup_Exists(t *testing.T) {
	m := map[string]int{"x": 42}
	val, ok := MapLookup(m, "x")
	if !ok || val != 42 {
		t.Errorf("MapLookup(m, \"x\") = (%d, %v), want (42, true)", val, ok)
	}
}

func TestMapLookup_Missing(t *testing.T) {
	m := map[string]int{"x": 42}
	val, ok := MapLookup(m, "missing")
	if ok || val != 0 {
		t.Errorf("MapLookup(m, \"missing\") = (%d, %v), want (0, false)", val, ok)
	}
}

func TestIterationOrder(t *testing.T) {
	m := map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5,
		"f": 6, "g": 7, "h": 8, "i": 9, "j": 10,
	}
	results := IterationOrder(m, 10)

	// Each iteration must contain all keys.
	for idx, keys := range results {
		if len(keys) != len(m) {
			t.Errorf("iteration %d: got %d keys, want %d", idx, len(keys), len(m))
		}
		sorted := make([]string, len(keys))
		copy(sorted, keys)
		sort.Strings(sorted)
		want := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
		for i, k := range sorted {
			if k != want[i] {
				t.Errorf("iteration %d: missing key %q", idx, want[i])
			}
		}
	}
}

func TestMapGrowth(t *testing.T) {
	const n = 10000
	got := MapGrowth(n)
	if got != n {
		t.Errorf("MapGrowth(%d) = %d, want %d", n, got, n)
	}
}

func TestNilMapRead(t *testing.T) {
	got := NilMapRead("anything")
	if got != 0 {
		t.Errorf("NilMapRead(\"anything\") = %d, want 0", got)
	}
}

func TestNilMapWrite_Panics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("NilMapWrite did not panic")
		}
	}()
	NilMapWrite("key", 1)
}
