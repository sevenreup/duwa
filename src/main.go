package main

import (
	"github.com/sevenreup/chewa/src/evaluator"
	"github.com/sevenreup/chewa/src/object"
	"github.com/sevenreup/chewa/src/utils"
	"log"
	"os"

	"github.com/sevenreup/chewa/src/lexer"
	"github.com/sevenreup/chewa/src/parser"
)

func main() {
	file, err := os.ReadFile("./examples/sorting.ny")
	if err != nil {
		log.Fatal(err)
	}
	l := lexer.New(file)
	p := parser.New(l)
	env := object.NewEnvironment()
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		utils.PrintParserErrors(os.Stdout, p.Errors())
	}
	evaluator.Eval(program, env)
}
