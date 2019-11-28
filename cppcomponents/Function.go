package cppcomponents

import (
	"fmt"
	"strings"

	"github.com/emloughl/CppCodeGenerator/util/configurations"
)

// Function ...
type Function struct {
	Name        string
	ReturnType  string
	Parameters  []Parameter
	ConstString string
}

// NewFunction .. Constructor
func NewFunction(pureVirtualFunctionLine string) *Function {
	f := Function{}

	// Remove "virtual " from string
	pureVirtualFunctionLine = strings.Split(pureVirtualFunctionLine, "virtual ")[1]

	// Parse function Name and return type
	returnTypeAndName := strings.Split(pureVirtualFunctionLine, "(")[0]
	returnTypeAndNameSlice := strings.Split(returnTypeAndName, " ")

	f.Name = returnTypeAndNameSlice[len(returnTypeAndNameSlice)-1]
	f.ReturnType = strings.Join(returnTypeAndNameSlice[:len(returnTypeAndNameSlice)-1], " ")

	// Parse parameter list
	rawParameters := strings.Split(strings.Split(pureVirtualFunctionLine, ")")[0], "(")[1]
	f.parseParameters(rawParameters)

	// Parse function const-ness
	if strings.Contains(strings.Split(pureVirtualFunctionLine, ")")[1], "const") {
		f.ConstString = " const"
	}

	return &f
}

func (f *Function) parseParameters(rawParametersString string) {
	var parameterStrings []string

	// Split into individual parameters of format "varType varName"
	if len(rawParametersString) > 0 {
		bracketCount := 0
		splitIndex := 0
		for pos, char := range rawParametersString {
			switch(char) {
			case '<':
				bracketCount++
			case '>':
				bracketCount--
			case ',':
				if bracketCount == 0 {
					extractedParameter := rawParametersString[splitIndex:pos + 1]
					extractedParameter = strings.TrimLeft(extractedParameter, ", ")
					extractedParameter = strings.TrimRight(extractedParameter, ", ")
					parameterStrings = append(parameterStrings, extractedParameter)
					splitIndex = pos + 1
				}
			}

			if pos == len(rawParametersString) - 1 {
				extractedParameter := rawParametersString[splitIndex:]
				extractedParameter = strings.TrimLeft(extractedParameter, ", ")
				extractedParameter = strings.TrimRight(extractedParameter, ", ")
				parameterStrings = append(parameterStrings, extractedParameter)
			}
		}

		for _, parameterString := range(parameterStrings) {
			f.Parameters = append(f.Parameters, *NewParameter(parameterString))
		}
	}
}

func (f Function) Declaration() string {
	return fmt.Sprintf("%v%v %v(%v)%v override;\n", configurations.Config.Syntax.Tab, f.ReturnType, f.Name, f.allParameters(), f.ConstString)
}

func (f Function) Definition(classScope string) string {
	return fmt.Sprintf("%v %v::%v(%v)%v\n{\n}\n\n", f.ReturnType, classScope, f.Name, f.allParameters(), f.ConstString)
}

func (f Function) allParameters() string {
	parametersString := ""
	separator := ""
	for i, p := range f.Parameters {
		if i > 0 {
			separator = ", "
		}
		parametersString += fmt.Sprintf("%v%v", separator, p.ToString())
	}
	return parametersString
}
