package util

import (
	"os"
	"path/filepath"
)

// TODO: These directories break "go run" command. Find way to make it work.

// Directories
var ExecutablePath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
var ProjectRoot = filepath.Join(ExecutablePath, "../")
var ResourcesDir = filepath.Join(ProjectRoot, "resources")
var TemplatesDir = filepath.Join(ResourcesDir, "templates")
var IncludeListsDir = filepath.Join(ResourcesDir, "include-lists")
var QtClassesPath = filepath.Join(IncludeListsDir, "qt-includes.txt")
var StdTypesPath = filepath.Join(IncludeListsDir, "std-types.txt")
var ConfigurationsPath = filepath.Join(ResourcesDir, "config.json")
