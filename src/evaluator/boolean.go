package evaluator

import (
	"github.com/sevenreup/chewa/src/object"
	"github.com/sevenreup/chewa/src/values"
)

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return values.TRUE
	}
	return values.FALSE
}
