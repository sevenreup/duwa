package evaluator

import "github.com/sevenreup/chewa/src/object"

func evalIndexExpression(left, index object.Object) object.Object {
	switch {
	case left.Type() == object.ARRAY_OBJ && index.Type() == object.INTEGER_OBJ:
		return evalArrayIndexExpression(left, index)
	default:
		return newError("index operator not supported: %s", left.Type())
	}
}

func evalArrayIndexExpression(array, index object.Object) object.Object {
	arrayObject := array.(*object.Array)
	idx := index.(*object.Integer).Value.IntPart()
	maxValue := int64(len(arrayObject.Elements) - 1)
	if idx < 0 || idx > maxValue {
		return NULL
	}
	return arrayObject.Elements[idx]
}
