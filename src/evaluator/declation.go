package evaluator

import (
	"github.com/sevenreup/chewa/src/ast"
	"github.com/sevenreup/chewa/src/object"
)

// TODO: Handle type
func evaluateDeclaration(node *ast.VariableDeclarationStatement, env *object.Environment) object.Object {
	val := Eval(node.Value, env)
	if isError(val) {
		return val
	}
	env.Set(node.Identifier.Value, val)
	return nil
}
