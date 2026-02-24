package panicrecover

import (
	"errors"
	"testing"
)

func TestSafeExecute_CatchesPanic(t *testing.T) {
	err := SafeExecute(func() {
		panic("something went wrong")
	})
	if err == nil {
		t.Fatal("SafeExecute: expected non-nil error for panicking function")
	}
	want := "panic: something went wrong"
	if err.Error() != want {
		t.Errorf("SafeExecute error = %q, want %q", err.Error(), want)
	}
}

func TestSafeExecute_NormalExecution(t *testing.T) {
	err := SafeExecute(func() {
		// no panic
	})
	if err != nil {
		t.Errorf("SafeExecute: expected nil error, got %v", err)
	}
}

func TestSafeExecute_PanicWithError(t *testing.T) {
	original := errors.New("original error")
	err := SafeExecute(func() {
		panic(original)
	})
	if err != original {
		t.Errorf("SafeExecute: expected original error, got %v", err)
	}
}

func TestRecoverInDefer(t *testing.T) {
	got := RecoverInDefer(func() {
		panic("caught")
	})
	if got != "caught" {
		t.Errorf("RecoverInDefer = %v, want \"caught\"", got)
	}
}

func TestRecoverInDefer_NoPanic(t *testing.T) {
	got := RecoverInDefer(func() {})
	if got != nil {
		t.Errorf("RecoverInDefer = %v, want nil", got)
	}
}

func TestRecoverOutsideDefer_Panics(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Error("RecoverOutsideDefer did not panic")
		}
	}()
	RecoverOutsideDefer()
}

func TestPanicToError_String(t *testing.T) {
	err := PanicToError("boom")
	want := "panic: boom"
	if err.Error() != want {
		t.Errorf("PanicToError(\"boom\") = %q, want %q", err.Error(), want)
	}
}

func TestPanicToError_Error(t *testing.T) {
	original := errors.New("wrapped")
	err := PanicToError(original)
	if err != original {
		t.Errorf("PanicToError(error) = %v, want %v", err, original)
	}
}

func TestPanicToError_Other(t *testing.T) {
	err := PanicToError(42)
	want := "panic: 42"
	if err.Error() != want {
		t.Errorf("PanicToError(42) = %q, want %q", err.Error(), want)
	}
}
