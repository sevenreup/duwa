package object

import (
	"fmt"
	"github.com/shopspring/decimal"
)

const INTEGER_OBJ = "INTEGER"

type Integer struct {
	Value decimal.Decimal
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

func (i *Integer) Method(method string, args []Object) (Object, bool) {
	//TODO implement me
	panic("implement me")
}
