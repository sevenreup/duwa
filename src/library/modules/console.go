package modules

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sevenreup/chewa/src/object"
	"github.com/sevenreup/chewa/src/token"
)

var ConsoleMethods = map[string]*object.LibraryFunction{}

func init() {
	RegisterMethod(ConsoleMethods, "lemba", consolePrint)
}

func consolePrint(env *object.Environment, tok token.Token, args ...object.Object) object.Object {
	values := make([]string, 0)

	for _, value := range args {
		values = append(values, value.Inspect())
	}

	libPrint(values)

	return nil
}

func libPrint(values []string) {
	if len(values) > 0 {
		str := make([]string, 0)

		str = append(str, values...)
		strRaw, _ := strconv.Unquote(`"` + strings.Join(str, " ") + `"`)
		fmt.Print(strRaw)
	}
}
