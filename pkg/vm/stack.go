package vm

import (
	"fmt"
	"strconv"
)

// Stack is a representation of the stack data structure
type Stack[T any] struct {
	values []T
	size   int
}

// NewStack returns an initialized stack of integers
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{size: 0}
}

// Push adds new element on top of the stack
func (s *Stack[T]) Push(value T) {
	s.values = append(s.values[:s.size], value)
	s.size++
}

// Peek returns the top element in stack
func (s *Stack[T]) Peek() T {
	return s.values[s.size-1]
}

// Size returns the number of elements in the stack
func (s *Stack[T]) Size() int {
	return s.size
}

// Pop removes the top element and returns its value
func (s *Stack[T]) Pop() T {
	s.size--
	popValue := s.values[s.size]
	s.values = s.values[:s.size]
	return popValue
}

// AtIndex returns a value which is at provided index
func (s *Stack[T]) AtIndex(index int) T {
	return s.values[index]
}

func (s *Stack[T]) ToString() string {
	value := "\033[92m--------------- Stack ---------------\033[00m\n"
	for i := s.size - 1; i >= 0; i-- {
		value += "\033[93m" + strconv.Itoa(i) + "|\t" + fmt.Sprintf("%v", s.values[i]) + "\033[00m\t\n"
	}
	value += "\033[92m-------------------------------------\033[00m\n"
	return value
}
