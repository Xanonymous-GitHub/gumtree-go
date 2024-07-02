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

func TestUpdateParent(t *testing.T) {
	t.Run("TestUpdateParent_1_to_2", func(t *testing.T) {
		t.Parallel()

		givenOldIdxToParent := 1
		givenNewIdxToParent := 2

		givenParentNode1, _ := ast.NewNode(ast.NodeParentInfo{Parent: nil, IdxToParent: -1}, "parent1-label", "parent1-value")
		givenParentNode2, _ := ast.NewNode(ast.NodeParentInfo{Parent: nil, IdxToParent: -1}, "parent2-label", "parent2-value")
		givenTargetNode, _ := ast.NewNode(ast.NodeParentInfo{Parent: givenParentNode1, IdxToParent: givenOldIdxToParent}, "target-label", "target-value")
		givenNewParentInfo := ast.NodeParentInfo{Parent: givenParentNode2, IdxToParent: givenNewIdxToParent}

		err := ast.UpdateParent(givenTargetNode, givenNewParentInfo)
		if err != nil {
			t.Errorf("error updating parent: %v", err)
		}

		if _, ok := givenParentNode1.Children[givenOldIdxToParent]; ok {
			t.Errorf("givenParentNode1.Children[%d] found, should be deleted after UpdateParent", givenOldIdxToParent)
		}
		if _, ok := givenParentNode2.Children[givenNewIdxToParent]; !ok {
			t.Errorf("givenParentNode2.Children[%d] not found, should be saved after UpdateParent", givenNewIdxToParent)
		}

		if givenTargetNode.Parent != givenParentNode2 {
			t.Errorf("givenTargetNode.Parent = %v, want %v", givenTargetNode.Parent, givenParentNode2)
		}
	})

	t.Run("TestUpdateParent_1_to_root", func(t *testing.T) {
		t.Parallel()

		givenOldIdxToParent := 1
		givenNewIdxToParent := -1

		givenParentNode1, _ := ast.NewNode(ast.NodeParentInfo{Parent: nil, IdxToParent: -1}, "parent1-label", "parent1-value")
		givenTargetNode, _ := ast.NewNode(ast.NodeParentInfo{Parent: givenParentNode1, IdxToParent: givenOldIdxToParent}, "target-label", "target-value")
		givenNewParentInfo := ast.NodeParentInfo{Parent: nil, IdxToParent: givenNewIdxToParent}

		err := ast.UpdateParent(givenTargetNode, givenNewParentInfo)
		if err != nil {
			t.Errorf("error updating parent: %v", err)
		}

		if _, ok := givenParentNode1.Children[givenOldIdxToParent]; ok {
			t.Errorf("givenParentNode1.Children[%d] found, should be deleted after UpdateParent", givenOldIdxToParent)
		}

		if givenTargetNode.Parent != nil {
			t.Errorf("givenTargetNode.Parent = %v, want nil", givenTargetNode.Parent)
		}
	})

	t.Run("TestUpdateParent_root_to_1", func(t *testing.T) {
		t.Parallel()

		givenOldIdxToParent := -1
		givenNewIdxToParent := 1

		givenParentNode1, _ := ast.NewNode(ast.NodeParentInfo{Parent: nil, IdxToParent: -1}, "parent1-label", "parent1-value")
		givenTargetNode, _ := ast.NewNode(ast.NodeParentInfo{Parent: nil, IdxToParent: givenOldIdxToParent}, "target-label", "target-value")
		givenNewParentInfo := ast.NodeParentInfo{Parent: givenParentNode1, IdxToParent: givenNewIdxToParent}

		err := ast.UpdateParent(givenTargetNode, givenNewParentInfo)
		if err != nil {
			t.Errorf("error updating parent: %v", err)
		}

		if _, ok := givenParentNode1.Children[givenNewIdxToParent]; !ok {
			t.Errorf("givenParentNode1.Children[%d] not found, should be saved after UpdateParent", givenNewIdxToParent)
		}

		if givenTargetNode.Parent != givenParentNode1 {
			t.Errorf("givenTargetNode.Parent = %v, want %v", givenTargetNode.Parent, givenParentNode1)
		}
	})

	t.Run("TestUpdateParent_root_to_root", func(t *testing.T) {
		t.Parallel()

		givenOldIdxToParent := -1
		givenNewIdxToParent := -2

		givenTargetNode, _ := ast.NewNode(ast.NodeParentInfo{Parent: nil, IdxToParent: givenOldIdxToParent}, "target-label", "target-value")
		givenNewParentInfo := ast.NodeParentInfo{Parent: nil, IdxToParent: givenNewIdxToParent}

		err := ast.UpdateParent(givenTargetNode, givenNewParentInfo)
		if err != nil {
			t.Errorf("error updating parent: %v", err)
		}

		if givenTargetNode.Parent != nil {
			t.Errorf("givenTargetNode.Parent = %v, want nil", givenTargetNode.Parent)
		}
	})

	t.Run("TestUpdateParent_idx_occupied", func(t *testing.T) {
		t.Parallel()

		givenOldIdxToParent := 1
		givenOccupiedIdxOfParent2 := 23

		givenParentNode1, _ := ast.NewNode(ast.NodeParentInfo{Parent: nil, IdxToParent: -1}, "parent1-label", "parent1-value")
		givenParentNode2, _ := ast.NewNode(ast.NodeParentInfo{Parent: nil, IdxToParent: -1}, "parent2-label", "parent2-value")
		givenOccupationNode, _ := ast.NewNode(ast.NodeParentInfo{Parent: givenParentNode2, IdxToParent: givenOccupiedIdxOfParent2}, "node3-label", "node3-value")

		givenTargetNode, _ := ast.NewNode(ast.NodeParentInfo{Parent: givenParentNode1, IdxToParent: givenOldIdxToParent}, "target-label", "target-value")
		givenNewParentInfo := ast.NodeParentInfo{Parent: givenParentNode2, IdxToParent: givenOccupiedIdxOfParent2}

		err := ast.UpdateParent(givenTargetNode, givenNewParentInfo)
		if err == nil {
			t.Errorf("Expect error happened")
		}

		if _, ok := givenParentNode1.Children[givenOldIdxToParent]; !ok {
			t.Errorf("givenParentNode1.Children[%d] not found, should not changed after UpdateParent", givenOldIdxToParent)
		}
		if givenOccupationNode != givenParentNode2.Children[givenOccupiedIdxOfParent2] {
			t.Errorf("givenParentNode2.Children[%d] not found, should not changed after UpdateParent", givenOccupiedIdxOfParent2)
		}
		if givenTargetNode.Parent != givenParentNode1 {
			t.Errorf("givenTargetNode.Parent = %v, want %v", givenTargetNode.Parent, givenParentNode1)
		}
	})

	t.Run("TestUpdateParent_to_root_but_idx_positive", func(t *testing.T) {
		t.Parallel()

		givenOldIdxToParent := 1
		givenNewIdxToParent := 999

		givenParentNode1, _ := ast.NewNode(ast.NodeParentInfo{Parent: nil, IdxToParent: -1}, "parent1-label", "parent1-value")
		givenTargetNode, _ := ast.NewNode(ast.NodeParentInfo{Parent: givenParentNode1, IdxToParent: givenOldIdxToParent}, "target-label", "target-value")
		givenNewParentInfo := ast.NodeParentInfo{Parent: nil, IdxToParent: givenNewIdxToParent}

		err := ast.UpdateParent(givenTargetNode, givenNewParentInfo)
		if err == nil {
			t.Errorf("Expect error happened")
		}

		if _, ok := givenParentNode1.Children[givenOldIdxToParent]; !ok {
			t.Errorf("givenParentNode1.Children[%d] not found, should not changed after UpdateParent", givenOldIdxToParent)
		}

		if givenTargetNode.Parent != givenParentNode1 {
			t.Errorf("givenTargetNode.Parent = %v, want %v", givenTargetNode.Parent, givenParentNode1)
		}
	})

	t.Run("TestUpdateParent_to_non_root_but_idx_negative", func(t *testing.T) {
		t.Parallel()

		givenOldIdxToParent := 0
		givenNewIdxToParent := -999

		givenParentNode1, _ := ast.NewNode(ast.NodeParentInfo{Parent: nil, IdxToParent: -1}, "parent1-label", "parent1-value")
		givenTargetNode, _ := ast.NewNode(ast.NodeParentInfo{Parent: nil, IdxToParent: givenOldIdxToParent}, "target-label", "target-value")
		givenNewParentInfo := ast.NodeParentInfo{Parent: givenParentNode1, IdxToParent: givenNewIdxToParent}

		err := ast.UpdateParent(givenTargetNode, givenNewParentInfo)
		if err == nil {
			t.Errorf("Expect error happened")
		}
		if givenTargetNode != nil {
			t.Errorf("givenTargetNode = %v, want nil", givenTargetNode)
		}
	})
}
