package main

import (
	"fmt"
	"testing"
)

func Test_NewInclude(t *testing.T) {
	fmt.Println("Hello")
}

var Test_deriveDependency_TypeContainsPointers_WillRemovePointers_Data = []struct {
	input          string
	expectedOutput string
}{
	{"unsigned int *", "unsigned int"},
	{"char **", "char"},
	{"unsigned char", "unsigned char"},
}

func Test_deriveDependency_TypeContainsPointers_WillRemovePointers(t *testing.T) {
	for _, tt := range Test_deriveDependency_TypeContainsPointers_WillRemovePointers_Data {
		actualOutput := deriveDependency(tt.input)
		fmt.Println(actualOutput)
		if tt.expectedOutput != actualOutput {
			t.Errorf("DerivedDependency(%v): EXPECTED: %v, ACTUAL: %v", tt.input, tt.expectedOutput, actualOutput)
		}
	}
}
