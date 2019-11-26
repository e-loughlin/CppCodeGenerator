package io

import (
	"fmt"
	"flag"
)

// Prints file usage
func PrintUsage() {

	fmt.Printf(`
C++ Code Generator

Generates a C++ class of specified type.

        Type                       Notes                               Example Usage
---------------------------------------------------------------------------------------------------
	Interface  |  Requires --name of new class.          |    -t Interface -n MyNewClass
	Class      |  Requires path to existing --interface  |    -t Class -i ./path/to/I_MyNewClass.h
	Test       |  Requires --name of tested class        |    -t Test -n MyNewClass
	Mock       |  Requires path to existing --interface  |    -t Mock -i ./path/to/I_MyNewClass.h
	
Configurations
--------------
Configurations to policies, prefixes, suffixes, and file extensions can be made in config.json.

Templated Class Modifications
-----------------------------
Templated .txt files located in {ProjectRoot}/resources/templates provide the basis for all 
generated code, whereby tags indicated within "{{ }}" brackets are completed and filled.
The template files can be altered as needed by the user. Additionally, the config.json  
file indicates which template files are used, allowing the user to create multiple
templates, and simply modify the config.json when a different template should be used.

Valid Arguments
---------------
`)
	flag.PrintDefaults()
}

// TODO: Organize all print messages here.
