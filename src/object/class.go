package object

import "github.com/sevenreup/chewa/src/ast"

const CLASS_OBJ = "CLASS"

type Class struct {
	Object
	Name *ast.Identifier
	Env  *Environment
}

func (c *Class) Type() ObjectType { return CLASS_OBJ }

func (c *Class) Inspect() string {
	return "class " + c.Name.String()
}

func (i *Class) Method(method string, args []Object) (Object, bool) {
	switch method {
	case "new":
		instance := &Instance{Class: i, Env: NewEnclosedEnvironment(i.Env)}

		if ok := i.Env.Has("constructor"); ok {
			result := instance.Call("constructor", args)

			if result != nil && result.Type() == ERROR_OBJ {
				return result, false
			}
		}

		return instance, true
	}
	return nil, false
}
