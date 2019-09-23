package main

import (
	"path/filepath"
	"strings"
)

// Interface ...
type Interface struct {
	name      string
	functions []Function
	signals   []Function
	includes  []string
}

// NewInterface ... Constructor
func NewInterface(filePath string) *Interface {

	// interfaceContents := readContents(filePath)

	i := Interface{}
	i.name = strings.TrimSuffix(filePath, filepath.Ext(filePath))

	return &i
}

// func (cppInterface Interface) parseFunctions() []Function {

// 	return []
// }
