package object

import (
	"bytes"
	"fmt"
	"strings"
)

const MAP_OBJ = "MAP"

// type=mgwirizano altenatives=map
// This is a map object that is used to store key-value pairs.
type Map struct {
	Pairs map[MapKey]MapPair
	Object
}

type MapPair struct {
	Key   Object
	Value Object
}

func (m *Map) Type() ObjectType { return MAP_OBJ }

func (m *Map) String() string {
	var out bytes.Buffer
	pairs := []string{}
	for _, pair := range m.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s",
			pair.Key.String(), pair.Value.String()))
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}

func (m *Map) Method(method string, args []Object) (Object, bool) {
	//TODO implement me
	panic("implement me")
}
