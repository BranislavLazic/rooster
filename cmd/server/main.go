package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/branislavlazic/rooster"
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
			rooster.RunVM([]byte(r.PostFormValue("rcode")), make(map[string]interface{}), w)
		}
	case GET:
		fmt.Fprintf(w, "Rooster server up and running")
	}
}

func startServer(port int) {
	http.HandleFunc("/", handleExecution)
	done := make(chan bool)
	go http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	log.Printf("Server started at %d port", port)
	<-done
}

func main() {
	port := flag.Int("port", 8000, "HTTP server port")
	flag.Parse()
	startServer(*port)
}
