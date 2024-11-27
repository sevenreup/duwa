package object

import (
	"bytes"
	"strings"

	"github.com/shopspring/decimal"
)

const ARRAY_OBJ = "ARRAY"

// type=array alternative=Array
// The Array object represents a list of elements.
type Array struct {
	Elements []Object
	Object
}

func (ao *Array) Type() ObjectType { return ARRAY_OBJ }

func (ao *Array) String() string {
	var out bytes.Buffer
	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.String())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

func (list *Array) Method(method string, args []Object) (Object, bool) {
	switch method {
	case "kutalika":
		return list.methodLength(args)
	case "chotsaKumbuyo":
		return list.methodPop(args)
	case "Kankha":
		return list.methodPush(args)
	case "chotsaKutsogolo":
		return list.methodShift(args)
	}
	return nil, false
}

// method=kutalika args=[] return={number}
// Returns the number of elements in an array.
func (list *Array) methodLength(_ []Object) (Object, bool) {
	return &Integer{Value: decimal.NewFromInt(int64(len(list.Elements)))}, true
}

// method=chotsaKumbuyo args=[] return={any}
// Removes the last element from an array and returns that element. This method changes the length of the array.
func (list *Array) methodPop(_ []Object) (Object, bool) {
	if len(list.Elements) > 0 {
		pop, elements := list.Elements[len(list.Elements)-1], list.Elements[:len(list.Elements)-1]
		list.Elements = elements
		return pop, true
	}

	return &Null{}, true
}

// method=chotsaKutsogolo args=[] return={any}
// Removes the first element from an array and returns that element. This method changes the length of the array.
func (list *Array) methodShift(_ []Object) (Object, bool) {
	if len(list.Elements) > 0 {
		shift, elements := list.Elements[0], list.Elements[1:]
		list.Elements = elements
		return shift, true
	}

	return &Null{}, true
}

// method=Kankha args=[any] return={number}
// Adds one or more elements to the end of an array and returns the new length of the array.
func (list *Array) methodPush(args []Object) (Object, bool) {
	length := len(list.Elements)
	newLength := length + 1

	newElements := make([]Object, newLength)
	copy(newElements, list.Elements)
	newElements[length] = args[0]

	list.Elements = newElements

	return &Integer{Value: decimal.NewFromInt(int64(newLength))}, true
}
