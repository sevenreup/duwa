package evaluator

import (
	"github.com/sevenreup/chewa/src/ast"
	"github.com/sevenreup/chewa/src/object"
)

func evalProgram(program *ast.Program, env *object.Environment) object.Object {
	var result object.Object
	for _, statement := range program.Statements {
		result = Eval(statement, env)
		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}
	return result
}
