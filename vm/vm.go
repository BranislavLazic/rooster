package vm

import (
	"fmt"
)

// VM contains properties of CPU
type VM struct {
	stack              *Stack
	instructionPointer int
	program            []int
	globals            map[int]int
	frameStack         *FrameStack
	flags              map[string]interface{}
}

// NewVM initializes the virtual machine
func NewVM(program []int) *VM {
	return &VM{
		stack:              NewStack(),
		instructionPointer: -1,
		program:            program,
		globals:            make(map[int]int),
		frameStack:         NewFrameStack(),
		flags: map[string]interface{}{
			"printStack": false,
		},
	}
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
		case GLOAD:
			address := vm.fetch()
			globalValue := vm.globals[address]
			vm.stack.Push(globalValue)
			break
		case GSTORE:
			value := vm.stack.Pop()
			address := vm.fetch()
			vm.globals[address] = value
			break
		// Pop the value from the stack and store it in locals
		case STORE:
			value := vm.stack.Pop()
			address := vm.fetch()
			vm.frameStack.Peek().variables[address] = value
			break
		// Load value from locals
		case LOAD:
			address := vm.fetch()
			frame := vm.frameStack.Peek()
			vm.stack.Push(frame.variables[address])
			break
		// After CALL instruction has been processed, the next instruction is
		// set as an address of the instruction pointer (to perform "jump"),
		// then the instruction after that is being used to designate how many
		// values to "pop" from the stack (number of arguments). Values are "popped"
		// and stored in the local memory of procedures frame
		case CALL:
			jump := vm.fetch()
			argsToLoad := vm.fetch()
			frame := &Frame{returnAddress: vm.instructionPointer, variables: make(map[int]int)}
			for i := 0; i < argsToLoad; i++ {
				frame.variables[i] = vm.stack.Pop()
			}
			vm.frameStack.Push(frame)
			vm.instructionPointer = jump - 1
			break
		case RET:
			returnAddress := vm.frameStack.Peek().returnAddress
			vm.frameStack.Pop()
			vm.instructionPointer = returnAddress
		case PRINT:
			fmt.Println(vm.stack.Pop())
			break
		case HALT:
			return
		default:
			break
		}

		if vm.flags["printStack"].(bool) {
			fmt.Println(vm.stack.ToString())
		}
	}

}

// SetFlags sets the flags for virtual machine
func (vm *VM) SetFlags(flags map[string]interface{}) {
	vm.flags = flags
}

func (vm *VM) fetch() int {
	vm.instructionPointer++
	return vm.program[vm.instructionPointer]
}
