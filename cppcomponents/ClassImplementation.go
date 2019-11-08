package cppcomponents

import (
	"strings"
	"github.com/emloughl/CppCodeGenerator/configurations"
)

// ClassImplementation ... Implements File
type ClassImplementation struct {
	InheritedInterface Interface
	Name                string
	FileName		    string
	HeaderFileName      string
	Includes			string 
	FunctionDefinitions string
	QtSignalDefinitions string
}

func NewClassImplementation(InheritedInterface Interface) *ClassImplementation {
	c := ClassImplementation{}
	c.InheritedInterface = InheritedInterface
	c.Name = strings.TrimPrefix(c.InheritedInterface.Name, configurations.Config.Prefixes.Interface)
	c.Name = strings.TrimSuffix(c.Name, configurations.Config.Suffixes.Interface)
	c.HeaderFileName = c.Name + configurations.Config.FileExtensions.CppHeader
	c.FileName = c.Name + configurations.Config.FileExtensions.CppImplementation

	return &c
}


// Fields ... The fields within templates to be replaced.
func (c ClassImplementation) Fields() map[string]string {
	fields := make(map[string]string)
	fields["{{Class.Header.FileName}}"] = c.HeaderFileName
	fields["{{FileName}}"] = c.FileName
	fields["{{Class.Name}}"] = c.Name
	fields["{{Class.Implementation.Includes}}"] = c.Includes
	fields["{{Class.Implementation.FunctionDefinitions}}"] = c.FunctionDefinitions
	fields["{{Class.Implementation.QtSignalDefinitions}}"] = c.QtSignalDefinitions
	return fields
}