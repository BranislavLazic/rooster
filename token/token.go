package token

// Types
const (
	INT       = "INT"
	STRING    = "STRING"
	EOL       = "EOL"
	EOF       = "EOF"
	COMMENT   = "COMMENT"
	LabelName = "LABEL_NAME"
	LABEL     = "LABEL"
	ILLEGAL   = "ILLEGAL"
	// Instructions
	ICONST = "ICONST"
	SCONST = "SCONST"
	IADD   = "IADD"
	ISUB   = "ISUB"
	IMUL   = "IMUL"
	IDIV   = "IDIV"
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
	"IDUV":   IDIV,
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
