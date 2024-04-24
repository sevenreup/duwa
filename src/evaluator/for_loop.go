package evaluator

import (
	"github.com/sevenreup/chewa/src/ast"
	"github.com/sevenreup/chewa/src/object"
)

func evalForLoop(node *ast.ForExpression, env *object.Environment) object.Object {
	existingIdentifier, identifierExisted := env.Get(node.Identifier.Value)

	defer func() {
		if identifierExisted {
			env.Set(node.Identifier.Value, existingIdentifier)
		} else {
			env.Delete(node.Identifier.Value)
		}
	}()

	initializer := Eval(node.Initializer, env)

	if isError(initializer) {
		return initializer
	}

	loop := true

	for loop {
		condition := Eval(node.Condition, env)

		if isError(condition) {
			return condition
		}

		if isTruthy(condition) {
			evaluated := Eval(node.Block, env)

			if isTerminator(evaluated) {
				if evaluated.Type() == object.RETURN_VALUE_OBJ {
					return evaluated
				}
				switch val := evaluated.(type) {
				case *object.Error:
					return val
				case *object.Continue:
				case *object.Break:
					return nil
				}
			}

			evaluated = Eval(node.Increment, env)

			if isError(evaluated) {
				return evaluated
			}

			continue
		}

		loop = false
	}

	return NULL
}
