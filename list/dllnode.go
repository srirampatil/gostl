package list

import (
	"github.com/srirampatil/stl"
)

// DLLNode represents a node in a doubly linked list. It implements the Iterator
// interface.
type DLLNode struct {
	Value stl.Comparable
	N, P  *DLLNode
}

func (node *DLLNode) HasNext() bool {
	if node == nil {
		return false
	}
	return node.N != nil
}

func (node *DLLNode) Next() stl.Iterator {
	if node == nil {
		return nil
	}
	return node.N
}
