package token

const (
	IDENT   = "IDENT"
	INT     = "INT"
	EOL     = "EOL"
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"
	// Instructions
	ICONST = "ICONST"
	IADD   = "IADD"
	ISUB   = "ISUB"
	IMUL   = "IMUL"
	PRINT  = "PRINT"
	HALT   = "HALT"
)

var instructions = map[string]string{
	"ICONST": ICONST,
	"IADD":   IADD,
	"ISUB":   ISUB,
	"IMUL":   IMUL,
	"PRINT":  PRINT,
	"HALT":   HALT,
}

func LookupInstruction(instruction string) string {
	if tok, ok := instructions[instruction]; ok {
		return tok
	}
	return IDENT
}

type Token struct {
	Type    string
	Literal string
}
