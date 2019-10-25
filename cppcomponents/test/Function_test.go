package cppcomponents_test

import (
	"testing"

	"github.com/emloughl/CppCodeGenerator/cppcomponents"
)

func Test_FunctionConstructor(t *testing.T) {
	for _, tt := range Test_FunctionConstructor_Data {
		newFunction := cppcomponents.NewFunction(tt.pureVirtualFunctionLine)
		actualName := newFunction.Name
		actualReturnType := newFunction.ReturnType
		actualParameters := newFunction.Parameters
		actualConstString := newFunction.ConstString
		if tt.expectedName != actualName {
			t.Errorf("NewFunction(%v): name: EXPECTED: %v, ACTUAL: %v", tt.pureVirtualFunctionLine, tt.expectedName, actualName)
		}
		if tt.expectedReturnType != actualReturnType {
			t.Errorf("NewFunction(%v): returnType: EXPECTED: %v, ACTUAL: %v", tt.pureVirtualFunctionLine, tt.expectedReturnType, actualReturnType)
		}
		if !testEqualParameterSlices(tt.expectedParameters, actualParameters) {
			t.Errorf("NewFunction(%v): parameters: EXPECTED: %v, ACTUAL: %v", tt.pureVirtualFunctionLine, tt.expectedParameters, actualParameters)
		}
		if tt.expectedConstString != actualConstString {
			t.Errorf("NewFunction(%v): constString: EXPECTED: %v, ACTUAL: %v", tt.pureVirtualFunctionLine, tt.expectedConstString, actualConstString)
		}
	}
}

var Test_FunctionConstructor_Data = []struct {
	pureVirtualFunctionLine string
	expectedName            string
	expectedReturnType      string
	expectedParameters      []cppcomponents.Parameter
	expectedConstString     string
}{
	// Void return type, no arguments
	{"virtual void doSomething() = 0;", "doSomething", "void", nil, ""},

	// QString return type, single argument
	{"virtual QString name(int id) = 0", "name", "QString", []cppcomponents.Parameter{{VarType: "int", VarName: "id"}}, ""},

	// QString return type, two arguments
	{"virtual QString name(int id, QString department) = 0", "name", "QString", []cppcomponents.Parameter{{VarType: "int", VarName: "id"}, {VarType: "QString", VarName: "department"}}, ""},

	// Const function
	{"virtual QString name(int id) const = 0", "name", "QString", []cppcomponents.Parameter{{VarType: "int", VarName: "id"}}, " const"},
}

// Helper functions
func testEqualParameterSlices(a, b []cppcomponents.Parameter) bool {

	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
