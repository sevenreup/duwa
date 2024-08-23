package object

import (
	"bytes"
	"github.com/sevenreup/duwa/src/ast"
	"strings"
)

const FUNCTION_OBJ = "FUNCTION"

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }

func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("ndondomeko")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}

func (i *Function) Method(method string, args []Object) (Object, bool) {
	//TODO implement me
	panic("implement me")
}
