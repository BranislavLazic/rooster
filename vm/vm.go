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
			addResult := vm.stack.Pop().(int) + vm.stack.Pop().(int)
			vm.stack.Push(addResult)
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
