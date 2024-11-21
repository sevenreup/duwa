package native

import (
	"bufio"
	"errors"
	"os"
	"os/exec"
	"runtime"
)

type NativeConsole struct {
}

func NewConsole() *NativeConsole {
	return &NativeConsole{}
}

func (nc *NativeConsole) Read() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	val := scanner.Scan()

	if !val {
		return "", errors.New("failed to read from console")
	}

	return scanner.Text(), nil
}

func (nc *NativeConsole) Clear() error {
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
