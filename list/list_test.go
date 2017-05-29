package list

import (
	"fmt"
	"testing"

	"github.com/srirampatil/stl"
)

type IntComparable int

func (lhs IntComparable) Compare(r stl.Comparable) (int, error) {
	rhs, ok := r.(IntComparable)
	if !ok {
		return 0, stl.TypeMismatchError{&lhs, &rhs}
	}
	if lhs < rhs {
		return -1, nil
	} else if lhs > rhs {
		return 1, nil
	}
	return 0, nil
}

func TestList(t *testing.T) {
	list := NewDLL()
	if v := list.Front(); v != nil {
		t.Fatalf("Expected nil, received %v", v)
	}

	if v := list.Back(); v != nil {
		t.Fatalf("Expected nil, received %v", v)
	}

	if !list.Empty() {
		t.Fatalf("Expected empty list, received non-empty")
	}

	list.PushFront(IntComparable(20))
	list.PushBack(IntComparable(10))
	list.PushBack(IntComparable(30))
	list.PushFront(IntComparable(40))

	v := list.Front()
	if res, err := v.Compare(IntComparable(40)); res != 0 || err != nil {
		t.Fatalf("Expected 40, received %v", v)
	}

	v = list.Back()
	if res, err := v.Compare(IntComparable(30)); res != 0 || err != nil {
		t.Fatalf("Expected 30, received %v", v)
	}

	if list.Empty() {
		t.Fatalf("Expected non-empty list, received empty")
	}

	if size := list.Size(); size != 4 {
		t.Fatalf("Expected size 4, received %d", size)
	}

	list.PopFront()

	v = list.Front()
	if res, err := v.Compare(IntComparable(20)); res != 0 || err != nil {
		t.Fatalf("Expected 20, received %v", v)
	}

	list.PopBack()
	v = list.Back()
	if res, err := v.Compare(IntComparable(10)); res != 0 || err != nil {
		t.Fatalf("Expected 10, received %v", v)
	}

	for it := list.Begin(); it != nil; it = it.Next() {
		fmt.Println(it)
	}

	fmt.Println()
	list.Reverse()
	for it := list.Begin(); it != nil; it = it.Next() {
		fmt.Println(it)
	}

	if list.Empty() {
		t.Fatalf("Expected list, received empty")
	}

	list.PopBack()
	list.PopBack()

	list.PushBack(IntComparable(30))
	list.PopFront()
}
