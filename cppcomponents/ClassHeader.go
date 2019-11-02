package cppcomponents

import (
	"strings"

	// "github.com/emloughl/CppCodeGenerator/util"
	"github.com/emloughl/CppCodeGenerator/configurations"
)

// ClassHeader ... Implements File
type ClassHeader struct {
	InheritedInterface Interface
	Name                 string
	FileName             string
	DefineName		     string
	ForwardDeclares	     string
	FunctionDeclarations string
	QtSignalDeclarations string
}

func NewClassHeader(InheritedInterface Interface) *ClassHeader {
	c := ClassHeader{}
	c.InheritedInterface = InheritedInterface
	c.Name = strings.TrimPrefix(c.InheritedInterface.Name, configurations.Config.Prefixes.Interface)
	c.Name = strings.TrimSuffix(c.Name, configurations.Config.Suffixes.Interface)
	c.FileName = c.Name + configurations.Config.FileExtensions.CppHeader

	return &c
}

// Fields ... The fields within templates to be replaced.
func (c ClassHeader) Fields() map[string]string {
	fields := make(map[string]string)
	fields["{{Interface.FileName}}"] = c.InheritedInterface.FileName
	fields["{{FileName}}"] = c.FileName
	fields["{{Class.Name}}"] = c.Name
	fields["{{Class.Header.DefineName}}"] = c.DefineName
	fields["{{Class.Header.ForwardDeclares}}"] = c.ForwardDeclares
	fields["{{Class.Header.FunctionDeclarations}}"] = c.FunctionDeclarations
	fields["{{Class.Header.QtSignalDeclarations}}"] = c.QtSignalDeclarations
	return fields
}