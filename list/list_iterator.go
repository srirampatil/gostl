package list

import "fmt"

type IteratorDirection uint8

const (
	FORWARD IteratorDirection = iota + 1
	BACKWARD
	BOTH
	NONE
)

type ListIterator struct {
	ptr *listNode
	dir IteratorDirection
}

func (it *ListIterator) Next() *ListIterator {
	switch it.dir {
	case NONE:
		return it
	case BACKWARD:
		it.ptr = it.ptr.prev
	default:
		it.ptr = it.ptr.next
	}
	return it
}

func (it *ListIterator) Prev() *ListIterator {
	switch it.dir {
	case NONE:
		return it
	case BACKWARD:
		it.ptr = it.ptr.next
	default:
		it.ptr = it.ptr.prev
	}
	return it
}

func (it *ListIterator) Value() interface{} {
	return it.ptr.Value()
}

func (lhs *ListIterator) Equals(rhs *ListIterator) bool {
	if lhs.ptr != rhs.ptr {
		return false
	}

	if lhs.dir != rhs.dir && lhs.dir != NONE && rhs.dir != NONE {
		return false
	}
	return true
}

func (it *ListIterator) String() string {
	return fmt.Sprintf("%T(%v)", it.Value(), it.Value())
}
