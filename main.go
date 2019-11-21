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
	var codeType string
	var interfaceFilePath string
	var name string

	flag.StringVar(&codeType, "type", "", "Type of file to generate (class, interface, mock, or test)")
	flag.StringVar(&codeType, "t", "", "--type")
	flag.StringVar(&interfaceFilePath, "interface", "", "Filepath to interface")
	flag.StringVar(&interfaceFilePath, "i", "", "--interface")
	flag.StringVar(&name, "name", "", "Name of new class")
	flag.StringVar(&name, "n", "", "--name")

	// If no arguments, print usage.
	if len(os.Args) < 2 {
		io.PrintUsage()
		os.Exit(0)
	}
	flag.Parse()
	
	// Load configurations
	configurations.Config = configurations.ReadConfigurations()
	configurations.SetTemplateFilePathsFromConfiguration()

	generatedType := generatortypes.GetGeneratorType(codeType)

	if generatedType == generatortypes.Unknown {
		fmt.Println("Invalid type! You must specify a type to generate. Use -type or -t. \n Valid types: Interface, Class, Test, and Mock.")
		os.Exit(0)
	}

	// TODO: Fix guards

	
	// Copyright Block
	copyrightBlock := cppcomponents.NewCopyrightCommentBlock()

	//Interface
	if generatedType == generatortypes.Interface {
		if(name == ""){
	 	fmt.Println("Trying to generate a new interface, but no name was provided. Use --name or -n.")
	 	os.Exit(0)
		}

		interfaceFilePath = name
		if(!cppcomponents.IsValidInterfaceFilePath(interfaceFilePath)) {
			interfaceFilePath = configurations.Config.Prefixes.Interface + name + configurations.Config.Suffixes.Interface + configurations.Config.FileExtensions.CppHeader
		}

		//TODO: Refactor templateType usage (enum)
		interfaceContents := templates.ReadTemplate(templates.Interface)

		// TODO: Refactor Interface so that it takes contents rather than filepath
		io.WriteToDisk(interfaceFilePath, interfaceContents)
		i := cppcomponents.NewInterface(interfaceFilePath)

		// Fill the copyright block fields
		interfaceContents = fieldreplacer.ReplaceAllFields(interfaceContents, copyrightBlock.Fields())

		// Fill the Interface fields
		interfaceContents = fieldreplacer.ReplaceAllFields(interfaceContents, i.Fields())
		
		io.WriteToDisk(interfaceFilePath, interfaceContents)

		// Print Result
		fmt.Printf("Generated new Interface: \n\t%v\n", interfaceFilePath)
		os.Exit(0)
	}

	// Parse the Interface
	var inheritedInterface *cppcomponents.Interface
	if interfaceFilePath != "" {
		if !io.FileExists(interfaceFilePath) {
			fmt.Fprintf(os.Stderr, "Invalid path to interface: %s\n", interfaceFilePath)
			os.Exit(0)
		}

		inheritedInterface = cppcomponents.NewInterface(interfaceFilePath)
	}

	// Class
	if generatedType == generatortypes.Class {
		if interfaceFilePath == "" {
			fmt.Println("Error: To create a class, you must provide the path to an interface. Use --interface or -i.")
			os.Exit(0)
		}
		if !io.FileExists(interfaceFilePath) {
			fmt.Fprintf(os.Stderr, "Invalid path to interface: %s\n", interfaceFilePath)
			os.Exit(0)
		}

		// --------------
		// CLASS HEADER 
		// --------------
		classHeader := cppcomponents.NewClassHeader(*inheritedInterface)
		interfaceDir := filepath.Dir(interfaceFilePath)
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

		// Print Result
		fmt.Printf("Generated Class from Interface %v: \n\t%v\n\t%v\n", inheritedInterface.FileName, classHeader.FileName, classImplementation.FileName)
	}

	// Test
	if generatedType == generatortypes.Test {
		if name == "" {
			fmt.Println("Error: To create a test, you must specify the name of the concrete that it's testing.")
			fmt.Println("Use option -name or -n <CONCRETE_NAME>")
			os.Exit(0)
		}

		test := cppcomponents.NewTestByConcreteName(name)
		testContents := templates.ReadTemplate(templates.Test)
		testContents = fieldreplacer.ReplaceAllFields(testContents, copyrightBlock.Fields())
		testContents = fieldreplacer.ReplaceAllFields(testContents, test.Fields())
		cwd, _ := os.Getwd()
		testFilePath := filepath.Join(cwd, test.FileName)
		io.WriteToDisk(testFilePath, testContents)

		// Print Result
		fmt.Printf("Generated Test for concrete class %v: \n\t%v\n", name, test.FileName)
	}

	// Mock
	if generatedType == generatortypes.Mock {
		if interfaceFilePath == "" {
			fmt.Println("Error: To create a Mock, you must provide the path to an interface. Use --interface or -i.")
			os.Exit(0)
		}

		cwd, _ := os.Getwd()

		mock := cppcomponents.NewMock(*inheritedInterface)
		mockHeaderContents := templates.ReadTemplate(templates.MockHeader)
		mockHeaderContents = fieldreplacer.ReplaceAllFields(mockHeaderContents, copyrightBlock.Fields())
		mockHeaderContents = fieldreplacer.ReplaceAllFields(mockHeaderContents, mock.Fields())
		mockHeaderFilePath := filepath.Join(cwd, mock.HeaderFileName)
		io.WriteToDisk(mockHeaderFilePath, mockHeaderContents)
		
		mockImplementationContents := templates.ReadTemplate(templates.MockImplementation)
		mockImplementationContents = fieldreplacer.ReplaceAllFields(mockImplementationContents, copyrightBlock.Fields())
		mockImplementationContents = fieldreplacer.ReplaceAllFields(mockImplementationContents, mock.Fields())
		mockImplementationFilePath := filepath.Join(cwd, mock.ImplementationFileName)
		io.WriteToDisk(mockImplementationFilePath, mockImplementationContents)

		// Print Result
		fmt.Printf("Generated Mock from Interface %v: \n\t%v\n\t%v\n", inheritedInterface.FileName, mock.HeaderFileName, mock.ImplementationFileName)
	}

}