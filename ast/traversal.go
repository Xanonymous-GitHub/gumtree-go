package ast

func (a *astConcrete) PreOrderNodes() []*Node {
	return a.preOrder(a.Root())
}

func (a *astConcrete) preOrder(node *Node) []*Node {
	result := make([]*Node, 0)
	if node == nil {
		return result
	}

	result = append(result, node)
	for _, child := range node.OrderedChildren() {
		result = append(result, a.preOrder(child)...)
	}
	return result
}

func (a *astConcrete) PostOrderNodes() []*Node {
	return a.postOrder(a.Root())
}

func (a *astConcrete) postOrder(node *Node) []*Node {
	result := make([]*Node, 0)
	if node == nil {
		return result
	}

	for _, child := range node.OrderedChildren() {
		result = append(result, a.postOrder(child)...)
	}

	result = append(result, node)
	return result
}
