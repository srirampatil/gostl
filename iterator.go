package stl

// Iterator is a type which Iterable holds.
type Iterator interface {
	// HasNext returns true if there is a next object available in the Iterable
	// object.
	HasNext() bool

	// Next returns the next object in the Iterable object.
	Next() Iterator
}

// Iterable represents a type which can be iterated.
type Iterable interface {
	// Begin returns the Iterator holding the first object.
	Begin() Iterator

	// End returns the Iterator holding the last object.
	End() Iterator
	//	RBegin() ReverseIterator
	//	REnd() ReverseIterator
}
