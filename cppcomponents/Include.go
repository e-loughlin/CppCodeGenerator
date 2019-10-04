package cppcomponents

import (
	"fmt"
	"strings"

	util "github.com/emloughl/CppCodeGenerator/util"
)

// Include ... A C++ Include of the form `#include "MyClass.h"` or `#include  <CoreClass>`
type Include struct {
	dependency string
}

// NewInclude .. Constructor
func NewInclude(dataType string) *Include {
	i := Include{dependency: deriveDependency(dataType)}
	return &i
}

// toString ... Creates the `#include "MyType.h"` string
func (i Include) toString() string {

	if util.IsStdDataType(i.dependency) {
		return ""
	}

	leftEnclosure := `"`
	rightEnclosure := `"`
	extension := `.h`
	if util.IsQtClass(i.dependency) {
		leftEnclosure = `<`
		rightEnclosure = `>`
		extension = ""
	}
	return fmt.Sprintf(`#include %s%s%s%s`, leftEnclosure, i.dependency, extension, rightEnclosure)
}

// deriveDependency ... (Attempts to) convert a data type to its base type.
func deriveDependency(dataType string) string {
	dataType = strings.TrimLeft(dataType, "&* ")
	dataType = strings.TrimRight(dataType, "&* ")
	return dataType
}
