package cppcomponents

import (
	"path/filepath"
	"strings"
	"fmt"
	
	"github.com/emloughl/CppCodeGenerator/util"
	"github.com/emloughl/CppCodeGenerator/configurations"
	// "github.com/fatih/camelcase"
)

// Interface ... Implements File
type Interface struct {
	Name      string
	FileName string
	DefineName string
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
	i.parseDefineName(i.Name)
	fmt.Println(i.Functions)
	return &i
}

// parseFunctions ... Reads a slice of lines and parses Function structs from it.
func (cppInterface Interface) parseFunctions(contentLines []string) {
	for _, line := range contentLines {
		if(isPureVirtualDefinition(line)) {
			newFunction := NewFunction(line)
			cppInterface.Functions = append(cppInterface.Functions, *newFunction)
		}
	}
}

// parseDefineName ... 
func (cppInterface Interface) parseDefineName(name string) {
	cppInterface.DefineName = configurations.Config.Affixes.Prefixes.DefineName +
							  cppInterface.DefineName +
							  configurations.Config.Affixes.Suffixes.DefineName
	
	fmt.Println(cppInterface.DefineName)
}

// isPureVirtualDefinition ... Returns whether a function is pure virtual.
func isPureVirtualDefinition(line string) bool {
	line = strings.Replace(line, " ", "", -1)
	return (strings.Contains(line, "virtual") && strings.Contains(line, ")=0;"))
}


// Fields ... The fields within templates to be replaced.
func (cppInterface Interface) Fields() map[string]string {
	fields := make(map[string]string)
	fields["Interface.Name"] = cppInterface.Name
	fields["Interface.FileName"] = cppInterface.FileName
	fields["Interface.DefineName"] = cppInterface.DefineName
	return fields
}