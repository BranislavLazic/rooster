package vm

import (
	"fmt"
)

type VM struct {
	stack              *Stack
	instructionPointer int
	program            []int
}

// NewVM initializes the virtual machine
func NewVM(program []int) *VM {
	return &VM{stack: NewStack(), instructionPointer: -1, program: program}
}

// Run runs the virtual machine and interprets its
// program by executing instruction after instruction
func (vm *VM) Run() {

	for vm.instructionPointer < len(vm.program) {
		opcode := vm.fetch()
		switch opcode {
		case ICONST:
			value := vm.fetch()
			vm.stack.Push(value)
			break
		case IADD:
			addResult := vm.stack.Pop() + vm.stack.Pop()
			vm.stack.Push(addResult)
			break
		case ISUB:
			subResult := vm.stack.Pop() - vm.stack.Pop()
			vm.stack.Push(subResult)
			break
		case IMUL:
			mulResult := vm.stack.Pop() * vm.stack.Pop()
			vm.stack.Push(mulResult)
			break
		// "Jump" to instruction pointer unconditionally. E.g. JMP 2 sets the instruction
		// pointer to 2 and starts executing program from that point.
		case JMP:
			jump := vm.fetch()
			vm.instructionPointer = jump - 1
			break
		// "Jump" to instruction pointer conditionally. E.g. JMPT 2 will "pop" the value from
		// stack and check if it's 1. If the condition is satisfied, then the instruction pointer
		// will be set to 2. Otherwise, it will just continue with normal code execution.
		case JMPT:
			jump := vm.fetch()
			if vm.stack.Pop() == 1 {
				vm.instructionPointer = jump - 1
			}
			break
		case JMPF:
			jump := vm.fetch()
			if vm.stack.Pop() == 0 {
				vm.instructionPointer = jump - 1
			}
			break
		case ILT:
			if vm.stack.Pop() < vm.stack.Pop() {
				vm.stack.Push(1)
			} else {
				vm.stack.Push(0)
			}
			break
		case IEQ:
			if vm.stack.Pop() == vm.stack.Pop() {
				vm.stack.Push(1)
			} else {
				vm.stack.Push(0)
			}
			break
		case PRINT:
			fmt.Println(vm.stack.Pop())
			break
		case HALT:
			return
		default:
			break
		}
	}

}

func (vm *VM) fetch() int {
	vm.instructionPointer++
	return vm.program[vm.instructionPointer]
}
