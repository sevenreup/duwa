package object

const LIBRARY_MODULE = "LIBRARY_MODULE"

type LibraryModule struct {
	Name    string
	Methods map[string]*LibraryFunction
}

func (libraryModule *LibraryModule) String() string {
	return libraryModule.Name
}

func (libraryModule *LibraryModule) Type() ObjectType {
	return LIBRARY_MODULE
}

func (libraryModule *LibraryModule) Method(method string, args []Object) (Object, bool) {
	return nil, false
}
