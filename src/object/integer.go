package object

import (
	"github.com/shopspring/decimal"
)

const INTEGER_OBJ = "INTEGER"

type Integer struct {
	Value decimal.Decimal
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

func (i *Integer) Inspect() string { return i.Value.String() }

func (i *Integer) Method(method string, args []Object) (Object, bool) {
	//TODO implement me
	panic("implement me")
}
