package cppcomponents

import (
	"fmt"
	"strings"
)

// Parameter ...
type Parameter struct {
	VarType string
	VarName string
}

// NewParameter .. Constructor
func NewParameter(rawParameterLine string) *Parameter {
	p := Parameter{}

	parameterSlice := strings.Split(rawParameterLine, " ")
	p.VarType = strings.TrimSpace(strings.Join(parameterSlice[:len(parameterSlice)-1], " "))
	p.VarName = strings.TrimSpace(parameterSlice[len(parameterSlice)-1])

	return &p
}

func (p Parameter) toString() string {
	return fmt.Sprintf("%v %v", p.VarType, p.VarName)
}
