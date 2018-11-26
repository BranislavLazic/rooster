package vm

import (
	"os"
	"testing"
)

func TestICONST(t *testing.T) {
	program := []int{
		ICONST, 42,
		HALT,
	}

	vm := NewVM(program, make(map[int]interface{}))
	vm.Run(os.Stdout)
	if vm.stack.Peek() != 42 {
		t.Fatalf("incorrect value on the stack. got=%d", vm.stack.Peek())
	}
}

func TestPRINT(t *testing.T) {
	program := []int{
		ICONST, 42,
		PRINT,
		HALT,
	}
	vm := NewVM(program, make(map[int]interface{}))
	vm.Run(os.Stdout)

	if vm.stack.Size() != 0 {
		t.Fatalf("value is still present on stack. got=%d", vm.stack.Size())
	}
}

func TestPRINT_TwoIntegers(t *testing.T) {
	program := []int{
		ICONST, 42,
		ICONST, 43,
		PRINT,
		HALT,
	}
	vm := NewVM(program, make(map[int]interface{}))
	vm.Run(os.Stdout)

	if vm.stack.Size() != 1 {
		t.Fatalf("incorrect size of the stack. got=%d but it should be 1", vm.stack.Size())
	}
}

func TestIADD(t *testing.T) {
	program := []int{
		ICONST, 42,
		ICONST, 43,
		IADD,
		HALT,
	}
	vm := NewVM(program, make(map[int]interface{}))
	vm.Run(os.Stdout)

	if vm.stack.Size() != 1 {
		t.Fatalf("incorrect size of the stack. got=%d but it should be 1", vm.stack.Size())
	}

	if vm.stack.Peek() != 85 {
		t.Fatalf("incorrect result of addition. got=%d but it should be 85", vm.stack.Peek())
	}
}

func TestISUB(t *testing.T) {
	program := []int{
		ICONST, 2,
		ICONST, 4,
		ISUB,
		HALT,
	}
	vm := NewVM(program, make(map[int]interface{}))
	vm.Run(os.Stdout)

	if vm.stack.Size() != 1 {
		t.Fatalf("incorrect size of the stack. got=%d but it should be 1", vm.stack.Size())
	}

	if vm.stack.Peek() != 2 {
		t.Fatalf("incorrect result of substraction. got=%d but it should be 2", vm.stack.Peek())
	}
}

func TestIMUL(t *testing.T) {
	program := []int{
		ICONST, 2,
		ICONST, 4,
		IMUL,
		HALT,
	}
	vm := NewVM(program, make(map[int]interface{}))
	vm.Run(os.Stdout)

	if vm.stack.Size() != 1 {
		t.Fatalf("incorrect size of the stack. got=%d but it should be 1", vm.stack.Size())
	}

	if vm.stack.Peek() != 8 {
		t.Fatalf("incorrect result of multiplication. got=%d but it should be 8", vm.stack.Peek())
	}
}

func TestIDIV(t *testing.T) {
	program := []int{
		ICONST, 2,
		ICONST, 4,
		IDIV,
		HALT,
	}
	vm := NewVM(program, make(map[int]interface{}))
	vm.Run(os.Stdout)

	if vm.stack.Size() != 1 {
		t.Fatalf("incorrect size of the stack. got=%d but it should be 1", vm.stack.Size())
	}

	if vm.stack.Peek() != 2 {
		t.Fatalf("incorrect result of multiplication. got=%d but it should be 2", vm.stack.Peek())
	}
}

func TestILT(t *testing.T) {
	program := []int{
		ICONST, 4,
		ICONST, 2,
		ILT,
		HALT,
	}
	vm := NewVM(program, make(map[int]interface{}))
	vm.Run(os.Stdout)

	if vm.stack.Size() != 1 {
		t.Fatalf("incorrect size of the stack. got=%d but it should be 1", vm.stack.Size())
	}

	if vm.stack.Peek() != 1 {
		t.Fatalf("incorrect result. got=%d but it should be 1", vm.stack.Peek())
	}
}

