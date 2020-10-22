package datastruct

import (
	"fmt"
	"strconv"
)

type doubleNode struct {
	val  int
	prev *doubleNode
	next *doubleNode
}

type doublylinked struct {
	head *doubleNode
	size int
}

// return a new instance of a doubly linked list
func NewDoublyLinked() *doublylinked {
	return &doublylinked{nil, 0}
}
