package cppcomponents

import (
	"strings"
	"fmt"

	"github.com/emloughl/CppCodeGenerator/util"
	"github.com/emloughl/CppCodeGenerator/util/configurations"
)

// Test ... Implements File
type Mock struct {
	InheritedInterface Interface
	Name          			string
	FileName          		string
	GMockMacros 			string
}

func NewMock(inheritedInterface Interface) *Mock {
	m := Mock{}
	m.InheritedInterface = inheritedInterface
	m.Name = strings.TrimPrefix(m.InheritedInterface.Name, configurations.Config.Prefixes.Interface)
	m.Name = strings.TrimSuffix(m.Name, configurations.Config.Suffixes.Interface)
	m.Name = configurations.Config.Prefixes.Mock + m.Name + configurations.Config.Suffixes.Mock
	m.FileName = m.Name + configurations.Config.FileExtensions.CppHeader
	m.GMockMacros = util.GetGMockGeneratorFunctionRegistrations(m.InheritedInterface.FileName)
	fmt.Println(m.GMockMacros)
	return &m
}

// Fields ... The fields within templates to be replaced.
func (m Mock) Fields() map[string]string {
	fields := make(map[string]string)
	fields["{{GMockMacros}}"] = m.GMockMacros
	return fields
}
