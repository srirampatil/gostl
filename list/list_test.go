package list

import (
	"fmt"
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

func testNewDLL(t *testing.T, list **DLL) {
	*list = NewDLL()
	if *list == nil {
		t.Fatalf("Could not allocate DLL")
	}
}

func testEmpty(t *testing.T, list *DLL) {
	if !list.Empty() {
		t.Fatalf("Expected empty list, received non-empty")
	}
}

func testNotEmpty(t *testing.T, list *DLL) {
	if list.Empty() {
		t.Fatalf("Expected non-empty list, received empty")
	}
}

func testSize(t *testing.T, list *DLL, expected int) {
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

func testFront(t *testing.T, list *DLL, expected stl.Comparable) {
	v := list.Front()
	fmt.Printf("%v, %v", v, expected)
	testComparable(t, v, expected)
}

func testBack(t *testing.T, list *DLL, expected stl.Comparable) {
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

	var list *DLL
	testNewDLL(t, &list)
	testFront(t, list, nil)
	testBack(t, list, nil)
	testEmpty(t, list)

	for _, n := range numbers {
		list.PushBack(IntComparable(n))
	}

	testFront(t, list, IntComparable(numbers[0]))
	testBack(t, list, IntComparable(numbers[length-1]))
	testNotEmpty(t, list)
	testSize(t, list, length)

	list.PopFront()
	testFront(t, list, IntComparable(numbers[1]))

	list.PopBack()
	testBack(t, list, IntComparable(numbers[length-2]))

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

	testNotEmpty(t, list)
	list.Clear()
	testEmpty(t, list)
	list.PopBack()
	list.PopFront()

	list.PushFront(IntComparable(numbers[0]))
	testFront(t, list, IntComparable(numbers[0]))
	testNotEmpty(t, list)

	list.Reverse()
	testNotEmpty(t, list)
	list.PopBack()
	testEmpty(t, list)
	list.Reverse()
	testEmpty(t, list)

	var list2 *DLL
	testNewDLL(t, &list2)
	for _, n := range numbers {
		list2.PushBack(IntComparable(n))
	}

	testNotEmpty(t, list2)
	list.Swap(list2)
	testNotEmpty(t, list)
	testEmpty(t, list2)
}
