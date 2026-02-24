// Package stringbytes demonstrates the relationship between strings and byte
// slices in Go, including conversions, immutability, rune iteration, and the
// performance implications of each approach.
//
// In Go, a string is an immutable sequence of bytes. Converting between string
// and []byte creates a copy because strings are immutable and byte slices are
// mutable. The unsafe package can bypass this copy, but doing so breaks the
// immutability guarantee and can lead to undefined behavior.
package stringbytes

import "unicode/utf8"

// StringToBytes converts a string to a byte slice. This conversion allocates
// a new backing array and copies the string's bytes into it, because strings
// are immutable but byte slices are mutable.
func StringToBytes(s string) []byte {
	return []byte(s)
}

// BytesToString converts a byte slice to a string. This conversion allocates
// a new backing array and copies the bytes, ensuring the resulting string
// is truly immutable and independent of the original slice.
func BytesToString(b []byte) string {
	return string(b)
}

// StringImmutability demonstrates that strings are immutable in Go.
// You cannot assign to individual bytes of a string (s[i] = 'x' won't compile).
// To "modify" a string, you must convert to []byte, make changes, and convert
// back, which creates a new string.
func StringImmutability(s string, index int, newByte byte) string {
	b := []byte(s)
	if index >= 0 && index < len(b) {
		b[index] = newByte
	}
	return string(b)
}

// RuneIteration iterates over a string by runes (Unicode code points) and
// returns the slice of runes. Multi-byte characters such as Japanese kanji
// or emoji are each counted as a single rune, even though they occupy
// multiple bytes.
func RuneIteration(s string) []rune {
	runes := make([]rune, 0, utf8.RuneCountInString(s))
	for _, r := range s {
		runes = append(runes, r)
	}
	return runes
}

// RuneCount returns the number of runes (Unicode code points) in s.
// For ASCII strings, this equals len(s). For strings containing multi-byte
// characters, this is less than len(s).
func RuneCount(s string) int {
	return utf8.RuneCountInString(s)
}

// UnsafeConversion is documented here for educational purposes.
//
// Using unsafe.Pointer and unsafe.StringData/unsafe.SliceData, you can convert
// between string and []byte without copying. However, this breaks the
// immutability guarantee of strings: if the byte slice is mutated after
// conversion, the "immutable" string changes too, leading to undefined
// behavior. The standard library uses this technique internally (e.g.,
// strings.Builder) where it can guarantee safety, but application code should
// almost never use it.
//
// Example (DO NOT use in production):
//
//	import "unsafe"
//	b := []byte("hello")
//	s := unsafe.String(&b[0], len(b)) // no copy, shares memory
//	b[0] = 'H' // mutates s too -- undefined behavior
func UnsafeConversion() string {
	return "see package documentation for unsafe conversion details"
}
