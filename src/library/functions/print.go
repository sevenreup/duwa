package functions

import (
	"strings"

	"github.com/sevenreup/duwa/src/object"
	"github.com/sevenreup/duwa/src/token"
)

func Print(env *object.Environment, tok token.Token, args ...object.Object) object.Object {
	if len(args) > 0 {
		str := make([]string, 0)

		for _, value := range args {
			str = append(str, value.String())
		}
		env.Logger.Info(strings.Join(str, " "))
	}

	return nil
}

func PrintLine(env *object.Environment, tok token.Token, args ...object.Object) object.Object {
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
