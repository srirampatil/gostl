package deque

import "github.com/srirampatil/go-stl/list"

type Deque struct {
	list   *list.List
	values []interface{}
}

func NewDeque(cap int) *Deque {
	q := new(Deque)
	q.values = make([]interface{}, 0, cap)
	q.list = list.NewList()
	return q
}

func (q *Deque) Size() int {
	return len(q.values)
}

func (q *Deque) MaxSize() int {
	return cap(q.values)
}

func (q *Deque) Empty() bool {
	return len(q.values) == 0
}

func (q *Deque) Resize(newCap int) {
	if newCap != cap(q.values) {
		newValues := make([]interface{}, 0, newCap)
		for i := range q.values {
			if i >= newCap {
				break
			}
			newValues[i] = q.values[i]
		}
		q.values = newValues
		q.list.Clear()
		for _, v := range q.values {
			q.list.PushBack(v)
		}
	}
}

func (q *Deque) ShrinkToFit() {
	q.Resize(q.Size())
}

func (q *Deque) At(idx int) interface{} {
	if idx >= len(q.values) {
		return nil
	}
	return q.values[idx]
}

func (q *Deque) Front() interface{} {
	if !q.Empty() {
		return q.values[0]
	}
	return nil
}

func (q *Deque) Back() interface{} {
	if !q.Empty() {
		return q.values[q.Size()-1]
	}
	return nil
}
