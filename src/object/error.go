package object

import "fmt"

const ERROR_OBJ = "ERROR"

type Error struct {
	Object
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }

func (e *Error) String() string { return "ERROR: " + e.Message }

func (i *Error) Method(method string, args []Object) (Object, bool) {
	//TODO implement me
	panic("implement me")
}

func NewError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}
