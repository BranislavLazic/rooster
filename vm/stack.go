package vm

type Stack struct {
	values []int
	size   int
}

// NewStack returns an initialized Stack object
func NewStack() *Stack {
	return &Stack{size: 0}
}

// Push adds new element on top of the stack
func (s *Stack) Push(value int) {
	s.values = append(s.values[:s.size], value)
	s.size++
}

// Peek returns the top element in stack
func (s *Stack) Peek() int {
	return s.values[s.size-1]
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return s.size
}

// Pop removes the top element and returns its value
func (s *Stack) Pop() int {
	s.size--
	return s.values[s.size]
}
