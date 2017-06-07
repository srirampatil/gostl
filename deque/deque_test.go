package deque

import (
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
)

func TestDeque(t *testing.T) {
	var cases = []struct {
		op     OpType
		empty  bool
		size   int
		idx    int
		value  interface{}
		values []int
	}{
		{op: NEW},
		{op: EMPTY, empty: true},
		{op: SIZE, size: 0},
		{op: MAXSIZE, size: 10},
		{op: SHRINKTOFIT},
		{op: EMPTY, empty: true},
		{op: SIZE, size: 0},
		{op: MAXSIZE, size: 0},
		{op: FRONT, value: nil},
		{op: BACK, value: nil},
		{op: PUSHBACK, value: 1, values: []int{1}},
		{op: PUSHBACK, value: 2, values: []int{1, 2}},
		{op: PUSHBACK, value: 3, values: []int{1, 2, 3}},
		{op: PUSHFRONT, value: 4, values: []int{4, 1, 2, 3}},
		{op: SIZE, size: 4},
		{op: MAXSIZE, size: 4},
		{op: FRONT, value: 4},
		{op: BACK, value: 3},
		{op: AT, value: 1, idx: 1},
		{op: PUSHFRONT, value: 5, values: []int{5, 4, 1, 2, 3}},
		{op: SIZE, size: 5},
		{op: MAXSIZE, size: 8},
		{op: AT, value: 2, idx: 3},
		{op: POPBACK, values: []int{5, 4, 1, 2}},
		{op: SIZE, size: 4},
		{op: FRONT, value: 5},
		{op: BACK, value: 2},
		{op: POPFRONT, value: 5, values: []int{4, 1, 2}},
		{op: SIZE, size: 3},
		{op: FRONT, value: 4},
		{op: BACK, value: 2},
		{op: CLEAR},
		{op: EMPTY, empty: true},
		{op: AT, value: nil, idx: 100},
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
			common.CheckExpected(t, i, q.Empty(), c.empty)
		case SIZE:
			common.CheckExpected(t, i, q.Size(), c.size)
		case MAXSIZE:
			common.CheckExpected(t, i, q.MaxSize(), c.size)
		case SHRINKTOFIT:
			q.ShrinkToFit()
		case PUSHBACK:
			q.PushBack(c.value)
		case PUSHFRONT:
			q.PushFront(c.value)
		case FRONT:
			common.CheckExpected(t, i, q.Front(), c.value)
		case BACK:
			common.CheckExpected(t, i, q.Back(), c.value)
		case AT:
			common.CheckExpected(t, i, q.At(c.idx), c.value)
		case POPBACK:
			q.PopBack()
		case POPFRONT:
			q.PopFront()
		case CLEAR:
			q.Clear()
		}
	}
}
