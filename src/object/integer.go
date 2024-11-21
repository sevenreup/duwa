package object

import (
	"fmt"
	"hash/fnv"

	"github.com/shopspring/decimal"
)

const INTEGER_OBJ = "INTEGER"

type Integer struct {
	Object
	Mappable
	Value decimal.Decimal
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

func (i *Integer) String() string { return i.Value.String() }

func (i *Integer) Method(method string, args []Object) (Object, bool) {
	//TODO implement me
	panic("implement me")
}

func (s *Integer) MapKey() MapKey {
	h := fnv.New64a()
	h.Write([]byte(fmt.Sprint(s.Value)))
	return MapKey{Type: s.Type(), Value: h.Sum64()}
}
