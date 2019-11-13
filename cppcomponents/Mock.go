package cppcomponents

import (
	"strings"

	"github.com/emloughl/CppCodeGenerator/util/gmockgenrunner"
	"github.com/emloughl/CppCodeGenerator/util/configurations"
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
	return &m
}

// func (m Mock) getMockHelperFunctionDeclarations string {
	
// }

// Fields ... The fields within templates to be replaced.
func (m Mock) Fields() map[string]string {
	fields := make(map[string]string)
	fields["{{Mock.Name}}"] = m.Name
	fields["{{Mock.Header.FileName}}"] = m.HeaderFileName
	fields["{{GMockMacros}}"] = m.GMockMacros
	fields["{{Mock.InheritedInterface.Name}}"] = m.InheritedInterface.Name
	fields["{{Mock.InheritedInterface.FileName}}"] = m.InheritedInterface.FileName
	fields["{{Mock.HelperFunctions.Declarations}}"] = m.MockHelperFunctionDeclarations
	fields["{{Mock.HelperFunctions.Definition}}"] = m.MockHelperFunctionDefinitions
	return fields
}
