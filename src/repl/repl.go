package repl

import (
	"bufio"
	"fmt"
	"github.com/sevenreup/chewa/src/evaluator"
	"github.com/sevenreup/chewa/src/object"
	"github.com/sevenreup/chewa/src/parser"
	"github.com/sevenreup/chewa/src/utils"
	"io"

	"github.com/sevenreup/chewa/src/lexer"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New([]byte(line))
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			utils.PrintParserErrors(out, p.Errors())
			continue
		}
		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}
