package cppcomponents

// ClassHeader ... Implements File
type ClassHeader struct {
	InheritedInterface Interface
	Name                 string
	DefineName		     string
	ForwardDeclares	     string
	FunctionDeclarations string
	QtSignalDeclarations string
}

func (c ClassHeader) newClassHeader() *ClassHeader {

	return &c
}

// Fields ... The fields within templates to be replaced.
func (c ClassHeader) Fields() map[string]string {
	fields := make(map[string]string)
	fields["{{Interface.FileName}}"] = c.InheritedInterface.FileName
	fields["{{Class.Name}}"] = c.Name
	fields["{{Class.Header.DefineName}}"] = c.DefineName
	fields["{{Class.Header.ForwardDeclares}}"] = c.ForwardDeclares
	fields["{{Class.Header.FunctionDeclarations}}"] = c.FunctionDeclarations
	fields["{{Class.Header.QtSignalDeclarations}}"] = c.QtSignalDeclarations
	return fields
}