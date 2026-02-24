// Package writebarrier demonstrates pointer writes and heap references in Go,
// illustrating the conditions under which the GC's write barrier is triggered.
//
// What is a write barrier?
// A write barrier is a small piece of code injected by the compiler at every
// pointer write to a heap location. During concurrent GC marking, the write
// barrier notifies the collector when a pointer is stored, ensuring that the
// tri-color invariant is maintained: no black object points to a white object
// without the collector knowing about it.
//
// Why does it exist?
// Without a write barrier, the concurrent collector could miss newly created
// references from already-scanned (black) objects to unscanned (white) objects.
// This would cause live objects to be incorrectly collected. The write barrier
// prevents this by recording pointer writes that occur during the marking phase.
//
// Performance impact:
// The write barrier adds a small overhead (typically a few nanoseconds) to
// every pointer write on the heap while the GC is active. Pointer writes to
// stack-local variables do not need a write barrier because the collector
// scans stacks atomically during STW phases. This is one reason why stack
// allocations are faster than heap allocations.
package writebarrier

// Node is a linked-list node used to demonstrate heap pointer writes.
type Node struct {
	Value int
	Next  *Node
}

// PointerWrite creates a Node and writes a pointer to its Next field.
// When the Node escapes to the heap, this pointer write triggers the
// write barrier during GC's concurrent marking phase.
func PointerWrite(value, nextValue int) *Node {
	next := &Node{Value: nextValue}
	head := &Node{Value: value, Next: next}
	return head
}

// HeapReference creates a chain of n nodes that reference each other on the
// heap. Each node's Next pointer points to the previously created node,
// forming a linked list. Returns the head of the list.
func HeapReference(n int) *Node {
	if n <= 0 {
		return nil
	}
	head := &Node{Value: 0}
	current := head
	for i := 1; i < n; i++ {
		current.Next = &Node{Value: i}
		current = current.Next
	}
	return head
}

// ChainLength traverses a linked list from head and returns the number of
// nodes in the chain.
func ChainLength(head *Node) int {
	count := 0
	for head != nil {
		count++
		head = head.Next
	}
	return count
}

// StackVsHeap demonstrates the difference between pointer writes that stay
// on the stack and those that escape to the heap.
//
// stackOnly creates a Node on the stack (no escape) and returns its value.
// The compiler can determine that the node does not escape, so no write
// barrier is needed for pointer operations on it.
//
// heapEscape returns a pointer to a Node, forcing it to escape to the heap.
// Pointer writes to this node's fields will trigger the write barrier during
// GC. Use `go build -gcflags="-m"` to observe escape analysis decisions.
func StackVsHeap() (stackValue int, heapNode *Node) {
	// This node stays on the stack because it does not escape.
	local := Node{Value: 42}
	stackValue = local.Value

	// This node escapes to the heap because we return a pointer to it.
	heapNode = &Node{Value: 100}
	return stackValue, heapNode
}
