package deque

import "github.com/srirampatil/gostl/common"

type IteratorDirection uint8

const (
	FORWARD IteratorDirection = iota + 1
	BACKWARD
	BOTH
	NONE
)

type DequeIterator struct {
	direction IteratorDirection
	index     int
	deque     *Deque
}

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

func (itr *DequeIterator) Next() common.Iterator {
	switch itr.direction {
	case BACKWARD:
		return itr.moveToPrev()
	case FORWARD:
		return itr.moveToNext()
	}

	return itr
}

func (itr *DequeIterator) Value() interface{} {
	return itr.deque.At(itr.index)
}

func (lhs *DequeIterator) Equals(r common.Iterator) bool {
	rhs, ok := r.(*DequeIterator)
	if !ok {
		return false
	}

	return (lhs.deque == rhs.deque) && (lhs.index == rhs.index)
}

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

func (q *Deque) Size() int {
	return q.size
}

func (q *Deque) MaxSize() int {
	return cap(q.values)
}

func (q *Deque) Empty() bool {
	return q.size == 0
}

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

func (q *Deque) ShrinkToFit() {
	q.Resize(q.Size())
}

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

func (q *Deque) Front() interface{} {
	if !q.Empty() {
		return q.values[q.front]
	}
	return nil
}

func (q *Deque) Back() interface{} {
	if !q.Empty() {
		return q.values[q.back-1]
	}
	return nil
}

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

func (q *Deque) PopBack() {
	if !q.Empty() {
		q.back--
		q.size--
	}
}

func (q *Deque) PopFront() {
	if !q.Empty() {
		q.front = (q.front + 1) % cap(q.values)
		q.size--
	}
}

func (q *Deque) Clear() {
	q.front = 0
	q.back = 0
	q.size = 0
}

func (q *Deque) Begin() common.Iterator {
	itr := new(DequeIterator)
	itr.index = 0
	itr.direction = FORWARD
	itr.deque = q
	return itr
}

func (q *Deque) End() common.Iterator {
	return q.endItr
}
