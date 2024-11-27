package functions

import (
	"github.com/sevenreup/duwa/src/object"
	"github.com/sevenreup/duwa/src/token"
	"github.com/shopspring/decimal"
)

// type=builtin-func method=kuNambala args=[string{value}] return={integer}
// The kuNambala function converts a string to a number.
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
