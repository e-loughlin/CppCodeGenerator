package cppcomponents

import (
	"github.com/emloughl/CppCodeGenerator/util/configurations"
	"github.com/emloughl/CppCodeGenerator/util/parsers"
)

// Class ... Implements File
type Class struct {
	InheritedInterface Interface
	Name               		string
	
	// Header
	HeaderFileName      	string
	DefineName		     	string
	ForwardDeclares		  	string
	FunctionDeclarations 	string
	QtSignalDeclarations 	string

	// Implementation
	ImplementationFileName	string
	IncludesString			string 
	FunctionDefinitions 	string
	QtSignalDefinitions 	string
}

func NewClass(InheritedInterface Interface, className string) *Class {
	c := Class{}
	c.InheritedInterface = InheritedInterface
	c.Name = className

	// Header
	c.HeaderFileName = c.Name + configurations.Config.FileExtensions.CppHeader
	c.DefineName = parsers.GenerateDefineName(c.Name)
	c.FunctionDeclarations = c.parseFunctionDeclarations()
	c.ForwardDeclares = c.InheritedInterface.ForwardDeclaresString
	
	//Implementation
	c.ImplementationFileName = c.Name + configurations.Config.FileExtensions.CppImplementation
	c.FunctionDefinitions = c.parseFunctionDefinitions()
	c.IncludesString = c.InheritedInterface.IncludesString
	return &c
}

func (c Class) parseFunctionDefinitions() string {
	functionDefinitions := ""
	for _, function := range c.InheritedInterface.Functions {
		functionDefinitions += function.Definition(c.Name)
	}
	return functionDefinitions
}

func (c Class) parseFunctionDeclarations() string {
	functionDeclarations := ""
	for _, function := range c.InheritedInterface.Functions {
		functionDeclarations += function.Declaration()
	}
	return functionDeclarations
}

// Fields ... The fields within templates to be replaced.
func (c Class) Fields() map[string]string {
	fields := make(map[string]string)
	//Interface
	fields["{{Interface.FileName}}"] = c.InheritedInterface.FileName
	fields["{{Interface.Name}}"] = c.InheritedInterface.Name

	//Class
	fields["{{Class.Name}}"] = c.Name

	// Header
	fields["{{Class.Header.FileName}}"] = c.HeaderFileName
	fields["{{Class.Header.DefineName}}"] = c.DefineName
	fields["{{Class.Header.ForwardDeclares}}"] = c.ForwardDeclares
	fields["{{Class.Header.FunctionDeclarations}}"] = c.FunctionDeclarations
	fields["{{Class.Header.QtSignalDeclarations}}"] = c.QtSignalDeclarations

	//Implementation
	fields["{{Class.Implementation.FileName}}"] = c.ImplementationFileName
	fields["{{Class.Implementation.Includes}}"] = c.IncludesString
	fields["{{Class.Implementation.FunctionDefinitions}}"] = c.FunctionDefinitions
	fields["{{Class.Implementation.QtSignalDefinitions}}"] = c.QtSignalDefinitions
	return fields
}