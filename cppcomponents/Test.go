package cppcomponents

import (
	"github.com/emloughl/CppCodeGenerator/util/configurations"
)

// Test ... Implements File
type Test struct {
	TestName          		string
	FileName          		string
	ConcreteName			string
	ConcreteFileName		string
}

func NewTestByConcreteName(concreteName string) *Test {
	t := Test{}
	t.TestName = configurations.Config.Prefixes.Test + concreteName + configurations.Config.Suffixes.Test
	t.FileName = t.TestName + configurations.Config.FileExtensions.CppImplementation
	t.ConcreteName = concreteName
	t.ConcreteFileName = concreteName + configurations.Config.FileExtensions.CppHeader
	return &t
}

// Fields ... The fields within templates to be replaced.
func (t Test) Fields() map[string]string {
	fields := make(map[string]string)
	fields["{{Test.ConcreteFileName}}"] = t.ConcreteFileName
	fields["{{FileName}}"] = t.FileName
	fields["{{Test.ConcreteName}}"] = t.ConcreteName
	fields["{{Test.Name}}"] = t.TestName
	return fields
}