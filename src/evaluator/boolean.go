package evaluator

import "github.com/sevenreup/chewa/src/object"

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}
