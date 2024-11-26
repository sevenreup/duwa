package functions

import (
	"strings"

	"github.com/sevenreup/duwa/src/object"
	"github.com/sevenreup/duwa/src/token"
)

// type=builtin-func method=lemba args=[any{valueToPrint}] return={null}
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

// type=builtin-func method=lembanzr args=[any{valueToPrint}] return={null}
// The lembanzr function prints a value to the console and adds a newline.
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
