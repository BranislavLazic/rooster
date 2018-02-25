package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/BranislavLazic/rooster/parser"
	"github.com/BranislavLazic/rooster/vm"
)

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
	program := parser.Program(string(fileContent))
	virtualMachine := vm.NewVM(program)
	virtualMachine.Run()
}
