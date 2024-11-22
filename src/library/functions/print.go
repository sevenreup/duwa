package functions

import (
	"strings"

	"github.com/sevenreup/duwa/src/object"
	"github.com/sevenreup/duwa/src/token"
)

// type=builtin-func method=peza args=[any{valueToPrint}] return={null}
// The lemba function prints a value to the console.
func BuiltInPrint(env *object.Environment, tok token.Token, args ...object.Object) object.Object {
	if len(args) > 0 {
		str := make([]string, 0)

		for _, value := range args {
			str = append(str, value.String())
		}
		env.Logger.Info(strings.Join(str, " "))
	}

	return nil
}

func BuiltInPrintLine(env *object.Environment, tok token.Token, args ...object.Object) object.Object {
	if len(args) > 0 {
		str := make([]string, 0)

		for _, value := range args {
			str = append(str, value.String())
		}

		env.Logger.Info(strings.Join(str, " ") + "\n")
	} else {
		env.Logger.Info("\n")
	}

	return nil
}
