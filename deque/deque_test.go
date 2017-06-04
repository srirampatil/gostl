package deque

import (
	"fmt"
	"testing"
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
)

func checkExpected(t *testing.T, i int, r interface{}, e interface{}) {
	if r != e {
		t.Fatalf("Case %d => Expected %v, Received %v\n", i, e, r)
	} else {
		fmt.Printf("Case %d => Success!\n", i)
	}
}

func TestDeque(t *testing.T) {
	var cases = []struct {
		op     OpType
		empty  bool
		size   int
		cap    int
		value  interface{}
		values []interface{}
	}{
		{op: NEW},
		{op: EMPTY, empty: true},
		{op: SIZE, size: 0},
		{op: MAXSIZE, cap: 10},
		{op: SHRINKTOFIT},
		{op: EMPTY, empty: true},
		{op: SIZE, size: 0},
		{op: MAXSIZE, cap: 0},
	}

	var q *Deque = nil
	for i, c := range cases {
		switch c.op {
		case NEW:
			q = NewDeque(10)
			if q == nil {
				t.Fatalf("Case %d => Could not create the Deque", i)
			}
		case EMPTY:
			checkExpected(t, i, q.Empty(), c.empty)
		case SIZE:
			checkExpected(t, i, q.Size(), c.size)
		case MAXSIZE:
			checkExpected(t, i, q.MaxSize(), c.cap)
		case SHRINKTOFIT:
			q.ShrinkToFit()
		}
	}
}
