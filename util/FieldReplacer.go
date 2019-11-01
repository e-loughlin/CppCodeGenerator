package util

import (
	"strings"
)

// ReplaceFields ... 
func ReplaceFields(filePath string, key string, value string) {
	contents := ReadContents(filePath)
	contents = strings.Replace(contents, key, value, -1)
	bytesToWrite := []byte(contents)
	WriteToDisk(filePath, bytesToWrite)
}

// ReplaceAllFields ...
func ReplaceAllFields(filePath string, keyValuesMap map[string]string) {
	for key, value := range keyValuesMap {
		ReplaceFields(filePath, key, value)
	}
}