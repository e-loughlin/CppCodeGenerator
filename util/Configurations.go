package util

type Configurations struct {
	Affixes Affixes
	Suffixes Suffixes
	FileExtensions FileExtensions
}

type Affixes struct {
	Prefixes Prefixes
	Suffixes Suffixes
}

type Prefixes struct {
	Interface string
	Mock string
	SpyMock string
	Test string
}

type Suffixes struct {
	Interface string
	Mock string
	SpyMock string
	Test string
}

type FileExtensions struct {
	CppHeader string
	CppImplementation string
}