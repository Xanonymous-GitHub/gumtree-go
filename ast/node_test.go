package ast_test

import (
	"github.com/Xanonymous-GitHub/gumtree-go/ast"
	"math/rand/v2"
	"testing"
)

func TestNewOrphanNode(t *testing.T) {
	t.Run("TestNewOrphanNode", func(t *testing.T) {
		newNode, err := ast.NewOrphanNode()
		if err != nil {
			t.Errorf("error creating new orphan node: %v", err)
		}
		if newNode == nil {
			t.Error("newNode is nil")
		}
		if newNode.Parent != nil {
			t.Errorf("newNode.Parent = %v, want nil", newNode.Parent)
		}
	})
}

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

		givenParentNode1, _ := ast.NewOrphanNode()
		givenParentNode2, _ := ast.NewOrphanNode()
		givenTargetNode, _ := ast.NewNode(ast.NodeParentInfo{Parent: givenParentNode1, IdxToParent: givenOldIdxToParent}, "target-label", "target-value")
		givenNewParentInfo := ast.NodeParentInfo{Parent: givenParentNode2, IdxToParent: givenNewIdxToParent}

		err := givenTargetNode.UpdateParent(givenNewParentInfo)
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

		givenParentNode1, _ := ast.NewOrphanNode()
		givenTargetNode, _ := ast.NewNode(ast.NodeParentInfo{Parent: givenParentNode1, IdxToParent: givenOldIdxToParent}, "target-label", "target-value")
		givenNewParentInfo := ast.NodeParentInfo{Parent: nil, IdxToParent: givenNewIdxToParent}

		err := givenTargetNode.UpdateParent(givenNewParentInfo)
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

		givenParentNode1, _ := ast.NewOrphanNode()
		givenTargetNode, _ := ast.NewNode(ast.NodeParentInfo{Parent: nil, IdxToParent: givenOldIdxToParent}, "target-label", "target-value")
		givenNewParentInfo := ast.NodeParentInfo{Parent: givenParentNode1, IdxToParent: givenNewIdxToParent}

		err := givenTargetNode.UpdateParent(givenNewParentInfo)
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

		err := givenTargetNode.UpdateParent(givenNewParentInfo)
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

		givenParentNode1, _ := ast.NewOrphanNode()
		givenParentNode2, _ := ast.NewOrphanNode()
		givenOccupationNode, _ := ast.NewNode(ast.NodeParentInfo{Parent: givenParentNode2, IdxToParent: givenOccupiedIdxOfParent2}, "node3-label", "node3-value")

		givenTargetNode, _ := ast.NewNode(ast.NodeParentInfo{Parent: givenParentNode1, IdxToParent: givenOldIdxToParent}, "target-label", "target-value")
		givenNewParentInfo := ast.NodeParentInfo{Parent: givenParentNode2, IdxToParent: givenOccupiedIdxOfParent2}

		err := givenTargetNode.UpdateParent(givenNewParentInfo)
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

		givenParentNode1, _ := ast.NewOrphanNode()
		givenTargetNode, _ := ast.NewNode(ast.NodeParentInfo{Parent: givenParentNode1, IdxToParent: givenOldIdxToParent}, "target-label", "target-value")
		givenNewParentInfo := ast.NodeParentInfo{Parent: nil, IdxToParent: givenNewIdxToParent}

		err := givenTargetNode.UpdateParent(givenNewParentInfo)
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

		givenParentNode1, _ := ast.NewOrphanNode()
		givenTargetNode, _ := ast.NewNode(ast.NodeParentInfo{Parent: nil, IdxToParent: givenOldIdxToParent}, "target-label", "target-value")
		givenNewParentInfo := ast.NodeParentInfo{Parent: givenParentNode1, IdxToParent: givenNewIdxToParent}

		err := givenTargetNode.UpdateParent(givenNewParentInfo)
		if err == nil {
			t.Errorf("Expect error happened")
		}
		if givenTargetNode != nil {
			t.Errorf("givenTargetNode = %v, want nil", givenTargetNode)
		}
	})
}

