package utils

import "cmp"

// AllowOrdered is an extension interface for types that designed to be ordered.
type AllowOrdered[P cmp.Ordered] interface {
	ValueOfOrder() P
}

// Equatable is an extension interface for types that designed to be compared.
type Equatable interface {
	IsEqualTo(other interface{}) bool
}
