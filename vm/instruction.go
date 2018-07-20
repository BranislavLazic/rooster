package vm

// Instructions of assembly
const (
	_ = iota
	ICONST
	SCONST
	IADD
	ISUB
	IMUL
	JMP
	JMPT
	JMPF
	ILT
	IEQ
	COPY
	GLOAD
	GSTORE
	LOAD
	STORE
	CALL
	RET
	PRINT
	HALT
)
