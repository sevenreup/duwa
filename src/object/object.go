package object

type ObjectType string

type Object interface {
	HasMethods
	Type() ObjectType
	Inspect() string
}

type HasMethods interface {
	Method(method string, args []Object) (Object, bool)
}
