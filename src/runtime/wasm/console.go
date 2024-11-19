//go:build js && wasm

package wasm

import (
	"errors"
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

	js.Global().Set("goProcessInput", js.FuncOf(console.processInput))
	js.Global().Set("goConsoleReady", js.FuncOf(console.consoleReady))

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

func (wc *WasmConsole) processInput(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		wc.inputChan <- args[0].String()
	} else {
		wc.errorChan <- errors.New("no input received")
	}
	return nil
}

func (wc *WasmConsole) consoleReady(this js.Value, args []js.Value) interface{} {
	close(wc.ready)
	return nil
}
