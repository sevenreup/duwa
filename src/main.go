package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sevenreup/chewa/src/lexer"
	"github.com/sevenreup/chewa/src/parser"
)

func main() {
	file, err := os.ReadFile("./examples/variables.ny")
	if err != nil {
		log.Fatal(err)
	}
	l := lexer.New(file)
	p := parser.New(l)
	program := p.ParseProgram()
	for _, v := range program.Statements {
		fmt.Printf("%+v\n", v)
	}
}
