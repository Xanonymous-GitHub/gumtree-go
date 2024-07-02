package datastructures

import (
	"cmp"
	"log/slog"
)

type MinHeap[E minHeapElementType[P], P cmp.Ordered] interface {
	BinaryHeap[E, P]
}

type minHeapElementType[P cmp.Ordered] interface {
	binaryHeapElementType[P]
}

type minHeap[E minHeapElementType[P], P cmp.Ordered] struct {
	heap BinaryHeap[E, P]
}

func (m *minHeap[E, P]) Push(e E) {
	m.heap.Push(e)
}

func (m *minHeap[E, P]) Pop() {
	m.heap.Pop()
}

func (m *minHeap[E, P]) Top() E {
	return m.heap.Top()
}

func (m *minHeap[E, P]) Size() int {
	return m.heap.Size()
}

func (m *minHeap[E, P]) IsEmpty() bool {
	return m.heap.IsEmpty()
}

func NewMinHeap[E minHeapElementType[P], P cmp.Ordered](logger slog.Logger) MinHeap[E, P] {
	return &minHeap[E, P]{
		heap: NewBinaryHeap[E](Less, logger),
	}
}
