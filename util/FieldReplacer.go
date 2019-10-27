package util

import (
	"strings"
)

// replaceFields ... 
func replaceFields(filePath string, key string, value string) {
	contents := ReadContents(filePath)
	contents = strings.Replace(contents, key, value, -1)
	bytesToWrite := []byte(contents)
	WriteToDisk(filePath, bytesToWrite)
}