package adaptor

import "github.com/srirampatil/gostl/container"

type Stack struct {
	cType container.ContainerType
	c     container.Container
}

func NewStack(cType container.ContainerType) *Stack {
	s := new(Stack)
	s.cType = cType
	switch cType {
	case container.ContainerTypeForwardList:
		s.c = container.NewForwardList()
	case container.ContainerTypeList:
		s.c = container.NewList()
	case container.ContainerTypeDeque:
		s.c = container.NewDeque(10)
	}
	return s
}

func (s *Stack) Empty() bool {
	return s.c.Empty()
}

func (s *Stack) Size() int {
	return s.c.Size()
}

func (s *Stack) Top() interface{} {
	return s.c.Front()
}

func (s *Stack) Push(v interface{}) {
	s.c.PushFront(v)
}

func (s *Stack) Pop() {
	s.c.PopFront()
}
