package container

import (
	"fmt"
	"testing"

	"github.com/srirampatil/gostl/common"
)

func TestDeque(t *testing.T) {
	fmt.Println("-----Running Deque Tests-----")

	var cases = []struct {
		op     OpType
		empty  bool
		size   int
		idx    int
		value  interface{}
		values []interface{}
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
		{op: PUSHBACK, value: 1, values: []interface{}{1}},
		{op: PUSHBACK, value: 2, values: []interface{}{1, 2}},
		{op: PUSHBACK, value: 3, values: []interface{}{1, 2, 3}},
		{op: PUSHFRONT, value: 4, values: []interface{}{4, 1, 2, 3}},
		{op: SIZE, size: 4},
		{op: MAXSIZE, size: 4},
		{op: FRONT, value: 4},
		{op: BACK, value: 3},
		{op: AT, value: 1, idx: 1},
		{op: PUSHFRONT, value: 5, values: []interface{}{5, 4, 1, 2, 3}},
		{op: SIZE, size: 5},
		{op: MAXSIZE, size: 8},
		{op: AT, value: 2, idx: 3},
		{op: POPBACK, values: []interface{}{5, 4, 1, 2}},
		{op: SIZE, size: 4},
		{op: FRONT, value: 5},
		{op: BACK, value: 2},
		{op: POPFRONT, value: 5, values: []interface{}{4, 1, 2}},
		{op: SIZE, size: 3},
		{op: FRONT, value: 4},
		{op: BACK, value: 2},
		{op: ITERATE, values: []interface{}{4, 1, 2}},
		{op: REVERSE_ITERATE, values: []interface{}{2, 1, 4}},
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
			common.CheckIfEqual(t, i, common.Iterate(q.Begin(), q.End()), c.values)
		case PUSHFRONT:
			q.PushFront(c.value)
			common.CheckIfEqual(t, i, common.Iterate(q.Begin(), q.End()), c.values)
		case FRONT:
			common.CheckExpected(t, i, q.Front(), c.value)
		case BACK:
			common.CheckExpected(t, i, q.Back(), c.value)
		case AT:
			common.CheckExpected(t, i, q.At(c.idx), c.value)
		case POPBACK:
			q.PopBack()
			common.CheckIfEqual(t, i, common.Iterate(q.Begin(), q.End()), c.values)
		case POPFRONT:
			q.PopFront()
			common.CheckIfEqual(t, i, common.Iterate(q.Begin(), q.End()), c.values)
		case CLEAR:
			q.Clear()
		case ITERATE:
			common.CheckIfEqual(t, i, common.Iterate(q.Begin(), q.End()), c.values)
		case REVERSE_ITERATE:
			common.CheckIfEqual(t, i, common.Iterate(q.Rbegin(), q.Rend()), c.values)
		}
	}
}
