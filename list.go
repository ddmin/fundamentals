package main

import (
	"fmt"
	"strconv"
)

type list struct {
	// maximum capacity of array
	max int
	// end of current list (number of items in list)
	end int
	// array slice representing items in list
	arr []int
}

// create new list instance
func newList() *list {
	return &list{2, 0, make([]int, 2)}
}

// display list
func (l *list) display() {
	str := "["
	for i := 0; i < l.end; i++ {
		str += strconv.Itoa(l.arr[i])
		if i != l.end-1 {
			str += " "
		}
	}
	str += "]"
	fmt.Println(str)
}

// adjust max size of list
// copy list into longer array slice
func (l *list) adjustLength() {
	if l.max <= l.end {
		// create array slice with double capacity
		l.max *= 2
		temp := make([]int, l.max)
		// copy elements
		for c := 0; c < l.end; c++ {
			temp[c] = l.arr[c]
		}
		// point to new array slice
		l.arr = temp
	}
}

// insert i at end of list
func (l *list) append(i int) {
	// check for room
	l.adjustLength()
	// add to end of array slice
	l.arr[l.end] = i
	// increment end of array by 1
	l.end += 1
}

// insert i at position n
func (l *list) insert(i, n int) {
	// check for room
	l.adjustLength()
	// shift every element >= n up 1
	for c := l.end - 1; c >= n; c-- {
		l.arr[c+1] = l.arr[c]
	}
	// insert item
	l.arr[n] = i
	// increment end of array by 1
	l.end += 1
}

// remove element at position n
func (l *list) remove(n int) {
	// shift every element > n down 1
	for c := n + 1; c < l.end; c++ {
		l.arr[c-1] = l.arr[c]
	}
	// decrement end of array by 1
	l.end -= 1
}

func main() {
	l := newList()
	l.display()

	for i := 0; i <= 10; i++ {
		l.append(i)
		l.display()
	}

	for i := 7; i <= 10; i++ {
		l.remove(i)
		l.display()
	}

	for i := 0; i <= 7; i++ {
		l.insert(i, 2*i)
		l.display()
	}
}
