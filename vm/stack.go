package vm

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value interface{}
	next  *Element
}

// NewStack returns an initialized Stack object
func NewStack() *Stack {
	return new(Stack)
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return s.size
}

// Push adds new element on top of the stack
func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

// Peek returns the top element in stack
func (s *Stack) Peek() *Element {
	return s.top
}

// Pop removes the top element and returns its value
func (s *Stack) Pop() interface{} {
	element := s.top
	s.top = element.next
	s.size--
	return element.value
}
