package evaluator

import (
	"github.com/sevenreup/duwa/src/ast"
	"github.com/sevenreup/duwa/src/object"
)

func evaluateMethod(node *ast.MethodExpression, env *object.Environment) object.Object {
	left := Eval(node.Left, env)

	if isError(left) {
		return left
	}

	arguments := evalExpressions(node.Arguments, env)

	if len(arguments) == 1 && isError(arguments[0]) {
		return arguments[0]
	}

	result, _ := left.Method(node.Method.(*ast.Identifier).Value, arguments)

	if isError(result) {
		return result
	}

	switch left.(type) {
	case *object.LibraryModule:
		method := node.Method.(*ast.Identifier)
		module := left.(*object.LibraryModule)

		if function, ok := module.Methods[method.Value]; ok {
			return applyFunction(node.Token, function, arguments, env)
		}
	}

	return result
}
