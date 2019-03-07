package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/BranislavLazic/rooster/cmd/rooster/lexer"
	"github.com/BranislavLazic/rooster/cmd/rooster/parser"
	"github.com/BranislavLazic/rooster/cmd/rooster/vm"
)

func main() {
	// Read source file
	if len(os.Args) != 2 {
		log.Fatalf("provide source code file with rcode extension")
	}
	fileName := os.Args[1]
	debug := flag.Bool("debug", false, "Runs VM in debug mode")
	startServer := flag.Bool("server", false, "Start HTTP server for remove execution")
	serverPort := flag.Int("serverPort", 8000, "HTTP server port")
	flag.Parse()
	if !*startServer {
		if fileName == "" {
			flag.PrintDefaults()
			log.Fatalf("source file must be provided via sourceFile flag")
		}
		if !strings.HasSuffix(fileName, ".rcode") {
			log.Fatalf("invalid file name. extension must be rcode")
		}
		fileContent, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Fatalf("cannot read file")
		}

		// Read flags
		flags := map[string]interface{}{
			"debug": *debug,
		}
		lxr := lexer.NewLexer(string(fileContent))
		constantPool := make(map[int]interface{})
		program := parser.Program(lxr, constantPool)
		vMachine := vm.NewVM(program, constantPool)
		vMachine.SetFlags(flags)
		vMachine.Run(os.Stdout)
	} else {
		StartServer(*serverPort)
	}
}
