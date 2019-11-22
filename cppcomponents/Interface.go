package cppcomponents

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/emloughl/CppCodeGenerator/util/configurations"
	"github.com/emloughl/CppCodeGenerator/util/io"
	"github.com/emloughl/CppCodeGenerator/util/parsers"
	"github.com/emloughl/CppCodeGenerator/util/slice"
)

// Interface ... Implements File
type Interface struct {
	Name         string
	FilePath     string
	FileName     string
	DefineName   string
	Functions    []Function
	Signals      []Function
	Dependencies []string
}

// NewInterface ...
func NewInterface(filePath string) *Interface {
	var interfaceLines []string

	if io.FileExists(filePath) {
		interfaceLines = io.ReadLines(filePath)
	}
	if !IsValidInterfaceFilePath(filePath) {
		fmt.Println("Error: Interface does not have correct extension, prefix, or suffix. Refer to config.json for accepted formats.")
		os.Exit(0)
	}

	i := Interface{}
	i.FilePath = filePath
	filePath = strings.Replace(filePath, ":", "", -1)
	i.Name = strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
	i.Functions = i.parseFunctions(interfaceLines)
	i.DefineName = parsers.GenerateDefineName(i.Name)
	i.FileName = i.parseFileName(i.Name)
	i.parseDependencies()
	fmt.Println(i.Dependencies)
	return &i
}

// parseFunctions ... Reads a slice of lines and parses Function structs from it.
func (i Interface) parseFunctions(contentLines []string) []Function {
	var functions []Function
	for _, line := range contentLines {
		if isPureVirtualDefinition(line) {
			newFunction := NewFunction(line)
			functions = append(functions, *newFunction)
		}
	}
	return functions
}

// parseFileName ...
func (i Interface) parseFileName(name string) string {
	fileName := i.Name + configurations.Config.FileExtensions.CppHeader
	return fileName
}

// parseDependencies ...
func (i *Interface) parseDependencies() {
	var dependencies []string
	for _, function := range i.Functions {
		dependencies = append(dependencies, function.ReturnType)
		for _, parameter := range function.Parameters {
			dependencies = append(dependencies, parameter.VarType)
		}
	}
	i.Dependencies = slice.RemoveDuplicates(dependencies)
	i.Dependencies = slice.RemoveStdDataTypes(dependencies)
	sort.Strings(i.Dependencies)
}

// isPureVirtualDefinition ... Returns whether a function is pure virtual.
func isPureVirtualDefinition(line string) bool {
	line = strings.Replace(line, " ", "", -1)
	return (strings.Contains(line, "virtual") && strings.Contains(line, ")=0;"))
}

// isValidInterfacePath ...
func IsValidInterfaceFilePath(filePath string) bool {
	filePath = strings.Replace(filePath, ":", "", -1)
	hasCorrectExtension := (filepath.Ext(filePath) == ".h")
	fileName := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
	hasCorrectPrefix := strings.HasPrefix(fileName, configurations.Config.Prefixes.Interface)
	hasCorrectSuffix := strings.HasSuffix(fileName, configurations.Config.Suffixes.Interface)

	isValid := hasCorrectExtension && hasCorrectPrefix && hasCorrectSuffix
	return isValid
}

// Fields ... The fields within templates to be replaced.
func (i Interface) Fields() map[string]string {
	fields := make(map[string]string)
	fields["{{Interface.Name}}"] = i.Name
	fields["{{FileName}}"] = i.FileName
	fields["{{Interface.DefineName}}"] = i.DefineName
	return fields
}
