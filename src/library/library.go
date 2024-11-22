package library

import (
	"github.com/sevenreup/duwa/src/library/functions"
	"github.com/sevenreup/duwa/src/library/modules"
	"github.com/sevenreup/duwa/src/object"
)

var Modules = map[string]*object.LibraryModule{}
var Functions = map[string]*object.LibraryFunction{}

func init() {
	RegisterModule("Khonso", modules.ConsoleMethods)
	RegisterModule("Masamu", modules.MathMethods)

	RegisterFunction("lemba", functions.BuiltInPrintLine)
	RegisterFunction("lembanzr", functions.BuiltInPrintLine)
	RegisterFunction("kuNambala", functions.BuiltInParseStringToNumber)
}

func RegisterModule(name string, methods map[string]*object.LibraryFunction) {
	Modules[name] = &object.LibraryModule{Name: name, Methods: methods}
}

func RegisterFunction(name string, function object.GoFunction) {
	Functions[name] = &object.LibraryFunction{Name: name, Function: function}
}
