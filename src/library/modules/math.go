package modules

import (
	"github.com/sevenreup/duwa/src/object"
	"github.com/sevenreup/duwa/src/token"
	"github.com/shopspring/decimal"
)

var MathMethods = map[string]*object.LibraryFunction{}

func init() {
	RegisterMethod(MathMethods, "yochepa", methodMathMin)
	RegisterMethod(MathMethods, "sqrt", methodMathSqrt)
}

func methodMathMin(scope *object.Environment, tok token.Token, args ...object.Object) object.Object {
	if len(args) < 2 {
		panic("Masamu.yochepa requires at least two arguments")
	}

	if args[0].Type() != object.INTEGER_OBJ {
		return nil
	}

	if args[1].Type() != object.INTEGER_OBJ {
		return nil
	}

	number1 := args[0].(*object.Integer)
	number2 := args[1].(*object.Integer)

	if number1.Value.LessThan(number2.Value) {
		return number1
	}

	return number2
}

func methodMathSqrt(scope *object.Environment, tok token.Token, args ...object.Object) object.Object {
	if len(args) != 1 {
		panic("Masamu.sqrt requires one argument")
	}

	if args[0].Type() != object.INTEGER_OBJ {
		return nil
	}

	number := args[0].(*object.Integer)

	// NOTE: using pow(0.5) is the same as sqrt() but may not be accurate
	return &object.Integer{Value: number.Value.Pow(decimal.NewFromFloat(0.5))}
}
