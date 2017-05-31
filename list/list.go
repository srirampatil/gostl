// Package list contains basic sequential data structures.
package list

import "github.com/srirampatil/stl"

// IList declares operations for a sequential list.
type IList interface {
	// Front returns the first object in IList.
	Front() stl.Comparable

	// Back returns the last object in IList.
	Back() stl.Comparable

	// Empty returns true if IList is empty.
	Empty() bool

	// Size returns the number of objects in the IList.
	Size() int

	// PushBack adds a Comparable object at the end of IList.
	PushBack(v stl.Comparable)

	// PushFront adds a Comparable object at the start of the IList.
	PushFront(v stl.Comparable)

	// PopFront removes the first object from IList.
	PopFront()

	// PopBack removes the last object in IList.
	PopBack()

	// Reverse reverses the IList.
	Reverse()

	// Swap the lists. Size of the lists can be different.
	Swap(rhs IList) error

	// Clear removes all the elements from the list.
	Clear()
}
