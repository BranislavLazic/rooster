package vm

import (
	"testing"
)

func TestPush(t *testing.T) {
	stack := createStackWithElements()

	if stack.Size() != 3 {
		t.Fatalf("stack.Push does not contain 3 elements, got=%d", stack.Size())
	}
}

func TestPeek(t *testing.T) {
	stack := createStackWithElements()

	peekResult := *stack.Peek()

	if peekResult.value != 3 {
		t.Fatalf("stack.Peek did not returned expected element. got=%d", peekResult.value)
	}
}

func TestPop(t *testing.T) {
	stack := createStackWithElements()

	popResult := stack.Pop()

	if popResult != 3 {
		t.Fatalf("stack.Pop did not return an expected value. got=%d", popResult)
	}

	if stack.Size() > 2 {
		t.Fatalf("stack.Pop did not remove the element.")
	}
}

func createStackWithElements() *Stack {
	stack := NewStack(100)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	return stack
}
