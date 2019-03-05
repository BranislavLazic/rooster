package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BranislavLazic/rooster/cmd/rooster/lexer"
	"github.com/BranislavLazic/rooster/cmd/rooster/parser"
	"github.com/BranislavLazic/rooster/cmd/rooster/vm"
)

// HTTP methods
const (
	POST = "POST"
	GET  = "GET"
)

func handleExecution(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case POST:
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(405)
			fmt.Fprintf(w, "")
		} else {
			rcode := r.PostFormValue("rcode")
			lxr := lexer.NewLexer(rcode)
			constantPool := make(map[int]interface{})
			program := parser.Program(lxr, constantPool)
			virtualMachine := vm.NewVM(program, constantPool)
			virtualMachine.Run(w)
		}
	case GET:
		fmt.Fprintf(w, "Rooster server up and running")
	}
}

// StartServer starts HTTP server which executes rcode via
// HTTP request
func StartServer(port int) {
	http.HandleFunc("/", handleExecution)
	done := make(chan bool)
	go http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	log.Printf("Server started at %d port", port)
	<-done
}