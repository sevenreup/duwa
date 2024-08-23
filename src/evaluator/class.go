package evaluator

import (
	"github.com/sevenreup/chewa/src/ast"
	"github.com/sevenreup/chewa/src/object"
)

func evaluateClass(node *ast.ClassStatement, env *object.Environment) object.Object {
	classEnv := object.NewEnclosedEnvironment(env)

	class := &object.Class{
		Name: node.Name,
		Env:  classEnv,
	}

	result := Eval(node.Body, classEnv)

	if isError(result) {
		return result
	}

	env.Set(class.Name.Value, class)

	return class
}
