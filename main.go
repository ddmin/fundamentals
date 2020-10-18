package main

import (
	"github.com/ddmin/fundamentals/datastruct"
)

// test list
func testList() {
	l := datastruct.NewList()
	l.Display()

	for i := 0; i <= 10; i++ {
		l.Append(i)
		l.Display()
	}

	for i := 7; i <= 10; i++ {
		l.Remove(i)
		l.Display()
	}

	for i := 0; i <= 7; i++ {
		l.Insert(i, 2*i)
		l.Display()
	}
}

// test linked list
func testLinkedList() {
	l := datastruct.NewNode(1)
	l.Display()
}

func main() {
	// test list functionality
	testList()
	//
	testLinkedList()
}
