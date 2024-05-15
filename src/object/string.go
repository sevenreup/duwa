package object

import "hash/fnv"

const STRING_OBJ = "STRING"

type String struct {
	Object
	Mappable
	Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }

func (s *String) Inspect() string { return s.Value }

func (i *String) Method(method string, args []Object) (Object, bool) {
	//TODO implement me
	panic("implement me")
}

func (s *String) MapKey() MapKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return MapKey{Type: s.Type(), Value: h.Sum64()}
}
