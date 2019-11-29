package io

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

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
	ioutil.WriteFile(filePath, []byte(data), 0644)
}