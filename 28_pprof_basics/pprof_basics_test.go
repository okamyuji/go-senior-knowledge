package pprofbasics

import (
	"io"
	"os"
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{5, 5},
		{10, 55},
	}
	for _, tt := range tests {
		got := Fibonacci(tt.n)
		if got != tt.want {
			t.Errorf("Fibonacci(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestCPUIntensiveWork(t *testing.T) {
	// Sum of Fibonacci(0)..Fibonacci(9) = 0+1+1+2+3+5+8+13+21+34 = 88
	got := CPUIntensiveWork(10)
	if got != 88 {
		t.Errorf("CPUIntensiveWork(10) = %d, want 88", got)
	}
}

func TestMemoryIntensiveWork(t *testing.T) {
	got := MemoryIntensiveWork(100, 1024)
	want := 100 * 1024
	if got != want {
		t.Errorf("MemoryIntensiveWork(100, 1024) = %d, want %d", got, want)
	}
}

func TestStartStopCPUProfile(t *testing.T) {
	// Write to a temp file to verify no error occurs.
	f, err := os.CreateTemp(t.TempDir(), "cpu-*.prof")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = f.Close() }()

	if err := StartCPUProfile(f); err != nil {
		t.Fatalf("StartCPUProfile: %v", err)
	}
	_ = CPUIntensiveWork(15)
	StopCPUProfile()

	info, err := f.Stat()
	if err != nil {
		t.Fatal(err)
	}
	if info.Size() == 0 {
		t.Error("CPU profile file is empty")
	}
}

func TestStartCPUProfileToDiscard(t *testing.T) {
	if err := StartCPUProfile(io.Discard); err != nil {
		t.Fatalf("StartCPUProfile(io.Discard): %v", err)
	}
	_ = CPUIntensiveWork(10)
	StopCPUProfile()
}

func TestWriteHeapProfile(t *testing.T) {
	// Allocate some memory first.
	MemoryIntensiveWork(10, 4096)

	f, err := os.CreateTemp(t.TempDir(), "heap-*.prof")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = f.Close() }()

	if err := WriteHeapProfile(f); err != nil {
		t.Fatalf("WriteHeapProfile: %v", err)
	}

	info, err := f.Stat()
	if err != nil {
		t.Fatal(err)
	}
	if info.Size() == 0 {
		t.Error("heap profile file is empty")
	}
}

func TestWriteHeapProfileToDiscard(t *testing.T) {
	if err := WriteHeapProfile(io.Discard); err != nil {
		t.Fatalf("WriteHeapProfile(io.Discard): %v", err)
	}
}
