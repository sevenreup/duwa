package modules

import (
	"github.com/sevenreup/duwa/src/ast"
	"github.com/sevenreup/duwa/src/object"
)

var evaluate func(node ast.Node, env *object.Environment) object.Object

func RegisterMethod(methods map[string]*object.LibraryFunction, name string, function object.GoFunction) {
	methods[name] = &object.LibraryFunction{Name: name, Function: function}
}

func RegisterEvaluator(e func(node ast.Node, env *object.Environment) object.Object) {
	evaluate = e
}
