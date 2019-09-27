package main

import (
	"fmt"
	"strings"
)

// Include ...
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

	if IsStdDataType(i.dependency) {
		return ""
	}

	leftEnclosure := `"`
	rightEnclosure := `"`
	extension := `.h`
	if IsQtClass(i.dependency) {
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
