package cppcomponents

import (
	"strings"
	"path/filepath"

	"github.com/emloughl/CppCodeGenerator/util/gmockgenrunner"
	"github.com/emloughl/CppCodeGenerator/util/io"
	"github.com/emloughl/CppCodeGenerator/util/paths"
	"github.com/emloughl/CppCodeGenerator/util/configurations"
	"github.com/emloughl/CppCodeGenerator/util/fieldreplacer"
)

// Test ... Implements File
type Mock struct {
	InheritedInterface Interface
	Name          			string
	HeaderFileName          string
	ImplementationFileName  string
	GMockMacros 			string
	MockHelperFunctionDeclarations string
	MockHelperFunctionDefinitions string
}

func NewMock(inheritedInterface Interface) *Mock {
	m := Mock{}
	m.InheritedInterface = inheritedInterface
	m.Name = strings.TrimPrefix(m.InheritedInterface.Name, configurations.Config.Prefixes.Interface)
	m.Name = strings.TrimSuffix(m.Name, configurations.Config.Suffixes.Interface)
	m.Name = configurations.Config.Prefixes.Mock + m.Name + configurations.Config.Suffixes.Mock
	m.HeaderFileName = m.Name + configurations.Config.FileExtensions.CppHeader
	m.ImplementationFileName = m.Name + configurations.Config.FileExtensions.CppImplementation
	m.GMockMacros = gmockgenrunner.GetGMockGeneratorFunctionRegistrations(m.InheritedInterface.FileName)
	m.setMockHelperFunctions()
	return &m
}

func (m *Mock) setMockHelperFunctions() {
	var declarations string
	var definitions string

	fields := make(map[string]string)
	fields["{{Mock.Name}}"] = m.Name

	for _, function := range (m.InheritedInterface.Functions) {
		fields["{{Function.Name}}"] = function.Name
		fields["{{Function.UppercaseName}}"] = strings.Title(function.Name)
		fields["{{Function.ReturnType}}"] = function.ReturnType
		
		for _, declarationTemplatePath := range (paths.MockHelperFunctionDeclarationPaths) {
			// Hacky code: Prevent "makeFunctionReturn()" from being generated if function return type is void. TODO: Find a cleaner way.
			if(function.ReturnType == "void" && filepath.Base(declarationTemplatePath) == "mock_MakeFunctionReturn_declaration.txt") {
				continue
			}
			declarationTemplate := io.ReadContents(declarationTemplatePath)
			parsedDeclarations := fieldreplacer.ReplaceAllFields(declarationTemplate, fields)
			declarations += parsedDeclarations
		}
		declarations += "\n\n"

		for _, definitionTemplatePath := range (paths.MockHelperFunctionDefinitionPaths) {
			// Hacky code: Prevent "makeFunctionReturn()" from being generated if function return type is void. TODO: Find a cleaner way.
			if(function.ReturnType == "void" && filepath.Base(definitionTemplatePath) == "mock_MakeFunctionReturn_definition.txt") {
				continue
			}
			definitionTemplate := io.ReadContents(definitionTemplatePath)
			parsedDefinitions := fieldreplacer.ReplaceAllFields(definitionTemplate, fields)
			definitions += parsedDefinitions
		}
	}
	m.MockHelperFunctionDeclarations = declarations
	m.MockHelperFunctionDefinitions = definitions
}

// Fields ... The fields within templates to be replaced.
func (m Mock) Fields() map[string]string {
	fields := make(map[string]string)
	fields["{{Mock.Name}}"] = m.Name
	fields["{{Mock.Header.FileName}}"] = m.HeaderFileName
	fields["{{GMockMacros}}"] = m.GMockMacros
	fields["{{Mock.InheritedInterface.Name}}"] = m.InheritedInterface.Name
	fields["{{Mock.InheritedInterface.FileName}}"] = m.InheritedInterface.FileName
	fields["{{Mock.HelperFunctions.Declarations}}"] = m.MockHelperFunctionDeclarations
	fields["{{Mock.HelperFunctions.Definitions}}"] = m.MockHelperFunctionDefinitions
	return fields
}
