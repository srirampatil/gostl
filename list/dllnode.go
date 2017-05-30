package list

import (
	"github.com/srirampatil/stl"
)

// DLLNode represents a node in a doubly linked list. It implements the Iterator
// interface.
type DLLNode struct {
	value      stl.Comparable
	next, prev *DLLNode
}

func (node DLLNode) Value() stl.Comparable {
	return node.value
}

func (node *DLLNode) SetValue(v stl.Comparable) {
	node.value = v
}
