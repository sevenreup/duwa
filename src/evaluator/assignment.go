package evaluator

import (
	"github.com/sevenreup/chewa/src/ast"
	"github.com/sevenreup/chewa/src/object"
)

func evaluateAssigment(node *ast.AssigmentStatement, env *object.Environment) object.Object {
	val := Eval(node.Value, env)
	if isError(val) {
		return val
	}
	switch smt := node.Identifier.(type) {
	case *ast.Identifier:
		return evaluateIdentifierAssignment(smt, val, env)
	}
	return nil
}

func evaluateIdentifierAssignment(node *ast.Identifier, val object.Object, env *object.Environment) object.Object {
	env.Set(node.Value, val)
	return nil
}
