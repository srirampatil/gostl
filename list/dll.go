package list

import (
	"github.com/srirampatil/stl"
)

// DLL represents a doubly linked list and imlpements IList and Iterable
// interfaces.
type DLL struct {
	Head, Tail *DLLNode
	Count      int
}

// NewDLL creates a new doubly linked list (DLL).
func NewDLL() (list *DLL) {
	list = new(DLL)
	list.Head, list.Tail = nil, nil
	list.Count = 0
	return
}

// Front returns the first object in DLL.
// Complexity: O(1).
func (list DLL) Front() stl.Comparable {
	if list.Head != nil {
		return list.Head.Value
	}
	return nil
}

// Back returns the last object in DLL.
// Complexity: O(1).
func (list DLL) Back() stl.Comparable {
	if list.Tail != nil {
		return list.Tail.Value
	}
	return nil
}

// Empty returns true if DLL is empty.
// Complexity: O(1).
func (list DLL) Empty() bool {
	return list.Count == 0
}

// Size returns the number of objects in DLL.
// Complexity: O(1).
func (list DLL) Size() int {
	return list.Count
}

// PushBack adds an object at the end of DLL.
// Complexity: O(1)
func (list *DLL) PushBack(v stl.Comparable) {
	node := DLLNode{v, nil, nil}
	if list.Head == nil {
		list.Head = &node
	} else {
		list.Tail.N = &node
		node.P = list.Tail
	}
	list.Tail = &node
	list.Count++
}

// PushFront adds an object at the start of DLL.
// Complexity: O(1).
func (list *DLL) PushFront(v stl.Comparable) {
	node := DLLNode{v, nil, nil}
	if list.Head != nil {
		node.P = list.Head
	}

	node.N, list.Head = list.Head, &node
	if list.Tail == nil {
		list.Tail = &node
	}
	list.Count++
}

// PopFront removes the first object from DLL.
// Complexity: O(1).
func (list *DLL) PopFront() {
	if list.Head != nil {
		node := list.Head
		list.Head = node.N
		node.N, node.P = nil, nil

		if list.Head == nil {
			list.Tail = nil
		}
	}
}

// PopBack removes the last object from DLL.
// Complexity: O(1).
func (list *DLL) PopBack() {
	node := list.Tail
	if node != nil {
		list.Tail = node.P
		node.N, node.P = nil, nil

		if list.Tail == nil {
			list.Head = nil
		} else {
			list.Tail.N = nil
		}
	}
}

// Begin returns an Iterator holding the first object of DLL.
// Complexity: O(1).
func (list *DLL) Begin() stl.Iterator {
	return list.Head
}

// End returns an Iterator holding the last object of DLL.
// Complexity: O(1).
func (list *DLL) End() stl.Iterator {
	return list.Tail
}

// Reverse reverses the DLL.
// Complexity: Linear in DLL size
func (list *DLL) Reverse() {
	var nextNode, prevNode *DLLNode = nil, nil
	if list.Head != nil {
		nextNode = list.Head.N
	}
	for node := list.Head; node != nil; {
		node.N = prevNode
		node.P = nextNode
		prevNode = node

		node = nextNode
		if node == nil {
			break
		}
		nextNode = nextNode.N
	}
	list.Tail = list.Head
	list.Head = prevNode
}
