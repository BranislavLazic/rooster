package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/BranislavLazic/rooster/lexer"
	"github.com/BranislavLazic/rooster/parser"
	"github.com/BranislavLazic/rooster/vm"
)

func main() {
	// Read source file
	fileName := flag.String("sourceFile", "", "Source code file with the rcode extension")
	printStack := flag.Bool("printStack", false, "Prints the state of the stack")
	flag.Parse()
	if *fileName == "" {
		flag.PrintDefaults()
		log.Fatalf("source file must be provided via sourceFile flag")
		os.Exit(1)
	}
	if !strings.HasSuffix(*fileName, ".rcode") {
		log.Fatalf("invalid file name. extension must be rcode")
		os.Exit(1)
	}
	fileContent, err := ioutil.ReadFile(*fileName)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	// Read flags
	flags := map[string]interface{}{
		"printStack": *printStack,
	}

	lexer := lexer.NewLexer(string(fileContent))
	program := parser.Program(lexer)
	virtualMachine := vm.NewVM(program)
	virtualMachine.SetFlags(flags)
	virtualMachine.Run()
}
