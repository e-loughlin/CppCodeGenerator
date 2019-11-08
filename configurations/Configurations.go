package configurations

var Config Configurations = readConfigurations()

type Configurations struct {
	UserInfo UserInfo
	Prefixes Prefixes
	Suffixes Suffixes
	Syntax Syntax
	Policies Policies
	FileExtensions FileExtensions

}

type UserInfo struct {
	Author string
	CompanyName string
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

type Syntax struct {
	Tab string
}

type Policies struct {
	DefineNameAllCapsEnabled bool
	DefineNameCamelCaseSeparator string
	DateFormat string
}

type FileExtensions struct {
	CppHeader string
	CppImplementation string
}