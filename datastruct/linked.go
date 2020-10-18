package datastruct

import (
	"fmt"
)

type node struct {
	val  int
	next *node
}

func NewNode(i int) *node {
	return &node{i, nil}
}

func (n *node) Display() {
	fmt.Println(n)
}
