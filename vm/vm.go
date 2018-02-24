package vm

import (
	"fmt"
)

type VM struct {
	stack              *Stack
	instructionPointer int
	program            []int
}

func NewVM(program []int) *VM {
	return &VM{stack: NewStack(), instructionPointer: -1, program: program}
}

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
		case JMP:
			jump := vm.fetch()
			vm.instructionPointer = jump - 1
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