func TestDestroySubtree(t *testing.T) {
	t.Run("TestDestroySubtree_no_children", func(t *testing.T) {
		t.Parallel()

		givenRootNode, _ := ast.NewOrphanNode()
		givenRootNode.DestroySubtree()
	})

	t.Run("TestDestroySubtree_with_1_layer_children", func(t *testing.T) {
		t.Parallel()

		givenRootNode, _ := ast.NewOrphanNode()
		givenChildrenNum := 18
		for i := 0; i < givenChildrenNum; i++ {
			_, _ = ast.NewNode(ast.NodeParentInfo{Parent: givenRootNode, IdxToParent: i}, "child-label", "child-value")
		}
		givenRootNode.DestroySubtree()
		if len(givenRootNode.Children) != 0 {
			t.Errorf("len(givenRootNode.Children) = %d, want 0", len(givenRootNode.Children))
		}
	})

	t.Run("TestDestroySubtree_with_2_layers_children", func(t *testing.T) {
		t.Parallel()

		givenRootNode, _ := ast.NewOrphanNode()
		givenChildrenNum := 8

		firstLayerChildren := make([]*ast.Node, givenChildrenNum)
		for i := 0; i < givenChildrenNum; i++ {
			childNode, _ := ast.NewNode(ast.NodeParentInfo{Parent: givenRootNode, IdxToParent: i}, "child-label", "child-value")
			firstLayerChildren[i] = childNode
			for j := 0; j < givenChildrenNum; j++ {
				_, _ = ast.NewNode(ast.NodeParentInfo{Parent: childNode, IdxToParent: j}, "grandchild-label", "grandchild-value")
			}
		}
		givenRootNode.DestroySubtree()

		if len(givenRootNode.Children) != 0 {
			t.Errorf("len(givenRootNode.Children) = %d, want 0", len(givenRootNode.Children))
		}

		for _, firstLayerChild := range firstLayerChildren {
			if len(firstLayerChild.Children) != 0 {
				t.Errorf("len(firstLayerChild.Children) = %d, want 0", len(firstLayerChild.Children))
			}
		}
	})
}

func TestNode_Height(t *testing.T) {
	t.Run("TestNode_Height_no_children", func(t *testing.T) {
		t.Parallel()

		givenRootNode, _ := ast.NewOrphanNode()

		if givenRootNode.Height() != 0 {
			t.Errorf("givenRootNode.Height() = %d, want 0", givenRootNode.Height())
		}
	})

	t.Run("TestNode_Height_with_1_layer_children", func(t *testing.T) {
		t.Parallel()

		givenRootNode, _ := ast.NewOrphanNode()
		givenChildrenNum := 18
		for i := 0; i < givenChildrenNum; i++ {
			_, _ = ast.NewNode(ast.NodeParentInfo{Parent: givenRootNode, IdxToParent: i}, "child-label", "child-value")
		}

		if givenRootNode.Height() != 1 {
			t.Errorf("givenRootNode.Height() = %d, want 1", givenRootNode.Height())
		}
	})

	t.Run("TestNode_Height_with_2_layers_children", func(t *testing.T) {
		t.Parallel()

		givenRootNode, _ := ast.NewOrphanNode()
		givenChildrenNum := 8

		for i := 0; i < givenChildrenNum; i++ {
			childNode, _ := ast.NewNode(ast.NodeParentInfo{Parent: givenRootNode, IdxToParent: i}, "child-label", "child-value")

			shouldHaveChildren := rand.IntN(2) == 0 || i == 0
			if shouldHaveChildren {
				for j := 0; j < givenChildrenNum; j++ {
					_, _ = ast.NewNode(ast.NodeParentInfo{Parent: childNode, IdxToParent: j}, "grandchild-label", "grandchild-value")
				}
			}
		}

		if givenRootNode.Height() != 2 {
			t.Errorf("givenRootNode.Height() = %d, want 2", givenRootNode.Height())
		}
	})
}

