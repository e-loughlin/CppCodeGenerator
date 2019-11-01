/// Author: Evan Loughlin
/// Date: 2019

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/emloughl/CppCodeGenerator/cppcomponents"
	"github.com/emloughl/CppCodeGenerator/util"
)

func main() {
	// Command-line argument flags
	typeFlagPtr := flag.String("type", "", "Type of file to generate (class, interface, mock, or test).")
	interfaceFilepathFlagPtr := flag.String("interface", "", "Filepath of interface from which to base a generated derived class.")

	// If no arguments, print usage.
	if(len(os.Args) < 2) {
		util.PrintUsage()
		os.Exit(0)
	}
	flag.Parse()

	if(*typeFlagPtr == "") {
		fmt.Println("You must specify a type to generate!")
		os.Exit(0)
	}

	if *typeFlagPtr == "interface" {
		interfaceFilepath := *interfaceFilepathFlagPtr
		if(interfaceFilepath == "") {
			fmt.Println("You must specify either a path to an existing interface, or a path to where you'd like a new interface to be created. Use option -interface=<PATH_TO_INTERFACE>")
			os.Exit(0)
		}

		//TODO: Refactor templateType usage (enum)
		var templateType util.Template = util.InterfaceTemplate
		interfaceContents := util.ReadTemplate(templateType)

		// Parse the existing interface and replace fields
		i := cppcomponents.NewInterface(interfaceFilepath)

		if(!util.FileExists(interfaceFilepath)) {
			util.WriteToDisk(interfaceFilepath, []byte(interfaceContents))
		}

		util.ReplaceAllFields(interfaceFilepath, i.Fields())
		os.Exit(0)
	}

	// Parse the Interface
	var inheritedInterface *cppcomponents.Interface
	interfaceFilepath := *interfaceFilepathFlagPtr
	if (interfaceFilepath != "") {
		if !util.FileExists(interfaceFilepath) {
			fmt.Fprintf(os.Stderr, "Invalid path to interface: %s\n", interfaceFilepath)
			os.Exit(0)
		}
		inheritedInterface = cppcomponents.NewInterface(interfaceFilepath)
		if(false) { //DELETE THIS
			fmt.Printf(inheritedInterface.Name)
		}
	}

	if *typeFlagPtr == "class" {
		fmt.Println("Class!")
	}
}
