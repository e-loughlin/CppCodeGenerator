package configurations

var Config Configurations = readConfigurations()

type Configurations struct {
	Affixes Affixes
	Suffixes Suffixes
	Policies Policies
	FileExtensions FileExtensions
}

type Affixes struct {
	Prefixes Prefixes
	Suffixes Suffixes
}

type Prefixes struct {
	Interface string
	DefineName string
	Mock string
	SpyMock string
	Test string
}

type Suffixes struct {
	Interface string
	DefineName string
	Mock string
	SpyMock string
	Test string
}

type Policies struct {
	DefineNameAllCapsEnabled bool
	DefineNameCamelCaseSeparator string
}

type FileExtensions struct {
	CppHeader string
	CppImplementation string
}