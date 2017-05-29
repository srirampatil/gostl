// Package stl declares the common interfaces used by other data structures in
// this package
package stl

// A type, which allows comparing the values. All the data structures store
// objects of type Comparable.
type Comparable interface {
	// Compare compares two Comparable objects and should return 0 if they are
	// equal. It should return -1 if lhs < rhs and 1 if lhs > rhs
	Compare(rhs Comparable) (int, error)
}
