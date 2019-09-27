package main

import (
	"testing"
)

func Test_constructor(t *testing.T) {
	for _, tt := range Test_constructor_Data {
		newFunction := NewFunction(tt.pureVirtualFunctionLine)
		actualName := newFunction.name
		actualReturnType := newFunction.returnType
		actualParameters := newFunction.parameters
		actualConstString := newFunction.constString
		if tt.expectedName != actualName {
			t.Errorf("NewFunction(%v): EXPECTED: %v, ACTUAL: %v", tt.pureVirtualFunctionLine, tt.expectedName, actualName)
		}
		if tt.expectedReturnType != actualReturnType {
			t.Errorf("NewFunction(%v): EXPECTED: %v, ACTUAL: %v", tt.pureVirtualFunctionLine, tt.expectedReturnType, actualReturnType)
		}
		if !testEqualParameterSlices(tt.expectedParameters, actualParameters) {
			t.Errorf("NewFunction(%v): EXPECTED: %v, ACTUAL: %v", tt.pureVirtualFunctionLine, tt.expectedParameters, actualParameters)
		}
		if tt.expectedConstString != actualConstString {
			t.Errorf("NewFunction(%v): EXPECTED: %v, ACTUAL: %v", tt.pureVirtualFunctionLine, tt.expectedConstString, actualConstString)
		}
	}
}

var Test_constructor_Data = []struct {
	pureVirtualFunctionLine string
	expectedName            string
	expectedReturnType      string
	expectedParameters      []Parameter
	expectedConstString     string
}{
	{"virtual void doSomething() = 0;", "doSomething", "void", []Parameter{}, ""},
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
