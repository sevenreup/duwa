package modules

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
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
		values = append(values, value.Inspect())
	}

	libPrint(values)

	return nil
}

func consoleRead(scope *object.Environment, tok token.Token, args ...object.Object) object.Object {
	scanner := bufio.NewScanner(os.Stdin)

	if len(args) == 1 {
		prompt := args[0].(*object.String).Value

		fmt.Print(prompt)
	}

	val := scanner.Scan()

	if !val {
		return values.NULL
	}

	return &object.String{Value: scanner.Text()}
}

func consoleClear(scope *object.Environment, tok token.Token, args ...object.Object) object.Object {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

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
