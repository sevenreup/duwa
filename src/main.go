package main

import (
	"flag"
	"log"

	"github.com/sevenreup/chewa/src/chewa"
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

	chewa := chewa.New(file)
	chewa.Run()
}
