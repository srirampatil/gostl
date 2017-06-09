package adaptors

import "github.com/srirampatil/gostl/container"

type Queue struct {
	cType container.ContainerType
	c     container.Container
}

func NewQueue(cType container.ContainerType) *Queue {
	q := new(Queue)
	q.cType = cType
	switch cType {
	case container.ContainerTypeForwardList:
		panic("Queue does not support ForwardList container")
	case container.ContainerTypeList:
		q.c = container.NewList()
	case container.ContainerTypeDeque:
		q.c = container.NewDeque(10)
	}
	return q
}

func (q *Queue) Empty() bool {
	return q.c.Empty()
}

func (q *Queue) Size() int {
	return q.c.Size()
}

func (q *Queue) Front() interface{} {
	return q.c.Front()
}

func (q *Queue) Back() interface{} {
	return q.c.Back()
}

func (q *Queue) Push(v interface{}) {
	q.c.PushBack(v)
}

func (q *Queue) Pop() {
	q.c.PopFront()
}
