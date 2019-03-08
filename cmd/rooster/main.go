package main

import (
	"flag"
	"log"
	"os"

	"github.com/BranislavLazic/rooster"
)

func main() {
	// Read source file
	if len(os.Args) != 2 {
		log.Fatalf("provide source code file with rcode extension")
	}
	fileName := os.Args[1]
	debug := flag.Bool("debug", false, "Runs VM in debug mode")
	flag.Parse()
	fileContent, err := rooster.LoadRCodeFile(fileName)
	if err != nil {
		log.Fatalf("cannot read file")
	}
	// Read flags
	flags := map[string]interface{}{
		"debug": *debug,
	}
	rooster.RunVM(fileContent, flags, os.Stdout)
}
