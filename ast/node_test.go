package ast_test

import (
	"github.com/Xanonymous-GitHub/gumtree-go/ast"
	"testing"
)

func TestNewNode(t *testing.T) {
	givenRootNode := ast.Node{
		Label:    ast.NodeLabelType("root-label"),
		Value:    ast.NodeValueType("root-value"),
		Parent:   nil,
		Children: make(map[int]*ast.Node),
	}

	t.Run("TestNewNode", func(t *testing.T) {
		givenIdxToParent := 23
		givenLabel := ast.NodeLabelType("label")
		givenValue := ast.NodeValueType("value")

		newNode, err := ast.NewNode(
			ast.NodeParentInfo{Parent: &givenRootNode, IdxToParent: givenIdxToParent},
			givenLabel,
			givenValue,
		)
		if err != nil {
			t.Errorf("error creating new node: %v", err)
		}
		if newNode == nil {
			t.Error("newNode is nil")
		}

		if newNode.Label != givenLabel {
			t.Errorf("newNode.Label = %v, want %v", newNode.Label, givenLabel)
		}
		if newNode.Value != givenValue {
			t.Errorf("newNode.Value = %v, want %v", newNode.Value, givenValue)
		}
		if newNode.Parent != &givenRootNode {
			t.Errorf("newNode.Parent = %v, want %v", newNode.Parent, &givenRootNode)
		}
		if _, ok := newNode.Parent.Children[givenIdxToParent]; !ok {
			t.Errorf("newNode.Parent.Children[%d] not found", givenIdxToParent)
		}
		if newNode.Parent.Children[givenIdxToParent] != newNode {
			t.Errorf("newNode.Parent.Children[%d] = %v, want %v", givenIdxToParent, newNode.Parent.Children[givenIdxToParent], newNode)
		}
	})
}
