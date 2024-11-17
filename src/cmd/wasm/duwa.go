//go:build js && wasm

package main

import (
	"context"
	"fmt"
	"github.com/sevenreup/duwa/src/duwa"
	"github.com/sevenreup/duwa/src/object"
	"log/slog"
	"runtime"
	"syscall/js"
)

var compiler *duwa.Duwa

type WasmConsoleHandler struct {
}

func main() {
	// Ensure the garbage collector runs more frequently
	runtime.SetFinalizer(compiler, nil)
	logger := slog.New(&WasmConsoleHandler{})
	compiler = duwa.New(object.New(logger))

	// Register the function in the global scope
	js.Global().Set("runDuwa", js.FuncOf(wrapFunction(run)))

	fmt.Println("Duwa is ready")

	// Keep the program running
	select {}
}

func run(this js.Value, inputs []js.Value) interface{} {
	if len(inputs) < 1 {
		return wrap("Please provide a file to run")
	}

	file := inputs[0].String()
	result := compiler.Run(file)
	if result == nil {
		return wrap(nil)
	}
	return wrap(result.Inspect())
}

// wrap safely converts Go values to JavaScript values
func wrap(value interface{}) interface{} {
	switch val := value.(type) {
	case nil:
		return js.Null()
	case string:
		return js.ValueOf(val)
	case int, int32, int64, float32, float64:
		return js.ValueOf(val)
	case bool:
		return js.ValueOf(val)
	case []interface{}:
		arr := make([]interface{}, len(val))
		for i, v := range val {
			arr[i] = wrap(v)
		}
		return js.ValueOf(arr)
	case map[string]interface{}:
		obj := js.Global().Get("Object").New()
		for k, v := range val {
			obj.Set(k, wrap(v))
		}
		return obj
	default:
		return js.ValueOf(fmt.Sprint(val))
	}
}

// wrapFunction provides error handling for JavaScript functions
func wrapFunction(fn func(this js.Value, args []js.Value) interface{}) func(this js.Value, args []js.Value) interface{} {
	return func(this js.Value, args []js.Value) interface{} {
		defer func() {
			if r := recover(); r != nil {
				// Convert panic to JavaScript error object
				errorObj := make(map[string]interface{})
				errorObj["error"] = fmt.Sprint(r)
				wrap(errorObj)
			}
		}()
		return fn(this, args)
	}
}

func (h *WasmConsoleHandler) Enabled(_ context.Context, level slog.Level) bool {
	return true
}

func emitConsoleEvent(r slog.Record) {
	// Create a custom event
	eventInit := js.Global().Get("Object").New()
	eventInit.Set("detail", map[string]interface{}{
		"message": r.Message,
		"level":   slogLevelToConsoleLevel(r.Level),
	})

	logType := "runtime"
	for attr := range r.Attrs {
		if attr.Key == "type" {
			logType = attr.Value.String()
		}
	}
	eventInit.Get("detail").Set("type", logType)

	event := js.Global().Get("CustomEvent").New("goConsoleEvent", eventInit)

	// Dispatch the event
	js.Global().Get("window").Call("dispatchEvent", event)
}

func (h *WasmConsoleHandler) Handle(ctx context.Context, r slog.Record) error {
	emitConsoleEvent(r)
	return nil
}

func (h *WasmConsoleHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *WasmConsoleHandler) WithGroup(name string) slog.Handler {
	return h
}

func (h *WasmConsoleHandler) Handler() slog.Handler {
	return h
}

func slogLevelToConsoleLevel(level slog.Level) string {
	switch level {
	case slog.LevelDebug:
		return "debug"
	case slog.LevelInfo:
		return "info"
	case slog.LevelWarn:
		return "warn"
	case slog.LevelError:
		return "error"
	default:
		return "info"
	}
}
