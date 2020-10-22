package main

import (
	"fmt"
	"github.com/ddmin/fundamentals/datastruct"
)

// test list functionality
func testList() {
	l := datastruct.NewList()
	fmt.Println("Empty List")
	l.Display()

	fmt.Println("Prepending")
	for i := 0; i <= 5; i++ {
		l.Prepend(i)
		l.Display()
	}

	fmt.Println("Appending")
	for i := 0; i <= 5; i++ {
		l.Append(i)
		l.Display()
	}

	fmt.Println("Removing")
	for i := 0; i <= 5; i++ {
		l.Remove(i)
		l.Display()
	}

	fmt.Println("Inserting")
	for i := 0; i <= 5; i++ {
		l.Insert(i, 2*i)
		l.Display()
	}

	fmt.Println("Reverse")
	l.Reverse()
	l.Display()
}

// test linked list functionality
func testLinkedList() {
	l := datastruct.NewLinkedList()
	fmt.Println("Empty Linked List")
	l.Display()

	fmt.Println("Prepending")
	for i := 0; i <= 8; i++ {
		l.Prepend(i)
		l.Display()
	}

	fmt.Println("Appending")
	for i := 1; i <= 4; i++ {
		l.Append(i)
		l.Display()
	}

	fmt.Println("Inserting")
	for i := 2; i <= 8; i += 2 {
		l.Insert(i, i)
		l.Display()
	}

	fmt.Println("Remove First")
	for i := 0; i < 3; i++ {
		l.RemoveFirst()
		l.Display()
	}

	fmt.Println("Remove Last")
	for i := 0; i < 3; i++ {
		l.RemoveLast()
		l.Display()
	}

	fmt.Println("Remove")
	for i := 0; i < 3; i++ {
		l.Remove(1)
		l.Display()
	}

	fmt.Println("Get elements")
	for i := 0; i < 100; i++ {
		val, ok := l.Get(i)
		if ok {
			fmt.Print(val, " ")
		}
	}
	fmt.Println()
}

func main() {
	testList()
	// testLinkedList()
}
