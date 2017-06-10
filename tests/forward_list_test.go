package tests

import (
	"fmt"
	"testing"

	"github.com/srirampatil/gostl/container"
)

var globalT *testing.T = nil
var list *container.ForwardList = nil

func testNewForwardList(list **container.ForwardList) {
	*list = container.NewForwardList()
	if *list == nil {
		globalT.Fatalf("Could not create ForwardList")
	}
}

func TestForwardList(t *testing.T) {
	fmt.Println("-----Running ForwardList Tests-----")
	globalT = t
	testNewForwardList(&list)

	var cases = []struct {
		op    OpType
		value interface{}
		list  []interface{}
		empty bool
		size  int
	}{
		{op: EMPTY, empty: true},
		{op: PUSH, value: 1, list: []interface{}{1}},
		{op: PUSH, value: 2, list: []interface{}{2, 1}},
		{op: PUSH, value: 3, list: []interface{}{3, 2, 1}},
		{op: PUSH, value: 4, list: []interface{}{4, 3, 2, 1}},
		{op: PUSH, value: 5, list: []interface{}{5, 4, 3, 2, 1}},
		{op: POP, list: []interface{}{4, 3, 2, 1}},
		{op: EMPTY, empty: false},
		{op: SIZE, size: 4},
		{op: FRONT, value: 4},
		{op: POP, list: []interface{}{3, 2, 1}},
		{op: REVERSE, list: []interface{}{1, 2, 3}},
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
			CheckIfEqual(t, i, Iterate(list.Begin(), list.End()), c.list)
		case POP:
			list.PopFront()
			CheckIfEqual(t, i, Iterate(list.Begin(), list.End()), c.list)
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
