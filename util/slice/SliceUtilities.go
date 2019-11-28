package slice

import (
	"strings"

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
	list := []string{}
	for _, entry := range stringSlice {
		if !io.IsStdDataType(entry) {
			list = append(list, entry)
		}
	}
	return list
}

func RemoveConstSpecifiers(stringSlice []string) []string {
	list := []string{}
	for _, entry := range stringSlice {
		list = append(list, strings.TrimPrefix(entry, "const "))
	}
	return list
}

func RemovePointersAndReferences(stringSlice []string) []string {
	list := []string{}
	for _, entry := range stringSlice {
		entry = strings.TrimLeft(entry, "&* ")
		entry = strings.TrimRight(entry, "&* ")
		list = append(list, entry)
	}
	return list
}