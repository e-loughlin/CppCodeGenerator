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
	return fmt.Sprintf(`#include %s%s%s%s`, leftEnclosure, i.dependency, rightEnclosure, extension)
}

// deriveDependency ... (Attempts to) convert a data type to its base type.
func deriveDependency(dataType string) string {
	for {
		dataType = strings.TrimLeft(dataType, "&")
		dataType = strings.TrimLeft(dataType, "*")
		dataType = strings.TrimLeft(dataType, " ")
		dataType = strings.TrimRight(dataType, "&")
		dataType = strings.TrimRight(dataType, "*")
		dataType = strings.TrimRight(dataType, " ")

		first := rune(dataType[0])
		last := rune(dataType[len(dataType)-1])
		if first != '&' && last != '&' &&
			first != '*' && last != '*' &&
			first != ' ' && last != ' ' {
			break
		}
	}
	return dataType
}
