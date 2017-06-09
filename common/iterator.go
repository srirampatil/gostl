package common

type IteratorDirection uint8

const (
	FORWARD IteratorDirection = iota + 1
	BACKWARD
	BOTH
	NONE
)

// Iterator is a type which Iterable holds.
type Iterator interface {
	// Next returns an iterator pointing to the next object.
	Next() Iterator

	// Value returns the value of the object pointed by the Iterator.
	Value() interface{}

	// Equals returns true if the Iterators point to the same object.
	Equals(r Iterator) bool
}

// Iterable represents a type which can be iterated.
type Iterable interface {
	// Begin returns an Iterator pointing to the first object.
	Begin() Iterator

	// End returns an Iterator pointing to a theoretical past-the-end object.
	End() Iterator
}

type ReverseIterable interface {
	// Rbegin returns an Iterator pointing to the last object.
	Rbegin() Iterator

	// Rend returns an Iterator pointing to a theoretical before-the-start object.
	Rend() Iterator
}
