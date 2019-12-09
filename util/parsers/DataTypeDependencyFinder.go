package parsers

import (
	"strings"

	"github.com/emloughl/CppCodeGenerator/util/io"
	"github.com/emloughl/CppCodeGenerator/util/paths"
)

// List of all Qt Classes
var allQtClasses []string = loadQtClasses()

// List of all C++ std data types
var allStdTypes []string = loadStdClasses()

// List of all types that should be included in header rather than forward declared
var allIncludeInHeaderTypes []string = loadIncludeInHeaderClasses()

// Map of data types to dependencies as per mapped-includes.txt
var allMappedTypes map[string]string = loadMappedTypes()

// IsQtClass ... Returns whether a given string is a Qt class.
func IsQtClass(className string) bool {
	for _, qtClass := range allQtClasses {
		qtClass = strings.TrimRight(qtClass, "\n\r")
		if className == qtClass {
			return true
		}
	}
	return false
}

// IsStdDataType ... Returns whether a given string is a std C/C++ type
func IsStdDataType(dataType string) bool {
	for _, stdType := range allStdTypes {
		stdType = strings.TrimRight(stdType, "\n\r")
		if dataType == stdType {
			return true
		}
	}
	return false
}

// IsMappedLibraryType ... Returns whether dependency is in the second column of listed mapped dependencies
func IsMappedLibraryType(libraryDependency string) bool {
	for dataType := range allMappedTypes {
		if libraryDependency == allMappedTypes[dataType] {
			return true
		}
	}
	return false
}

// MapToListedDependency ... Maps a user-configured <data_type> to its <library_dependency>
// If no mapping is found, the original dataType is returned.
func MapToListedDependency(dataType string) string {
	returnValue := dataType
	if libraryDependency, ok := allMappedTypes[dataType]; ok {
		returnValue = libraryDependency
	}
	return returnValue
}

// ShouldBeIncludedInHeader .. Determines whether a type is configured to be included in header rather than forward declared
func ShouldBeIncludedInHeader(dataType string) bool {
	for _, includeInHeaderType:= range allIncludeInHeaderTypes {
		includeInHeaderType = strings.TrimRight(includeInHeaderType, "\n\r")
		if dataType == includeInHeaderType{
			return true
		}
	}
	return false
}

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
		if !IsStdDataType(entry) {
			list = append(list, entry)
		}
	}
	return list
}

func MapDataTypesToLibraryDependencies(stringSlice []string) []string {
	list := []string{}
	for _, entry := range stringSlice {
		list = append(list, MapToListedDependency(entry))
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

func filter(ss []string, test func(string) bool) (ret []string) {
    for _, s := range ss {
        if test(s) {
            ret = append(ret, s)
        }
    }
    return
}

// loadStdClasses ... Loads Std Classes txt file and returns list of its contents ignoring comments
func loadStdClasses() []string {
	stdClasses := strings.Split(io.ReadContents(paths.StdTypesPath), "\n")
	stdClasses = filter(stdClasses, isNotCommentLine)
	return stdClasses
}

// loadQtClasses ... Loads Qt Classes txt file and returns list of its contents ignoring comments
func loadQtClasses() []string {
	qtClasses := strings.Split(io.ReadContents(paths.QtClassesPath), "\n")
	qtClasses = filter(qtClasses, isNotCommentLine)
	return qtClasses
}

// loadIncludeInHeaderClasses ... Loads classes from file that should be included in header instead of forward declared
func loadIncludeInHeaderClasses() []string {
	includeInHeaderClasses := strings.Split(io.ReadContents(paths.IncludeInHeaderPath), "\n")
	includeInHeaderClasses = filter(includeInHeaderClasses, isNotCommentLine)
	return includeInHeaderClasses
}

func loadMappedTypes() map[string]string {
	mappedTypesContents := strings.Split(io.ReadContents(paths.MappedTypesPath), "\n")
	mappedTypesContents = filter(mappedTypesContents, isNotCommentLine)

	var mappedTypes map[string]string
	mappedTypes = make(map[string]string)

	for _, line := range(mappedTypesContents) {
		line = strings.TrimRight(line, "\n\r")
		line = strings.TrimSpace(line)
		if len(line) < 4 {
			continue
		}
		splitLine := strings.Split(line, " ")
		dataType := splitLine[0]
		libraryDependency := splitLine[1]
		mappedTypes[dataType] = libraryDependency
	}

	return mappedTypes
}

func isNotCommentLine(line string) bool {
	return !strings.ContainsRune(line, '#')
}