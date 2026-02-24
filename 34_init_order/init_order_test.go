package initorder

import (
	"testing"
)

func TestInitOrder(t *testing.T) {
	order := InitOrder()

	expected := []string{
		"var_first",
		"var_second",
		"init_1",
		"init_2",
		"init_3",
	}

	if len(order) != len(expected) {
		t.Fatalf("expected %d entries, got %d: %v", len(expected), len(order), order)
	}

	for i, want := range expected {
		if order[i] != want {
			t.Fatalf("index %d: expected %q, got %q\nfull order: %v", i, want, order[i], order)
		}
	}
}

func TestVariablesInitBeforeInit(t *testing.T) {
	order := InitOrder()

	// Verify that all var_ entries come before init_ entries.
	seenInit := false
	for _, entry := range order {
		if len(entry) > 4 && entry[:4] == "init" {
			seenInit = true
		}
		if seenInit && len(entry) > 3 && entry[:3] == "var" {
			t.Fatalf("variable %q initialized after init(), order: %v", entry, order)
		}
	}
}

func TestMultipleInitFunctions(t *testing.T) {
	order := InitOrder()

	// Count init entries.
	count := 0
	for _, entry := range order {
		if len(entry) > 4 && entry[:5] == "init_" {
			count++
		}
	}
	if count != 3 {
		t.Fatalf("expected 3 init functions, found %d in order: %v", count, order)
	}
}
