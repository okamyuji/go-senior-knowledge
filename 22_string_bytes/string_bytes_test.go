package stringbytes

import "testing"

func TestStringToBytes(t *testing.T) {
	s := "hello"
	b := StringToBytes(s)
	if string(b) != s {
		t.Errorf("StringToBytes(%q) = %q, want %q", s, b, s)
	}
	// Mutating b must not affect the original string.
	b[0] = 'H'
	if s != "hello" {
		t.Error("mutating byte slice affected original string")
	}
}

func TestBytesToString(t *testing.T) {
	b := []byte{0x47, 0x6f} // "Go"
	s := BytesToString(b)
	if s != "Go" {
		t.Errorf("BytesToString(%v) = %q, want %q", b, s, "Go")
	}
	// Mutating b must not affect the resulting string.
	b[0] = 0x00
	if s != "Go" {
		t.Error("mutating byte slice affected resulting string")
	}
}

func TestStringImmutability(t *testing.T) {
	original := "hello"
	modified := StringImmutability(original, 0, 'H')
	if modified != "Hello" {
		t.Errorf("StringImmutability(%q, 0, 'H') = %q, want %q", original, modified, "Hello")
	}
	if original != "hello" {
		t.Error("original string was modified")
	}
}

func TestRuneIteration(t *testing.T) {
	tests := []struct {
		input string
		count int
	}{
		{"hello", 5},
		{"Go言語", 4},
		{"", 0},
	}
	for _, tt := range tests {
		runes := RuneIteration(tt.input)
		if len(runes) != tt.count {
			t.Errorf("RuneIteration(%q) returned %d runes, want %d", tt.input, len(runes), tt.count)
		}
	}
}

func TestRuneCount(t *testing.T) {
	tests := []struct {
		input    string
		wantRune int
		wantByte int
	}{
		{"hello", 5, 5},
		{"Go言語", 4, 8},
		{"", 0, 0},
	}
	for _, tt := range tests {
		gotRune := RuneCount(tt.input)
		gotByte := len(tt.input)
		if gotRune != tt.wantRune {
			t.Errorf("RuneCount(%q) = %d, want %d", tt.input, gotRune, tt.wantRune)
		}
		if gotByte != tt.wantByte {
			t.Errorf("len(%q) = %d, want %d", tt.input, gotByte, tt.wantByte)
		}
	}
}
