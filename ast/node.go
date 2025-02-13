package ast

import (
	"fmt"
	"github.com/cespare/xxhash/v2"
	"github.com/google/uuid"
	"golang.org/x/exp/maps"
	"sort"
	"strconv"
	"sync"
)

var lock = sync.RWMutex{}

type Node struct {
	Label       NodeLabelType
	Value       NodeValueType
	Parent      *Node
	Children    map[int]*Node
	Id          NodeIdType
	idxToParent int
}

type NodeHashMemo map[NodeIdType]uint64

type NodeParentInfo struct {
	// The Parent of a Node.
	Parent *Node

	// The index of the Node in the Parent's children map.
	// If the Parent is nil, this field is ignored.
	IdxToParent int
}

func NewNode(parentInfo NodeParentInfo, label NodeLabelType, value NodeValueType) (*Node, error) {
	newUUId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	newIdStr := NodeIdType(newUUId.String())
	if newIdStr == "" {
		return nil, fmt.Errorf("newIdStr is empty during node creation")
	}

	newNode := Node{
		Label:       label,
		Value:       value,
		Parent:      nil,
		Children:    make(map[int]*Node),
		Id:          newIdStr,
		idxToParent: -1,
	}

	err = newNode.UpdateParent(parentInfo)
	if err != nil {
		return nil, err
	}

	return &newNode, nil
}

func NewOrphanNode() (*Node, error) {
	return NewNode(NodeParentInfo{Parent: nil, IdxToParent: -1}, "", "")
}

func (n *Node) UpdateParent(newParentInfo NodeParentInfo) error {
	if newParentInfo.Parent != nil {
		if newParentInfo.IdxToParent < 0 {
			return fmt.Errorf("IdxToParent should not be negative when Parent is not nil")
		}
		if _, ok := newParentInfo.Parent.Children[newParentInfo.IdxToParent]; ok {
			return fmt.Errorf("new Parent already has a child at index %d", newParentInfo.IdxToParent)
		}

		newParentInfo.Parent.Children[newParentInfo.IdxToParent] = n
	} else if newParentInfo.IdxToParent >= 0 {
		return fmt.Errorf("IdxToParent should be negative when Parent is nil")
	}

	if n.Parent != nil {
		delete(n.Parent.Children, n.idxToParent)
	}

	n.Parent = newParentInfo.Parent
	n.idxToParent = newParentInfo.IdxToParent

	return nil
}

func (n *Node) DestroySubtree() {
	if n == nil {
		panic("destroying node is nil")
	}

	for _, child := range n.Children {
		child.DestroySubtree()
		delete(n.Children, child.idxToParent)
	}
}

func (n *Node) Height() int {
	if len(n.Children) == 0 {
		return 0
	}

	maxHeight := 0
	for _, child := range n.Children {
		height := child.Height()
		if height > maxHeight {
			maxHeight = height
		}
	}

	return 1 + maxHeight
}

func (n *Node) Degree() int {
	return len(n.Children)
}

func (n *Node) ValueOfOrder() int {
	return n.Height()
}

func (n *Node) OrderedChildren() []*Node {
	children := maps.Values(n.Children)
	sort.Slice(children, func(i, j int) bool {
		return children[i].idxToParent < children[j].idxToParent
	})
	return children
}

// Isomorphic returns true if the Node is isomorphic to the other Node.
func (n *Node) Isomorphic(other *Node) bool {
	if n == nil || other == nil {
		return false
	}

	return n.HashValue(nil) == other.HashValue(nil)
}

func (n *Node) HashValue(memo *NodeHashMemo) uint64 {
	propertyStr := fmt.Sprintf("<%s>[%s]<", n.Label, n.Value)
	propertyHash := xxhash.Sum64String(propertyStr)

	if len(n.Children) == 0 {
		if memo != nil {
			lock.Lock()
			defer lock.Unlock()
			(*memo)[n.Id] = propertyHash
		}
		return propertyHash
	}

	children := n.OrderedChildren()

	var combinedChildrenHash xxhash.Digest
	_, _ = combinedChildrenHash.WriteString(strconv.FormatUint(propertyHash, 10))

	for _, child := range children {
		childHash := child.HashValue(memo)
		_, _ = combinedChildrenHash.WriteString(strconv.FormatUint(childHash, 10))
	}

	result := combinedChildrenHash.Sum64()
	if memo != nil {
		lock.Lock()
		defer lock.Unlock()
		(*memo)[n.Id] = result
	}
	return result
}

func (n *Node) IsEqualTo(other interface{}) bool {
	if other == nil {
		return false
	}

	if otherNode, ok := other.(*Node); ok {
		return n.Label == otherNode.Label && n.Value == otherNode.Value && n.Id == otherNode.Id
	}

	return false
}
