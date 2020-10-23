package main

import (
	"fmt"
	"github.com/ddmin/fundamentals/datastruct"
	"math"
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

	fmt.Println("Get elements")
	for i := 0; i < 100; i++ {
		val, ok := l.Get(i)
		if ok {
			fmt.Print(val, " ")
		}
	}
	fmt.Println()
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

	fmt.Println("Reverse")
	a := datastruct.NewLinkedList()
	for n := 0; n < 10; n++ {
		a.Display()
		a.Reverse()
		a.Display()
		a.Append(n)
	}

	l.Reverse()
	l.Display()

	fmt.Println("Get elements")
	for i := 0; i < 100; i++ {
		val, ok := l.Get(i)
		if ok {
			fmt.Print(val, " ")
		}
	}
	fmt.Println()

	fmt.Println("Recursive")
	l.RecurseReverse()
	l.Display()

	fmt.Println("Display reversed")
	l.DisplayReversed()
}

// test doubly linked list functionality
func testDoublyLinked() {
	d := datastruct.NewDoublyLinked()
	fmt.Println("Empty Doubly Linked List")
	d.Display()

	fmt.Println("Prepending")
	for i := 0; i < 3; i++ {
		d.Prepend(-i)
		d.Display()
	}

	fmt.Println("Appending")
	for i := 0; i < 3; i++ {
		d.Append(-i)
		d.Display()
	}

	fmt.Println("Inserting")
	for i := 0; i < 10; i++ {
		d.Insert(2*i, i)
		d.Display()
	}

	fmt.Println("Remove First")
	for i := 0; i < 4; i++ {
		d.RemoveFirst()
		d.Display()
	}

	fmt.Println("Remove Last")
	for i := 0; i < 7; i++ {
		d.RemoveLast()
		d.Display()
	}

	fmt.Println("Remove")
	d.Remove(0)
	d.Remove(4)
	d.Display()

	fmt.Println("Reverse")
	d.Reverse()
	d.Display()
}

// testing reversing using stack
func reverseInteger(n int) int {
	i := 0
	s := datastruct.NewStack()
	for int(math.Pow10(i)) <= n {
		i++
	}
	count := i - 1
	for i > 0 {
		item := n / int(math.Pow10(i-1))
		n -= item * int(math.Pow10(i-1))
		s.Push(item)
		i--
	}
	sum := 0
	for !s.IsEmpty() {
		val, _ := s.Top()
		sum += val * int(math.Pow10(count))
		s.Pop()
		count--
	}
	return sum
}

// test stack functionality
func testStack() {
	s := datastruct.NewStack()
	fmt.Println("Empty Stack")
	fmt.Println("Pushing")
	for i := 0; i <= 10; i++ {
		s.Push(i)
		n, ok := s.Top()
		if ok {
			fmt.Println("Top:", n)
		}
	}
	fmt.Println("Popping")
	for i := 0; i < 12; i++ {
		s.Pop()
		n, ok := s.Top()
		if ok {
			fmt.Println("Top:", n)
		} else {
			fmt.Println("Stack is empty!")
		}
	}
	fmt.Println("Reversing (using stacks)")
	for _, integer := range [...]int{1000, 12419, 28072, 5, 19, 10012} {
		fmt.Println(integer, "->", reverseInteger(integer))
	}
}

func main() {
	// list-like data structures
	testList()
	testLinkedList()
	testDoublyLinked()

	testStack()
}
