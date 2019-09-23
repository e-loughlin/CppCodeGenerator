package main

import (
	"fmt"
	"strings"
)

// Function ...
type Function struct {
	name       string
	returnType string
	parameters []Parameter
}

// NewFunction .. Constructor
func NewFunction(rawFunctionLine string) *Function {
	f := Function{}

	returnTypeAndName := strings.Split(rawFunctionLine, "(")[0]
	returnTypeAndNameSlice := strings.Split(returnTypeAndName, " ")

	f.name = returnTypeAndNameSlice[len(returnTypeAndNameSlice)-1]
	f.returnType = strings.Join(returnTypeAndNameSlice[:len(returnTypeAndNameSlice)-1], " ")

	rawParameters := strings.Split(strings.Split(rawFunctionLine, ")")[0], "(")[1]

	rawParametersSlice := strings.Split(rawParameters, ",")
	for _, rawParameterString := range rawParametersSlice {
		f.parameters = append(f.parameters, *NewParameter(rawParameterString))
	}

	return &f
}

func (f Function) declaration() string {
	parametersString := ""
	separator := ""
	for i, p := range f.parameters {
		if i > 0 {
			separator = ", "
		}
		parametersString += fmt.Sprintf("%v%v", separator, p.toString())
	}
	return fmt.Sprintf("\t%v %v(%v);", f.returnType, f.name, parametersString)
}
