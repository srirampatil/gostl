package tests

import (
	"fmt"
	"testing"

	"github.com/srirampatil/gostl/adaptor"
	"github.com/srirampatil/gostl/container"
)

func TestStack(t *testing.T) {
	var cases = []struct {
		op    OpType
		value interface{}
	}{
		{op: NEW, value: container.NewForwardList()},
		{op: EMPTY, value: true},
		{op: SIZE, value: 0},
		{op: PUSH, value: 1},
		{op: TOP, value: 1},
		{op: POP},
		{op: EMPTY, value: true},
		{op: SIZE, value: 0},
		{op: PUSH, value: 1},
		{op: PUSH, value: 2},
		{op: TOP, value: 2},
		{op: POP},
		{op: TOP, value: 1},
		{op: PUSH, value: 3},
		{op: SIZE, value: 2},
		{op: EMPTY, value: false},
		{op: TOP, value: 3},
		{op: NEW, value: container.NewList()},
		{op: EMPTY, value: true},
		{op: SIZE, value: 0},
		{op: PUSH, value: 1},
		{op: TOP, value: 1},
		{op: POP},
		{op: EMPTY, value: true},
		{op: SIZE, value: 0},
		{op: PUSH, value: 1},
		{op: PUSH, value: 2},
		{op: TOP, value: 2},
		{op: POP},
		{op: TOP, value: 1},
		{op: PUSH, value: 3},
		{op: SIZE, value: 2},
		{op: EMPTY, value: false},
		{op: TOP, value: 3},
		{op: NEW, value: container.NewDeque(2)},
		{op: EMPTY, value: true},
		{op: SIZE, value: 0},
		{op: PUSH, value: 1},
		{op: TOP, value: 1},
		{op: POP},
		{op: EMPTY, value: true},
		{op: SIZE, value: 0},
		{op: PUSH, value: 1},
		{op: PUSH, value: 2},
		{op: TOP, value: 2},
		{op: POP},
		{op: TOP, value: 1},
		{op: PUSH, value: 3},
		{op: SIZE, value: 2},
		{op: EMPTY, value: false},
		{op: TOP, value: 3},
	}

	var s *adaptor.Stack
	for i, c := range cases {
		switch c.op {
		case NEW:
			fmt.Printf("-----Running Stack tests with container %T-----\n", c.value)
			s = adaptor.NewStack(c.value.(adaptor.StackContainer))
		case EMPTY:
			CheckExpected(t, i, s.Empty(), c.value)
		case SIZE:
			CheckExpected(t, i, s.Size(), c.value)
		case TOP:
			CheckExpected(t, i, s.Top(), c.value)
		case PUSH:
			s.Push(c.value)
		case POP:
			s.Pop()
		}
	}
}
