package cppcomponents

import (
	"path/filepath"
	"strings"
	"os"
	"fmt"
	
	"github.com/emloughl/CppCodeGenerator/util"
	"github.com/emloughl/CppCodeGenerator/configurations"
	"github.com/fatih/camelcase"
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

// NewInterface ... Constructor taking in a filepath to an existing interface.
func NewInterface(filePath string) *Interface {
	var interfaceLines []string

	if(util.FileExists(filePath)) {
		interfaceLines = util.ReadLines(filePath)
	}
	if(!isValidInterfacePath(filePath)) {
		fmt.Println("Error: Interface does not have correct extension, prefix, or suffix. Refer to config.json for accepted formats.")
		os.Exit(0)
	}
	
	i := Interface{}
	filePath = strings.Replace(filePath, ":", "", -1)
	i.Name = strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
	i.parseFunctions(interfaceLines)
	i.DefineName = i.parseDefineName(i.Name)
	i.FileName = i.parseFileName(i.Name)
	return &i
}

// parseFunctions ... Reads a slice of lines and parses Function structs from it.
func (i Interface) parseFunctions(contentLines []string) {
	for _, line := range contentLines {
		if(isPureVirtualDefinition(line)) {
			newFunction := NewFunction(line)
			i.Functions = append(i.Functions, *newFunction)
		}
	}
}

// parseDefineName ... 
func (i Interface) parseDefineName(name string) string {

	name = strings.Replace(name, "_", "", -1)
	splitName := camelcase.Split(name)
	name = strings.Join(splitName, configurations.Config.Policies.DefineNameCamelCaseSeparator)
	name = strings.ToUpper(name)

	defineName := configurations.Config.Affixes.Prefixes.DefineName +
	name +
	configurations.Config.Affixes.Suffixes.DefineName
	return defineName
}

// parseFileName ...
func (i Interface) parseFileName(name string) string {
	fileName :=  i.Name + configurations.Config.FileExtensions.CppHeader
	return fileName
}

// isPureVirtualDefinition ... Returns whether a function is pure virtual.
func isPureVirtualDefinition(line string) bool {
	line = strings.Replace(line, " ", "", -1)
	return (strings.Contains(line, "virtual") && strings.Contains(line, ")=0;"))
}

// Fields ... The fields within templates to be replaced.
func (i Interface) Fields() map[string]string {
	fields := make(map[string]string)
	fields["{{Interface.Name}}"] = i.Name
	fields["{{Interface.FileName}}"] = i.FileName
	fields["{{Interface.DefineName}}"] = i.DefineName
	return fields
}

// isValidInterfacePath ...
func isValidInterfacePath(filePath string) bool {
	filePath = strings.Replace(filePath, ":", "", -1)
	hasCorrectExtension := (filepath.Ext(filePath) == ".h")
	fileName := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
	hasCorrectPrefix := strings.HasPrefix(fileName, configurations.Config.Affixes.Prefixes.Interface)
	hasCorrectSuffix := strings.HasSuffix(fileName, configurations.Config.Affixes.Suffixes.Interface)

	isValid := hasCorrectExtension && hasCorrectPrefix && hasCorrectSuffix
	return isValid
}