package list

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func testNewList(t *testing.T, list **List) {
	*list = NewList()
	if *list == nil {
		t.Fatalf("Could not allocate List")
	}
}

func testEmpty(t *testing.T, list *List) {
	if !list.Empty() {
		t.Fatalf("Expected empty list, received non-empty")
	}
}

func testNotEmpty(t *testing.T, list *List) {
	if list.Empty() {
		t.Fatalf("Expected non-empty list, received empty")
	}
}

func testSize(t *testing.T, list *List, expected int) {
	if list.Size() != expected {
		t.Fatalf("Expected size %d, received size %d", expected, list.Size())
	}
}

func testComparable(t *testing.T, v interface{}, expected interface{}) {
	if (v == nil && expected != nil) || (v != nil && expected == nil) {
		t.Fatalf("Expected %v, received %v", expected, v)
	} else {
		return
	}

	if v != expected {
		t.Fatalf("Expected %v, received %v", expected, v)
	}
}

func testFront(t *testing.T, list *List, expected interface{}) {
	v := list.Front()
	fmt.Printf("%v, %v", v, expected)
	testComparable(t, v, expected)
}

func testBack(t *testing.T, list *List, expected interface{}) {
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

	var list *List
	testNewList(t, &list)
	testFront(t, list, nil)
	testBack(t, list, nil)
	testEmpty(t, list)

	for _, n := range numbers {
		list.PushBack(n)
	}

	testFront(t, list, numbers[0])
	testBack(t, list, numbers[length-1])
	testNotEmpty(t, list)
	testSize(t, list, length)

	list.PopFront()
	testFront(t, list, numbers[1])

	list.PopBack()
	testBack(t, list, numbers[length-2])

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

	list.PushFront(numbers[0])
	testFront(t, list, numbers[0])
	testNotEmpty(t, list)

	list.Reverse()
	testNotEmpty(t, list)
	list.PopBack()
	testEmpty(t, list)
	list.Reverse()
	testEmpty(t, list)

	var list2 *List
	testNewList(t, &list2)
	for _, n := range numbers {
		list2.PushBack(n)
	}

	testNotEmpty(t, list2)
	list.Swap(list2)
	testNotEmpty(t, list)
	testEmpty(t, list2)
}
