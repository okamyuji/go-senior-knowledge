package writebarrier

import "testing"

func TestPointerWrite(t *testing.T) {
	head := PointerWrite(1, 2)
	if head.Value != 1 {
		t.Errorf("head.Value = %d, want 1", head.Value)
	}
	if head.Next == nil {
		t.Fatal("head.Next is nil, want non-nil")
	}
	if head.Next.Value != 2 {
		t.Errorf("head.Next.Value = %d, want 2", head.Next.Value)
	}
}

func TestHeapReference(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{5, 5},
		{10, 10},
	}
	for _, tt := range tests {
		head := HeapReference(tt.n)
		got := ChainLength(head)
		if got != tt.want {
			t.Errorf("ChainLength(HeapReference(%d)) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestHeapReferenceValues(t *testing.T) {
	head := HeapReference(5)
	for i := range 5 {
		if head == nil {
			t.Fatalf("node at position %d is nil", i)
		}
		if head.Value != i {
			t.Errorf("node[%d].Value = %d, want %d", i, head.Value, i)
		}
		head = head.Next
	}
}

func TestStackVsHeap(t *testing.T) {
	stackValue, heapNode := StackVsHeap()
	if stackValue != 42 {
		t.Errorf("stackValue = %d, want 42", stackValue)
	}
	if heapNode == nil {
		t.Fatal("heapNode is nil, want non-nil")
	}
	if heapNode.Value != 100 {
		t.Errorf("heapNode.Value = %d, want 100", heapNode.Value)
	}
}

func TestChainLength(t *testing.T) {
	if got := ChainLength(nil); got != 0 {
		t.Errorf("ChainLength(nil) = %d, want 0", got)
	}
	single := &Node{Value: 1}
	if got := ChainLength(single); got != 1 {
		t.Errorf("ChainLength(single) = %d, want 1", got)
	}
}
