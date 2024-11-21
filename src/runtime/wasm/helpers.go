//go:build js && wasm

package wasm

import (
	"fmt"
	"syscall/js"
)

// wrapFunction provides error handling for JavaScript functions
func WrapFunction(fn func(this js.Value, args []js.Value) interface{}) func(this js.Value, args []js.Value) interface{} {
	return func(this js.Value, args []js.Value) interface{} {
		defer func() {
			if r := recover(); r != nil {
				// Convert panic to JavaScript error object
				errorObj := make(map[string]interface{})
				errorObj["error"] = fmt.Sprint(r)
				Wrap(errorObj)
			}
		}()
		return fn(this, args)
	}
}

// wrap safely converts Go values to JavaScript values
func Wrap(value interface{}) interface{} {
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
			arr[i] = Wrap(v)
		}
		return js.ValueOf(arr)
	case map[string]interface{}:
		obj := js.Global().Get("Object").New()
		for k, v := range val {
			obj.Set(k, Wrap(v))
		}
		return obj
	default:
		return js.ValueOf(fmt.Sprint(val))
	}
}

func DispatchEvent(name string, detail js.Value) {
	eventData := js.Global().Get("Object").New()
	eventData.Set("detail", map[string]interface{}{
		"detail": detail,
		"type":   name,
	})
	event := js.Global().Get("CustomEvent").New("goConsoleEvent", eventData)
	js.Global().Get("window").Call("dispatchEvent", event)
}
