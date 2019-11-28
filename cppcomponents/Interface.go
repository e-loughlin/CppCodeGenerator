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

// parseDependencies ... The term "dependency" is used here to refer to any
// data type that may require an include or forward declare. 
func (i *Interface) parseDependencies() {
	var dependencies []string
	for _, function := range i.Functions {

		// "expanded" refers to creating a slice from a templated type, i.e "QMap <int, QString>" becomes [QMap int QString]
		expandedReturnType := strings.FieldsFunc(function.ReturnType, templatedTypeSeparators) 
		for _, dataType := range(expandedReturnType) {
			dependencies = append(dependencies, strings.TrimSpace(dataType))
		}

		for _, parameter := range function.Parameters {
			expandedParameter := strings.FieldsFunc(parameter.Type, templatedTypeSeparators)
			for _, innerParameter := range expandedParameter {
				dependencies = append(dependencies, strings.TrimSpace(innerParameter))
			} 
		}
	}
	i.Dependencies = dependencies
	i.Dependencies = slice.RemoveDuplicates(i.Dependencies)
	i.Dependencies = slice.RemoveStdDataTypes(i.Dependencies)
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

// templatedTypeSeparators ... Used to expand templated types such as QMap<QString, QMap<QString, std::string>>
func templatedTypeSeparators (r rune) bool {
	return r == '<' || r == '>' || r == ','
}