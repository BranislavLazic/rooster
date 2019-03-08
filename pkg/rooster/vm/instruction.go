package vm

// Instructions of assembly
const (
	_ = iota
	ICONST
	SCONST
	FCONST
	IADD
	ISUB
	IMUL
	IDIV
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
	// Use to print value from stack
	PRINT
	// Use to print value from constant pool which value is
	// identified by an index from stack
	PRINTC
	HALT
)
