package evaluator

import (
	"github.com/sevenreup/chewa/src/lexer"
	"github.com/sevenreup/chewa/src/object"
	"github.com/sevenreup/chewa/src/parser"
	"testing"
)

func testEval(input string) object.Object {
	l := lexer.New([]byte(input))
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	return Eval(program, env)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d",
			result.Value, expected)
		return false
	}
	return true
}

func testStringObject(t *testing.T, obj object.Object, expected string) bool {
	result, ok := obj.(*object.String)
	if !ok {
		t.Errorf("object is not String. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%s, want=%s",
			result.Value, expected)
		return false
	}
	return true
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t",
			result.Value, expected)
		return false
	}
	return true
}

func testLiteralExpression(
	t *testing.T,
	obj object.Object,
	expected interface{},
) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerObject(t, obj, int64(v))
	case int64:
		return testIntegerObject(t, obj, v)
	case bool:
		return testBooleanObject(t, obj, v)
	case string:
		return testStringObject(t, obj, v)
	}
	t.Errorf("type of exp not handled. got=%T", expected)
	return false
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)
		return false
	}
	return true
}

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5- 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 +-50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 *-10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 +-10", 50},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"zoona", true},
		{"bodza", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"zoona == zoona", true},
		{"bodza == bodza", true},
		{"zoona == bodza", false},
		{"zoona != bodza", true},
		{"bodza != zoona", true},
		{"(1 < 2) == zoona", true},
		{"(1 < 2) == bodza", false},
		{"(1 > 2) == zoona", false},
		{"(1 > 2) == bodza", true},
		{`("foo" == "bar") == bodza`, true},
		{`("foo" == "foo") == zoona`, true},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!zoona", false},
		{"!bodza", true},
		{"!5", false},
		{"!!zoona", true},
		{"!!bodza", false},
		{"!!5", true}}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"ngati (zoona) { 10 }", 10},
		{"ngati (bodza) { 10 }", nil},
		{"ngati (1) { 10 }", 10},
		{"ngati (1 < 2) { 10 }", 10},
		{"ngati (1 > 2) { 10 }", nil},
		{"ngati (1 > 2) { 10 } kapena { 20 }", 20},
		{"ngati (1 < 2) { 10 } kapena { 20 }", 10},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"bweza 10;", 10},
		{"bweza 10; 9;", 10},
		{"bweza 2 * 5; 9;", 10},
		{"9; bweza 2 * 5; 9;", 10},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"5 + zoona;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"5 + zoona; 5;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"-zoona",
			"unknown operator:-BOOLEAN",
		},
		{
			"zoona + bodza;",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"5; zoona + bodza; 5",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"ngati (10 > 1) { zoona + bodza; }",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			`
					ngati (10 > 1) {
 						ngati (10 > 1) {
 							bweza zoona + bodza;
 						}
 						bweza 1;
 					}
				`,
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"foobar",
			"identifier not found: foobar",
		},
		{
			`"Hello"- "World"`,
			"unknown operator: STRING - STRING",
		},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)",
				evaluated, evaluated)
			continue
		}
		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q",
				tt.expectedMessage, errObj.Message)
		}
	}
}

func TestAssignmentStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"nambala a = 5; a;", 5},
		{"nambala a = 5 * 5; a;", 25},
		{"nambala a = 5; nambala b = a; b;", 5},
		{"nambala a = 5; nambala b = a; nambala c = a + b + 5; c;", 15},
		{`mawu a = "b"; a;`, "b"},
		{`mawu a = "5"; mawu b = a; mawu c = a + b + "5"; c;`, "555"},
	}
	for _, tt := range tests {
		testLiteralExpression(t, testEval(tt.input), tt.expected)
	}
}

func TestFunctionObject(t *testing.T) {
	input := "ndondomeko phatikizaZiwiri(x) { x + 2; };"
	evaluated := testEval(input)
	fn, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("object is not Function. got=%T (%+v)", evaluated, evaluated)
	}
	if len(fn.Parameters) != 1 {
		t.Fatalf("function has wrong parameters. Parameters=%+v",
			fn.Parameters)
	}
	if fn.Parameters[0].String() != "x" {
		t.Fatalf("parameter is not 'x'. got=%q", fn.Parameters[0])
	}
	expectedBody := "(x + 2)"
	if fn.Body.String() != expectedBody {
		t.Fatalf("body is not %q. got=%q", expectedBody, fn.Body.String())
	}
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"ndondomeko identity(x) { x; }; identity(5);", 5},
		{"ndondomeko  identity(x) { bweza x; }; identity(5);", 5},
		{"ndondomeko double(x){ x * 2; }; double(5);", 10},
		{"ndondomeko  add(x, y) { x + y; }; add(5, 5);", 10},
		{"ndondomeko add(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
		{"ndondomeko zina(x) { x; }(5)", 5},
	}
	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestClosures(t *testing.T) {
	input := `
 	ndondomeko newAdder(x) {
 		ndondomeko temp(y) { x + y };
 	};
 	nambala addTwo = newAdder(2);
 	addTwo(2);`
	testIntegerObject(t, testEval(input), 4)
}

func TestStringLiteral(t *testing.T) {
	input := `"Hello World!"`
	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}
	if str.Value != "Hello World!" {
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}

func TestStringConcatenation(t *testing.T) {
	input := `"Hello" + " " + "World!"`
	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}
	if str.Value != "Hello World!" {
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}
