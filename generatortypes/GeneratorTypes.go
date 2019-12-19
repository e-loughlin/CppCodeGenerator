// Copyright (c) 2019 Evan Loughlin
//
// MIT License

package generatortypes

import (
	"strings"
)

type GeneratorType int 

const (
	Interface GeneratorType = iota
	Class
	Test
	Mock
	SpyMock
	Unknown
)

func GetGeneratorType(inputString string) GeneratorType {
	inputString = strings.ToLower(inputString)
	switch(inputString) {
	case "interface":
		return Interface
	case "class":
		return Class
	case "test":
		return Test
	case "mock":
		return Mock
	case "spymock":
		return SpyMock
	default:
		return Unknown
	}
}