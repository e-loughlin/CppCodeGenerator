package cppcomponents

import (
	"path/filepath"
	"strings"
)

// Interface ... Implements File
type Interface struct {
	Name      string
	Functions []Function
	Signals   []Function
	Includes  []string
}

// NewInterface ... Constructor
func NewInterface(filePath string) *Interface {

	// interfaceContents := readContents(filePath)

	i := Interface{}
	filePath = strings.Replace(filePath, ":", "", -1)
	i.Name = strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))

	return &i
}

// func (cppInterface Interface) parseFunctions() []Function {

// 	return []
// }

func (i Interface) generate() {

}
