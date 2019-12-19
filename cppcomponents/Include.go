// Copyright (c) 2019 Evan Loughlin
//
// MIT License

package cppcomponents

import (
	"fmt"
	"strings"

	"github.com/emloughl/CppCodeGenerator/util/parsers"
)

// Include ... A C++ Include of the form `#include "MyClass.h"` or `#include  <CoreClass>`
type Include struct {
	Dependency string
}

// NewInclude .. Constructor
func NewInclude(dataType string) *Include {
	i := Include{Dependency: DeriveDependency(dataType)}
	return &i
}

// ToString ... Creates the `#include "MyType.h"` string
func (i Include) ToString() string {

	if parsers.IsStdDataType(i.Dependency) {
		return ""
	}

	leftEnclosure := `"`
	rightEnclosure := `"`
	extension := `.h`
	if parsers.IsQtClass(i.Dependency) || parsers.IsMappedLibraryType(i.Dependency) {
		leftEnclosure = `<`
		rightEnclosure = `>`
		extension = ""
	}

	return fmt.Sprintf(`#include %s%s%s%s`, leftEnclosure, i.Dependency, extension, rightEnclosure)
}

// DeriveDependency ... (Attempts to) convert a data type to its base type.
func DeriveDependency(dataType string) string {
	dataType = strings.TrimLeft(dataType, "&* ")
	dataType = strings.TrimRight(dataType, "&* ")
	return dataType
}
