package datastructures

import (
	"fmt"
)

type Pair[A, B interface{}] interface {
	Left() A
	Right() B
	String() string
}

type pair[A, B interface{}] struct {
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

func NewPair[A, B interface{}](first A, second B) Pair[A, B] {
	return &pair[A, B]{first, second}
}
