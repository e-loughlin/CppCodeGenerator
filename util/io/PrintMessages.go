package io

import (
	"fmt"
	"flag"
	"os"
)

// Prints file usage
func PrintUsage() {

	fmt.Printf(`
	C++ Code Generator
	NewClass.go: Generates a class of specified type from a given interface.
	Or, if generating an interface, writes a new interface file with the given
	INTERFACE_PATH as a filename.

	If specifying 
   
	Usage:
	  NewClass -type=<CLASS_TYPE> -interface=<INTERFACE_PATH>
	
	Optional:
	  -iname=<CLASS_NAME>: Used when creating a new interface. Prefix of "I_" or "I" not required.
   
	CLASS_TYPE   |                    Notes                    |
	------------------------------------------------------------
	  interface  |
	  class      |    Generates .h and .cpp of concrete implementation
	  test       |    In Progress (Requires testing framework to be in place)
	  mock       |    In Progress (Generates Mocks and SpyMocks)
	`)
	fmt.Println("")
	fmt.Printf("Usage of %v:\n", os.Args[0])

	flag.PrintDefaults()
}

// TODO: Organize all print messages here.
