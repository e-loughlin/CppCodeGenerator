package util

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// List of all Qt Classes
var allQtClasses = strings.Split(ReadContents(QtClassesPath), "\n")

// List of all C++ std data types
var allStdTypes = strings.Split(ReadContents(StdTypesPath), "\n")

// IsValidDirectory ... Checks whether a directory exists by creating and deleting a temporary file.
func IsValidDirectory(directory string) bool {
	fileName := "validity_test.txt"
	filePath := filepath.Join(directory, fileName)

	if FileExists(filePath) {
		return true
	}

	// Attempt to create a file
	var data []byte
	if err := ioutil.WriteFile(filePath, data, 0644); err == nil {
		os.Remove(filePath) // And delete it
		return true
	}
	return false
}

// FileExists ... Returns whether a file exists at a given @filePath
func FileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	}
	return false
}

// ReadContents ...
func ReadContents(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	contents, err := ioutil.ReadAll(file)
	return string(contents)
}

// ReadLines ...
func ReadLines(filePath string) []string {
	contents := ReadContents(filePath)
	return strings.Split(contents, "\n")
}

// WriteToDisk ...
func WriteToDisk(filePath string, data string) {
	file, err := os.Create(filePath)
	Check(err)
	defer file.Close()
	file.Write([]byte(data))
}

// IsQtClass ... Returns whether a given string is a Qt class.
func IsQtClass(className string) bool {
	for _, qtClass := range allQtClasses {
		qtClass = strings.TrimRight(qtClass, "\n")
		qtClass = strings.TrimRight(qtClass, "\r")
		if className == qtClass {
			return true
		}
	}
	return false
}

// IsStdDataType ... Returns whether a given string is a std C/C++ type
func IsStdDataType(dataType string) bool {
	for _, stdType := range allStdTypes {
		stdType = strings.TrimRight(stdType, "\n")
		stdType = strings.TrimRight(stdType, "\r")
		if dataType == stdType {
			return true
		}
	}
	return false
}