package receivers

import "testing"

func TestCounter_Increment(t *testing.T) {
	c := Counter{N: 0}
	c.Increment()
	c.Increment()
	if c.N != 2 {
		t.Errorf("after 2 increments: N = %d, want 2", c.N)
	}
}

func TestCounter_Value(t *testing.T) {
	c := Counter{N: 42}
	if got := c.Value(); got != 42 {
		t.Errorf("Value() = %d, want 42", got)
	}
}

func TestCounter_ValueReceiverGetsCopy(t *testing.T) {
	c := Counter{N: 10}

	// Value() uses a value receiver, so it gets a copy.
	// Even if the value receiver tried to modify N, the original
	// would remain unchanged. We verify the original is intact.
	_ = c.Value()
	if c.N != 10 {
		t.Errorf("N changed after Value(): got %d, want 10", c.N)
	}
}

func TestCounter_PointerReceiverModifiesOriginal(t *testing.T) {
	c := Counter{N: 0}
	c.Increment()
	if c.N != 1 {
		t.Errorf("after Increment: N = %d, want 1", c.N)
	}
}

func TestInterfaceSatisfaction(t *testing.T) {
	if !InterfaceSatisfaction() {
		t.Error("InterfaceSatisfaction returned false")
	}
}

func TestValuerInterface_Value(t *testing.T) {
	c := Counter{N: 5}

	// Counter value satisfies Valuer.
	var v Valuer = c
	if got := v.Value(); got != 5 {
		t.Errorf("Valuer.Value() = %d, want 5", got)
	}
}

func TestValuerInterface_Pointer(t *testing.T) {
	c := Counter{N: 7}

	// *Counter also satisfies Valuer.
	var v Valuer = &c
	if got := v.Value(); got != 7 {
		t.Errorf("Valuer.Value() = %d, want 7", got)
	}
}

func TestIncrementerInterface(t *testing.T) {
	c := Counter{N: 0}

	// Only *Counter satisfies Incrementer.
	var inc Incrementer = &c
	inc.Increment()
	if c.N != 1 {
		t.Errorf("after Incrementer.Increment: N = %d, want 1", c.N)
	}
}

func TestPointerVsValueSemantics(t *testing.T) {
	pResult, vResult := PointerVsValueSemantics()

	if pResult != 3 {
		t.Errorf("pointer result = %d, want 3", pResult)
	}
	// valueReceiverDemo modifies a copy, so original stays at 3.
	if vResult != 3 {
		t.Errorf("value result = %d, want 3", vResult)
	}
}

func TestCounter_String(t *testing.T) {
	c := Counter{N: 42}
	want := "Counter(42)"
	if got := c.String(); got != want {
		t.Errorf("String() = %q, want %q", got, want)
	}
}
