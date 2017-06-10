package tests

import (
	"fmt"
	"testing"

	"github.com/srirampatil/gostl/adaptor"
	"github.com/srirampatil/gostl/container"
)

func TestQueue(t *testing.T) {
	var cases = []struct {
		op    OpType
		value interface{}
	}{
		{op: NEW, value: container.ContainerTypeList},
		{op: EMPTY, value: true},
		{op: SIZE, value: 0},
		{op: PUSH, value: 1},
		{op: FRONT, value: 1},
		{op: BACK, value: 1},
		{op: SIZE, value: 1},
		{op: POP},
		{op: EMPTY, value: true},
		{op: SIZE, value: 0},
		{op: PUSH, value: 1},
		{op: PUSH, value: 2},
		{op: EMPTY, value: false},
		{op: SIZE, value: 2},
		{op: FRONT, value: 1},
		{op: BACK, value: 2},
		{op: POP},
		{op: FRONT, value: 2},
		{op: BACK, value: 2},
		{op: NEW, value: container.ContainerTypeDeque},
		{op: EMPTY, value: true},
		{op: SIZE, value: 0},
		{op: PUSH, value: 1},
		{op: FRONT, value: 1},
		{op: BACK, value: 1},
		{op: SIZE, value: 1},
		{op: POP},
		{op: EMPTY, value: true},
		{op: SIZE, value: 0},
		{op: PUSH, value: 1},
		{op: PUSH, value: 2},
		{op: EMPTY, value: false},
		{op: SIZE, value: 2},
		{op: FRONT, value: 1},
		{op: BACK, value: 2},
		{op: POP},
		{op: FRONT, value: 2},
		{op: BACK, value: 2},
	}

	var q *adaptor.Queue
	for i, c := range cases {
		switch c.op {
		case NEW:
			cType := c.value.(container.ContainerType)
			fmt.Printf("-----Running Queue tests with container %v-----\n", cType)
			q = adaptor.NewQueue(cType)
		case EMPTY:
			CheckExpected(t, i, q.Empty(), c.value)
		case SIZE:
			CheckExpected(t, i, q.Size(), c.value)
		case FRONT:
			CheckExpected(t, i, q.Front(), c.value)
		case BACK:
			CheckExpected(t, i, q.Back(), c.value)
		case PUSH:
			q.Push(c.value)
		case POP:
			q.Pop()
		}
	}
}
