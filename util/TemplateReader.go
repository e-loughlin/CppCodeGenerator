package util

import (
	"log"
	"path/filepath"
)

// Template ... An enumeration for possible template files.
type Template int

const (
	ClassImplementationTemplate Template = iota
	ClassHeaderTemplate
	CommentBlockTemplate
	CopyrightTemplate
	InterfaceTemplate
)

// Filepaths for various templates
var classImplementionTemplatePath = filepath.Join(TemplatesDir, "class_cpp.txt")
var classHeaderTemplatePath = filepath.Join(TemplatesDir, "class_header.txt")
var commentBlockTemplatePath = filepath.Join(TemplatesDir, "comment_block_function.txt")
var copyrightTemplatePath = filepath.Join(TemplatesDir, "copyright.txt")
var interfaceTemplatePath = filepath.Join(TemplatesDir, "interface.txt")

// ReadTemplate ... Returns the contents of a given template file
func ReadTemplate(templateType Template) string {

	switch templateType {
	case ClassImplementationTemplate:
		return ReadContents(classImplementionTemplatePath)
	case ClassHeaderTemplate:
		return ReadContents(classHeaderTemplatePath)
	case CommentBlockTemplate:
		return ReadContents(commentBlockTemplatePath)
	case CopyrightTemplate:
		return ReadContents(copyrightTemplatePath)
	case InterfaceTemplate:
		return ReadContents(interfaceTemplatePath)
	default:
		log.Fatal("Attempted to read a non-existent template file.")
		return ""
	}
}
