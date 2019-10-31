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
	typePtr := flag.String("type", "", "Type of file to generate (class, interface, mock, or test).")
	interfaceFilepathPtr := flag.String("interface", "", "Filepath of interface from which to base a generated derived class.")

	// If no arguments, print usage.
	if(len(os.Args) < 2) {
		util.PrintUsage()
	}
	flag.Parse()

	if *typePtr == "class" {
		fmt.Println("Class!")
	}

	var inheritedInterface *cppcomponents.Interface
	interfaceFilepath := *interfaceFilepathPtr
	if interfaceFilepath != "" {
		if !util.IsValidInterface(interfaceFilepath) {
			fmt.Fprintf(os.Stderr, "Invalid path to interface: %s\n", interfaceFilepath)
			os.Exit(0)
		}
		inheritedInterface = cppcomponents.NewInterface(interfaceFilepath)
		fmt.Println(inheritedInterface.Name)
	}
}