func TestNode_Degree(t *testing.T) {
	t.Run("TestNode_Degree_no_children", func(t *testing.T) {
		t.Parallel()

		givenRootNode, _ := ast.NewOrphanNode()

		if givenRootNode.Degree() != 0 {
			t.Errorf("givenRootNode.Degree() = %d, want 0", givenRootNode.Degree())
		}
	})

	t.Run("TestNode_Degree_with_1_layer_children", func(t *testing.T) {
		t.Parallel()

		givenRootNode, _ := ast.NewOrphanNode()
		givenChildrenNum := 18
		for i := 0; i < givenChildrenNum; i++ {
			_, _ = ast.NewNode(ast.NodeParentInfo{Parent: givenRootNode, IdxToParent: i}, "child-label", "child-value")
		}

		if givenRootNode.Degree() != givenChildrenNum {
			t.Errorf("givenRootNode.Degree() = %d, want %d", givenRootNode.Degree(), givenChildrenNum)
		}
	})
}

func TestNode_Isomorphic(t *testing.T) {
	t.Parallel()

	t.Run("two root nodes, no children, different labels and values", func(t *testing.T) {
		root1, _ := ast.NewOrphanNode()
		root2, _ := ast.NewNode(ast.NodeParentInfo{Parent: nil, IdxToParent: -1}, "diff-label", "diff-value")
		if root1.Isomorphic(root2) {
			t.Errorf("expected false, got true")
		}
	})

	t.Run("two root nodes, no children, same labels and values", func(t *testing.T) {
		root1, _ := ast.NewOrphanNode()
		root2, _ := ast.NewOrphanNode()
		if !root1.Isomorphic(root2) {
			t.Errorf("expected true, got false")
		}
	})

	t.Run("same children structure, 1 layer, one of them have different labels and values", func(t *testing.T) {
		root1, _ := ast.NewOrphanNode()
		root2, _ := ast.NewOrphanNode()

		const givenSameLabel ast.NodeLabelType = "same-label"
		const givenSameValue ast.NodeValueType = "same-value"

		for i := 0; i < 5; i++ {
			_, _ = ast.NewNode(ast.NodeParentInfo{Parent: root1, IdxToParent: i}, givenSameLabel, givenSameValue)

			if i == 3 {
				_, _ = ast.NewNode(ast.NodeParentInfo{Parent: root2, IdxToParent: i}, "diff-label", "diff-value")
			} else {
				_, _ = ast.NewNode(ast.NodeParentInfo{Parent: root2, IdxToParent: i}, givenSameLabel, givenSameValue)
			}
		}
		if root1.Isomorphic(root2) {
			t.Errorf("expected false, got true")
		}
	})

	t.Run("same children structure, 1 layer, same labels and values", func(t *testing.T) {
		root1, _ := ast.NewOrphanNode()
		root2, _ := ast.NewOrphanNode()

		const givenSameLabel ast.NodeLabelType = "same-label"
		const givenSameValue ast.NodeValueType = "same-value"

		for i := 0; i < 5; i++ {
			_, _ = ast.NewNode(ast.NodeParentInfo{Parent: root1, IdxToParent: i}, givenSameLabel, givenSameValue)
			_, _ = ast.NewNode(ast.NodeParentInfo{Parent: root2, IdxToParent: i}, givenSameLabel, givenSameValue)
		}
		if !root1.Isomorphic(root2) {
			t.Errorf("expected true, got false")
		}
	})

	t.Run("same children structure, 2 layers, one of them have different labels and values", func(t *testing.T) {
		root1, _ := ast.NewOrphanNode()
		root2, _ := ast.NewOrphanNode()

		const givenSameLabel ast.NodeLabelType = "same-label"
		const givenSameValue ast.NodeValueType = "same-value"

		for i := 0; i < 5; i++ {
			child1, _ := ast.NewNode(ast.NodeParentInfo{Parent: root1, IdxToParent: i}, givenSameLabel, givenSameValue)
			child2, _ := ast.NewNode(ast.NodeParentInfo{Parent: root2, IdxToParent: i}, givenSameLabel, givenSameValue)

			for j := 0; j < 5; j++ {
				_, _ = ast.NewNode(ast.NodeParentInfo{Parent: child1, IdxToParent: j}, givenSameLabel, givenSameValue)

				if j == 3 {
					_, _ = ast.NewNode(ast.NodeParentInfo{Parent: child2, IdxToParent: j}, "diff-label", "diff-value")
				} else {
					_, _ = ast.NewNode(ast.NodeParentInfo{Parent: child2, IdxToParent: j}, givenSameLabel, givenSameValue)
				}
			}
		}
		if root1.Isomorphic(root2) {
			t.Errorf("expected false, got true")
		}
	})

	t.Run("same children structure, 2 layers, same labels and values", func(t *testing.T) {
		root1, _ := ast.NewOrphanNode()
		root2, _ := ast.NewOrphanNode()

		const givenSameLabel ast.NodeLabelType = "same-label"
		const givenSameValue ast.NodeValueType = "same-value"

		for i := 0; i < 5; i++ {
			child1, _ := ast.NewNode(ast.NodeParentInfo{Parent: root1, IdxToParent: i}, givenSameLabel, givenSameValue)
			child2, _ := ast.NewNode(ast.NodeParentInfo{Parent: root2, IdxToParent: i}, givenSameLabel, givenSameValue)

			for j := 0; j < 5; j++ {
				_, _ = ast.NewNode(ast.NodeParentInfo{Parent: child1, IdxToParent: j}, givenSameLabel, givenSameValue)
				_, _ = ast.NewNode(ast.NodeParentInfo{Parent: child2, IdxToParent: j}, givenSameLabel, givenSameValue)
			}
		}
		if !root1.Isomorphic(root2) {
			t.Errorf("expected true, got false")
		}
	})

	t.Run("different children structure, 1 layer", func(t *testing.T) {
		root1, _ := ast.NewOrphanNode()
		root2, _ := ast.NewOrphanNode()

		const givenSameLabel ast.NodeLabelType = "same-label"
		const givenSameValue ast.NodeValueType = "same-value"
		const givenRoot1ChildrenNum = 5
		const givenRoot2ChildrenNum = 6

		for i := 0; i < givenRoot1ChildrenNum; i++ {
			_, _ = ast.NewNode(ast.NodeParentInfo{Parent: root1, IdxToParent: i}, givenSameLabel, givenSameValue)
		}
		for i := 0; i < givenRoot2ChildrenNum; i++ {
			_, _ = ast.NewNode(ast.NodeParentInfo{Parent: root2, IdxToParent: i}, givenSameLabel, givenSameValue)
		}

		if root1.Isomorphic(root2) {
			t.Errorf("expected false, got true")
		}
	})

	t.Run("different children structure, layer-2 is different", func(t *testing.T) {
		root1, _ := ast.NewOrphanNode()
		root2, _ := ast.NewOrphanNode()

		const givenSameLabel ast.NodeLabelType = "same-label"
		const givenSameValue ast.NodeValueType = "same-value"

		for i := 0; i < 5; i++ {
			child1, _ := ast.NewNode(ast.NodeParentInfo{Parent: root1, IdxToParent: i}, givenSameLabel, givenSameValue)
			child2, _ := ast.NewNode(ast.NodeParentInfo{Parent: root2, IdxToParent: i}, givenSameLabel, givenSameValue)

			for j := 0; j < 5; j++ {
				_, _ = ast.NewNode(ast.NodeParentInfo{Parent: child1, IdxToParent: j}, givenSameLabel, givenSameValue)
			}
			for j := 0; j < 6; j++ {
				_, _ = ast.NewNode(ast.NodeParentInfo{Parent: child2, IdxToParent: j}, givenSameLabel, givenSameValue)
			}
		}
		if root1.Isomorphic(root2) {
			t.Errorf("expected false, got true")
		}
	})
}
