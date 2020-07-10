package parser

import (
	"log"
	"strconv"

	"github.com/branislavlazic/rooster/pkg/lexer"
	"github.com/branislavlazic/rooster/pkg/token"
	"github.com/branislavlazic/rooster/pkg/vm"
)

// Program converts source code to array of instructions
func Program(lxr *lexer.Lexer, constantPool map[int]interface{}) []int {
	var instructions []int

	// Collect all tokens
	var tokens []token.Token
	for tok := lxr.NextToken(); tok.Type != token.EOF; tok = lxr.NextToken() {
		tokens = append(tokens, tok)
	}

	// Filter and transfer them to opcode instructions
	for idx, tok := range tokens {
		tok = findLabel(tok, tokens)
		if tok.Type == token.ILLEGAL {
			log.Fatalf("Line %d: %s is not a valid syntax", tok.LineNumber, tok.Literal)
		}

		// If the value should be in the constant pool, set the instruction as
		// the index of that value and push the value into the constant pool
		if tok.Type == token.STRING || tok.Type == token.FLOAT {
			constantPool[idx] = tok.Literal
			instructions = append(instructions, idx)
			continue
		}

		if isValidToken(tok) {
			instructions = append(instructions, tokenToInstruction(tok))
		}
	}
	return instructions
}

func tokenToInstruction(t token.Token) int {
	var instruction int
	switch t.Literal {
	case token.ICONST:
		instruction = vm.ICONST
	case token.SCONST:
		instruction = vm.SCONST
	case token.FCONST:
		instruction = vm.FCONST
	case token.IADD:
		instruction = vm.IADD
	case token.ISUB:
		instruction = vm.ISUB
	case token.IMUL:
		instruction = vm.IMUL
	case token.IDIV:
		instruction = vm.IDIV
	case token.JMP:
		instruction = vm.JMP
	case token.JMPT:
		instruction = vm.JMPT
	case token.JMPF:
		instruction = vm.JMPF
	case token.IEQ:
		instruction = vm.IEQ
	case token.ILT:
		instruction = vm.ILT
	case token.COPY:
		instruction = vm.COPY
	case token.GLOAD:
		instruction = vm.GLOAD
	case token.GSTORE:
		instruction = vm.GSTORE
	case token.LOAD:
		instruction = vm.LOAD
	case token.STORE:
		instruction = vm.STORE
	case token.CALL:
		instruction = vm.CALL
	case token.RET:
		instruction = vm.RET
	case token.PRINT:
		instruction = vm.PRINT
	case token.PRINTC:
		instruction = vm.PRINTC
	case token.HALT:
		instruction = vm.HALT
	default:
		instruction, _ = strconv.Atoi(t.Literal)
	}
	return instruction
}

func isValidToken(tok token.Token) bool {
	return tok.Type != token.EOL && tok.Type != token.EOF && tok.Type != token.COMMENT && tok.Type != token.LABEL
}

// Find label names and replace them with opcode index
func findLabel(tok token.Token, tokens []token.Token) token.Token {
	if tok.Type == token.LabelName {
		for _, tk := range tokens {
			if tk.Type == token.LABEL && tok.Literal == tk.Literal {
				tok.Type = token.INT
				tok.Literal = strconv.Itoa(tk.Index)
				break
			}
		}
	}
	return tok
}
