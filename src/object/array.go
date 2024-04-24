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
	case "pop":
		return list.pop(args)
	case "push":
		return list.push(args)
	case "shift":
		return list.shift(args)
	}
	return nil, false
}

func (list *Array) length(args []Object) (Object, bool) {
	return &Integer{Value: decimal.NewFromInt(int64(len(list.Elements)))}, true
}

func (list *Array) pop(args []Object) (Object, bool) {
	if len(list.Elements) > 0 {
		pop, elements := list.Elements[len(list.Elements)-1], list.Elements[:len(list.Elements)-1]
		list.Elements = elements
		return pop, true
	}

	return &Null{}, true
}

func (list *Array) shift(args []Object) (Object, bool) {
	if len(list.Elements) > 0 {
		shift, elements := list.Elements[0], list.Elements[1:]
		list.Elements = elements
		return shift, true
	}

	return &Null{}, true
}

func (list *Array) push(args []Object) (Object, bool) {
	length := len(list.Elements)
	newLength := length + 1

	newElements := make([]Object, newLength)
	copy(newElements, list.Elements)
	newElements[length] = args[0]

	list.Elements = newElements

	return &Integer{Value: decimal.NewFromInt(int64(newLength))}, true
}
