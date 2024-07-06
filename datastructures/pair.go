package datastructures

import (
	"fmt"
	"github.com/Xanonymous-GitHub/gumtree-go/utils"
)

type Pair[A, B comparable] interface {
	Left() A
	Right() B
	String() string

	utils.Equatable
}

type pair[A, B comparable] struct {
	first  A
	second B
}

func (p *pair[A, B]) Left() A {
	return p.first
}

func (p *pair[A, B]) Right() B {
	return p.second
}

func (p *pair[A, B]) String() string {
	return fmt.Sprintf("(%v, %v)", p.first, p.second)
}

func (p *pair[A, B]) IsEqualTo(other interface{}) bool {
	if other == nil {
		return false
	}

	if otherPair, ok := other.(Pair[A, B]); ok {
		return p.Left() == otherPair.Left() && p.Right() == otherPair.Right()
	}

	return false
}

// CrossPairOf returns a slice of pairs of elements from two collections.
// Similar to the Cartesian product of two sets.
func CrossPairOf[T comparable](collection1, collection2 []T) []Pair[T, T] {
	if collection1 == nil || collection2 == nil {
		return nil
	}

	length1 := len(collection1)
	length2 := len(collection2)
	if length1 == 0 || length2 == 0 {
		return nil
	}

	pairs := make([]Pair[T, T], 0, length1*length2)
	for _, item1 := range collection1 {
		for _, item2 := range collection2 {
			pairs = append(pairs, NewPair(item1, item2))
		}
	}

	return pairs
}

func PairOf[T comparable](collection1, collection2 []T) []Pair[T, T] {
	if collection1 == nil || collection2 == nil {
		return nil
	}

	length1 := len(collection1)
	length2 := len(collection2)
	if length1 == 0 || length2 == 0 {
		return nil
	}

	if length1 != length2 {
		return nil
	}

	pairs := make([]Pair[T, T], 0, length1)
	for i := 0; i < length1; i++ {
		pairs = append(pairs, NewPair(collection1[i], collection2[i]))
	}

	return pairs
}

func NewPair[A, B comparable](first A, second B) Pair[A, B] {
	return &pair[A, B]{first, second}
}
