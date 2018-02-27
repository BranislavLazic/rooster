package vm

import "strconv"

// Stack is a representation of stack indexed collection
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
	popValue := s.values[s.size]
	s.values = s.values[:s.size]
	return popValue
}

// AtIndex returns a value which is at provided index
func (s *Stack) AtIndex(index int) int {
	return s.values[index]
}

func (s *Stack) ToString() string {
	value := "Stack:\n"
	for i := s.size - 1; i >= 0; i-- {
		value += "| " + strconv.Itoa(i) + "\t|\t" + strconv.Itoa(s.values[i]) + " |\n"
	}
	return value
}
