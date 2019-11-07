/// Author: Evan Loughlin
/// Date: 2019

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/emloughl/CppCodeGenerator/cppcomponents"
	"github.com/emloughl/CppCodeGenerator/util"
	"github.com/emloughl/CppCodeGenerator/generatortypes"
)

func main() {
	// Command-line argument flags
	typeFlagPtr := flag.String("type", "", "Type of file to generate (class, interface, mock, or test).")
	interfaceFilepathFlagPtr := flag.String("interface", "", "Filepath of interface from which to base a generated derived class.")

	// If no arguments, print usage.
	if len(os.Args) < 2 {
		util.PrintUsage()
		os.Exit(0)
	}
	flag.Parse()

	generatedType := generatortypes.GetGeneratorType(*typeFlagPtr)

	if generatedType == generatortypes.Unknown {
		fmt.Println("Invalid type! You must specify a type to generate. Use -type=<TYPE>")
		os.Exit(0)
	}

	interfaceFilepath := *interfaceFilepathFlagPtr
	if interfaceFilepath == "" {
		fmt.Println("You must specify either a path to an existing interface, or a path (including name) to where you'd like a new interface to be created. Use option -interface=<PATH_TO_INTERFACE>")
		os.Exit(0)
	}

	//Interface
	if generatedType == generatortypes.Interface {
		//TODO: Refactor templateType usage (enum)
		var templateType util.Template = util.InterfaceTemplate
		interfaceContents := util.ReadTemplate(templateType)

		// TODO: Refactor Interface so that it takes contents rather than filepath
		util.WriteToDisk(interfaceFilepath, interfaceContents)
		i := cppcomponents.NewInterface(interfaceFilepath)

		// Fill the copyright block fields
		copyrightBlock := cppcomponents.NewCopyrightCommentBlock()
		interfaceContents = util.ReplaceAllFields(interfaceContents, copyrightBlock.Fields())

		// Fill the Interface fields
		interfaceContents = util.ReplaceAllFields(interfaceContents, i.Fields())
		
		util.WriteToDisk(interfaceFilepath, interfaceContents)
		os.Exit(0)
	}

	// Parse the Interface
	var inheritedInterface *cppcomponents.Interface
	if interfaceFilepath != "" {
		if !util.FileExists(interfaceFilepath) {
			fmt.Fprintf(os.Stderr, "Invalid path to interface: %s\n", interfaceFilepath)
			os.Exit(0)
		}

		inheritedInterface = cppcomponents.NewInterface(interfaceFilepath)
	}

	// Class
	if generatedType == generatortypes.Class {
		if interfaceFilepath == "" {
			fmt.Println("Error: To create a class, you must specify an interface.")
			os.Exit(0)
		}
		if !util.FileExists(interfaceFilepath) {
			fmt.Fprintf(os.Stderr, "Invalid path to interface: %s\n", interfaceFilepath)
			os.Exit(0)
		}

		classHeader := cppcomponents.NewClassHeader(*inheritedInterface)
		interfaceDir := filepath.Dir(interfaceFilepath)
		classHeaderFilePath := filepath.Join(interfaceDir, classHeader.FileName)

		// Read Template File
		var templateType util.Template = util.ClassHeaderTemplate
		classHeaderContents := util.ReadTemplate(templateType)

		// Write template to disk
		util.WriteToDisk(classHeaderFilePath, classHeaderContents)

		// Fill the copyright block fields
		copyrightBlock := cppcomponents.NewCopyrightCommentBlock()
		util.ReplaceAllFields(interfaceFilepath, copyrightBlock.Fields())
		util.ReplaceAllFields(interfaceFilepath, copyrightBlock.Fields())
	}
}
