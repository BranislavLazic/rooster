package token

const (
	// Types
	INT        = "INT"
	STRING     = "STRING"
	EOL        = "EOL"
	EOF        = "EOF"
	COMMENT    = "COMMENT"
	LABEL_NAME = "LABEL_NAME"
	LABEL      = "LABEL"
	ILLEGAL    = "ILLEGAL"
	// Instructions
	ICONST = "ICONST"
	SCONST = "SCONST"
	IADD   = "IADD"
	ISUB   = "ISUB"
	IMUL   = "IMUL"
	JMP    = "JMP"
	JMPT   = "JMPT"
	JMPF   = "JMPF"
	IEQ    = "IEQ"
	ILT    = "ILT"
	COPY   = "COPY"
	GLOAD  = "GLOAD"
	GSTORE = "GSTORE"
	LOAD   = "LOAD"
	STORE  = "STORE"
	CALL   = "CALL"
	RET    = "RET"
	PRINT  = "PRINT"
	PRINTC = "PRINTC"
	HALT   = "HALT"
)

var instructions = map[string]string{
	"ICONST": ICONST,
	"SCONST": SCONST,
	"IADD":   IADD,
	"ISUB":   ISUB,
	"IMUL":   IMUL,
	"JMP":    JMP,
	"JMPT":   JMPT,
	"JMPF":   JMPF,
	"IEQ":    IEQ,
	"ILT":    ILT,
	"COPY":   COPY,
	"GLOAD":  GLOAD,
	"GSTORE": GSTORE,
	"LOAD":   LOAD,
	"STORE":  STORE,
	"CALL":   CALL,
	"RET":    RET,
	"PRINT":  PRINT,
	"PRINTC": PRINTC,
	"HALT":   HALT,
}

func LookupInstruction(instruction string) string {
	if tok, ok := instructions[instruction]; ok {
		return tok
	}
	return ILLEGAL
}

type Token struct {
	Type       string
	Literal    string
	LineNumber int
	Index      int
}
