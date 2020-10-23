package datastruct

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
