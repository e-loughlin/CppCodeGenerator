package cppcomponents

import (
	"fmt"
	"strings"
)

// Parameter ...
type Parameter struct {
	varType string
	varName string
}

// NewParameter .. Constructor
func NewParameter(rawParameterLine string) *Parameter {
	p := Parameter{}

	parameterSlice := strings.Split(rawParameterLine, " ")
	p.varType = strings.TrimSpace(strings.Join(parameterSlice[:len(parameterSlice)-1], " "))
	p.varName = strings.TrimSpace(parameterSlice[len(parameterSlice)-1])

	return &p
}

func (p Parameter) toString() string {
	return fmt.Sprintf("%v %v", p.varType, p.varName)
}
