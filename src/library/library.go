package library

import (
	"github.com/sevenreup/chewa/src/library/functions"
	"github.com/sevenreup/chewa/src/library/modules"
	"github.com/sevenreup/chewa/src/object"
)

var Modules = map[string]*object.LibraryModule{}
var Functions = map[string]*object.LibraryFunction{}

func init() {
	RegisterModule("console", modules.ConsoleMethods)
	RegisterModule("math", modules.MathMethods)

	RegisterFunction("lemba", functions.Print)
	RegisterFunction("lembanzr", functions.PrintLine)
	RegisterFunction("kuNambala", functions.ParseStringToNumber)
}

func RegisterModule(name string, methods map[string]*object.LibraryFunction) {
	Modules[name] = &object.LibraryModule{Name: name, Methods: methods}
}

func RegisterFunction(name string, function object.GoFunction) {
	Functions[name] = &object.LibraryFunction{Name: name, Function: function}
}
