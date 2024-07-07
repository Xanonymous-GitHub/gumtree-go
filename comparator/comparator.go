package comparator

import (
	"github.com/Xanonymous-GitHub/gumtree-go/ast"
	"log/slog"
)

type Comparator interface {
}

type comparator struct {
	tree1, tree2       *ast.AST
	list1, list2       HeightIndexedPriorityList
	candidateMappings  mappingsType
	uniqueMappings     mappingsType
	minDice            float64
	minHeight, maxSize int
	logger             slog.Logger
}

func NewComparator(
	tree1, tree2 *ast.AST,
	minHeight, maxSize int,
	minDice float64,
	logger slog.Logger,
) Comparator {
	if tree1 == nil || tree2 == nil {
		panic("trees cannot be nil")
	}
	if minHeight < 0 {
		panic("minHeight cannot be negative")
	}

	return &comparator{
		tree1:             tree1,
		tree2:             tree2,
		list1:             NewHeightIndexedPriorityList(logger),
		list2:             NewHeightIndexedPriorityList(logger),
		candidateMappings: make(mappingsType, 0),
		uniqueMappings:    make(mappingsType, 0),
		minDice:           minDice,
		minHeight:         minHeight,
		maxSize:           maxSize,
	}
}
