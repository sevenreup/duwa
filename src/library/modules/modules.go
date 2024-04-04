package modules

import (
	"github.com/sevenreup/chewa/src/ast"
	"github.com/sevenreup/chewa/src/object"
)

var evaluate func(node ast.Node, scope *object.Environment) object.Object

func RegisterMethod(methods map[string]*object.LibraryFunction, name string, function object.GoFunction) {
	methods[name] = &object.LibraryFunction{Name: name, Function: function}
}

func RegisterEvaluator(e func(node ast.Node, scope *object.Environment) object.Object) {
	evaluate = e
}
