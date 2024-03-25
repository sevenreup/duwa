package object

const STRING_OBJ = "STRING"

type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }

func (s *String) Inspect() string { return s.Value }

func (i *String) Method(method string, args []Object) (Object, bool) {
	//TODO implement me
	panic("implement me")
}
