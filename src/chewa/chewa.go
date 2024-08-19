package chewa

import (
	"log"
	"os"

	"github.com/sevenreup/chewa/src/evaluator"
	"github.com/sevenreup/chewa/src/object"
	"github.com/sevenreup/chewa/src/utils"

	"github.com/sevenreup/chewa/src/lexer"
	"github.com/sevenreup/chewa/src/parser"
)

type Chewa struct {
	file        string
	Environment *object.Environment
}

func New(file string) *Chewa {
	chewa := &Chewa{
		file:        file,
		Environment: object.NewEnvironment(),
	}
	chewa.registerEvaluator()
	return chewa
}

func (c *Chewa) Run() {
	file, err := os.ReadFile(c.file)
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

func (c *Chewa) registerEvaluator() {
	object.RegisterEvaluator(evaluator.Eval)
}
