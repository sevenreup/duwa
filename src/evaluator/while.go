package evaluator

import (
	"github.com/sevenreup/chewa/src/ast"
	"github.com/sevenreup/chewa/src/object"
)

func evaluateWhile(node *ast.WhileExpression, env *object.Environment) object.Object {
	for {
		condition := Eval(node.Condition, env)

		if isError(condition) {
			return condition
		}

		if isTruthy(condition) {
			evaluated := Eval(node.Consequence, env)

			if isTerminator(evaluated) {
				switch val := evaluated.(type) {
				case *object.Error:
					return val
				case *object.Continue:
					//
				case *object.Break:
					return nil
				}
			}
		} else {
			break
		}
	}

	return nil
}
