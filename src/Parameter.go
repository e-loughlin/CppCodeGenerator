package main

// Parameter ...
type Parameter struct {
	varName string
	varType string
}

// NewParameter .. Constructor
func NewParameter(rawFunctionLine string) *Parameter {
	p := Parameter{}
	return &p
}
