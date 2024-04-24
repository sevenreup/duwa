package object

import (
	"fmt"
)

const LIBRARY_FUNCTION = "LIBRARY_FUNCTION"

type LibraryFunction struct {
	Name     string
	Function GoFunction
}

func (libraryFunction *LibraryFunction) Inspect() string {
	return fmt.Sprintf("library function {%s}", libraryFunction.Name)
}

func (libraryFunction *LibraryFunction) Type() ObjectType {
	return LIBRARY_FUNCTION
}

func (libraryFunction *LibraryFunction) Method(method string, args []Object) (Object, bool) {
	return nil, false
}
