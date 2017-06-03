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

func (itr *Iterator) Next() *Iterator {
	itr.node = itr.node.next
	return itr
}

func (itr *Iterator) Value() interface{} {
	return itr.node.value
}

func (lhs *Iterator) Equals(rhs *Iterator) bool {
	if lhs.list != rhs.list {
		return false
	}

	if lhs.node == rhs.node {
		return true
	}
	return false
}

func (list *ForwardList) Empty() bool {
	return (list.size == 0)
}

func (list *ForwardList) Size() int {
	return list.size
}

func (list *ForwardList) Front() interface{} {
	if !list.Empty() {
		return list.sentinel.next.value
	}
	return nil
}

func (list *ForwardList) PushFront(v interface{}) {
	node := new(listNode)
	node.value = v
	node.next = list.sentinel.next
	list.sentinel.next = node
	list.size++
}

func (list *ForwardList) PopFront() {
	if !list.Empty() {
		node := list.sentinel.next
		list.sentinel.next = node.next
		node.next = nil
		list.size--
	}
}

func (list *ForwardList) Clear() {
	list.sentinel.next = list.sentinel
	list.size = 0
}

func (list *ForwardList) Begin() *Iterator {
	itr := new(Iterator)
	itr.node = list.sentinel.next
	itr.list = list
	return itr
}

func (list *ForwardList) End() *Iterator {
	return list.endItr
}

func (list *ForwardList) Reverse() {
	var prev, curr, next *listNode = list.sentinel, list.sentinel.next, list.sentinel.next.next

	for prev != next {
		curr.next = prev
		prev = curr
		curr = next
		next = curr.next
	}
}
