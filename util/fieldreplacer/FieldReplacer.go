package fieldreplacer

import (
	"strings"
)

// ReplaceAllFields ...
func ReplaceAllFields(contents string, keyValuesMap map[string]string) string {
	
	// TODO Refactor so that this doesn't happen twice. (Stupid implementation due to nested structure of some fields.)
	for key, value := range keyValuesMap {
		contents = strings.Replace(contents, key, value, -1)
	}
	for key, value := range keyValuesMap {
		contents = strings.Replace(contents, key, value, -1)
	}
	return contents
}