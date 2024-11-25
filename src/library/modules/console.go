package modules

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sevenreup/duwa/src/object"
	"github.com/sevenreup/duwa/src/token"
	"github.com/sevenreup/duwa/src/values"
)

// library=Khonso
// This is the console module
// It contains functions that interact with the console
// It is used to read and write to the console
var ConsoleMethods = map[string]*object.LibraryFunction{}

func init() {
	RegisterMethod(ConsoleMethods, "lemba", methodConsolePrint)
	RegisterMethod(ConsoleMethods, "fufuta", methodConsoleClear)
	RegisterMethod(ConsoleMethods, "landira", methodConsoleRead)
}

// method=lemba args=[...] return={null}
// This method prints the arguments to the console
func methodConsolePrint(env *object.Environment, tok token.Token, args ...object.Object) object.Object {
	values := make([]string, 0)

	for _, value := range args {
		values = append(values, value.String())
	}

	libPrint(env, values)

	return nil
}

// method=landira args=[string{mawu}] return={string}
// This method reads a string from the console
func methodConsoleRead(scope *object.Environment, tok token.Token, args ...object.Object) object.Object {
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

// method=fufuta args=[] return={null}
// This method clears the console
func methodConsoleClear(scope *object.Environment, tok token.Token, args ...object.Object) object.Object {
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
