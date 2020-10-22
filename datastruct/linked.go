package datastruct

import (
	"fmt"
	"strconv"
)

type node struct {
	val  int
	next *node
}

type linkedlist struct {
	head *node
	size int
}

// return a new instance of the head node
func NewLinkedList() *linkedlist {
	return &linkedlist{nil, 0}
}

// display linked list
func (l *linkedlist) Display() {
	// empty list
	if l.head == nil {
		fmt.Println("[ ]")
		return
	}

	str := "[ "
	temp := l.head
	str += strconv.Itoa(temp.val) + " "

	// traverse list until end
	for temp.next != nil {
		temp = temp.next
		str += strconv.Itoa(temp.val) + " "
	}

	str += "]"
	fmt.Println(str)
}

// get ith item of list
func (l *linkedlist) Get(i int) (n int, ok bool) {
	if i < l.size {
		temp := l.head
		for count := 0; count < i; count++ {
			temp = temp.next
		}
		return temp.val, true
	} else {
		return 0, false
	}
}

// change first element in list
func (l *linkedlist) Prepend(i int) {
	// save original head to temp
	temp := l.head
	// set head to new node
	l.head = &node{i, temp}
	l.size++
}

// insert i at the end of the list
func (l *linkedlist) Append(i int) {
	// if list is empty, set head to point to first node
	if l.head == nil {
		l.head = &node{i, nil}
		return
	}

	temp := l.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = &node{i, nil}
	l.size++
}

// insert i at position n
func (l *linkedlist) Insert(i, n int) {
	if n == 0 {
		l.Prepend(i)
	} else if l.size == n {
		l.Append(i)
	} else if l.size > n {
		count := 0
		prev := l.head
		for count < n-1 {
			prev = prev.next
			count++
		}
		next := prev.next
		prev.next = &node{i, next}
		l.size++
	} else {
		return
	}
}

// remove head node
func (l *linkedlist) RemoveFirst() {
	temp := l.head
	l.head = temp.next
	l.size--
}

// remove tail node
func (l *linkedlist) RemoveLast() {
	temp := l.head
	for temp.next.next != nil {
		temp = temp.next
	}
	temp.next = nil
	l.size--
}

// remove node at position n
func (l *linkedlist) Remove(n int) {
	if n == 0 {
		l.RemoveFirst()
	} else if l.size == n-1 {
		l.RemoveLast()
	} else if l.size > n {
		count := 0
		temp := l.head
		for count < n-1 {
			temp = temp.next
			count++
		}
		next := temp.next.next
		temp.next = next
		l.size--
	} else {
		return
	}
}
