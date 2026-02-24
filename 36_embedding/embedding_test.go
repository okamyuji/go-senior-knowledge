package embedding

import "testing"

func TestPromotion(t *testing.T) {
	got := Promotion()
	want := "[SVC] started"
	if got != want {
		t.Errorf("Promotion() = %q, want %q", got, want)
	}
}

func TestPromotionCounter(t *testing.T) {
	s := Service{}
	if got := s.Increment(); got != 1 {
		t.Errorf("first Increment() = %d, want 1", got)
	}
	if got := s.Increment(); got != 2 {
		t.Errorf("second Increment() = %d, want 2", got)
	}
}

func TestAmbiguityResolution(t *testing.T) {
	got := AmbiguityResolution()
	want := "reader+writer"
	if got != want {
		t.Errorf("AmbiguityResolution() = %q, want %q", got, want)
	}
}

func TestAmbiguityExplicitAccess(t *testing.T) {
	rw := ReadWriter{}
	if got := rw.Reader.Describe(); got != "reader" {
		t.Errorf("rw.Reader.Describe() = %q, want %q", got, "reader")
	}
	if got := rw.Writer.Describe(); got != "writer" {
		t.Errorf("rw.Writer.Describe() = %q, want %q", got, "writer")
	}
}

func TestZeroValueEmbedding(t *testing.T) {
	got := ZeroValueEmbedding()
	want := "hello"
	if got != want {
		t.Errorf("ZeroValueEmbedding() = %q, want %q", got, want)
	}
}

func TestZeroValueCounter(t *testing.T) {
	var s Service
	if s.N != 0 {
		t.Errorf("zero value Service.N = %d, want 0", s.N)
	}
	s.Increment()
	if s.N != 1 {
		t.Errorf("after Increment(), Service.N = %d, want 1", s.N)
	}
}

func TestInterfaceSatisfaction(t *testing.T) {
	got := InterfaceSatisfaction()
	want := "Alice"
	if got != want {
		t.Errorf("InterfaceSatisfaction() = %q, want %q", got, want)
	}
}

func TestPersonImplementsStringer(t *testing.T) {
	p := Person{Name: Name{Value: "Bob"}, Age: 25}
	var _ Stringer = p // compile-time check
	if got := p.String(); got != "Bob" {
		t.Errorf("p.String() = %q, want %q", got, "Bob")
	}
}
