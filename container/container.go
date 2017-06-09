package container

type Container interface {
	Empty() bool
	Size() int
	Front() interface{}
	Back() interface{}
	PushFront(v interface{})
	PushBack(v interface{})
	PopFront()
	PopBack()
}

type ContainerType uint8

const (
	ContainerTypeForwardList ContainerType = iota + 1
	ContainerTypeList
	ContainerTypeDeque
)
