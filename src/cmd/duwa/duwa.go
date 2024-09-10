package main

import (
	"flag"
	"log"

	"github.com/sevenreup/duwa/src/duwa"
)

var (
	file string
)

func init() {
	flag.StringVar(&file, "f", "", "Source file")
}

func main() {
	flag.Parse()

	if file == "" {
		log.Fatal("Please provide a file to run")
	}

	duwa := duwa.New()
	duwa.RunFile(file)
}
