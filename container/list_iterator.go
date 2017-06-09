package container

import (
	"fmt"

	"github.com/srirampatil/gostl/common"
)

type ListIterator struct {
	ptr *listNode
	dir common.IteratorDirection
}

func (it *ListIterator) Next() common.Iterator {
	switch it.dir {
	case common.NONE:
		return it
	case common.BACKWARD:
		it.ptr = it.ptr.prev
	default:
		it.ptr = it.ptr.next
	}
	return it
}

func (it *ListIterator) Value() interface{} {
	return it.ptr.Value()
}

func (lhs *ListIterator) Equals(r common.Iterator) bool {
	rhs := r.(*ListIterator)
	if lhs.ptr != rhs.ptr {
		return false
	}

	if lhs.dir != rhs.dir && lhs.dir != common.NONE && rhs.dir != common.NONE {
		return false
	}
	return true
}

func (it *ListIterator) String() string {
	return fmt.Sprintf("%T(%v)", it.Value(), it.Value())
}
