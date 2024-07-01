package ast

import (
	"fmt"
	"github.com/google/uuid"
)

type Node struct {
	Label       NodeLabelType
	Value       NodeValueType
	Parent      *Node
	Children    map[int]*Node
	id          string
	idxToParent int
}

type nodeParentInfo struct {
	// The parent of a Node.
	parent *Node

	// The index of the Node in the parent's children map.
	// If the parent is nil, this field is ignored.
	idxToParent int
}

func newNode(parentInfo nodeParentInfo, label NodeLabelType, value NodeValueType) (*Node, error) {
	newId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	newIdStr := newId.String()
	if newIdStr == "" {
		return nil, fmt.Errorf("newIdStr is empty during node creation")
	}

	newNode := Node{
		Label:       label,
		Value:       value,
		Parent:      nil,
		Children:    make(map[int]*Node),
		id:          newIdStr,
		idxToParent: -1,
	}

	err = updateParent(&newNode, parentInfo)
	if err != nil {
		return nil, err
	}

	return &newNode, nil
}

func updateParent(n *Node, newParentInfo nodeParentInfo) error {
	if newParentInfo.parent != nil {
		if newParentInfo.idxToParent < 0 {
			return fmt.Errorf("idxToParent should not be negative when parent is not nil")
		}
		if _, ok := newParentInfo.parent.Children[newParentInfo.idxToParent]; ok {
			return fmt.Errorf("new parent already has a child at index %d", newParentInfo.idxToParent)
		}

		newParentInfo.parent.Children[newParentInfo.idxToParent] = n
	} else if newParentInfo.idxToParent >= 0 {
		return fmt.Errorf("idxToParent should be negative when parent is nil")
	}

	if n.Parent != nil {
		delete(n.Parent.Children, n.idxToParent)
	}

	n.Parent = newParentInfo.parent
	n.idxToParent = newParentInfo.idxToParent

	return nil
}

func destroySubtree(n *Node) {
	if n == nil {
		panic("destroying node is nil")
	}

	for _, child := range n.Children {
		destroySubtree(child)
		delete(n.Children, child.idxToParent)
	}
}
