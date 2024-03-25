package object

const NULL_OBJ = "NULL"

type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }

func (n *Null) Inspect() string { return "null" }

func (i *Null) Method(method string, args []Object) (Object, bool) {
	//TODO implement me
	panic("implement me")
}
