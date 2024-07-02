package datastructures

import (
	"cmp"
	"github.com/Xanonymous-GitHub/gumtree-go/utils"
	"log/slog"
)

type BinaryHeap[E binaryHeapElementType[P], P cmp.Ordered] interface {
	// Push pushes the element `e` onto the heap.
	Push(e E)

	// Pop removes the minimum element (according to Less) from the heap.
	Pop()

	// Top returns the minimum element (according to Less) from the heap.
	Top() E

	// Size returns the size of the heap.
	Size() int

	// IsEmpty returns the heap is empty or not.
	IsEmpty() bool
}

type binaryHeapElementType[P cmp.Ordered] interface {
	any
	utils.AllowOrdered[P]
}

type binaryHeap[E binaryHeapElementType[P], P cmp.Ordered] struct {
	logger   slog.Logger
	elements []E
	lessFunc LessFunc[P]
}

type LessFunc[E cmp.Ordered] func(a, b E) bool

func Less[E cmp.Ordered](a, b E) bool {
	return cmp.Less(a, b)
}

func Greater[E cmp.Ordered](a, b E) bool {
	return cmp.Less(b, a)
}

func (b *binaryHeap[E, P]) Push(e E) {
	b.elements = append(b.elements, e)
	b.up(len(b.elements) - 1)
}

func (b *binaryHeap[E, P]) Pop() {
	size := len(b.elements)
	if size <= 1 {
		b.elements = make([]E, 0)
		return
	}

	lastIdx := size - 1
	b.swap(0, lastIdx)
	b.elements = b.elements[:lastIdx]
	b.down(0)
}

func (b *binaryHeap[E, P]) Top() E {
	return b.elements[0]
}

func (b *binaryHeap[E, P]) Size() int {
	return len(b.elements)
}

func (b *binaryHeap[E, P]) IsEmpty() bool {
	return len(b.elements) == 0
}

// swap swaps the elements with indices i and j.
func (b *binaryHeap[E, P]) swap(i, j int) {
	b.elements[i], b.elements[j] = b.elements[j], b.elements[i]
}

// up moves the element at index `childIdx` up to its correct position.
func (b *binaryHeap[E, P]) up(childIdx int) {
	if childIdx <= 0 {
		return
	}

	parentIdx := (childIdx - 1) >> 1
	if !b.lessFunc(b.elements[childIdx].ValueOfOrder(), b.elements[parentIdx].ValueOfOrder()) {
		return
	}

	b.swap(childIdx, parentIdx)
	b.up(parentIdx)
}

// down moves the element at index `parentIdx` down to its correct position.
func (b *binaryHeap[E, P]) down(parentIdx int) {
	leastIdx := parentIdx
	lChildIdx := (parentIdx << 1) + 1
	rChildIdx := lChildIdx + 1

	heapSize := len(b.elements)
	if lChildIdx < heapSize && b.lessFunc(b.elements[lChildIdx].ValueOfOrder(), b.elements[leastIdx].ValueOfOrder()) {
		leastIdx = lChildIdx
	}
	if rChildIdx < heapSize && b.lessFunc(b.elements[rChildIdx].ValueOfOrder(), b.elements[leastIdx].ValueOfOrder()) {
		leastIdx = rChildIdx
	}

	if leastIdx != parentIdx {
		b.swap(parentIdx, leastIdx)
		b.down(leastIdx)
	}
}

func NewBinaryHeap[E binaryHeapElementType[P], P cmp.Ordered](
	lessFunc LessFunc[P],
	logger slog.Logger,
) BinaryHeap[E, P] {
	return &binaryHeap[E, P]{
		logger:   logger,
		elements: make([]E, 0),
		lessFunc: lessFunc,
	}
}
