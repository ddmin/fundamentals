package main

import (
	"fmt"
	"github.com/ddmin/fundamentals/datastruct"
	"math"
)

// to differentiate text
func titleize(s string) {
	bars := "======================================="
	fmt.Print("\n", bars, " ", s, " ", bars, "\n\n")
}

// to differentiate different data structures
func separate(s string) {
	bars := "###################################################################"
	fmt.Print("\n", bars, " ", s, " ", bars, "\n\n")
}

// test list functionality
func testList() {
	l := datastruct.NewList()
	separate("Init List")
	l.Display()

	titleize("Prepending")
	for i := 0; i <= 5; i++ {
		l.Prepend(i)
		l.Display()
	}

	titleize("Appending")
	for i := 0; i <= 5; i++ {
		l.Append(i)
		l.Display()
	}

	titleize("Removing")
	for i := 0; i <= 5; i++ {
		l.Remove(i)
		l.Display()
	}

	titleize("Inserting")
	for i := 0; i <= 5; i++ {
		l.Insert(i, 2*i)
		l.Display()
	}

	titleize("Reverse")
	l.Reverse()
	l.Display()

	titleize("Get elements")
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
	separate("Init Linked List")
	l.Display()

	titleize("Prepending")
	for i := 0; i <= 8; i++ {
		l.Prepend(i)
		l.Display()
	}

	titleize("Appending")
	for i := 1; i <= 4; i++ {
		l.Append(i)
		l.Display()
	}

	titleize("Inserting")
	for i := 2; i <= 8; i += 2 {
		l.Insert(i, i)
		l.Display()
	}

	titleize("Remove First")
	for i := 0; i < 3; i++ {
		l.RemoveFirst()
		l.Display()
	}

	titleize("Remove Last")
	for i := 0; i < 3; i++ {
		l.RemoveLast()
		l.Display()
	}

	titleize("Remove")
	for i := 0; i < 3; i++ {
		l.Remove(1)
		l.Display()
	}

	titleize("Reverse")
	a := datastruct.NewLinkedList()
	for n := 0; n < 10; n++ {
		a.Display()
		a.Reverse()
		a.Display()
		a.Append(n)
	}

	l.Reverse()
	l.Display()

	titleize("Get elements")
	for i := 0; i < 100; i++ {
		val, ok := l.Get(i)
		if ok {
			fmt.Print(val, " ")
		}
	}
	fmt.Println()

	titleize("Recursive")
	l.RecurseReverse()
	l.Display()

	titleize("Display reversed")
	l.DisplayReversed()
}

// test doubly linked list functionality
func testDoublyLinked() {
	d := datastruct.NewDoublyLinked()
	separate("Init Doubly Linked")
	d.Display()

	titleize("Prepending")
	for i := 0; i < 3; i++ {
		d.Prepend(-i)
		d.Display()
	}

	titleize("Appending")
	for i := 0; i < 3; i++ {
		d.Append(-i)
		d.Display()
	}

	titleize("Inserting")
	for i := 0; i < 10; i++ {
		d.Insert(2*i, i)
		d.Display()
	}

	titleize("Remove First")
	for i := 0; i < 4; i++ {
		d.RemoveFirst()
		d.Display()
	}

	titleize("Remove Last")
	for i := 0; i < 7; i++ {
		d.RemoveLast()
		d.Display()
	}

	titleize("Remove")
	d.Remove(0)
	d.Remove(4)
	d.Display()

	titleize("Reverse")
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

// stack implementation of bracket pair check
func checkBracketPairs(exp string) bool {
	s := datastruct.NewStack()
	for _, c := range exp {
		switch c {
		case 40, 91, 123:
			s.Push(int(c))
		case 41:
			if val, _ := s.Top(); val == 40 {
				s.Pop()
			} else {
				return false
			}
		case 93:
			if val, _ := s.Top(); val == 91 {
				s.Pop()
			} else {
				return false
			}
		case 125:
			if val, _ := s.Top(); val == 123 {
				s.Pop()
			} else {
				return false
			}
		}
	}
	return true
}

// test stack functionality
func testStack() {
	s := datastruct.NewStack()
	separate("Init Stack")
	titleize("Pushing")
	for i := 0; i <= 10; i++ {
		s.Push(i)
		n, ok := s.Top()
		if ok {
			fmt.Println("Top:", n)
		}
	}
	titleize("Popping")
	for i := 0; i < 12; i++ {
		s.Pop()
		n, ok := s.Top()
		if ok {
			fmt.Println("Top:", n)
		} else {
			fmt.Println("Stack is empty!")
		}
	}
	titleize("Reversing (using stacks)")
	for _, integer := range [...]int{1000, 12419, 28072, 5, 19, 10012} {
		fmt.Println(integer, "->", reverseInteger(integer))
	}
	titleize("Check for matching braces")
	function := "func return5() {int x:=5; return x}"
	gibberish := "(())))"
	gibberish2 := "{[something((interesting[]))]}0"
	fmt.Println(function, "->", checkBracketPairs(function))
	fmt.Println(gibberish, "->", checkBracketPairs(gibberish))
	fmt.Println(gibberish2, "->", checkBracketPairs(gibberish2))
}

func main() {
	// list-like data structures
	testList()
	testLinkedList()
	testDoublyLinked()

	testStack()
}
