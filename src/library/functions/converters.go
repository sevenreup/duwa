package functions

import (
	"github.com/sevenreup/duwa/src/object"
	"github.com/sevenreup/duwa/src/token"
	"github.com/shopspring/decimal"
)

func BuiltInParseStringToNumber(env *object.Environment, tok token.Token, args ...object.Object) object.Object {
	if len(args) != 1 {
		// TODO: Return error dont panic
		panic("parse requires exactly one argument")
	}

	if args[0].Type() != object.STRING_OBJ {
		return nil
	}

	str := args[0].(*object.String).Value
	number, _ := decimal.NewFromString(str)

	return &object.Integer{Value: number}
}
