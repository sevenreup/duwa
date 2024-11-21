package modules

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sevenreup/duwa/src/object"
	"github.com/sevenreup/duwa/src/token"
	"github.com/sevenreup/duwa/src/values"
)

var ConsoleMethods = map[string]*object.LibraryFunction{}

func init() {
	RegisterMethod(ConsoleMethods, "lemba", consolePrint)
	RegisterMethod(ConsoleMethods, "fufuta", consoleClear)
	RegisterMethod(ConsoleMethods, "landira", consoleRead)
}

func consolePrint(env *object.Environment, tok token.Token, args ...object.Object) object.Object {
	values := make([]string, 0)

	for _, value := range args {
		values = append(values, value.String())
	}

	libPrint(env, values)

	return nil
}

func consoleRead(scope *object.Environment, tok token.Token, args ...object.Object) object.Object {
	if len(args) == 1 {
		prompt := args[0].(*object.String).Value

		fmt.Print(prompt)
	}

	val, err := scope.Console.Read()

	if err != nil {
		return values.NULL
	}

	return &object.String{Value: val}
}

func consoleClear(scope *object.Environment, tok token.Token, args ...object.Object) object.Object {
	err := scope.Console.Clear()
	if err != nil {
		return values.NULL
	}

	return values.NULL
}

func libPrint(env *object.Environment, values []string) {
	if len(values) > 0 {
		str := make([]string, 0)

		str = append(str, values...)
		strRaw, _ := strconv.Unquote(`"` + strings.Join(str, " ") + `"`)
		env.Logger.Info(strRaw)
	}
}
