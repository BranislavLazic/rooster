package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/BranislavLazic/rooster/server"

	"github.com/BranislavLazic/rooster/lexer"
	"github.com/BranislavLazic/rooster/parser"
	"github.com/BranislavLazic/rooster/vm"
)

func main() {
	// Read source file
	fileName := flag.String("sourceFile", "", "Source code file with the rcode extension")
	debug := flag.Bool("debug", false, "Runs VM in debug mode")
	startServer := flag.Bool("server", false, "Start HTTP server for remove execution")
	serverPort := flag.Int("serverPort", 8000, "HTTP server port")
	flag.Parse()
	if !*startServer {
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
			"debug": *debug,
		}
		lexer := lexer.NewLexer(string(fileContent))
		constantPool := make(map[int]interface{})
		program := parser.Program(lexer, constantPool)
		virtualMachine := vm.NewVM(program, constantPool)
		virtualMachine.SetFlags(flags)
		virtualMachine.Run(os.Stdout)
	} else {
		server.StartServer(*serverPort)
	}
}
