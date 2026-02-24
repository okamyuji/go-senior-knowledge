package jsonpitfalls

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestNumberPrecisionLoss(t *testing.T) {
	// 2^53 + 1 = 9007199254740993 cannot be represented exactly as float64.
	input := `{"value": 9007199254740993}`

	val, err := NumberPrecision(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	f, ok := val.(float64)
	if !ok {
		t.Fatalf("expected float64, got %T", val)
	}

	// float64 rounds 9007199254740993 to 9007199254740992.
	// Direct int literal comparison would also round, so compare strings.
	str := fmt.Sprintf("%.0f", f)
	if str == "9007199254740993" {
		t.Fatal("expected precision loss, but the number was preserved exactly")
	}
	if str != "9007199254740992" {
		t.Fatalf("expected 9007199254740992, got %s", str)
	}
}

func TestUseDecoderPreservesPrecision(t *testing.T) {
	input := `{"value": 9007199254740993}`

	num, err := UseDecoder(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if num.String() != "9007199254740993" {
		t.Fatalf("expected 9007199254740993, got %s", num.String())
	}

	// Convert to int64 to verify the value is preserved.
	n, err := num.Int64()
	if err != nil {
		t.Fatalf("unexpected error converting to int64: %v", err)
	}
	if n != 9007199254740993 {
		t.Fatalf("expected 9007199254740993, got %d", n)
	}
}

func TestOmitemptyWithZeroValues(t *testing.T) {
	d := OmitemptyDemo{} // all zero values

	result, err := OmitemptyBehavior(d)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Fields with omitempty and zero values should be absent.
	if strings.Contains(result, `"name"`) {
		t.Error("name with omitempty should be omitted when empty")
	}
	if strings.Contains(result, `"age"`) {
		t.Error("age with omitempty should be omitted when zero")
	}
	if strings.Contains(result, `"active"`) {
		t.Error("active with omitempty should be omitted when false")
	}
	if strings.Contains(result, `"score"`) {
		t.Error("score with omitempty should be omitted when nil")
	}

	// Fields without omitempty should always be present.
	if !strings.Contains(result, `"address"`) {
		t.Error("address without omitempty should be present even when empty")
	}
	if !strings.Contains(result, `"balance"`) {
		t.Error("balance without omitempty should be present even when zero")
	}
}

func TestOmitemptyNilVsZeroPointer(t *testing.T) {
	zero := 0
	d := OmitemptyDemo{Score: &zero}

	result, err := OmitemptyBehavior(d)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// A pointer to zero is NOT nil, so it should be present despite omitempty.
	if !strings.Contains(result, `"score"`) {
		t.Error("score pointing to 0 should NOT be omitted (pointer is non-nil)")
	}

	var parsed map[string]json.RawMessage
	if err := json.Unmarshal([]byte(result), &parsed); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}
	if string(parsed["score"]) != "0" {
		t.Fatalf("expected score=0, got %s", string(parsed["score"]))
	}
}

func TestUnexportedFieldsIgnored(t *testing.T) {
	result, err := MarshalUnexported("visible", "hidden")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(result, `"exported"`) {
		t.Error("exported field should be present in JSON")
	}
	if strings.Contains(result, "hidden") {
		t.Error("unexported field value should not appear in JSON output")
	}
	if strings.Contains(result, "unexported") {
		t.Error("unexported field key should not appear in JSON output")
	}
}
