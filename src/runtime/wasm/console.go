//go:build js && wasm

package wasm

import (
	"errors"
	"fmt"
	"syscall/js"
)

type WasmConsole struct {
	inputChan chan string
	errorChan chan error
	ready     chan struct{}
}

func NewConsole() *WasmConsole {
	console := &WasmConsole{
		inputChan: make(chan string),
		errorChan: make(chan error),
		ready:     make(chan struct{}),
	}

	js.Global().Set("duwaConsoleProcessInput", js.FuncOf(WrapFunction(console.processInput)))
	js.Global().Set("duwaConsoleReady", js.FuncOf(WrapFunction(console.consoleReady)))

	<-console.ready

	return console
}

func (wc *WasmConsole) Read() (string, error) {
	select {
	case input := <-wc.inputChan:
		return input, nil
	case err := <-wc.errorChan:
		return "", err
	}
}

func (wc *WasmConsole) Clear() error {
	eventInit := js.Global().Get("Object").New()
	eventInit.Set("command", "clear")
	DispatchEvent("duwaConsoleCommandEvent", eventInit)
	return nil
}

func (wc *WasmConsole) processInput(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		fmt.Println("Received input:", args[0].String())
		wc.inputChan <- args[0].String()
	} else {
		wc.errorChan <- errors.New("no input received")
	}
	return nil
}

func (wc *WasmConsole) consoleReady(this js.Value, args []js.Value) interface{} {
	close(wc.ready)
	fmt.Println("Console emulator is ready")
	return nil
}
