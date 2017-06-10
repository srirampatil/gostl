package tests

import (
	"fmt"
	"testing"

	"github.com/srirampatil/gostl/common"
)

type OpType uint8

const (
	NEW OpType = iota + 1
	EMPTY
	SIZE
	AT
	FRONT
	BACK
	MAXSIZE
	RESIZE
	SHRINKTOFIT
	PUSHBACK
	PUSHFRONT
	POPBACK
	POPFRONT
	CLEAR
	ITERATE
	REVERSE_ITERATE
	PUSH
	POP
	REVERSE
	TOP
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
			t.Fatalf("Case %d => Idx: %d, Expected %v, Received %v\n", testNo, i, e[i], r[i])
		}
	}

	fmt.Printf("Case %d => Success!\n", testNo)
}

func Iterate(begin common.Iterator, end common.Iterator) []interface{} {
	var s []interface{}
	for itr := begin; !itr.Equals(end); itr.Next() {
		s = append(s, itr.Value())
	}
	return s
}
