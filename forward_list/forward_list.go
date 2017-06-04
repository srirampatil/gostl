package forward_list

type listNode struct {
	value interface{}
	next  *listNode
}

type Iterator struct {
	node *listNode
	list *ForwardList
}

type ForwardList struct {
	sentinel *listNode
	endItr   *Iterator
	size     int
}

// NewForwardList creates a new ForwardList and returns a pointer to the same.
func NewForwardList() *ForwardList {
	list := new(ForwardList)
	list.size = 0
	list.sentinel = new(listNode)
	list.sentinel.next = list.sentinel
	list.endItr = new(Iterator)
	list.endItr.node = list.sentinel
	list.endItr.list = list
	return list
}

// Next moves the Iterator to the next node.
func (itr *Iterator) Next() *Iterator {
	itr.node = itr.node.next
	return itr
}

// Value returns the value stored in the node pointed by the Iterator.
func (itr *Iterator) Value() interface{} {
	return itr.node.value
}

// Equals returns true if both Iterators are pointing to the same node.
func (lhs *Iterator) Equals(rhs *Iterator) bool {
	if lhs.list != rhs.list {
		return false
	}

	if lhs.node == rhs.node {
		return true
	}
	return false
}

// Empty return true if the ForwardList is empty.
func (list *ForwardList) Empty() bool {
	return (list.size == 0)
}

// Size returns the number of nodes in the ForwardList.
func (list *ForwardList) Size() int {
	return list.size
}

// Front returns the values stores in the first node of the ForwardList if it is
// not empty.
func (list *ForwardList) Front() interface{} {
	if !list.Empty() {
		return list.sentinel.next.value
	}
	return nil
}

// PushFront creates a new node containing value v and adds it to the start of
// the ForwardList.
func (list *ForwardList) PushFront(v interface{}) {
	node := new(listNode)
	node.value = v
	node.next = list.sentinel.next
	list.sentinel.next = node
	list.size++
}

// PopFront removed the first node from the ForwardList.
func (list *ForwardList) PopFront() {
	if !list.Empty() {
		node := list.sentinel.next
		list.sentinel.next = node.next
		node.next = nil
		list.size--
	}
}

// Clear removes all the nodes from the ForwardList.
func (list *ForwardList) Clear() {
	list.sentinel.next = list.sentinel
	list.size = 0
}

// Begin returns an Iterator pointing to the first node of the ForwardList.
func (list *ForwardList) Begin() *Iterator {
	itr := new(Iterator)
	itr.node = list.sentinel.next
	itr.list = list
	return itr
}

// End returns an Iterator pointing to the node past-the-end of the ForwardList.
func (list *ForwardList) End() *Iterator {
	return list.endItr
}

// Reverse reverses the ForwardList in place.
func (list *ForwardList) Reverse() {
	var prev, curr, next *listNode = list.sentinel, list.sentinel.next, list.sentinel.next.next

	for prev != next {
		curr.next = prev
		prev = curr
		curr = next
		next = curr.next
	}
}
