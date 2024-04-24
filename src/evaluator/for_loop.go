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
			err := Eval(node.Block, env)

			if err.Type() == object.RETURN_VALUE_OBJ {
				return err
			}

			if isTerminator(err) {
				switch val := err.(type) {
				case *object.Error:
					return val
				case *object.Continue:
				case *object.Break:
					return nil
				}
			}

			err = Eval(node.Increment, env)

			if isError(err) {
				return err
			}

			continue
		}

		loop = false
	}

	return NULL
}
