package list

import (
	"math/rand"
	"testing"
	"time"

	"github.com/srirampatil/stl"
)

type IntComparable int

func (lhs IntComparable) Compare(r stl.Comparable) (int, error) {
	if r == nil {
		return 1, nil
	}

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

var list *DLL

func testNewDLL(t *testing.T) {
	list = NewDLL()
	if list == nil {
		t.Fatalf("Could not allocate DLL")
	}
}

func testEmpty(t *testing.T) {
	if !list.Empty() {
		t.Fatalf("Expected empty list, received non-empty")
	}
}

func testNotEmpty(t *testing.T) {
	if list.Empty() {
		t.Fatalf("Expected non-empty list, received empty")
	}
}

func testSize(t *testing.T, expected int) {
	if list.Size() != expected {
		t.Fatalf("Expected size %d, received size %d", expected, list.Size())
	}
}

func testComparable(t *testing.T, v stl.Comparable, expected stl.Comparable) {
	if (v == nil && expected != nil) || (v != nil && expected == nil) {
		t.Fatalf("Expected %v, received %v", expected, v)
	} else {
		return
	}

	if res, err := v.Compare(expected); res != 0 || err != nil {
		t.Fatalf("Expected %v, received %v", expected, v)
	}
}

func testFront(t *testing.T, expected stl.Comparable) {
	v := list.Front()
	testComparable(t, v, expected)
}

func testBack(t *testing.T, expected stl.Comparable) {
	v := list.Back()
	testComparable(t, v, expected)
}

func TestList(t *testing.T) {
	var length int = 10
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, 10, 10)
	var direction bool = false
	for i := 0; i < length; i++ {
		if direction {
			numbers[i] = rand.Intn(100)
		} else {
			numbers[length-i-1] = rand.Intn(100)
		}
		direction = !direction
	}

	testNewDLL(t)
	testFront(t, nil)
	testBack(t, nil)
	testEmpty(t)

	for _, n := range numbers {
		list.PushBack(IntComparable(n))
	}

	testFront(t, IntComparable(numbers[0]))
	testBack(t, IntComparable(numbers[length-1]))
	testNotEmpty(t)
	testSize(t, length)

	list.PopFront()
	testFront(t, IntComparable(numbers[1]))

	list.PopBack()
	testBack(t, IntComparable(numbers[length-2]))

	/*
		for it := list.Begin(); it != nil; it = it.Next() {
			fmt.Println(it)
		}

		fmt.Println()
		list.Reverse()
		for it := list.Begin(); it != nil; it = it.Next() {
			fmt.Println(it)
		}
	*/

	testNotEmpty(t)
	for !list.Empty() {
		list.PopBack()
	}

	testEmpty(t)
	list.PopBack()
	list.PopFront()

	list.PushFront(IntComparable(numbers[0]))
	testFront(t, IntComparable(numbers[0]))
	testNotEmpty(t)

	list.Reverse()
	testNotEmpty(t)
	list.PopBack()
	testEmpty(t)
	list.Reverse()
	testEmpty(t)
}
