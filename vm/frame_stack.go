package vm

type Frame struct {
	variables     map[int]int
	returnAddress int
}

// FrameStack is a representation of FrameStack indexed collection
type FrameStack struct {
	values []Frame
	size   int
}

// NewFrameStack returns an initialized FrameStack object
func NewFrameStack() *FrameStack {
	return &FrameStack{size: 0}
}

// Push adds new element on top of the FrameStack
func (s *FrameStack) Push(value *Frame) {
	s.values = append(s.values[:s.size], *value)
	s.size++
}

// Peek returns the top element in FrameStack
func (s *FrameStack) Peek() *Frame {
	return &s.values[0]
}

// Size returns the number of elements in the FrameStack
func (s *FrameStack) Size() int {
	return s.size
}

// Pop removes the top element and returns its value
func (s *FrameStack) Pop() *Frame {
	s.size--
	popValue := s.values[s.size]
	s.values = s.values[:s.size]
	return &popValue
}

// AtIndex returns a value which is at provided index
func (s *FrameStack) AtIndex(index int) *Frame {
	return &s.values[index]
}
