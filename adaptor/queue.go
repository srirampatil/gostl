// Package adaptor implements container adaptors. Container adaptors are the
// the classes which use an underlying container which implements a specific
// set of functions which satisfies the adaptors behaviour.
package adaptor

// QueueContainer specifies a contract to be met by the container used by Queue.
type QueueContainer interface {
	Empty() bool
	Size() int
	Front() interface{}
	Back() interface{}
	PushBack(v interface{})
	PopFront()
}

// Queue implements a FIFO (first in first out) adaptor. The elements are
// inserted into one end and extracted from another. The underlying container
// can be one of the standard container classes or a custom class which
// implements QueueContainer.
type Queue struct {
	c QueueContainer
}

// NewQueue returns a new Queue object initialized with the given Container object.
func NewQueue(c QueueContainer) *Queue {
	q := new(Queue)
	q.c = c
	return q
}

// Empty returns true if Queue is empty.
func (q *Queue) Empty() bool {
	return q.c.Empty()
}

// Size returns the number of elements stored in the Queue.
func (q *Queue) Size() int {
	return q.c.Size()
}

// Front returns the first element in the Queue.
func (q *Queue) Front() interface{} {
	return q.c.Front()
}

// Back returns the last element in the Queue. This is the recently added element.
func (q *Queue) Back() interface{} {
	return q.c.Back()
}

// Push adds an element at the end of the Queue.
func (q *Queue) Push(v interface{}) {
	q.c.PushBack(v)
}

// Pop removes the first element from the Queue.
func (q *Queue) Pop() {
	q.c.PopFront()
}
