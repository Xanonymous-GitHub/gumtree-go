package comparator_test

import (
	"github.com/Xanonymous-GitHub/gumtree-go/ast"
	"github.com/Xanonymous-GitHub/gumtree-go/comparator"
	"log/slog"
	"math/rand/v2"
	"testing"
)

func TestHeightIndexedPriorityList_Push(t *testing.T) {
	t.Parallel()

	t.Run("Test pushing only root node", func(t *testing.T) {
		list := comparator.NewHeightIndexedPriorityList(slog.Logger{})
		node, _ := ast.NewOrphanNode()
		list.Push(node)
	})

	t.Run("Test pushing root node and a child node", func(t *testing.T) {
		list := comparator.NewHeightIndexedPriorityList(slog.Logger{})
		root, _ := ast.NewOrphanNode()
		child, _ := ast.NewNode(ast.NodeParentInfo{Parent: root, IdxToParent: 0}, "child", "child")
		list.Push(root)
		list.Push(child)
	})
}

func TestHeightIndexedPriorityList_Pop(t *testing.T) {
	t.Parallel()

	t.Run("Test popping from an empty list", func(t *testing.T) {
		list := comparator.NewHeightIndexedPriorityList(slog.Logger{})
		nodes := list.Pop()
		if nodes.Cardinality() != 0 {
			t.Errorf("Expected empty set, got %d", nodes.Cardinality())
		}
	})

	t.Run("Test popping from a list with only root node", func(t *testing.T) {
		list := comparator.NewHeightIndexedPriorityList(slog.Logger{})
		root, _ := ast.NewOrphanNode()
		list.Push(root)
		nodes := list.Pop()
		if nodes.Cardinality() != 1 {
			t.Errorf("Expected set with 1 element, got %d", nodes.Cardinality())
		}
		if !nodes.Contains(root) {
			t.Errorf("Expected set to contain root node")
		}
	})

	t.Run("Test popping from a list with root node and a child node", func(t *testing.T) {
		list := comparator.NewHeightIndexedPriorityList(slog.Logger{})
		root, _ := ast.NewOrphanNode()
		child, _ := ast.NewNode(ast.NodeParentInfo{Parent: root, IdxToParent: 0}, "child", "child")
		list.Push(root)
		list.Push(child)
		nodes := list.Pop()
		if nodes.Cardinality() != 1 {
			t.Errorf("Expected set with 1 element, got %d", nodes.Cardinality())
		}
		if !nodes.Contains(root) {
			t.Errorf("Expected set to contain root node")
		}
		if nodes.Contains(child) {
			t.Errorf("Expected set to not contain child node since it has smaller height than root node")
		}
	})

	t.Run("Test popping from a list with many nodes which has same height", func(t *testing.T) {
		list := comparator.NewHeightIndexedPriorityList(slog.Logger{})

		givenPushNum := 3
		for i := 0; i < givenPushNum; i++ {
			root, _ := ast.NewOrphanNode()
			child, _ := ast.NewNode(ast.NodeParentInfo{Parent: root, IdxToParent: 0}, "child", "child")
			list.Push(root)
			list.Push(child)
		}

		nodes := list.Pop()
		if nodes.Cardinality() != givenPushNum {
			t.Errorf("Expected set with %d elements, got %d", givenPushNum, nodes.Cardinality())
		}

		if anyOfNode, ok := nodes.Pop(); !ok {
			t.Errorf("Expected set to contain nodes")
		} else {
			if anyOfNode.Height() != 1 {
				t.Errorf("Expected set to contain node with height 1, got %d", anyOfNode.Height())
			}
		}
	})

	t.Run("Test popping from a list with many nodes which has different height", func(t *testing.T) {
		list := comparator.NewHeightIndexedPriorityList(slog.Logger{})

		givenHeight3Num := 4
		givenHeight2Num := 3
		givenHeight1Num := 2

		height3Nodes := make([]*ast.Node, givenHeight3Num)
		height2Nodes := make([]*ast.Node, givenHeight2Num)
		height1Nodes := make([]*ast.Node, givenHeight1Num)

		for i := 0; i < givenHeight3Num; i++ {
			node, _ := ast.NewOrphanNode()
			child1, _ := ast.NewNode(ast.NodeParentInfo{Parent: node, IdxToParent: 0}, "child1", "child1")
			child2, _ := ast.NewNode(ast.NodeParentInfo{Parent: child1, IdxToParent: 1}, "child2", "child2")
			_, _ = ast.NewNode(ast.NodeParentInfo{Parent: child2, IdxToParent: 2}, "child3", "child3")
			height3Nodes[i] = node
		}

		for i := 0; i < givenHeight2Num; i++ {
			node, _ := ast.NewOrphanNode()
			child1, _ := ast.NewNode(ast.NodeParentInfo{Parent: node, IdxToParent: 0}, "child1", "child1")
			_, _ = ast.NewNode(ast.NodeParentInfo{Parent: child1, IdxToParent: 1}, "child2", "child2")
			height2Nodes[i] = node
		}

		for i := 0; i < givenHeight1Num; i++ {
			node, _ := ast.NewOrphanNode()
			_, _ = ast.NewNode(ast.NodeParentInfo{Parent: node, IdxToParent: 0}, "child1", "child1")
			height1Nodes[i] = node
		}

		allPreparedNodes := append(append(height3Nodes, height2Nodes...), height1Nodes...)
		rand.Shuffle(len(allPreparedNodes), func(i, j int) {
			allPreparedNodes[i], allPreparedNodes[j] = allPreparedNodes[j], allPreparedNodes[i]
		})

		for _, node := range allPreparedNodes {
			list.Push(node)
		}

		nodes := list.Pop()
		if nodes.Cardinality() != givenHeight3Num {
			t.Errorf("Expected set with %d elements, got %d", givenHeight3Num, nodes.Cardinality())
		}
		if anyOfNode, ok := nodes.Pop(); !ok {
			t.Errorf("Expected set to contain nodes")
		} else {
			if anyOfNode.Height() != 3 {
				t.Errorf("Expected set to contain node with height 3, got %d", anyOfNode.Height())
			}
		}

		nodes = list.Pop()
		if nodes.Cardinality() != givenHeight2Num {
			t.Errorf("Expected set with %d elements, got %d", givenHeight2Num, nodes.Cardinality())
		}
		if anyOfNode, ok := nodes.Pop(); !ok {
			t.Errorf("Expected set to contain nodes")
		} else {
			if anyOfNode.Height() != 2 {
				t.Errorf("Expected set to contain node with height 2, got %d", anyOfNode.Height())
			}
		}

		nodes = list.Pop()
		if nodes.Cardinality() != givenHeight1Num {
			t.Errorf("Expected set with %d elements, got %d", givenHeight1Num, nodes.Cardinality())
		}
		if anyOfNode, ok := nodes.Pop(); !ok {
			t.Errorf("Expected set to contain nodes")
		} else {
			if anyOfNode.Height() != 1 {
				t.Errorf("Expected set to contain node with height 1, got %d", anyOfNode.Height())
			}
		}
	})
}

