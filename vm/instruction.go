package vm

// Instructions of assembly
const (
	_ = iota
	ICONST
	IADD
	ISUB
	IMUL
	JMP
	JMPT
	JMPF
	IEQ
	ILT
	GLOAD
	GSTORE
	LOAD
	PRINT
	HALT
)
