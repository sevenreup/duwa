package evaluator

import (
	"github.com/sevenreup/duwa/src/ast"
	"github.com/sevenreup/duwa/src/object"
	"github.com/sevenreup/duwa/src/values"
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

	switch receiver := left.(type) {
	case *object.LibraryModule:
		method := node.Method.(*ast.Identifier)

		if function, ok := receiver.Methods[method.Value]; ok {
			return applyFunction(node.Token, function, arguments, env)
		}
	case *object.Instance:
		method := node.Method.(*ast.Identifier)
		evaluated := evaluateInstanceMethod(node, receiver, method.Value, arguments)

		if isError(evaluated) {
			return evaluated
		}

		return unwrapReturn(evaluated)
	}

	return result
}

func evaluateInstanceMethod(node *ast.MethodExpression, receiverInstance *object.Instance, name string, arguments []object.Object) object.Object {
	method, ok := receiverInstance.Class.Env.Get(name)

	if !ok {
		return newError("%d:%d:%s: runtime error: undefined method %s for class %s", node.Token.Pos.Line, node.Token.Pos.Column, node.Token.File, name, receiverInstance.Class.Name.Value)
	}

	if method, ok := method.(*object.Function); ok {
		extendedEnv := extendFunctionEnv(method, arguments)
		return Eval(method.Body, extendedEnv)
	} else {
		return newError("not a method: %s", name)
	}
}

func unwrapReturn(obj object.Object) object.Object {
	switch value := obj.(type) {
	case *object.Error:
		return obj
	case *object.ReturnValue:
		return value.Value
	}

	return values.NULL
}
