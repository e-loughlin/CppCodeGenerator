package util

import (
	"path/filepath"
	"runtime"
)

// TODO: These directories break "go run" command. Find way to make it work.

// Directories
var _, b, _, _ = runtime.Caller(0)
var ExecutablePath = filepath.Dir(b)
var ProjectRoot = filepath.Join(ExecutablePath, "../")
var ResourcesDir = filepath.Join(ProjectRoot, "resources")
var TemplatesDir = filepath.Join(ResourcesDir, "templates")
var GMockGeneratorPath = filepath.Join(ResourcesDir, "gmock_generator", "gmock_gen_wrapper.py")
var IncludeListsDir = filepath.Join(ResourcesDir, "include-lists")
var QtClassesPath = filepath.Join(IncludeListsDir, "qt-includes.txt")
var StdTypesPath = filepath.Join(IncludeListsDir, "std-types.txt")
var ConfigurationsPath = filepath.Join(ResourcesDir, "config.json")
