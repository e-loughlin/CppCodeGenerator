package main

import (
	"fmt"
	"strings"
)

// Function ...
type Function struct {
	name        string
	returnType  string
	parameters  []Parameter
	constString string
}

// NewFunction .. Constructor
func NewFunction(pureVirtualFunctionLine string) *Function {
	f := Function{}

	// Remove "virtual " from string
	pureVirtualFunctionLine = strings.TrimLeft(pureVirtualFunctionLine, "virtual ")

	// Parse function name and return type
	returnTypeAndName := strings.Split(pureVirtualFunctionLine, "(")[0]
	returnTypeAndNameSlice := strings.Split(returnTypeAndName, " ")

	f.name = returnTypeAndNameSlice[len(returnTypeAndNameSlice)-1]
	f.returnType = strings.Join(returnTypeAndNameSlice[:len(returnTypeAndNameSlice)-1], " ")

	// Parse parameter list
	rawParameters := strings.Split(strings.Split(pureVirtualFunctionLine, ")")[0], "(")[1]
	rawParametersSlice := strings.Split(rawParameters, ",")
	for _, rawParameterString := range rawParametersSlice {
		f.parameters = append(f.parameters, *NewParameter(rawParameterString))
	}

	// Parse function const-ness
	f.constString = ""
	if strings.Contains(strings.Split(pureVirtualFunctionLine, ")")[1], "const") {
		f.constString = " const"
	}

	return &f
}

/// TODO: Make \t resource configurable (3 spaces, 4 spaces?)
func (f Function) declaration() string {
	return fmt.Sprintf("\t%v %v(%v)%v override;", f.returnType, f.name, f.allParameters(), f.constString)
}

/// TODO: Allow for const function keyword
func (f Function) definition(classScope string) string {
	return fmt.Sprintf("%v %v::%v(%v)%v\n{\n}", f.returnType, classScope, f.name, f.allParameters(), f.constString)
}

func (f Function) allParameters() string {
	parametersString := ""
	separator := ""
	for i, p := range f.parameters {
		if i > 0 {
			separator = ", "
		}
		parametersString += fmt.Sprintf("%v%v", separator, p.toString())
	}
	return parametersString
}
