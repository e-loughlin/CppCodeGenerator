package util

import (
	"os"
	"path/filepath"
)

// TODO: These directories break "go run" command. Find way to make it work.

// Directories
var executablePath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
var projectRoot = filepath.Join(executablePath, "../")
var resourcesDir = filepath.Join(projectRoot, "resources")
var templatesDir = filepath.Join(resourcesDir, "templates")
var includeListsDir = filepath.Join(resourcesDir, "include-lists")
var qtClassesPath = filepath.Join(includeListsDir, "qt-includes.txt")
var stdTypesPath = filepath.Join(includeListsDir, "std-types.txt")
