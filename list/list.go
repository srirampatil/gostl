package list

// ListNode represents a node in a doubly linked list. It implements the Iterator
// interface.
type ListNode struct {
	value      interface{}
	next, prev *ListNode
}

func (node ListNode) Value() interface{} {
	return node.value
}

// List represents a doubly linked list and imlpements IList and Iterable
// interfaces.
type List struct {
	sentinel *ListNode
	size     int
}

// NewList creates a new doubly linked list (List).
func NewList() (list *List) {
	list = new(List)
	list.size = 0
	list.sentinel = new(ListNode)
	list.sentinel.next = list.sentinel
	list.sentinel.prev = list.sentinel
	return
}

// Front returns the first object in List.
// Complexity: O(1).
func (list List) Front() interface{} {
	if list.sentinel.next != list.sentinel {
		return list.sentinel.next.Value()
	}
	return nil
}

// Back returns the last object in List.
// Complexity: O(1).
func (list List) Back() interface{} {
	if list.sentinel.prev != list.sentinel {
		return list.sentinel.prev.Value()
	}
	return nil
}

// Empty returns true if List is empty.
// Complexity: O(1).
func (list List) Empty() bool {
	return list.size == 0
}

// Size returns the number of objects in List.
// Complexity: O(1).
func (list List) Size() int {
	return list.size
}

func insertAfter(node *ListNode, after *ListNode) {
	node.prev = after
	node.next = after.next

	after.next.prev = node
	after.next = node
}

func removeAfter(after *ListNode) {
	nextNode := after.next
	after.next = nextNode.next
	nextNode.next.prev = after

	nextNode.next, nextNode.prev = nil, nil
}

func removeBefore(before *ListNode) {
	prevNode := before.prev
	before.prev = prevNode.prev
	prevNode.prev.next = before

	prevNode.prev, prevNode.next = nil, nil
}

// PushBack adds an object at the end of List.
// Complexity: O(1)
func (list *List) PushBack(v interface{}) {
	node := ListNode{v, nil, nil}
	insertAfter(&node, list.sentinel.prev)
	list.size++
}

// PushFront adds an object at the start of List.
// Complexity: O(1).
func (list *List) PushFront(v interface{}) {
	node := ListNode{v, nil, nil}
	insertAfter(&node, list.sentinel)
	list.size++
}

// PopFront removes the first object from List.
// Complexity: O(1).
func (list *List) PopFront() {
	if list.Empty() {
		return
	}
	removeAfter(list.sentinel)
	list.size--
}

// PopBack removes the last object from List.
// Complexity: O(1).
func (list *List) PopBack() {
	if list.Empty() {
		return
	}
	removeBefore(list.sentinel)
	list.size--
}

/*
// Begin returns an Iterator holding the first object of List.
// Complexity: O(1).
func (list *List) Begin() stl.Iterator {
	return list.head
}

// End returns an Iterator holding the last object of List.
// Complexity: O(1).
func (list *List) End() stl.Iterator {
	return list.tail
}
*/

// Reverse reverses the List.
// Complexity: Linear in size of List
func (list *List) Reverse() {
	if list.Empty() {
		return
	}

	node := list.sentinel
	var nextNode, prevNode *ListNode = node.next, node.prev
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

// Swap swaps two Lists.
// Complexity: O(1).
func (lhs *List) Swap(rhs *List) error {
	lhs.sentinel, rhs.sentinel = rhs.sentinel, lhs.sentinel
	lhs.size, rhs.size = rhs.size, lhs.size
	return nil
}

// Clear removes all the elements from List.
// Complexity: Linear in size of List
func (list *List) Clear() {
	for !list.Empty() {
		list.PopBack()
	}
}
