package vm

import "strconv"

// IntStack is a representation of the stack of integers
type IntStack struct {
	values []int
	size   int
}

// NewIntStack returns an initialized stack of integers
func NewIntStack() *IntStack {
	return &IntStack{size: 0}
}

// Push adds new element on top of the stack
func (s *IntStack) Push(value int) {
	s.values = append(s.values[:s.size], value)
	s.size++
}

// Peek returns the top element in stack
func (s *IntStack) Peek() int {
	return s.values[s.size-1]
}

// Size returns the number of elements in the stack
func (s *IntStack) Size() int {
	return s.size
}

// Pop removes the top element and returns its value
func (s *IntStack) Pop() int {
	s.size--
	popValue := s.values[s.size]
	s.values = s.values[:s.size]
	return popValue
}

// AtIndex returns a value which is at provided index
func (s *IntStack) AtIndex(index int) int {
	return s.values[index]
}

func (s *IntStack) ToString() string {
	value := "\033[92m--------------- Stack ---------------\033[00m\n"
	for i := s.size - 1; i >= 0; i-- {
		value += "\033[93m" + strconv.Itoa(i) + "|\t" + strconv.Itoa(s.values[i]) + "\033[00m\t\n"
	}
	value += "\033[92m-------------------------------------\033[00m\n"
	return value
}
