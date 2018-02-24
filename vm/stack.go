package vm

import (
	"log"
)

type Stack struct {
	top     *Element
	maxSize int
	size    int
}

type Element struct {
	value int
	next  *Element
}

// NewStack returns an initialized Stack object
func NewStack(maxSize int) *Stack {
	return &Stack{maxSize: maxSize, size: 0}
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return s.size
}

// Push adds new element on top of the stack
func (s *Stack) Push(value int) {
	s.size++
	if s.maxSize < s.size {
		log.Fatalf("stack overflow")
	} else {
		s.top = &Element{value, s.top}
	}
}

// Peek returns the top element in stack
func (s *Stack) Peek() *Element {
	return s.top
}

// Pop removes the top element and returns its value
func (s *Stack) Pop() int {
	element := s.top
	s.top = element.next
	s.size--
	return element.value
}
