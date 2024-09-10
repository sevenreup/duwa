//go:build js && wasm

package main

import (
	"fmt"

	"syscall/js"

	"github.com/sevenreup/duwa/src/duwa"
)

var compiler *duwa.Duwa

func main() {
	compiler = duwa.New()
	js.Global().Set("runDuwa", js.FuncOf(jsRecover(run)))
	fmt.Println("Duwa is ready")
	<-make(chan bool)
}

func run(this js.Value, inputs []js.Value) interface{} {
	if len(inputs) < 1 {
		fmt.Println("Please provide a file to run")
		return nil
	}

	file := inputs[0].String()
	result := compiler.Run(file)
	return js.ValueOf(result.Inspect())
}

// jsRecover wraps a handler function to recover from panics and return an error message in case of a panic.
// It ensures that the function's result is compatible with js.ValueOf, is called by the browser when a handler is being executed.
func jsRecover(fn func(this js.Value, args []js.Value) any) func(this js.Value, args []js.Value) any {
	return func(this js.Value, args []js.Value) (result any) {
		defer func() {
			if r := recover(); r != nil {
				result = map[string]any{
					"error": fmt.Sprintf("Internal error: %v", r),
				}
			}
		}()

		result = fn(this, args)
		js.ValueOf(result)
		return
	}
}
