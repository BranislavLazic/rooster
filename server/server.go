package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BranislavLazic/rooster/lexer"
	"github.com/BranislavLazic/rooster/parser"
	"github.com/BranislavLazic/rooster/vm"
)

func handleExecution(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		err := r.ParseForm()
		if err != nil {
			log.Fatalf("failed to parse body")
		}
		rcode := r.PostFormValue("rcode")
		lexer := lexer.NewLexer(rcode)
		constantPool := make(map[int]interface{})
		program := parser.Program(lexer, constantPool)
		virtualMachine := vm.NewVM(program, constantPool)
		virtualMachine.Run(w)
	case "GET":
		fmt.Fprintf(w, "Rooster server up and running")
	}
}

// StartServer starts HTTP server which executes rcode via
// simple HTTP request
func StartServer(port int) {
	http.HandleFunc("/", handleExecution)
	done := make(chan bool)
	go http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	log.Printf("Server started at %d port", port)
	<-done
}
