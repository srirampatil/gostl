package common

import (
	"fmt"
	"testing"
)

func CheckExpected(t *testing.T, i int, r interface{}, e interface{}) {
	if r != e {
		t.Fatalf("Case %d => Expected %v, Received %v\n", i, e, r)
	} else {
		fmt.Printf("Case %d => Success!\n", i)
	}
}

func CheckIfEqual(t *testing.T, testNo int, r []interface{}, e []interface{}) {
	for i := range r {
		if r[i] != e[i] {
			t.Fatalf("Case %d => Idx: %d, Expected %v, Received %v\n", i, testNo, e[i], r[i])
		}
	}

	fmt.Printf("Case %d => Success!\n", testNo)
}

func Iterate(begin Iterator, end Iterator) []interface{} {
	var s []interface{}
	for itr := begin; !itr.Equals(end); itr.Next() {
		s = append(s, itr.Value())
	}
	return s
}
