package datastructures_test

import (
	. "github.com/Xanonymous-GitHub/gumtree-go/datastructures"
	"testing"
)

func TestCrossPairOf(t *testing.T) {
	t.Parallel()

	t.Run("two collection non empty same size", func(t *testing.T) {
		c1 := []int{1, 2, 3}
		c2 := []int{4, 5, 6}

		pairs := CrossPairOf(c1, c2)

		if len(pairs) != 9 {
			t.Errorf("expected 3 pairs, got %d", len(pairs))
		}

		expectResult := []Pair[int, int]{
			NewPair(1, 4),
			NewPair(1, 5),
			NewPair(1, 6),
			NewPair(2, 4),
			NewPair(2, 5),
			NewPair(2, 6),
			NewPair(3, 4),
			NewPair(3, 5),
			NewPair(3, 6),
		}

		for i, pair := range pairs {
			if !pair.IsEqualTo(expectResult[i]) {
				t.Errorf("expected %v, got %v", expectResult[i], pair)
			}
		}
	})
}
