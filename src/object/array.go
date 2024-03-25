package object

import (
	"bytes"
	"strings"

	"github.com/shopspring/decimal"
)

const ARRAY_OBJ = "ARRAY"

type Array struct {
	Elements []Object
}

func (ao *Array) Type() ObjectType { return ARRAY_OBJ }

func (ao *Array) Inspect() string {
	var out bytes.Buffer
	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

func (list *Array) Method(method string, args []Object) (Object, bool) {
	switch method {
	case "length":
		return list.length(args)
	}
	return nil, false
}

func (list *Array) length(args []Object) (Object, bool) {
	return &Integer{Value: decimal.NewFromInt(int64(len(list.Elements)))}, true
}
