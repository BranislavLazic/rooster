package rooster

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"strings"

	"github.com/branislavlazic/rooster/pkg/lexer"
	"github.com/branislavlazic/rooster/pkg/parser"
	"github.com/branislavlazic/rooster/pkg/vm"
)

// LoadRCodeFile will try to load a file with the rcode extension.
func LoadRCodeFile(fileName string) ([]byte, error) {
	if fileName == "" {
		flag.PrintDefaults()
		log.Fatalf("source file must be provided via sourceFile flag")
	}
	if !strings.HasSuffix(fileName, ".rcode") {
		log.Fatalf("invalid file name. extension must be rcode")
	}
	return ioutil.ReadFile(fileName)
}

// RunVM reads a file and executes it on the virtual machine
func RunVM(fileContent []byte, flags map[string]interface{}, w io.Writer) {
	lxr := lexer.NewLexer(string(fileContent))
	constantPool := make(map[int]interface{})
	program := parser.Program(lxr, constantPool)
	vMachine := vm.NewVM(program, constantPool)
	vMachine.SetFlags(flags)
	vMachine.Run(w)
}
