package comparator

import (
	"github.com/Xanonymous-GitHub/gumtree-go/ast"
	. "github.com/Xanonymous-GitHub/gumtree-go/datastructures"
	"github.com/samber/lo"
)

func (c *comparator) topDown() {
	tree1HashMemo := (*c.tree1).MakeHashMemo()
	tree2HashMemo := (*c.tree2).MakeHashMemo()

	c.list1.Push((*c.tree1).Root())
	c.list2.Push((*c.tree2).Root())

	for {
		maxHeightOfL1 := c.list1.PeekMax()
		maxHeightOfL2 := c.list2.PeekMax()

		if lo.Min([]int{maxHeightOfL1, maxHeightOfL2}) <= c.minHeight {
			break
		}

		if maxHeightOfL1 != maxHeightOfL2 {
			if maxHeightOfL1 > maxHeightOfL2 {
				nodes := c.list1.Pop()
				for _, n := range nodes.ToSlice() {
					c.list1.Open(n)
				}
			} else {
				nodes := c.list2.Pop()
				for _, n := range nodes.ToSlice() {
					c.list2.Open(n)
				}
			}
		} else {
			h1 := c.list1.Pop()
			h2 := c.list2.Pop()

			sameHeightNodePairs := CrossPairOf(h1.ToSlice(), h2.ToSlice())
			mappings := NewIsomorphicMappings(tree1HashMemo, tree2HashMemo, sameHeightNodePairs)

			for _, uniqueIsomorphicMapping := range mappings.UniqueIsomorphicMappings() {
				// The length of the left and right sets should be 1, since they are unique.
				uniquePair := PairOf(uniqueIsomorphicMapping.Left().ToSlice(), uniqueIsomorphicMapping.Right().ToSlice())[0]
				forEachIsomorphicNodesPairOf(uniquePair, func(pair Pair[*ast.Node, *ast.Node]) {
					c.uniqueMappings = append(c.uniqueMappings, pair)
				})
			}

			for _, nonUniqueIsomorphicMapping := range mappings.NonUniqueIsomorphicMappings() {
				// Only add mappings that are isomorphic but not unique.
				c.candidateMappings = append(
					c.candidateMappings,
					CrossPairOf(
						nonUniqueIsomorphicMapping.Left().ToSlice(),
						nonUniqueIsomorphicMapping.Right().ToSlice(),
					)...,
				)
			}

			for _, nonIsomorphicMapping := range mappings.NonIsomorphicMappings() {
				for _, node := range nonIsomorphicMapping.Left().ToSlice() {
					c.list1.Open(node)
				}
				for _, node := range nonIsomorphicMapping.Right().ToSlice() {
					c.list2.Open(node)
				}
			}
		}
	}

	c.handleCandidateMappings()
}
