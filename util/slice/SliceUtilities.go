package slice

import (
	"github.com/emloughl/CppCodeGenerator/util/io"
)

func RemoveDuplicates(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func RemoveStdDataTypes(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if !io.IsStdDataType(entry) {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}