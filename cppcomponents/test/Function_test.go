package cppcomponents_test

import (
	"testing"
)

func Test_FunctionConstructor(t *testing.T) {
	for _, tt := range Test_FunctionConstructor_Data {
		newFunction := NewFunction(tt.pureVirtualFunctionLine)
		actualName := newFunction.name
		actualReturnType := newFunction.returnType
		actualParameters := newFunction.parameters
		actualConstString := newFunction.constString
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
	expectedParameters      []Parameter
	expectedConstString     string
}{
	// Void return type, no arguments
	{"virtual void doSomething() = 0;", "doSomething", "void", nil, ""},

	// QString return type, single argument
	{"virtual QString name(int id) = 0", "name", "QString", []Parameter{{varType: "int", varName: "id"}}, ""},

	// QString return type, two arguments
	{"virtual QString name(int id, QString department) = 0", "name", "QString", []Parameter{{varType: "int", varName: "id"}, {varType: "QString", varName: "department"}}, ""},

	// Const function
	{"virtual QString name(int id) const = 0", "name", "QString", []Parameter{{varType: "int", varName: "id"}}, " const"},
}

// Helper functions
func testEqualParameterSlices(a, b []Parameter) bool {

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
