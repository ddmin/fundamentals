package datastruct

import (
	"math"
)

type item struct {
	val   int
	below *item
}

type stack struct {
	top *item
}

// return new instance of stack
func NewStack() *stack {
	return &stack{nil}
}

// return top of stack
func (s *stack) Top() (i int, ok bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		return s.top.val, true
	}
}

// check if stack is empty
func (s *stack) IsEmpty() bool {
	return s.top == nil
}

// push i on top of stack
func (s *stack) Push(i int) {
	s.top = &item{i, s.top}
}

// pop the top off of stack
func (s *stack) Pop() {
	if s.IsEmpty() {
		return
	}
	s.top = s.top.below
}

// not struct methods, but related to structs

// testing reversing using stack
func ReverseInteger(n int) int {
	i := 0
	s := NewStack()
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
func CheckBracketPairs(exp string) bool {
	s := NewStack()
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
