package comparator

import (
	"github.com/Xanonymous-GitHub/gumtree-go/ast"
	. "github.com/Xanonymous-GitHub/gumtree-go/datastructures"
	. "github.com/deckarep/golang-set/v2"
	"github.com/samber/lo"
	"golang.org/x/exp/maps"
)

type mappingsType []Pair[*ast.Node, *ast.Node]
type isomorphicNodesType Pair[Set[*ast.Node], Set[*ast.Node]]

type IsomorphicMappings interface {
	UniqueIsomorphicMappings() []isomorphicNodesType
	NonUniqueIsomorphicMappings() []isomorphicNodesType
	NonIsomorphicMappings() []isomorphicNodesType
}

type isomorphicMappings struct {
	hashToNodePairs map[uint64]isomorphicNodesType
}

func (i *isomorphicMappings) UniqueIsomorphicMappings() []isomorphicNodesType {
	return lo.Filter(maps.Values(i.hashToNodePairs), func(sameHashPair isomorphicNodesType, idx int) bool {
		return sameHashPair.Left().Cardinality() == 1 && sameHashPair.Right().Cardinality() == 1
	})
}

func (i *isomorphicMappings) NonUniqueIsomorphicMappings() []isomorphicNodesType {
	return lo.Filter(maps.Values(i.hashToNodePairs), func(sameHashPair isomorphicNodesType, idx int) bool {
		return (sameHashPair.Left().Cardinality() > 1 && sameHashPair.Right().Cardinality() >= 1) ||
			(sameHashPair.Left().Cardinality() >= 1 && sameHashPair.Right().Cardinality() > 1)
	})
}

func (i *isomorphicMappings) NonIsomorphicMappings() []isomorphicNodesType {
	return lo.Filter(maps.Values(i.hashToNodePairs), func(sameHashPair isomorphicNodesType, idx int) bool {
		return sameHashPair.Left().Cardinality() == 0 || sameHashPair.Right().Cardinality() == 0
	})
}

func NewIsomorphicMappings(memo1, memo2 ast.NodeHashMemo, mappings mappingsType) IsomorphicMappings {
	hashToNodePairs := make(map[uint64]isomorphicNodesType)

	for _, mapping := range mappings {
		if mapping.Left() == nil || mapping.Right() == nil {
			panic("mapping should not contain nil nodes")
		}

		hashOfLeft, ok1 := memo1[mapping.Left().Id]
		hashOfRight, ok2 := memo2[mapping.Right().Id]
		if !ok1 || !ok2 {
			panic("mapping should contain nodes in the memo")
		}

		if _, ok := hashToNodePairs[hashOfLeft]; !ok {
			hashToNodePairs[hashOfLeft] = NewPair(NewSet[*ast.Node](), NewSet[*ast.Node]())
		}
		if _, ok := hashToNodePairs[hashOfRight]; !ok {
			hashToNodePairs[hashOfRight] = NewPair(NewSet[*ast.Node](), NewSet[*ast.Node]())
		}

		hashToNodePairs[hashOfLeft].Left().Add(mapping.Left())
		hashToNodePairs[hashOfLeft].Right().Add(mapping.Right())
	}

	return &isomorphicMappings{
		hashToNodePairs: hashToNodePairs,
	}
}

func forEachIsomorphicNodesPairOf(p Pair[*ast.Node, *ast.Node], f func(Pair[*ast.Node, *ast.Node])) {
	f(p)
	childrenPairs := PairOf(p.Left().OrderedChildren(), p.Right().OrderedChildren())
	for _, childrenPair := range childrenPairs {
		forEachIsomorphicNodesPairOf(childrenPair, f)
	}
}
