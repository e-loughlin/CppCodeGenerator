package cppcomponents_test

import (
	"testing"
)

func Test_InterfaceConstructorWillDeriveInterfaceNameFromFilepath(t *testing.T) {
	for _, tt := range Test_InterfaceConstructorWillDeriveInterfaceNameFromFilepath_Data {
		newInterface := NewInterface(tt.filepath)
		actualName := newInterface.name
		if tt.expectedName != actualName {
			t.Errorf("NewFunction(%v): name: EXPECTED: %v, ACTUAL: %v", tt.filepath, tt.expectedName, actualName)
		}
	}
}

var Test_InterfaceConstructorWillDeriveInterfaceNameFromFilepath_Data = []struct {
	filepath     string
	expectedName string
}{
	// No difference in relative path
	{"I_MyInterface1.h", "I_MyInterface1"},

	// UNIX
	{"c/ws/MyProject/I_MyInterface2.h", "I_MyInterface2"},

	// Windows
	{"C:\\ws\\MyProject\\I_MyInterface3.h", "I_MyInterface3"},
}
