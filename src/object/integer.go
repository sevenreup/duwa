package object

import (
	"fmt"
	"hash/fnv"

	"github.com/shopspring/decimal"
)

const INTEGER_OBJ = "INTEGER"

// type=Nambala alternative=Number
// The Integer object represents a number.
type Integer struct {
	Object
	Mappable
	Value decimal.Decimal
}

func (int *Integer) Type() ObjectType { return INTEGER_OBJ }

func (int *Integer) String() string { return int.Value.String() }

func (int *Integer) Method(method string, args []Object) (Object, bool) {
	switch method {
	case "kuMawu":
		return int.methodToString(args)
	}

	return nil, false
}

func (int *Integer) MapKey() MapKey {
	h := fnv.New64a()
	h.Write([]byte(fmt.Sprint(int.Value)))
	return MapKey{Type: int.Type(), Value: h.Sum64()}
}

func (int *Integer) methodToString(_ []Object) (Object, bool) {
	return &String{Value: int.Value.String()}, true
}
