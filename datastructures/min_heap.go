package datastructures

import "log/slog"

type MinHeap[E minHeapElementType] interface {
	BinaryHeap[E]
}

type minHeapElementType interface {
	binaryHeapElementType
}

type minHeap[E minHeapElementType] struct {
	heap BinaryHeap[E]
}

func (m *minHeap[E]) Push(e E) {
	m.heap.Push(e)
}

func (m *minHeap[E]) Pop() {
	m.heap.Pop()
}

func (m *minHeap[E]) Top() E {
	return m.heap.Top()
}

func (m *minHeap[E]) Size() int {
	return m.heap.Size()
}

func (m *minHeap[E]) IsEmpty() bool {
	return m.heap.IsEmpty()
}

func NewMinHeap[E minHeapElementType](logger slog.Logger) MinHeap[E] {
	return &minHeap[E]{
		heap: NewBinaryHeap[E](Less, logger),
	}
}
