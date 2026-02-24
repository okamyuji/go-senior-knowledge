// Package jsonpitfalls demonstrates common pitfalls when working with
// encoding/json in Go, including float64 precision loss for large integers,
// the UseNumber decoder option, omitempty behavior with zero values and nil
// pointers, and the fact that unexported struct fields are invisible to the
// JSON encoder/decoder.
package jsonpitfalls

import (
	"bytes"
	"encoding/json"
)

// NumberPrecision unmarshals a JSON string containing a number field into
// an interface{}. Because JSON numbers are mapped to float64 by default,
// large integers (beyond 2^53) lose precision. The function returns the
// raw value stored in the interface{}.
//
// Example: the integer 9007199254740993 (2^53 + 1) cannot be represented
// exactly as float64, so it rounds to 9007199254740992.
func NumberPrecision(jsonStr string) (any, error) {
	var result map[string]any
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return nil, err
	}
	return result["value"], nil
}

// UseDecoder unmarshals the same JSON using json.NewDecoder with UseNumber
// enabled. Numbers are decoded as json.Number (a string type) instead of
// float64, preserving the original textual representation without any
// precision loss.
func UseDecoder(jsonStr string) (json.Number, error) {
	dec := json.NewDecoder(bytes.NewReader([]byte(jsonStr)))
	dec.UseNumber()

	var result map[string]any
	if err := dec.Decode(&result); err != nil {
		return "", err
	}
	num, ok := result["value"].(json.Number)
	if !ok {
		return "", nil
	}
	return num, nil
}

// OmitemptyDemo holds fields that demonstrate omitempty behavior.
// With omitempty, zero values (0, "", false, nil) are omitted from the
// JSON output. Without omitempty, zero values are included.
type OmitemptyDemo struct {
	Name    string  `json:"name,omitempty"`
	Age     int     `json:"age,omitempty"`
	Active  bool    `json:"active,omitempty"`
	Score   *int    `json:"score,omitempty"`
	Address string  `json:"address"`
	Balance float64 `json:"balance"`
}

// OmitemptyBehavior marshals the given struct to JSON and returns the
// result as a string. Fields tagged with omitempty are omitted when they
// hold their zero value, while fields without omitempty are always present.
//
// Note that for pointer fields, nil is the zero value (omitted), but a
// pointer to zero (e.g., new(int) pointing to 0) is NOT omitted because
// the pointer itself is non-nil.
func OmitemptyBehavior(d OmitemptyDemo) (string, error) {
	b, err := json.Marshal(d)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// UnexportedFieldsDemo contains both exported and unexported fields.
// The JSON encoder ignores unexported fields entirely. This is a common
// source of confusion when struct fields are accidentally lowercase.
//
//	type UnexportedFieldsDemo struct {
//	    Exported   string `json:"exported"`
//	    unexported string `json:"unexported"` // IGNORED by encoding/json
//	}
//
// Even with a json tag, unexported fields are never included in the
// output and never populated during decoding.
type UnexportedFieldsDemo struct {
	Exported   string `json:"exported"`
	unexported string //nolint // ignored by encoding/json
}

// MarshalUnexported marshals the demo struct. The unexported field is
// silently dropped from the output.
func MarshalUnexported(exported, unexported string) (string, error) {
	d := UnexportedFieldsDemo{
		Exported:   exported,
		unexported: unexported,
	}
	b, err := json.Marshal(d)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
