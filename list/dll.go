package list

import (
	"github.com/srirampatil/stl"
)

// DLL represents a doubly linked list and imlpements IList and Iterable
// interfaces.
type DLL struct {
	sentinel *DLLNode
	size     int
}

// NewDLL creates a new doubly linked list (DLL).
func NewDLL() (list *DLL) {
	list = new(DLL)
	list.size = 0
	list.sentinel = new(DLLNode)
	list.sentinel.next = list.sentinel
	list.sentinel.prev = list.sentinel
	return
}

// Front returns the first object in DLL.
// Complexity: O(1).
func (list DLL) Front() stl.Comparable {
	if list.sentinel.next != list.sentinel {
		return list.sentinel.next.Value()
	}
	return nil
}

// Back returns the last object in DLL.
// Complexity: O(1).
func (list DLL) Back() stl.Comparable {
	if list.sentinel.prev != list.sentinel {
		return list.sentinel.prev.Value()
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

func insertAfter(node *DLLNode, after *DLLNode) {
	node.prev = after
	node.next = after.next

	after.next.prev = node
	after.next = node
}

func removeAfter(after *DLLNode) {
	nextNode := after.next
	after.next = nextNode.next
	nextNode.next.prev = after

	nextNode.next, nextNode.prev = nil, nil
}

func removeBefore(before *DLLNode) {
	prevNode := before.prev
	before.prev = prevNode.prev
	prevNode.prev.next = before

	prevNode.prev, prevNode.next = nil, nil
}

// PushBack adds an object at the end of DLL.
// Complexity: O(1)
func (list *DLL) PushBack(v stl.Comparable) {
	node := DLLNode{v, nil, nil}
	insertAfter(&node, list.sentinel.prev)
	list.size++
}

// PushFront adds an object at the start of DLL.
// Complexity: O(1).
func (list *DLL) PushFront(v stl.Comparable) {
	node := DLLNode{v, nil, nil}
	insertAfter(&node, list.sentinel)
	list.size++
}

// PopFront removes the first object from DLL.
// Complexity: O(1).
func (list *DLL) PopFront() {
	if list.Empty() {
		return
	}
	removeAfter(list.sentinel)
	list.size--
}

// PopBack removes the last object from DLL.
// Complexity: O(1).
func (list *DLL) PopBack() {
	if list.Empty() {
		return
	}
	removeBefore(list.sentinel)
	list.size--
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
// Complexity: Linear in size of DLL
func (list *DLL) Reverse() {
	if list.Empty() {
		return
	}

	node := list.sentinel
	var nextNode, prevNode *DLLNode = node.next, node.prev
	for {
		node.next = prevNode
		node.prev = nextNode

		prevNode = node
		node = nextNode
		nextNode = nextNode.next

		if node == list.sentinel {
			break
		}
	}
}

// Swap swaps two DLLs.
// Complexity: O(1).
func (lhs *DLL) Swap(r IList) error {
	rhs, ok := r.(*DLL)
	if !ok {
		return stl.TypeMismatchError{lhs, r}
	}

	lhs.sentinel, rhs.sentinel = rhs.sentinel, lhs.sentinel
	lhs.size, rhs.size = rhs.size, lhs.size
	return nil
}

// Clear removes all the elements from DLL.
// Complexity: Linear in size of DLL
func (list *DLL) Clear() {
	for !list.Empty() {
		list.PopBack()
	}
}
