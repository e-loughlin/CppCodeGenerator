package parsers

import (
	"strings"
	"github.com/fatih/camelcase"
	"github.com/emloughl/CppCodeGenerator/util/configurations"	
)

// GenerateDefineName ... 
func GenerateDefineName(name string) string {

	name = strings.Replace(name, "_", "", -1)
	splitName := camelcase.Split(name)
	name = strings.Join(splitName, configurations.Config.Policies.DefineNameCamelCaseSeparator)
	
	if(configurations.Config.Policies.DefineNameAllCapsEnabled) {
		name = strings.ToUpper(name)
	}

	defineName := configurations.Config.Prefixes.DefineName +
	name +
	configurations.Config.Suffixes.DefineName
	return defineName
}