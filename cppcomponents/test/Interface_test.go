// Copyright (c) 2019 Evan Loughlin
//
// MIT License

package cppcomponents_test

import (
	"testing"

	"github.com/emloughl/CppCodeGenerator/cppcomponents"
)

func Test_InterfaceConstructorWillDeriveInterfaceNameFromFilepath(t *testing.T) {
	for _, tt := range Test_InterfaceConstructorWillDeriveInterfaceNameFromFilepath_Data {
		newInterface := cppcomponents.NewInterfaceFromExistingFile(tt.filepath)
		actualName := newInterface.Name
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
	// Note: path/filepath package depends on GOOS - (i.e. This test will fail if not run in Windows)
	// {"C:\\ws\\MyProject\\I_MyInterface3.h", "I_MyInterface3"},
}
