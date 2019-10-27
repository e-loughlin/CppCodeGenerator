package cppcomponents

import (
	"path/filepath"
	"strings"
	"fmt"
	
	"github.com/emloughl/CppCodeGenerator/util"
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
	fmt.Println(i.Functions)
	i.parseFunctions(interfaceLines)
	fmt.Println(i.Functions)
	return &i
}

func (cppInterface Interface) parseFunctions(contentLines []string) {
	for _, line := range contentLines {
		if(isPureVirtualDefinition(line)) {
			newFunction := NewFunction(line)
			cppInterface.Functions = append(cppInterface.Functions, *newFunction)
		}
	}
}

func isPureVirtualDefinition(line string) bool {
	line = strings.Replace(line, " ", "", -1)
	return (strings.Contains(line, "virtual") && strings.Contains(line, ")=0;"))
}

