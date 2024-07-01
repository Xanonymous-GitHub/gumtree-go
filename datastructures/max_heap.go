package datastructures

import "log/slog"

type MaxHeap[E maxHeapElementType] interface {
	BinaryHeap[E]
}

type maxHeapElementType interface {
	binaryHeapElementType
}

type maxHeap[E maxHeapElementType] struct {
	heap BinaryHeap[E]
}

func (m *maxHeap[E]) Push(e E) {
	m.heap.Push(e)
}

func (m *maxHeap[E]) Pop() {
	m.heap.Pop()
}

func (m *maxHeap[E]) Top() E {
	return m.heap.Top()
}

func (m *maxHeap[E]) Size() int {
	return m.heap.Size()
}

func (m *maxHeap[E]) IsEmpty() bool {
	return m.heap.IsEmpty()
}

func (m *maxHeap[E]) swap(i, j int) {
	m.heap.swap(i, j)
}

func (m *maxHeap[E]) up(childIdx int) {
	m.heap.up(childIdx)
}

func (m *maxHeap[E]) down(parentIdx int) {
	m.heap.down(parentIdx)
}

func NewMaxHeap[E maxHeapElementType](logger slog.Logger) MaxHeap[E] {
	return &maxHeap[E]{
		heap: NewBinaryHeap[E](Greater, logger),
	}
}
