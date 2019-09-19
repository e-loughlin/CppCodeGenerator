package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// isValidDirectory ... Checks whether a directory exists by creating and deleting a temporary file.
func isValidDirectory(directory string) bool {
	fileName := "validity_test.txt"
	filePath := filepath.Join(directory, fileName)

	if fileExists(filePath) {
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

// fileExists ... Returns whether a file exists at a given @filePath
func fileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	}
	return false
}

// isValidInterface ...
func isValidInterface(filePath string) bool {
	return (fileExists(filePath) && (filepath.Ext(filePath) == ".h"))
}

// writeToDisk ...
func writeToDisk(fp string, data []byte) {
	f, err := os.Create(fp)
	check(err)
	defer f.Close()
	f.Write(data)
}
