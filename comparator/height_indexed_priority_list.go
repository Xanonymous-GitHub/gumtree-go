package comparator

import (
	"github.com/Xanonymous-GitHub/gumtree-go/ast"
	. "github.com/deckarep/golang-set/v2"
	"golang.org/x/exp/maps"
	"log/slog"
	"sort"
)

type HeightIndexedPriorityList interface {
	// Push inserts the node n in the list.
	Push(n *ast.Node)

	// PeekMax returns the greatest height of the list
	PeekMax() int

	// Pop returns and removes the set of all nodes which has a height equals to PeekMax()
	Pop() Set[*ast.Node]

	// Open inserts all the children of n in the list.
	Open(n *ast.Node)
}

type heightIndexedPriorityList struct {
	nodes  map[int][]*ast.Node
	logger slog.Logger
}

func (h *heightIndexedPriorityList) Push(n *ast.Node) {
	if n == nil {
		msg := "Pushing nil node"
		h.logger.Error(msg)
		panic(msg)
	}

	height := n.Height()
	if _, ok := h.nodes[height]; !ok {
		h.nodes[height] = make([]*ast.Node, 1)
		h.nodes[height][0] = n
	} else {
		h.nodes[height] = append(h.nodes[height], n)
	}
}

func (h *heightIndexedPriorityList) PeekMax() int {
	if len(h.nodes) == 0 {
		return 0
	}

	heights := maps.Keys(h.nodes)
	sort.Ints(heights)
	return heights[len(heights)-1]
}

func (h *heightIndexedPriorityList) Pop() Set[*ast.Node] {
	if len(h.nodes) == 0 {
		return NewSet[*ast.Node]()
	}

	maxHeight := h.PeekMax()
	outputNodes := NewSet(h.nodes[maxHeight]...)
	delete(h.nodes, maxHeight)
	return outputNodes
}

func (h *heightIndexedPriorityList) Open(n *ast.Node) {
	if n == nil {
		msg := "Opening nil node"
		h.logger.Error(msg)
		panic(msg)
	}

	for _, child := range n.Children {
		h.Push(child)
	}
}

func NewHeightIndexedPriorityList(logger slog.Logger) HeightIndexedPriorityList {
	return &heightIndexedPriorityList{
		logger: logger,
		nodes:  make(map[int][]*ast.Node),
	}
}
