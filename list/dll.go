package list

import (
	"github.com/srirampatil/stl"
)

// DLL represents a doubly linked list and imlpements IList and Iterable
// interfaces.
type DLL struct {
	head, tail *DLLNode
	size       int
}

// NewDLL creates a new doubly linked list (DLL).
func NewDLL() (list *DLL) {
	list = new(DLL)
	list.head, list.tail = nil, nil
	list.size = 0
	return
}

// Front returns the first object in DLL.
// Complexity: O(1).
func (list DLL) Front() stl.Comparable {
	if list.head != nil {
		return list.head.Value()
	}
	return nil
}

// Back returns the last object in DLL.
// Complexity: O(1).
func (list DLL) Back() stl.Comparable {
	if list.tail != nil {
		return list.tail.Value()
	}
	return nil
}

// Empty returns true if DLL is empty.
// Complexity: O(1).
func (list DLL) Empty() bool {
	return list.size == 0
}

// Size returns the number of objects in DLL.
// Complexity: O(1).
func (list DLL) Size() int {
	return list.size
}

// PushBack adds an object at the end of DLL.
// Complexity: O(1)
func (list *DLL) PushBack(v stl.Comparable) {
	node := DLLNode{v, nil, nil}
	if list.head == nil {
		list.head = &node
	} else {
		list.tail.next = &node
		node.prev = list.tail
	}
	list.tail = &node
	list.size++
}

// PushFront adds an object at the start of DLL.
// Complexity: O(1).
func (list *DLL) PushFront(v stl.Comparable) {
	node := DLLNode{v, nil, nil}
	if list.head != nil {
		node.prev = list.head
	}

	node.next = list.head
	list.head = &node
	if list.tail == nil {
		list.tail = &node
	}
	list.size++
}

// PopFront removes the first object from DLL.
// Complexity: O(1).
func (list *DLL) PopFront() {
	if list.head != nil {
		node := list.head
		list.head = node.next
		node.next = nil
		node.prev = nil

		if list.head == nil {
			list.tail = nil
		}
	}
}

// PopBack removes the last object from DLL.
// Complexity: O(1).
func (list *DLL) PopBack() {
	node := list.tail
	if node != nil {
		list.tail = node.prev
		node.next = nil
		node.prev = nil

		if list.tail == nil {
			list.head = nil
		} else {
			list.tail.next = nil
		}
	}
}

/*
// Begin returns an Iterator holding the first object of DLL.
// Complexity: O(1).
func (list *DLL) Begin() stl.Iterator {
	return list.head
}

// End returns an Iterator holding the last object of DLL.
// Complexity: O(1).
func (list *DLL) End() stl.Iterator {
	return list.tail
}
*/

// Reverse reverses the DLL.
// Complexity: Linear in DLL size
func (list *DLL) Reverse() {
	var nextNode, prevNode *DLLNode = nil, nil
	if list.head != nil {
		nextNode = list.head.next
	}
	for node := list.head; node != nil; {
		node.next = prevNode
		node.prev = nextNode
		prevNode = node

		node = nextNode
		if node == nil {
			break
		}
		nextNode = nextNode.next
	}
	list.tail = list.head
	list.head = prevNode
}
