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
