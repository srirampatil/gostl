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

func CheckIfEqual(t *testing.T, testNo int, r []int, e []int) {
	for i := range r {
		if r[i] != e[i] {
			t.Fatalf("Case %d => Idx: %d, Expected %v, Received %v\n", i, testNo, e[i], r[i])
		}
	}

	fmt.Printf("Case %d => Success!\n", testNo)
}

func Iterate(begin Iterator, end Iterator) []int {
	var s []int
	for itr := begin; !itr.Equals(end); itr.Next() {
		v := itr.Value().(int)
		s = append(s, v)
	}
	return s
}
