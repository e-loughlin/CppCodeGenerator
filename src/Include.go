package main

// Include ...
type Include struct {
	value string
}

// NewInclude .. Constructor
func NewInclude(rawIncludeLine string) *Include {
	f := Include{}
	return &f
}
