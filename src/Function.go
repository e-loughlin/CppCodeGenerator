package main

// Function ...
type Function struct {
	name       string
	returnType string
	arguments  []Parameter
}

// NewFunction .. Constructor
func NewFunction(rawFunctionLine string) *Function {
	f := Function{}
	return &f
}
