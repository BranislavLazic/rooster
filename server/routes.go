package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/BranislavLazic/rooster/lexer"
	"github.com/BranislavLazic/rooster/parser"
	"github.com/BranislavLazic/rooster/vm"
)

type code struct {
	RCode string `json:rcode`
}

func executeRCode(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var c code
		err := decoder.Decode(&c)
		if err != nil {
			log.Fatalf("failed to decode body")
		}
		lexer := lexer.NewLexer(string(c.RCode))
		constantPool := make(map[int]interface{})
		program := parser.Program(lexer, constantPool)
		virtualMachine := vm.NewVM(program, constantPool)
		virtualMachine.Run(w)
	}
}

// StartServer starts HTTP server which executes rcode via
// simple HTTP request
func StartServer() {
	http.HandleFunc("/", executeRCode)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
