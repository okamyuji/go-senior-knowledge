package emptyiface

import "testing"

func TestTypeAssertSafe_Success(t *testing.T) {
	val, ok := TypeAssertSafe("hello")
	if !ok {
		t.Fatal("TypeAssertSafe(\"hello\"): expected ok=true")
	}
	if val != "hello" {
		t.Errorf("TypeAssertSafe(\"hello\") = %q, want %q", val, "hello")
	}
}

func TestTypeAssertSafe_Failure(t *testing.T) {
	val, ok := TypeAssertSafe(42)
	if ok {
		t.Fatal("TypeAssertSafe(42): expected ok=false for non-string")
	}
	if val != "" {
		t.Errorf("TypeAssertSafe(42) val = %q, want zero value \"\"", val)
	}
}

func TestTypeSwitch(t *testing.T) {
	tests := []struct {
		input any
		want  string
	}{
		{42, "int:42"},
		{"go", "string:go"},
		{true, "bool:true"},
		{3.14, "unknown:float64"},
	}
	for _, tt := range tests {
		got := TypeSwitch(tt.input)
		if got != tt.want {
			t.Errorf("TypeSwitch(%v) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestAssertionPanic_Panics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("AssertionPanic(\"not int\") did not panic")
		}
	}()
	AssertionPanic("not int")
}

func TestAssertionPanic_Success(t *testing.T) {
	got := AssertionPanic(99)
	if got != 99 {
		t.Errorf("AssertionPanic(99) = %d, want 99", got)
	}
}

func TestStoreAny(t *testing.T) {
	got := StoreAny(7)
	if got != 7 {
		t.Errorf("StoreAny(7) = %v, want 7", got)
	}

	got = StoreAny("boxed")
	if got != "boxed" {
		t.Errorf("StoreAny(\"boxed\") = %v, want \"boxed\"", got)
	}
}
