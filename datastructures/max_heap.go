package datastructures

import (
	"cmp"
	"log/slog"
)

type MaxHeap[E maxHeapElementType[P], P cmp.Ordered] interface {
	BinaryHeap[E, P]
}

type maxHeapElementType[P cmp.Ordered] interface {
	binaryHeapElementType[P]
}

type maxHeap[E maxHeapElementType[P], P cmp.Ordered] struct {
	heap BinaryHeap[E, P]
}

func (m *maxHeap[E, P]) Push(e E) {
	m.heap.Push(e)
}

func (m *maxHeap[E, P]) Pop() {
	m.heap.Pop()
}

func (m *maxHeap[E, P]) Top() E {
	return m.heap.Top()
}

func (m *maxHeap[E, P]) Size() int {
	return m.heap.Size()
}

func (m *maxHeap[E, P]) IsEmpty() bool {
	return m.heap.IsEmpty()
}

func NewMaxHeap[E maxHeapElementType[P], P cmp.Ordered](logger slog.Logger) MaxHeap[E, P] {
	return &maxHeap[E, P]{
		heap: NewBinaryHeap[E](Greater, logger),
	}
}
