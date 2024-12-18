package evaluator

import (
	"github.com/sevenreup/duwa/src/ast"
	"github.com/sevenreup/duwa/src/object"
)

func evaluateContinue(node *ast.ContinueStatement, env *object.Environment) object.Object {
	return &object.Continue{}
}
