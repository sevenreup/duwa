package object

const ERROR_OBJ = "ERROR"

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }

func (e *Error) Inspect() string { return "ERROR: " + e.Message }

func (i *Error) Method(method string, args []Object) (Object, bool) {
	//TODO implement me
	panic("implement me")
}
