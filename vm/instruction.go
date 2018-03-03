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
	ILT
	IEQ
	DUPL
	GLOAD
	GSTORE
	LOAD
	STORE
	CALL
	RET
	PRINT
	HALT
)
