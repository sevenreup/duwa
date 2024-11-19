package native

import (
	"bufio"
	"errors"
	"os"

	"github.com/sevenreup/duwa/src/library/runtime"
)

type NativeConsole struct {
	runtime.Console
}

func NewConsole() *NativeConsole {
	return &NativeConsole{}
}

func (nc *NativeConsole) Read() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	val := scanner.Scan()

	if !val {
		return "", errors.New("Failed to read from console")
	}

	return scanner.Text(), nil
}
