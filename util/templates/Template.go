package templates

// Template ... An enumeration for possible template files.
type Template int

const (
	ClassImplementation Template = iota
	ClassHeader
	CommentBlock
	Copyright
	Interface
)