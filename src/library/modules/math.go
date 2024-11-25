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
	RegisterMethod(MathMethods, "round", methodRound)
	RegisterMethod(MathMethods, "pansi", methodFloor)
}

// method=yochepa args=[number{number1}, number{number2}] return={number}
// This method returns the smaller of two numbers.
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

// method=sqrt args=[number{number1}] return={number}
// This method returns the square root of a number.
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

// method=round args=[number{number1}, number{number2}] return={number}
// This method rounds a number to a specified number of decimal places.
func methodRound(scope *object.Environment, tok token.Token, args ...object.Object) object.Object {
	if len(args) < 2 {
		panic("Masamu.yochepa requires at least two arguments")
	}

	if args[0].Type() != object.INTEGER_OBJ {
		return nil
	}

	if args[1].Type() != object.INTEGER_OBJ {
		return nil
	}

	number := args[0].(*object.Integer)
	places := args[1].(*object.Integer)

	return &object.Integer{Value: number.Value.Round(int32(places.Value.IntPart()))}
}

// method=pansi args=[number{number1}] return={number}
// This method returns the largest integer less than or equal to a number.
func methodFloor(scope *object.Environment, tok token.Token, args ...object.Object) object.Object {
	if args[0].Type() != object.INTEGER_OBJ {
		return nil
	}
	number1 := args[0].(*object.Integer)
	return &object.Integer{Value: number1.Value.Floor()}
}
