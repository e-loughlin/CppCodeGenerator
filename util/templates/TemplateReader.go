package templates

import (
	"log"

	"github.com/emloughl/CppCodeGenerator/util/paths"
	"github.com/emloughl/CppCodeGenerator/util/io"
)

// ReadTemplate ... Returns the contents of a given template file
func ReadTemplate(templateType Template) string {

	switch templateType {
	case ClassImplementation:
		return io.ReadContents(paths.ClassImplementionTemplatePath)
	case ClassHeader:
		return io.ReadContents(paths.ClassHeaderTemplatePath)
	case CommentBlock:
		return io.ReadContents(paths.CommentBlockTemplatePath)
	case Copyright:
		return io.ReadContents(paths.CopyrightTemplatePath)
	case Interface:
		return io.ReadContents(paths.InterfaceTemplatePath)
	case Test:
		return io.ReadContents(paths.TestTemplatePath)
	case MockImplementation:
		return io.ReadContents(paths.MockImplementationTemplatePath)
	case MockHeader:
		return io.ReadContents(paths.MockHeaderTemplatePath)
	default:
		log.Fatal("Attempted to read a non-existent template file.")
		return ""
	}
}


// TODO: Refactor replacement of fields. Causes a race condition sometimes due to constantly reading/saving file in loop.