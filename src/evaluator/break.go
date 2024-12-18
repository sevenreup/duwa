package evaluator

import (
	"github.com/sevenreup/duwa/src/ast"
	"github.com/sevenreup/duwa/src/object"
)

func evaluateBreak(node *ast.BreakStatement, env *object.Environment) object.Object {
	return &object.Break{}
}
