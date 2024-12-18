package evaluator

import (
	"github.com/sevenreup/duwa/src/ast"
	"github.com/sevenreup/duwa/src/object"
	"github.com/sevenreup/duwa/src/values"
)

func evaluateNull(node *ast.NullLiteral, env *object.Environment) object.Object {
	return values.NULL
}
