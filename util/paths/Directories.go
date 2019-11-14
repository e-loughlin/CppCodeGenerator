package paths

import (
	"path/filepath"
	"os"
)

// TODO: These directories break "go run" command. Find way to make it work.
// Getting the Executable path directly in order to find relative paths is super 
// hacky... This is a serious TODO.

// Directories
var ex, _ = os.Executable()
var ExecutablePath = filepath.Dir(ex)
var ProjectRoot = filepath.Join(ExecutablePath, "../")
var ResourcesDir = filepath.Join(ProjectRoot, "resources")
var TemplatesDir = filepath.Join(ResourcesDir, "templates")
var GMockGeneratorPath = filepath.Join(ResourcesDir, "gmock_generator", "gmock_gen_wrapper.py")
var IncludeListsDir = filepath.Join(ResourcesDir, "include-lists")
var QtClassesPath = filepath.Join(IncludeListsDir, "qt-includes.txt")
var StdTypesPath = filepath.Join(IncludeListsDir, "std-types.txt")
var ConfigurationsPath = filepath.Join(ResourcesDir, "config.json")