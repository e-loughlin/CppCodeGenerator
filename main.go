/// Author: Evan Loughlin
/// Date: 2019
///
/// C++ Code Generator
/// NewClass.go: Generates a class of specified type from a given interface.
/// Or, if generating an interface, writes a new interface file with the given
/// INTERFACE_PATH as a filename.
///
/// Usage:
///   NewClass.exe <CLASS_TYPE> <INTERFACE_PATH>
///
/// CLASS_TYPE   |                    Notes                    |
/// ------------------------------------------------------------
///   interface  |
///   class      |    Generates .h and .cpp of concrete implementation
///   test       |    In Progress (Requires testing framework to be in place)
///   mock       |    In Progress (Generates Mocks and SpyMocks)

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

	testFunctionDeclaration := "virtual QString vehicleNumber(int id, std::string name, VehicleManufacturer manufacturer) const = 0"
	f := cppcomponents.NewFunction(testFunctionDeclaration)
	fmt.Println(f.Declaration())
	fmt.Println(f.Definition("AutonomousDrone"))
}
