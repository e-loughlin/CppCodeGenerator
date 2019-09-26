package main

// ClassHeader ... Implements File
type ClassHeader struct {
	inheritedInterface Interface
	name               string
}

func (c ClassHeader) newClassHeader() *ClassHeader {

	return &c
}

func (c ClassHeader) generate() {

}
