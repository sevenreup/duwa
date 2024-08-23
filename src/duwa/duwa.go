package duwa

import (
	"log"
	"os"

	"github.com/sevenreup/duwa/src/evaluator"
	"github.com/sevenreup/duwa/src/object"
	"github.com/sevenreup/duwa/src/utils"

	"github.com/sevenreup/duwa/src/lexer"
	"github.com/sevenreup/duwa/src/parser"
)

type Duwa struct {
	file        string
	Environment *object.Environment
}

func New(file string) *Duwa {
	duwa := &Duwa{
		file:        file,
		Environment: object.NewEnvironment(),
	}
	duwa.registerEvaluator()
	return duwa
}

func (c *Duwa) Run() {
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

func (c *Duwa) registerEvaluator() {
	object.RegisterEvaluator(evaluator.Eval)
}
