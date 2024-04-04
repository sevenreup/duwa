package object

import "github.com/sevenreup/chewa/src/token"

type ObjectType string

type Object interface {
	HasMethods
	Type() ObjectType
	Inspect() string
}

type HasMethods interface {
	Method(method string, args []Object) (Object, bool)
}

type GoFunction func(scope *Environment, tok token.Token, args ...Object) Object
type GoProperty func(scope *Environment, tok token.Token) Object
