package datastructures

import (
	"cmp"
	"log/slog"
)

type PriorityQueue[E priorityQueueElementType[P], P cmp.Ordered] interface {
	Push(e E)
	Pop()
	Front() E
	Size() int
}

type priorityQueueElementType[P cmp.Ordered] interface {
	maxHeapElementType[P]
}

type priorityQueue[E priorityQueueElementType[P], P cmp.Ordered] struct {
	heap MaxHeap[E, P]
}

func (p *priorityQueue[E, P]) Push(e E) {
	p.heap.Push(e)
}

func (p *priorityQueue[E, P]) Pop() {
	p.heap.Pop()
}

func (p *priorityQueue[E, P]) Front() E {
	return p.heap.Top()
}

func (p *priorityQueue[E, P]) Size() int {
	return p.heap.Size()
}

func NewPriorityQueue[E priorityQueueElementType[P], P cmp.Ordered](logger slog.Logger) PriorityQueue[E, P] {
	return &priorityQueue[E, P]{
		heap: NewMaxHeap[E](logger),
	}
}
