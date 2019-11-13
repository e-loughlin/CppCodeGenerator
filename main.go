/// Author: Evan Loughlin
/// Date: 2019

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/emloughl/CppCodeGenerator/cppcomponents"
	"github.com/emloughl/CppCodeGenerator/util/templates"
	"github.com/emloughl/CppCodeGenerator/generatortypes"
	"github.com/emloughl/CppCodeGenerator/util/io"
	"github.com/emloughl/CppCodeGenerator/util/fieldreplacer"
	"github.com/emloughl/CppCodeGenerator/util/configurations"
)

func main() {
	// Command-line argument flags
	typeFlagPtr := flag.String("type", "", "Type of file to generate (class, interface, mock, or test).")
	interfaceFilepathFlagPtr := flag.String("interface", "", "Filepath of interface from which to base a generated derived class.")
	namePtr := flag.String("name", "", "Name of concrete class.")

	// If no arguments, print usage.
	if len(os.Args) < 2 {
		io.PrintUsage()
		os.Exit(0)
	}
	flag.Parse()

	// Load configurations
	configurations.Config = configurations.ReadConfigurations()
	configurations.SetTemplateFilePathsFromConfiguration()

	generatedType := generatortypes.GetGeneratorType(*typeFlagPtr)

	if generatedType == generatortypes.Unknown {
		fmt.Println("Invalid type! You must specify a type to generate. Use -type=<TYPE>")
		os.Exit(0)
	}

	// TODO: Fix guards
	interfaceFilepath := *interfaceFilepathFlagPtr
	// if interfaceFilepath == "" {
	// 	fmt.Println("You must specify either a path to an existing interface, or a path (including name) to where you'd like a new interface to be created. Use option -interface=<PATH_TO_INTERFACE>")
	// 	os.Exit(0)
	// }
	
	// Copyright Block
	copyrightBlock := cppcomponents.NewCopyrightCommentBlock()

	//Interface
	if generatedType == generatortypes.Interface {
		//TODO: Refactor templateType usage (enum)
		interfaceContents := templates.ReadTemplate(templates.Interface)

		// TODO: Refactor Interface so that it takes contents rather than filepath
		io.WriteToDisk(interfaceFilepath, interfaceContents)
		i := cppcomponents.NewInterface(interfaceFilepath)

		// Fill the copyright block fields
		interfaceContents = fieldreplacer.ReplaceAllFields(interfaceContents, copyrightBlock.Fields())

		// Fill the Interface fields
		interfaceContents = fieldreplacer.ReplaceAllFields(interfaceContents, i.Fields())
		
		io.WriteToDisk(interfaceFilepath, interfaceContents)
		os.Exit(0)
	}

	// Parse the Interface
	var inheritedInterface *cppcomponents.Interface
	if interfaceFilepath != "" {
		if !io.FileExists(interfaceFilepath) {
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
		if !io.FileExists(interfaceFilepath) {
			fmt.Fprintf(os.Stderr, "Invalid path to interface: %s\n", interfaceFilepath)
			os.Exit(0)
		}

		// --------------
		// CLASS HEADER 
		// --------------
		classHeader := cppcomponents.NewClassHeader(*inheritedInterface)
		interfaceDir := filepath.Dir(interfaceFilepath)
		classHeaderFilePath := filepath.Join(interfaceDir, classHeader.FileName)

		// Read Template File
		classHeaderContents := templates.ReadTemplate(templates.ClassHeader)

		// Fill the copyright block fields
		classHeaderContents = fieldreplacer.ReplaceAllFields(classHeaderContents, copyrightBlock.Fields())
		classHeaderContents = fieldreplacer.ReplaceAllFields(classHeaderContents, classHeader.Fields())

		// Write to disk
		io.WriteToDisk(classHeaderFilePath, classHeaderContents)

		// ----------------------
		// CLASS IMPLEMENTATION 
		// ----------------------
		classImplementation := cppcomponents.NewClassImplementation(*inheritedInterface)
		classImplementationFilePath := filepath.Join(interfaceDir, classImplementation.FileName)
		
		// Read Template File
		classImplementationContents := templates.ReadTemplate(templates.ClassImplementation)

		// Fill the copyright block fields
		classImplementationContents = fieldreplacer.ReplaceAllFields(classImplementationContents, copyrightBlock.Fields())
		classImplementationContents = fieldreplacer.ReplaceAllFields(classImplementationContents, classImplementation.Fields())

		// Write to disk
		io.WriteToDisk(classImplementationFilePath, classImplementationContents)
	}

	// Test
	if generatedType == generatortypes.Test {
		if *namePtr == "" {
			fmt.Println("Error: To create a test, you must specify the name of the concrete that it's testing.")
			os.Exit(0)
		}

		test := cppcomponents.NewTestByConcreteName(*namePtr)
		testContents := templates.ReadTemplate(templates.Test)
		testContents = fieldreplacer.ReplaceAllFields(testContents, copyrightBlock.Fields())
		testContents = fieldreplacer.ReplaceAllFields(testContents, test.Fields())
		cwd, _ := os.Getwd()
		testFilePath := filepath.Join(cwd, test.FileName)
		io.WriteToDisk(testFilePath, testContents)
	}

	// Mock
	if generatedType == generatortypes.Mock {
		if interfaceFilepath == "" {
			fmt.Println("Error: To create a Mock, you must provide the path to an interface.")
			os.Exit(0)
		}

		mock := cppcomponents.NewMock(*inheritedInterface)
		mockContents := templates.ReadTemplate(templates.MockHeader)
		mockContents = fieldreplacer.ReplaceAllFields(mockContents, copyrightBlock.Fields())
		mockContents = fieldreplacer.ReplaceAllFields(mockContents, mock.Fields())
		cwd, _ := os.Getwd()
		mockFilePath := filepath.Join(cwd, mock.HeaderFileName)
		io.WriteToDisk(mockFilePath, mockContents)
	}

}

// TODO: Consider creating a factory for each file type

// TODO: Feature: Command-line arguments to modify settings (Such as username)