func TestILT_NotLessThan(t *testing.T) {
	program := []int{
		ICONST, 4,
		ICONST, 7,
		ILT,
		HALT,
	}
	vm := NewVM(program, make(map[int]interface{}))
	vm.Run(os.Stdout)

	if vm.stack.Size() != 1 {
		t.Fatalf("incorrect size of the stack. got=%d but it should be 1", vm.stack.Size())
	}

	if vm.stack.Peek() != 0 {
		t.Fatalf("incorrect result. got=%d but it should be 0", vm.stack.Peek())
	}
}

func TestIEQ(t *testing.T) {
	program := []int{
		ICONST, 4,
		ICONST, 4,
		IEQ,
		HALT,
	}
	vm := NewVM(program, make(map[int]interface{}))
	vm.Run(os.Stdout)

	if vm.stack.Size() != 1 {
		t.Fatalf("incorrect size of the stack. got=%d but it should be 1", vm.stack.Size())
	}

	if vm.stack.Peek() != 1 {
		t.Fatalf("incorrect result. got=%d but it should be 1", vm.stack.Peek())
	}
}

func TestIEQ_NotEqual(t *testing.T) {
	program := []int{
		ICONST, 4,
		ICONST, 6,
		IEQ,
		HALT,
	}
	vm := NewVM(program, make(map[int]interface{}))
	vm.Run(os.Stdout)

	if vm.stack.Size() != 1 {
		t.Fatalf("incorrect size of the stack. got=%d but it should be 1", vm.stack.Size())
	}

	if vm.stack.Peek() != 0 {
		t.Fatalf("incorrect result. got=%d but it should be 0", vm.stack.Peek())
	}
}

func TestJMP(t *testing.T) {
	// Expect that stack size will be 2 since ICONST 7 and ICONST 8 will be skipped
	program := []int{
		ICONST, 4,
		ICONST, 6,
		JMP, 10,
		ICONST, 7,
		ICONST, 8,
		HALT,
	}
	vm := NewVM(program, make(map[int]interface{}))
	vm.Run(os.Stdout)

	if vm.stack.Size() != 2 {
		t.Fatalf("incorrect size of the stack. got=%d but it should be 2", vm.stack.Size())
	}
}

func TestJMPT(t *testing.T) {
	// Expect that stack size will be 0 since the ICONST 7 and ICONST 8 will be skipped
	program := []int{
		ICONST, 4,
		ICONST, 4,
		IEQ,
		JMPT, 11,
		ICONST, 7,
		ICONST, 8,
		HALT,
	}
	vm := NewVM(program, make(map[int]interface{}))
	vm.Run(os.Stdout)

	if vm.stack.Size() != 0 {
		t.Fatalf("incorrect size of the stack. got=%d but it should be 0", vm.stack.Size())
	}
}

func TestJMPT_NotTrue(t *testing.T) {
	program := []int{
		ICONST, 4,
		ICONST, 3,
		IEQ,
		JMPT, 11,
		ICONST, 7,
		ICONST, 8,
		HALT,
	}
	vm := NewVM(program, make(map[int]interface{}))
	vm.Run(os.Stdout)

	if vm.stack.Size() != 2 {
		t.Fatalf("incorrect size of the stack. got=%d but it should be 2", vm.stack.Size())
	}
}

func TestJMPF(t *testing.T) {
	program := []int{
		ICONST, 4,
		ICONST, 3,
		IEQ,
		JMPF, 11,
		ICONST, 7,
		ICONST, 8,
		HALT,
	}
	vm := NewVM(program, make(map[int]interface{}))
	vm.Run(os.Stdout)

	if vm.stack.Size() != 0 {
		t.Fatalf("incorrect size of the stack. got=%d but it should be 0", vm.stack.Size())
	}
}

func TestJMPF_NotFalse(t *testing.T) {
	program := []int{
		ICONST, 4,
		ICONST, 4,
		IEQ,
		JMPF, 11,
		ICONST, 7,
		ICONST, 8,
		HALT,
	}
	vm := NewVM(program, make(map[int]interface{}))
	vm.Run(os.Stdout)

	if vm.stack.Size() != 2 {
		t.Fatalf("incorrect size of the stack. got=%d but it should be 2", vm.stack.Size())
	}
}

