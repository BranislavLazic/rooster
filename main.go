package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/branislavlazic/rooster/lexer"
	"github.com/branislavlazic/rooster/token"
	"github.com/branislavlazic/rooster/vm"
)

func tokenToInstruction(t token.Token) int {
	var instruction int
	switch t.Literal {
	case token.ICONST:
		instruction = vm.ICONST
		break
	case token.IADD:
		instruction = vm.IADD
		break
	case token.ISUB:
		instruction = vm.ISUB
		break
	case token.IMUL:
		instruction = vm.IMUL
		break
	case token.JMP:
		instruction = vm.JMP
		break
	case token.JMPT:
		instruction = vm.JMPT
		break
	case token.JMPF:
		instruction = vm.JMPF
		break
	case token.IEQ:
		instruction = vm.IEQ
		break
	case token.ILT:
		instruction = vm.ILT
		break
	case token.PRINT:
		instruction = vm.PRINT
		break
	case token.HALT:
		instruction = vm.HALT
		break
	default:
		instruction, _ = strconv.Atoi(t.Literal)
	}
	return instruction
}

func isSyntaxError(tok token.Token) bool {
	return tok.Type == token.ILLEGAL
}

func isValidToken(tok token.Token) bool {
	return tok.Type != token.EOL && tok.Type != token.EOF && tok.Type != token.COMMENT
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("please provide an assembly file name")
		os.Exit(1)
	}
	fileName := os.Args[1]
	if !strings.HasSuffix(fileName, ".rcode") {
		log.Fatalf("invalid file name. extension must be rcode")
		os.Exit(1)
	}
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	program := string(fileContent)
	var instructions []int
	lxr := lexer.NewLexer(program)
	for tok := lxr.NextToken(); tok.Type != token.EOF; tok = lxr.NextToken() {
		if isSyntaxError(tok) {
			log.Fatalf("%s is not a valid syntax", tok.Literal)
			os.Exit(1)
		}
		if isValidToken(tok) {
			instructions = append(instructions, tokenToInstruction(tok))
		}

	}
	virtualMachine := vm.NewVM(instructions)
	virtualMachine.Run()
}
