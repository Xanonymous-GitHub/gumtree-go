package comparator

import (
	"github.com/Xanonymous-GitHub/gumtree-go/ast"
	. "github.com/Xanonymous-GitHub/gumtree-go/datastructures"
	"github.com/samber/lo"
	"sort"
)

func GetDiceValueOf(pair Pair[*ast.Node, *ast.Node]) float64 {
	n1Degree := pair.Left().Degree()
	n2Degree := pair.Right().Degree()

	if n1Degree == 0 && n2Degree == 0 {
		// If both nodes are leaves, return 1.0
		// This means they are structure-wise equal.
		return 1.0
	}

	return float64(2.0*n1Degree) / float64(n1Degree+n2Degree)
}

func (c *comparator) handleCandidateMappings() {
	// Sort the candidate mappings by their dice values in descending order.
	sort.SliceStable(c.candidateMappings, func(i, j int) bool {
		return GetDiceValueOf(c.candidateMappings[i]) > GetDiceValueOf(c.candidateMappings[j])
	})

	for len(c.candidateMappings) > 0 {
		// Pop the mapping with the highest dice value.
		mapping := c.candidateMappings[0]
		c.candidateMappings = c.candidateMappings[1:]

		forEachIsomorphicNodesPairOf(mapping, func(pair Pair[*ast.Node, *ast.Node]) {
			c.uniqueMappings = append(c.uniqueMappings, pair)
		})

		c.candidateMappings = lo.Filter(c.candidateMappings, func(pair Pair[*ast.Node, *ast.Node], idx int) bool {
			return !pair.Left().IsEqualTo(mapping.Left()) && !pair.Right().IsEqualTo(mapping.Right())
		})
	}
}
