package ast

import (
	"fmt"
	"log/slog"
)

// AST is the interface that wraps the basic operations of an Abstract Syntax Tree.
// It is a labeled & ordered rooted tree where nodes may have a string value.
// Labels of nodes correspond to the name of their production rule in the grammar, i.e., they encode the structure.
// Values of the nodes correspond to the actual tokens in the code.
// Please refer to https://hal.science/hal-01054552.
type AST interface {
	// Add adds a new node in the AST.
	// If `parent` is not nil and `i` is specified,
	// then the new node will be the `i`th child of `parent`.
	// Otherwise, the new node is the new root node and has the previous root node as its only child.
	// Finally, `label` is the label of the new node and `value` is the value of the new node.
	// If `parent` is nil, then `i` should be less than zero.
	Add(parent *Node, i int, label NodeLabelType, value NodeValueType) (*Node, error)

	// Move moves a node `n` and make it the ith child of `newParent`.
	// Note that all children of `n` are moved as well,
	// therefore this actions moves a whole subtree.
	// If the `newParent` is nil, then `n` becomes the new root node.
	// In this case, `i` should be less than zero.
	Move(n, newParent *Node, i int) error

	// Delete deletes a node `n` from the AST.
	Delete(n *Node) error

	// Root returns the root node of the AST.
	// If the AST is empty, then it returns nil.
	Root() *Node

	// UpdateValue updates the label of a node `n` to `newValue`.
	UpdateValue(n *Node, newValue NodeValueType) error

	// UpdateLabel updates the value of a node `n` to `newLabel`.
	UpdateLabel(n *Node, newLabel NodeLabelType) error

	// MakeHashMemo creates a new hash memo for the entire AST.
	MakeHashMemo() *NodeHashMemo
}

type astConcrete struct {
	// nodes is a map of node IDs to nodes.
	nodes map[NodeIdType]*Node

	// root is the root node of the AST.
	root *Node

	// logger is the logger of the AST.
	logger slog.Logger
}

func (a *astConcrete) Add(parent *Node, i int, label NodeLabelType, value NodeValueType) (*Node, error) {
	newNode, err := NewNode(NodeParentInfo{Parent: parent, IdxToParent: i}, label, value)
	if err != nil {
		a.logger.Error("error creating new node")
		return nil, err
	}
	if newNode == nil {
		msg := "newNode is nil"
		a.logger.Error(msg)
		return nil, fmt.Errorf(msg)
	}
	if newNode.Parent == nil {
		if a.root != nil {
			msg := "the root node already exists in the AST"
			a.logger.Error(msg)
			return nil, fmt.Errorf(msg)
		}
		a.root = newNode
	}

	a.nodes[newNode.id] = newNode
	return a.nodes[newNode.id], nil
}

func (a *astConcrete) Move(n, newParent *Node, i int) error {
	return n.UpdateParent(NodeParentInfo{Parent: newParent, IdxToParent: i})
}

func (a *astConcrete) Delete(n *Node) error {
	if n == nil {
		msg := "node is nil"
		a.logger.Error(msg)
		return fmt.Errorf(msg)
	}

	n.DestroySubtree()
	delete(a.nodes, n.id)
	return nil
}

func (a *astConcrete) Root() *Node {
	return a.root
}

func (a *astConcrete) UpdateValue(n *Node, newValue NodeValueType) error {
	if n == nil {
		msg := "node is nil"
		a.logger.Error(msg)
		return fmt.Errorf(msg)
	}

	n.Value = newValue
	return nil
}

func (a *astConcrete) UpdateLabel(n *Node, newLabel NodeLabelType) error {
	if n == nil {
		msg := "node is nil"
		a.logger.Error(msg)
		return fmt.Errorf(msg)
	}

	n.Label = newLabel
	return nil
}

func (a *astConcrete) MakeHashMemo() *NodeHashMemo {
	if a.root == nil {
		return nil
	}

	memo := make(NodeHashMemo)
	_ = a.root.HashValue(&memo)
	return &memo
}

// NewAST creates a new AST.
func NewAST(logger slog.Logger) AST {
	return &astConcrete{
		nodes:  make(map[NodeIdType]*Node),
		root:   nil,
		logger: logger,
	}
}