func TestHeightIndexedPriorityList_Open(t *testing.T) {
	t.Parallel()

	t.Run("Test opening node with no child", func(t *testing.T) {
		list := comparator.NewHeightIndexedPriorityList(slog.Logger{})
		root, _ := ast.NewOrphanNode()
		list.Open(root)

		nodes := list.Pop()
		if nodes.Cardinality() != 0 {
			t.Errorf("Expected empty set, got %d", nodes.Cardinality())
		}
	})

	t.Run("Test opening node with one child", func(t *testing.T) {
		list := comparator.NewHeightIndexedPriorityList(slog.Logger{})
		root, _ := ast.NewOrphanNode()
		child, _ := ast.NewNode(ast.NodeParentInfo{Parent: root, IdxToParent: 0}, "child", "child")
		list.Open(root)

		nodes := list.Pop()
		if nodes.Cardinality() != 1 {
			t.Errorf("Expected set with 1 element, got %d", nodes.Cardinality())
		}
		if !nodes.Contains(child) {
			t.Errorf("Expected set to contain child node")
		}
		if nodes.Contains(root) {
			t.Errorf("Expected set to not contain root node since it has smaller height than child node")
		}
	})
}

func TestHeightIndexedPriorityList_PeekMax(t *testing.T) {
	t.Parallel()

	t.Run("Test peeking from an empty list", func(t *testing.T) {
		list := comparator.NewHeightIndexedPriorityList(slog.Logger{})
		if height := list.PeekMax(); height != -1 {
			t.Errorf("Expected -1, got %d", height)
		}
	})

	t.Run("Test peeking from a list with only root node", func(t *testing.T) {
		list := comparator.NewHeightIndexedPriorityList(slog.Logger{})
		root, _ := ast.NewOrphanNode()
		list.Push(root)
		if height := list.PeekMax(); height != 0 {
			t.Errorf("Expected 0, got %d", height)
		}
	})

	t.Run("Test peeking from a list with root node and a child node", func(t *testing.T) {
		list := comparator.NewHeightIndexedPriorityList(slog.Logger{})
		root, _ := ast.NewOrphanNode()
		child, _ := ast.NewNode(ast.NodeParentInfo{Parent: root, IdxToParent: 0}, "child", "child")
		list.Push(root)
		list.Push(child)
		if height := list.PeekMax(); height != 1 {
			t.Errorf("Expected 1, got %d", height)
		}
	})
}
