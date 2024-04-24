package evaluator

import (
	"github.com/sevenreup/chewa/src/ast"
	"github.com/sevenreup/chewa/src/library"
	"github.com/sevenreup/chewa/src/object"
)

func evalIdentifier(
	node *ast.Identifier,
	env *object.Environment,
) object.Object {
	if libraryModule, ok := library.Modules[node.Value]; ok {
		return libraryModule
	}

	if libraryFunction, ok := library.Functions[node.Value]; ok {
		return libraryFunction
	}

	val, ok := env.Get(node.Value)
	if !ok {
		return newError("identifier not found: " + node.Value)
	}

	return val
}
