package vm

import "testing"

func TestVM_ICONST(t *testing.T) {
	program := []int{
		ICONST, 42,
		HALT,
	}

	vm := NewVM(program)
	vm.Run()
	if vm.stack.Peek().value != 42 {
		t.Fatalf("incorrect value on the stack. got=%d", vm.stack.Peek().value)
	}
}

func TestVM_PRINT(t *testing.T) {
	program := []int{
		ICONST, 42,
		PRINT,
		HALT,
	}
	vm := NewVM(program)
	vm.Run()

	if vm.stack.Size() != 0 {
		t.Fatalf("value is still present on stack. stack size is %d", vm.stack.Size())
	}
}

func TestVM_PRINT_TwoIntegers(t *testing.T) {
	program := []int{
		ICONST, 42,
		ICONST, 43,
		PRINT,
		HALT,
	}
	vm := NewVM(program)
	vm.Run()

	if vm.stack.Size() != 1 {
		t.Fatalf("incorrect size of the stack. stack size is %d but it should be 1", vm.stack.Size())
	}
}

func TestVM_IADD(t *testing.T) {
	program := []int{
		ICONST, 42,
		ICONST, 43,
		IADD,
		HALT,
	}
	vm := NewVM(program)
	vm.Run()

	if vm.stack.Size() != 1 {
		t.Fatalf("incorrect size of the stack. stack size is %d but it should be 1", vm.stack.Size())
	}

	if vm.stack.Peek().value != 85 {
		t.Fatalf("incorrect result. result is %d but it should be 85", vm.stack.Peek().value)
	}
}

func TestVM_ISUB(t *testing.T) {
	program := []int{
		ICONST, 2,
		ICONST, 4,
		ISUB,
		HALT,
	}
	vm := NewVM(program)
	vm.Run()

	if vm.stack.Size() != 1 {
		t.Fatalf("incorrect size of the stack. stack size is %d but it should be 1", vm.stack.Size())
	}

	if vm.stack.Peek().value != 2 {
		t.Fatalf("incorrect result. result is %d but it should be 2", vm.stack.Peek().value)
	}
}

func TestVM_IMUL(t *testing.T) {
	program := []int{
		ICONST, 2,
		ICONST, 4,
		IMUL,
		HALT,
	}
	vm := NewVM(program)
	vm.Run()

	if vm.stack.Size() != 1 {
		t.Fatalf("incorrect size of the stack. stack size is %d but it should be 1", vm.stack.Size())
	}

	if vm.stack.Peek().value != 8 {
		t.Fatalf("incorrect result. result is %d but it should be 8", vm.stack.Peek().value)
	}
}

func TestVM_ILT(t *testing.T) {
	program := []int{
		ICONST, 4,
		ICONST, 2,
		ILT,
		HALT,
	}
	vm := NewVM(program)
	vm.Run()

	if vm.stack.Size() != 1 {
		t.Fatalf("incorrect size of the stack. stack size is %d but it should be 1", vm.stack.Size())
	}

	if vm.stack.Peek().value != 1 {
		t.Fatalf("incorrect result. result is %d but it should be 1", vm.stack.Peek().value)
	}
}

func TestVM_ILT_NotLessThan(t *testing.T) {
	program := []int{
		ICONST, 4,
		ICONST, 7,
		ILT,
		HALT,
	}
	vm := NewVM(program)
	vm.Run()

	if vm.stack.Size() != 1 {
		t.Fatalf("incorrect size of the stack. stack size is %d but it should be 1", vm.stack.Size())
	}

	if vm.stack.Peek().value != 0 {
		t.Fatalf("incorrect result. result is %d but it should be 0", vm.stack.Peek().value)
	}
}

func TestVM_IEQ(t *testing.T) {
	program := []int{
		ICONST, 4,
		ICONST, 4,
		IEQ,
		HALT,
	}
	vm := NewVM(program)
	vm.Run()

	if vm.stack.Size() != 1 {
		t.Fatalf("incorrect size of the stack. stack size is %d but it should be 1", vm.stack.Size())
	}

	if vm.stack.Peek().value != 1 {
		t.Fatalf("incorrect result. result is %d but it should be 1", vm.stack.Peek().value)
	}
}

func TestVM_IEQ_NotEqual(t *testing.T) {
	program := []int{
		ICONST, 4,
		ICONST, 6,
		IEQ,
		HALT,
	}
	vm := NewVM(program)
	vm.Run()

	if vm.stack.Size() != 1 {
		t.Fatalf("incorrect size of the stack. stack size is %d but it should be 1", vm.stack.Size())
	}

	if vm.stack.Peek().value != 0 {
		t.Fatalf("incorrect result. result is %d but it should be 0", vm.stack.Peek().value)
	}
}

func TestVM_JMP(t *testing.T) {
	// Expect that stack size will be 2 since ICONST 7 and ICONST 8 will be skipped
	program := []int{
		ICONST, 4,
		ICONST, 6,
		JMP, 10,
		ICONST, 7,
		ICONST, 8,
		HALT,
	}
	vm := NewVM(program)
	vm.Run()

	if vm.stack.Size() != 2 {
		t.Fatalf("incorrect size of the stack. stack size is %d but it should be 2", vm.stack.Size())
	}
}

func TestVM_JMPT(t *testing.T) {
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
	vm := NewVM(program)
	vm.Run()

	if vm.stack.Size() != 0 {
		t.Fatalf("incorrect size of the stack. stack size is %d but it should be 0", vm.stack.Size())
	}
}

func TestVM_JMPT_NotTrue(t *testing.T) {
	program := []int{
		ICONST, 4,
		ICONST, 3,
		IEQ,
		JMPT, 11,
		ICONST, 7,
		ICONST, 8,
		HALT,
	}
	vm := NewVM(program)
	vm.Run()

	if vm.stack.Size() != 2 {
		t.Fatalf("incorrect size of the stack. stack size is %d but it should be 2", vm.stack.Size())
	}
}

func TestVM_JMPF(t *testing.T) {
	program := []int{
		ICONST, 4,
		ICONST, 3,
		IEQ,
		JMPF, 11,
		ICONST, 7,
		ICONST, 8,
		HALT,
	}
	vm := NewVM(program)
	vm.Run()

	if vm.stack.Size() != 0 {
		t.Fatalf("incorrect size of the stack. stack size is %d but it should be 0", vm.stack.Size())
	}
}

func TestVM_JMPF_NotFalse(t *testing.T) {
	program := []int{
		ICONST, 4,
		ICONST, 4,
		IEQ,
		JMPF, 11,
		ICONST, 7,
		ICONST, 8,
		HALT,
	}
	vm := NewVM(program)
	vm.Run()

	if vm.stack.Size() != 2 {
		t.Fatalf("incorrect size of the stack. stack size is %d but it should be 2", vm.stack.Size())
	}
}
