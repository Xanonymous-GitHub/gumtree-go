package datastructures_test

import (
	"github.com/Xanonymous-GitHub/gumtree-go/datastructures"
	"log/slog"
	"testing"
)

type integerQueueType int

func (t integerQueueType) ValueOfOrder() integerQueueType {
	return t
}

func TestPriorityQueue_Front(t *testing.T) {
	logger := slog.Logger{}
	pq := datastructures.NewPriorityQueue[integerQueueType](logger)

	t.Run("TestPriorityQueue_Front", func(t *testing.T) {
		pq.Push(1)

		if pq.Front() != 1 {
			t.Errorf("Expected 1, got %d", pq.Front())
		}
	})
}

func TestPriorityQueue_Pop(t *testing.T) {
	logger := slog.Logger{}
	pq := datastructures.NewPriorityQueue[integerQueueType](logger)

	t.Run("TestPriorityQueue_Pop", func(t *testing.T) {
		pq.Push(1)
		pq.Pop()

		if pq.Size() != 0 {
			t.Errorf("Expected 0, got %d", pq.Size())
		}
	})
}

func TestPriorityQueue_Push(t *testing.T) {
	logger := slog.Logger{}
	pq := datastructures.NewPriorityQueue[integerQueueType](logger)

	t.Run("TestPriorityQueue_Push", func(t *testing.T) {
		pq.Push(1)

		if pq.Size() != 1 {
			t.Errorf("Expected 1, got %d", pq.Size())
		}
	})
}
