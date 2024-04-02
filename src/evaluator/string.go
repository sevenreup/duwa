package evaluator

import "github.com/sevenreup/chewa/src/object"

func evalStringInfixExpression(
	operator string,
	left, right object.Object,
) object.Object {
	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value
	switch operator {
	case "+":
		return &object.String{Value: leftVal + rightVal}
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
}
