package adaptor

// StackContainer specifies a contract to be met by the container used by Stack.
type StackContainer interface {
	Empty() bool
	Size() int
	Front() interface{}
	PushFront(v interface{})
	PopFront()
}

// Stack implements a LIFO (last in first out) adaptor. The elements are inserted
// into adn extracted from the same end. The underlying container can be one of
// the standard containers or a custom container which implements StackContainer.
type Stack struct {
	c StackContainer
}

// NewStack returns a new Stack initialized with given Container.
func NewStack(c StackContainer) *Stack {
	s := new(Stack)
	s.c = c
	return s
}

// Empty returns true with the Stack is empty.
func (s *Stack) Empty() bool {
	return s.c.Empty()
}

// Size returns the number of elements stored in the Stack.
func (s *Stack) Size() int {
	return s.c.Size()
}

// Top returns top element of the Stack. This is a recently pushed element.
func (s *Stack) Top() interface{} {
	return s.c.Front()
}

// Push adds an element in the Stack.
func (s *Stack) Push(v interface{}) {
	s.c.PushFront(v)
}

// Pop removes the top element from the Stack.
func (s *Stack) Pop() {
	s.c.PopFront()
}
