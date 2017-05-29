package stl

import "fmt"

// An error, used when the expected and received types do not match for a
// generic function.
type TypeMismatchError struct {
	// Expeted value
	Expected interface{}

	// Recevived value
	Received interface{}
}

func (e TypeMismatchError) Error() string {
	return fmt.Sprintf("TypeMismatchError: Expected %T, Received %T\n", e.Expected, e.Received)
}
