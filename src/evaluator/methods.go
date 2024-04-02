package evaluator

import (
	"github.com/sevenreup/chewa/src/ast"
	"github.com/sevenreup/chewa/src/object"
)

func evaluateMethod(node *ast.MethodExpression, scope *object.Environment) object.Object {
	left := Eval(node.Left, scope)

	if isError(left) {
		return left
	}

	arguments := evalExpressions(node.Arguments, scope)

	if len(arguments) == 1 && isError(arguments[0]) {
		return arguments[0]
	}

	result, _ := left.Method(node.Method.(*ast.Identifier).Value, arguments)

	if isError(result) {
		return result
	}

	return result
}
