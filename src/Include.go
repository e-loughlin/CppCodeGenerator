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
	leftEnclosure := `"`
	rightEnclosure := `"`
	extension := ".h"
	if IsQtClass(i.dependency) {
		leftEnclosure = `<`
		rightEnclosure = `>`
	}
	return fmt.Sprintf(`#include %s%s%s`, leftEnclosure, i.dependency, rightEnclosure)
}

// deriveDependency ... (Attempts to) convert a data type to its base type.
func deriveDependency(dataType string) string {
	derivedType := dataType
	for {
		derivedType = strings.TrimLeft(dataType, "&")
		derivedType = strings.TrimLeft(dataType, "*")
		derivedType = strings.TrimLeft(dataType, " ")
		derivedType = strings.TrimRight(dataType, "&")
		derivedType = strings.TrimRight(dataType, "*")
		derivedType = strings.TrimRight(dataType, " ")

		first := derivedType[0]
		last := derivedType[len(derivedType)-1]
		if first != '&' && last != '&' &&
			first != '*' && last != '*' &&
			first != ' ' && last != ' ' {
			break
		}
	}
	return derivedType
}