func TestCOPY(t *testing.T) {
	program := []int{
		ICONST, 42,
		COPY,
		PRINT,
		HALT,
	}
	vm := NewVM(program, make(map[int]interface{}))
	vm.Run(os.Stdout)

	if vm.stack.Size() != 1 {
		t.Fatalf("incorrect size of the stack. got=%d but it should be 1", vm.stack.Size())
	}
}

func TestGSTORE(t *testing.T) {
	// Size of the globals space should be 1 since value at 0 address is set for both 42 and 43
	program := []int{
		ICONST, 42,
		GSTORE, 0,
		ICONST, 43,
		GSTORE, 0,
		HALT,
	}
	vm := NewVM(program, make(map[int]interface{}))
	vm.Run(os.Stdout)

	if vm.stack.Size() != 0 {
		t.Fatalf("incorrect size of the stack. got=%d but it should be 0", vm.stack.Size())
	}

	if len(vm.globals) != 1 {
		t.Fatalf("incorrect size of globals space. got=%d but is should be 1", len(vm.globals))
	}

	if vm.globals[0] != 43 {
		t.Fatalf("incorrect value at 0 address. got=%d but it should be 43", vm.globals[0])
	}
}

func TestGLOAD(t *testing.T) {
	program := []int{
		ICONST, 42,
		GSTORE, 0,
		GLOAD, 0,
		HALT,
	}
	vm := NewVM(program, make(map[int]interface{}))
	vm.Run(os.Stdout)

	if vm.stack.Size() != 1 {
		t.Fatalf("incorrect size of the stack. got=%d but it should be 1", vm.stack.Size())
	}

	if vm.stack.Peek() != 42 {
		t.Fatalf("incorrect value on stack. got=%d but it should be 42", vm.stack.Peek())
	}
}

func TestCALL(t *testing.T) {
	// First procedure will load 50 and 43 and add them,
	// then the second procedure will be called from the body
	// of the first procedure which will load 42 and just print.
	program := []int{
		ICONST, 42,
		ICONST, 43,
		ICONST, 50,
		CALL, 10, 2,
		HALT,
		// First procedure
		LOAD, 0,
		LOAD, 1,
		IADD,
		PRINT,
		CALL, 20, 1,
		RET,
		// Second procedure
		LOAD, 0,
		PRINT,
		RET,
	}
	vm := NewVM(program, make(map[int]interface{}))
	vm.Run(os.Stdout)

	if vm.frameStack.Size() != 0 {
		t.Fatalf("incorrect size of the frame stack. got=%d but it should be 0", vm.frameStack.Size())
	}

	if vm.stack.Size() != 0 {
		t.Fatalf("incorrect size of the stack. got=%d but it should be 0", vm.stack.Size())
	}
}

func TestCALL_with_STORE(t *testing.T) {
	program := []int{
		ICONST, 42,
		CALL, 6, 0,
		HALT,
		// First procedure
		STORE, 0,
		LOAD, 0,
		PRINT,
		RET,
	}
	vm := NewVM(program, make(map[int]interface{}))
	vm.Run(os.Stdout)
	if vm.frameStack.Size() != 0 {
		t.Fatalf("incorrect size of the frame stack. got=%d but it should be 0", vm.frameStack.Size())
	}

	if vm.stack.Size() != 0 {
		t.Fatalf("incorrect size of the stack. got=%d but it should be 0", vm.stack.Size())
	}
}

func TestPRINTC(t *testing.T) {
	program := []int{
		SCONST, 1,
		PRINTC,
		HALT,
	}
	constantPool := make(map[int]interface{})
	constantPool[1] = "Hello world!"
	vm := NewVM(program, constantPool)
	vm.Run(os.Stdout)

	if vm.stack.Size() != 0 {
		t.Fatalf("value is still present on stack. got=%d", vm.stack.Size())
	}
}
