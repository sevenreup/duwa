package functions

import (
	"fmt"
	"github.com/sevenreup/duwa/src/object"
	"github.com/sevenreup/duwa/src/token"
	"os"
	"strings"
)

func Print(env *object.Environment, tok token.Token, args ...object.Object) object.Object {
	if len(args) > 0 {
		str := make([]string, 0)

		for _, value := range args {
			str = append(str, value.Inspect())
		}

		fmt.Print(strings.Join(str, " "))
	}

	return nil
}

func PrintLine(env *object.Environment, tok token.Token, args ...object.Object) object.Object {
	if len(args) > 0 {
		str := make([]string, 0)

		for _, value := range args {
			str = append(str, value.Inspect())
		}

		fmt.Fprintln(os.Stdout, strings.Join(str, " "))
	} else {
		fmt.Fprintln(os.Stdout)
	}

	return nil
}
