package templates

import (
	"log"
	"path/filepath"

	"github.com/emloughl/CppCodeGenerator/util"
	"github.com/emloughl/CppCodeGenerator/util/configurations"
)

// Filepaths for various templates
var interfaceTemplatePath = filepath.Join(util.TemplatesDir, configurations.Config.TemplateFileNames.Interface)
var classImplementionTemplatePath = filepath.Join(util.TemplatesDir, configurations.Config.TemplateFileNames.ClassImplementation)
var classHeaderTemplatePath = filepath.Join(util.TemplatesDir, configurations.Config.TemplateFileNames.ClassHeader)
var commentBlockTemplatePath = filepath.Join(util.TemplatesDir, configurations.Config.TemplateFileNames.FunctionCommentBlock)
var copyrightTemplatePath = filepath.Join(util.TemplatesDir, configurations.Config.TemplateFileNames.CopyrightBlock)
var testTemplatePath = filepath.Join(util.TemplatesDir, configurations.Config.TemplateFileNames.Test)
var mockTemplatePath = filepath.Join(util.TemplatesDir, configurations.Config.TemplateFileNames.Mock)

// ReadTemplate ... Returns the contents of a given template file
func ReadTemplate(templateType Template) string {

	switch templateType {
	case ClassImplementation:
		return util.ReadContents(classImplementionTemplatePath)
	case ClassHeader:
		return util.ReadContents(classHeaderTemplatePath)
	case CommentBlock:
		return util.ReadContents(commentBlockTemplatePath)
	case Copyright:
		return util.ReadContents(copyrightTemplatePath)
	case Interface:
		return util.ReadContents(interfaceTemplatePath)
	case Test:
		return util.ReadContents(testTemplatePath)
	case Mock:
		return util.ReadContents(mockTemplatePath)
	default:
		log.Fatal("Attempted to read a non-existent template file.")
		return ""
	}
}


// TODO: Refactor replacement of fields. Causes a race condition sometimes due to constantly reading/saving file in loop.