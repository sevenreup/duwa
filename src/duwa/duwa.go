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
	Environment *object.Environment
}

func New(env *object.Environment) *Duwa {
	duwa := &Duwa{
		Environment: env,
	}
	duwa.registerEvaluator()
	return duwa
}

func (c *Duwa) RunFile(filePath string) object.Object {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return c.run(file)
}

func (c *Duwa) Run(data string) object.Object {
	return c.run([]byte(data))
}

func (c *Duwa) run(data []byte) object.Object {
	l := lexer.New(data)
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		utils.PrintParserErrors(os.Stdout, p.Errors())
	}
	return evaluator.Eval(program, c.Environment)
}

func (c *Duwa) registerEvaluator() {
	object.RegisterEvaluator(evaluator.Eval)
}
