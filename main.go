/// Author: Evan Loughlin
/// Date: 2019

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/emloughl/CppCodeGenerator/cppcomponents"
	"github.com/emloughl/CppCodeGenerator/util"
	"github.com/emloughl/CppCodeGenerator/configurations"
)

func main() {
	// Command-line argument flags
	typeFlagPtr := flag.String("type", "", "Type of file to generate (class, interface, mock, or test).")
	interfaceFilepathFlagPtr := flag.String("interface", "", "Filepath of interface from which to base a generated derived class.")
	interfaceNameFlagPtr := flag.String("iname", "", "Name of new interface. (Do not add a prefix / suffix). Used in conjunction with -type=interface.")

	// If no arguments, print usage.
	if(len(os.Args) < 2) {
		util.PrintUsage()
	}
	flag.Parse()

	// Read configurations
	config := configurations.ReadConfigurations()

	// 
	if *typeFlagPtr == "interface" {
		interfaceName := *interfaceNameFlagPtr
		interfaceName = config.Affixes.Prefixes.Interface + interfaceName + config.Affixes.Suffixes.Interface + config.FileExtensions.CppHeader
		fmt.Printf("Creating a new interface: %v\n", interfaceName)
	}

	// Parse the Interface
	var inheritedInterface *cppcomponents.Interface
	interfaceFilepath := *interfaceFilepathFlagPtr
	if (interfaceFilepath != "") {
		if !util.IsValidInterface(interfaceFilepath) {
			fmt.Fprintf(os.Stderr, "Invalid path to interface: %s\n", interfaceFilepath)
			os.Exit(0)
		}
		inheritedInterface = cppcomponents.NewInterface(interfaceFilepath)
		fmt.Println(inheritedInterface.Name)
	}

	if *typeFlagPtr == "class" {
		fmt.Println("Class!")
	}
}
