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
	return Eval(program)
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
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
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
			"if (10 > 1) { zoona + bodza; }",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			`
					if (10 > 1) {
 						if (10 > 1) {
 							return zoona + bodza;
 						}
 						bweza 1;
 					}
				`,
			"unknown operator: BOOLEAN + BOOLEAN",
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
