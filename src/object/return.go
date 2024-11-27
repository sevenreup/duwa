package object

const RETURN_VALUE_OBJ = "RETURN_VALUE"

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }

func (rv *ReturnValue) String() string { return rv.Value.String() }

func (i *ReturnValue) Method(method string, args []Object) (Object, bool) {
	return nil, false
}
