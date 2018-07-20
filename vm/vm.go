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
	constantPool       map[int]interface{}
	frameStack         *FrameStack
	flags              map[string]interface{}
}

// NewVM initializes the virtual machine
func NewVM(program []int, constantPool map[int]interface{}) *VM {
	return &VM{
		stack:              NewStack(),
		instructionPointer: -1,
		program:            program,
		globals:            make(map[int]int),
		constantPool:       constantPool,
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
		case SCONST:
			value := vm.fetch()
			vm.stack.Push(value)
		case IADD:
			addResult := vm.stack.Pop() + vm.stack.Pop()
			vm.stack.Push(addResult)
		case ISUB:
			subResult := vm.stack.Pop() - vm.stack.Pop()
			vm.stack.Push(subResult)
		case IMUL:
			mulResult := vm.stack.Pop() * vm.stack.Pop()
			vm.stack.Push(mulResult)
		// "Jump" to instruction pointer unconditionally. E.g. JMP 2 sets the instruction
		// pointer to 2 and starts executing program from that point.
		case JMP:
			jump := vm.fetch()
			vm.instructionPointer = jump - 1
		// "Jump" to instruction pointer conditionally. E.g. JMPT 2 will "pop" the value from
		// stack and check if it's 1. If the condition is satisfied, then the instruction pointer
		// will be set to 2. Otherwise, it will just continue with normal code execution.
		case JMPT:
			jump := vm.fetch()
			if vm.stack.Pop() == 1 {
				vm.instructionPointer = jump - 1
			}
		case JMPF:
			jump := vm.fetch()
			if vm.stack.Pop() == 0 {
				vm.instructionPointer = jump - 1
			}
		case ILT:
			if vm.stack.Pop() < vm.stack.Pop() {
				vm.stack.Push(1)
			} else {
				vm.stack.Push(0)
			}
		case IEQ:
			if vm.stack.Pop() == vm.stack.Pop() {
				vm.stack.Push(1)
			} else {
				vm.stack.Push(0)
			}
		// Makes a duplicate of stack top value and puts it
		// back on the stack as a new top value
		case COPY:
			topValue := vm.stack.Peek()
			vm.stack.Push(topValue)
		case GLOAD:
			address := vm.fetch()
			globalValue := vm.globals[address]
			vm.stack.Push(globalValue)
		case GSTORE:
			value := vm.stack.Pop()
			address := vm.fetch()
			vm.globals[address] = value
		// Pop the value from the stack and store it in the local memory
		// of the current frame
		case STORE:
			value := vm.stack.Pop()
			address := vm.fetch()
			vm.frameStack.Peek().variables[address] = value
		// Load value from the local memory of the current frame
		case LOAD:
			address := vm.fetch()
			frame := vm.frameStack.Peek()
			vm.stack.Push(frame.variables[address])
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
		case RET:
			returnAddress := vm.frameStack.Peek().returnAddress
			vm.frameStack.Pop()
			vm.instructionPointer = returnAddress
		case PRINT:
			fmt.Println(vm.stack.Pop())
		case HALT:
			return
		default:
			return
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
