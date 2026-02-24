package generics

import (
	"strconv"
	"testing"
)

func TestMinInt(t *testing.T) {
	if got := Min(3, 5); got != 3 {
		t.Fatalf("Min(3, 5) = %d, want 3", got)
	}
	if got := Min(10, 2); got != 2 {
		t.Fatalf("Min(10, 2) = %d, want 2", got)
	}
	if got := Min(-1, -5); got != -5 {
		t.Fatalf("Min(-1, -5) = %d, want -5", got)
	}
}

func TestMinFloat(t *testing.T) {
	if got := Min(3.14, 2.71); got != 2.71 {
		t.Fatalf("Min(3.14, 2.71) = %f, want 2.71", got)
	}
}

func TestMinString(t *testing.T) {
	if got := Min("apple", "banana"); got != "apple" {
		t.Fatalf("Min(apple, banana) = %s, want apple", got)
	}
	if got := Min("zebra", "alpha"); got != "alpha" {
		t.Fatalf("Min(zebra, alpha) = %s, want alpha", got)
	}
}

func TestMinEqual(t *testing.T) {
	if got := Min(7, 7); got != 7 {
		t.Fatalf("Min(7, 7) = %d, want 7", got)
	}
}

func TestMapIntToString(t *testing.T) {
	input := []int{1, 2, 3}
	result := Map(input, strconv.Itoa)

	expected := []string{"1", "2", "3"}
	if len(result) != len(expected) {
		t.Fatalf("length mismatch: got %d, want %d", len(result), len(expected))
	}
	for i, v := range result {
		if v != expected[i] {
			t.Fatalf("index %d: got %q, want %q", i, v, expected[i])
		}
	}
}

func TestMapSquare(t *testing.T) {
	input := []int{1, 2, 3, 4}
	result := Map(input, func(n int) int { return n * n })

	expected := []int{1, 4, 9, 16}
	for i, v := range result {
		if v != expected[i] {
			t.Fatalf("index %d: got %d, want %d", i, v, expected[i])
		}
	}
}

func TestMapEmptySlice(t *testing.T) {
	result := Map([]int{}, func(n int) string { return "" })
	if len(result) != 0 {
		t.Fatalf("expected empty result, got length %d", len(result))
	}
}

func TestFilterEven(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	result := Filter(input, func(n int) bool { return n%2 == 0 })

	expected := []int{2, 4, 6}
	if len(result) != len(expected) {
		t.Fatalf("length mismatch: got %d, want %d", len(result), len(expected))
	}
	for i, v := range result {
		if v != expected[i] {
			t.Fatalf("index %d: got %d, want %d", i, v, expected[i])
		}
	}
}

func TestFilterNoneMatch(t *testing.T) {
	input := []int{1, 3, 5}
	result := Filter(input, func(n int) bool { return n%2 == 0 })
	if len(result) != 0 {
		t.Fatalf("expected empty result, got %v", result)
	}
}

func TestFilterAllMatch(t *testing.T) {
	input := []string{"go", "generics", "great"}
	result := Filter(input, func(s string) bool { return len(s) > 0 })
	if len(result) != 3 {
		t.Fatalf("expected 3 results, got %d", len(result))
	}
}

func TestStackPushPop(t *testing.T) {
	var s Stack[int]
	s.Push(10)
	s.Push(20)
	s.Push(30)

	if s.Len() != 3 {
		t.Fatalf("expected length 3, got %d", s.Len())
	}

	val, ok := s.Pop()
	if !ok || val != 30 {
		t.Fatalf("expected (30, true), got (%d, %v)", val, ok)
	}

	val, ok = s.Pop()
	if !ok || val != 20 {
		t.Fatalf("expected (20, true), got (%d, %v)", val, ok)
	}

	val, ok = s.Pop()
	if !ok || val != 10 {
		t.Fatalf("expected (10, true), got (%d, %v)", val, ok)
	}

	_, ok = s.Pop()
	if ok {
		t.Fatal("expected Pop on empty stack to return false")
	}
}

func TestStackPeek(t *testing.T) {
	var s Stack[string]

	_, ok := s.Peek()
	if ok {
		t.Fatal("expected Peek on empty stack to return false")
	}

	s.Push("hello")
	s.Push("world")

	val, ok := s.Peek()
	if !ok || val != "world" {
		t.Fatalf("expected (world, true), got (%s, %v)", val, ok)
	}

	// Peek should not remove the element.
	if s.Len() != 2 {
		t.Fatalf("Peek should not change length, got %d", s.Len())
	}
}

func TestStackStringType(t *testing.T) {
	var s Stack[string]
	s.Push("a")
	s.Push("b")

	val, ok := s.Pop()
	if !ok || val != "b" {
		t.Fatalf("expected (b, true), got (%s, %v)", val, ok)
	}
}
