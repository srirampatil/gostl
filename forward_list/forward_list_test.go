package forward_list

import (
	"testing"

	"github.com/srirampatil/gostl/common"
)

var globalT *testing.T = nil
var list *ForwardList = nil

func testNewForwardList(list **ForwardList) {
	*list = NewForwardList()
	if *list == nil {
		globalT.Fatalf("Could not create ForwardList")
	}
}

type OpType uint8

const (
	PUSH OpType = iota + 1
	POP
	EMPTY
	SIZE
	REVERSE
	FRONT
	CLEAR
)

func TestForwardList(t *testing.T) {
	globalT = t
	testNewForwardList(&list)

	var cases = []struct {
		op    OpType
		value interface{}
		list  []int
		empty bool
		size  int
	}{
		{op: EMPTY, empty: true},
		{op: PUSH, value: 1, list: []int{1}},
		{op: PUSH, value: 2, list: []int{2, 1}},
		{op: PUSH, value: 3, list: []int{3, 2, 1}},
		{op: PUSH, value: 4, list: []int{4, 3, 2, 1}},
		{op: PUSH, value: 5, list: []int{5, 4, 3, 2, 1}},
		{op: POP, list: []int{4, 3, 2, 1}},
		{op: EMPTY, empty: false},
		{op: SIZE, size: 4},
		{op: FRONT, value: 4},
		{op: POP, list: []int{3, 2, 1}},
		{op: REVERSE, list: []int{1, 2, 3}},
		{op: FRONT, value: 1},
		{op: CLEAR},
		{op: FRONT, value: nil},
		{op: EMPTY, empty: true},
	}

	for i, c := range cases {
		switch c.op {
		case EMPTY:
			if list.Empty() != c.empty {
				t.Fatalf("Expected: %v, Received: %v", c.empty, list.Empty())
			}
		case SIZE:
			if list.Size() != c.size {
				t.Fatalf("Expected: %d, Received: %d", c.size, list.Size())
			}
		case PUSH:
			list.PushFront(c.value)
			common.CheckIfEqual(t, i, common.Iterate(list.Begin(), list.End()), c.list)
		case POP:
			list.PopFront()
			common.CheckIfEqual(t, i, common.Iterate(list.Begin(), list.End()), c.list)
		case FRONT:
			v := list.Front()
			if v != c.value {
				t.Fatalf("Expected: %v, Received: %v", c.value, v)
			}
		case CLEAR:
			list.Clear()
		case REVERSE:
			list.Reverse()
		}
	}
}
