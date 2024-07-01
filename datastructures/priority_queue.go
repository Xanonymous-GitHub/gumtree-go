package datastructures

import "log/slog"

type PriorityQueue[E priorityQueueElementType] interface {
	Push(e E)
	Pop()
	Front() E
	Size() int
}

type priorityQueueElementType interface {
	maxHeapElementType
}

type priorityQueue[E priorityQueueElementType] struct {
	heap MaxHeap[E]
}

func (p priorityQueue[E]) Push(e E) {
	p.heap.Push(e)
}

func (p priorityQueue[E]) Pop() {
	p.heap.Pop()
}

func (p priorityQueue[E]) Front() E {
	return p.heap.Top()
}

func (p priorityQueue[E]) Size() int {
	return p.heap.Size()
}

func NewPriorityQueue[E priorityQueueElementType](logger slog.Logger) PriorityQueue[E] {
	return &priorityQueue[E]{
		heap: NewMaxHeap[E](logger),
	}
}
