package evaluator

import (
	"github.com/sevenreup/duwa/src/object"
	"github.com/sevenreup/duwa/src/values"
)

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return values.TRUE
	}
	return values.FALSE
}
