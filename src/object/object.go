package object

import "github.com/sevenreup/chewa/src/token"

type ObjectType string

type Object interface {
	HasMethods
	Type() ObjectType
	Inspect() string
}

type MapKey struct {
	Type  ObjectType
	Value uint64
}

type Mappable interface {
	MapKey() MapKey
}

type HasMethods interface {
	Method(method string, args []Object) (Object, bool)
}

type GoFunction func(env *Environment, tok token.Token, args ...Object) Object
type GoProperty func(env *Environment, tok token.Token) Object
