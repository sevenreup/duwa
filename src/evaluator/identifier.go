package evaluator

import (
	"github.com/sevenreup/chewa/src/ast"
	"github.com/sevenreup/chewa/src/object"
)

func evalIdentifier(
	node *ast.Identifier,
	env *object.Environment,
) object.Object {
	val, ok := env.Get(node.Value)
	if !ok {
		return newError("identifier not found: " + node.Value)
	}
	return val
}
