package container

import "github.com/srirampatil/gostl/common"

type IteratorDirection uint8

const (
	FORWARD IteratorDirection = iota + 1
	BACKWARD
	BOTH
	NONE
)

// Iterator for Deque.
type DequeIterator struct {
	direction IteratorDirection
	index     int
	deque     *Deque
}

// Deque is a double ended queue. Deque supports push and pop from both ends.
type Deque struct {
	front, back int
	size        int
	values      []interface{}
	endItr      *DequeIterator
}

func (itr *DequeIterator) moveToPrev() common.Iterator {
	itr.index = itr.index - 1
	return itr
}

func (itr *DequeIterator) moveToNext() common.Iterator {
	itr.index = itr.index + 1
	if itr.index >= itr.deque.Size() {
		itr.index = -1
	}
	return itr
}

// Next returns an Iterator pointing to the next element in Deque.
func (itr *DequeIterator) Next() common.Iterator {
	switch itr.direction {
	case BACKWARD:
		return itr.moveToPrev()
	case FORWARD:
		return itr.moveToNext()
	}

	return itr
}

// Value returns the element pointed by Iterator.
func (itr *DequeIterator) Value() interface{} {
	return itr.deque.At(itr.index)
}

// Equals returns true if both the Iterators are pointing to the same element.
func (lhs *DequeIterator) Equals(r common.Iterator) bool {
	rhs, ok := r.(*DequeIterator)
	if !ok {
		return false
	}

	return (lhs.deque == rhs.deque) && (lhs.index == rhs.index)
}

// NewDeque returns a new Deque object.
func NewDeque(cap int) *Deque {
	q := new(Deque)
	q.values = make([]interface{}, cap, cap)
	q.front = 0
	q.back = 0
	q.size = 0
	q.endItr = new(DequeIterator)
	q.endItr.index = -1
	q.endItr.deque = q
	return q
}

// Size returns the number of elements stored in Deque.
func (q *Deque) Size() int {
	return q.size
}

// MaxSize returns the current capacity of Deque.
func (q *Deque) MaxSize() int {
	return cap(q.values)
}

// Empty returns true if Deque is empty.
func (q *Deque) Empty() bool {
	return q.size == 0
}

// Resize changes the capacity of Deque. MaxSize is affected after Resize. Size
// remains the same. If new Cap is less than the current MaxSize of then Deque
// is truncated to newCap.
func (q *Deque) Resize(newCap int) {
	if newCap != cap(q.values) {
		newValues := make([]interface{}, newCap, newCap)
		var j int = 0
		for i := q.front; j < q.size && j < newCap; i, j = (i+1)%cap(q.values), j+1 {
			newValues[j] = q.values[i]
		}
		q.values = newValues
		q.front = 0
		q.back = j
	}
}

// ShrinkToFit resizes the Deque to Deque.Size().
func (q *Deque) ShrinkToFit() {
	q.Resize(q.Size())
}

// At returns the element at index idx.
func (q *Deque) At(idx int) interface{} {
	if idx >= len(q.values) {
		return nil
	}
	actualIdx := idx
	if q.front != 0 {
		actualIdx = (idx + q.front) % cap(q.values)
	}
	return q.values[actualIdx]
}

// Front returns the first elemnt in the Deque.
func (q *Deque) Front() interface{} {
	if !q.Empty() {
		return q.values[q.front]
	}
	return nil
}

// Back returns the Last element in the Deque.
func (q *Deque) Back() interface{} {
	if !q.Empty() {
		return q.values[q.back-1]
	}
	return nil
}

// PushBack adds v to the end of the Deque.
func (q *Deque) PushBack(v interface{}) {
	if q.size == cap(q.values) {
		newCap := 2
		if cap(q.values) > 0 {
			newCap = cap(q.values) * 2
		}
		q.Resize(newCap)
	}

	q.values[q.back] = v
	q.size++
	q.back = (q.back + 1) % cap(q.values)
}

// PushFront adds v to the start of the Deque.
func (q *Deque) PushFront(v interface{}) {
	if q.size == cap(q.values) {
		newCap := 2
		if cap(q.values) > 0 {
			newCap = cap(q.values) * 2
		}
		q.Resize(newCap)
	}

	q.front--
	if q.front < 0 {
		q.front = cap(q.values) - 1
	}

	q.values[q.front] = v
	q.size++
}

// PopBack removes an element from the end of Deque.
func (q *Deque) PopBack() {
	if !q.Empty() {
		q.back--
		q.size--
	}
}

// PopFront removes an element from the start of the Deque.
func (q *Deque) PopFront() {
	if !q.Empty() {
		q.front = (q.front + 1) % cap(q.values)
		q.size--
	}
}

// Clear removes everything from the Deque.
func (q *Deque) Clear() {
	q.front = 0
	q.back = 0
	q.size = 0
}

// Begin returns an Iterator pointing to the first element of the Deque.
func (q *Deque) Begin() common.Iterator {
	itr := new(DequeIterator)
	itr.index = 0
	itr.direction = FORWARD
	itr.deque = q
	return itr
}

// End returns the Iterator pointing to a theoretical element past-the-end of Deque.
func (q *Deque) End() common.Iterator {
	return q.endItr
}

// Rbegin returns an Iterator poiting to the last element of the Deque, which
// moves in reverse direction.
func (q *Deque) Rbegin() common.Iterator {
	itr := new(DequeIterator)
	itr.index = q.Size() - 1
	itr.direction = BACKWARD
	itr.deque = q
	return itr
}

// Rend returns the Iterator pointing to a theoretical elemtn before-the-start of Deque.
func (q *Deque) Rend() common.Iterator {
	return q.endItr
}
