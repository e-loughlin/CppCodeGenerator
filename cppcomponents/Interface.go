package cppcomponents

import (
	"path/filepath"
	"strings"
	"fmt"
)

// Interface ... Implements File
type Interface struct {
	Name      string
	Functions []Function
	Signals   []Function
	Includes  []string
}

// NewInterface ... Constructor
func NewInterface(filePath string) *Interface {

	interfaceLines := util.ReadLines(filePath)

	i := Interface{}
	filePath = strings.Replace(filePath, ":", "", -1)
	i.Name = strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
	return &i
}

func (cppInterface Interface) parseFunctions(contentLines string[]) []Function {
	for _, line := range contentLines {
		if(isVirtualFunctionLine(line)) {
			newFunction := NewFunction(line)
			cppInterface.Functions = append(cppInterface.Functions, newFunction)
		}
	}

}

func (line string) isVirtualFunctionLine bool {
	line = strings.Replace(line, " ", "", -1)
	return (strings.Contains(line, "virtual") && strings.Contains(line, "=0;"))
}

func (i Interface) generate() {

}
