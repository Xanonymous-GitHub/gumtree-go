package ast

type NodeHashMemo map[NodeIdType]uint64

func (m *NodeHashMemo) IsIsomorphicBetween(n1, n2 NodeIdType) bool {
	return (*m)[n1] == (*m)[n2]
}
