package object

const BREAK_OBJ = "BREAK"

type Break struct{
	Object
}

func (obj *Break) Type() ObjectType {
	return BREAK_OBJ
}

func (obj *Break) Inspect() string {
	return BREAK_OBJ
}

func (obj *Break) Method(method string, args []Object) (Object, bool) {
	return nil, false
}
