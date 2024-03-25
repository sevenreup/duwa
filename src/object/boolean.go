package object

import "fmt"

const BOOLEAN_OBJ = "BOOLEAN"

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

func (i *Boolean) Method(method string, args []Object) (Object, bool) {
	//TODO implement me
	panic("implement me")
}
