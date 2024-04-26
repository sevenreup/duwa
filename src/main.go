package main

import (
	"flag"
	"log"
	"os"

	"github.com/sevenreup/chewa/src/evaluator"
	"github.com/sevenreup/chewa/src/object"
	"github.com/sevenreup/chewa/src/utils"

	"github.com/sevenreup/chewa/src/lexer"
	"github.com/sevenreup/chewa/src/parser"
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

	file, err := os.ReadFile(file)
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
