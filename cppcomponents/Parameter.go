package cppcomponents

import (
	"fmt"
	"strings"
)

// Parameter ...
type Parameter struct {
	Type string
	Name string
}

// NewParameter .. Constructor
func NewParameter(rawParameterLine string) *Parameter {
	p := Parameter{}
	p.parseNameAndType(rawParameterLine)
	p.Type = strings.TrimSpace(p.Type)
	p.Name = strings.TrimSpace(p.Name)
	return &p
}

func (p Parameter) ToString() string {
	return fmt.Sprintf("%v %v", p.Type, p.Name)
}

// parseNameAndType ... 
func (p *Parameter) parseNameAndType(rawParameterLine string) {
	// Temporarily remove "const "
	constString := ""
	if strings.Contains(rawParameterLine, "const "){
		constString = "const "
		rawParameterLine = strings.TrimPrefix(rawParameterLine, constString)
	}

	// Split on index of first space " " not within < angle brackets >
	bracketCount := 0
	for pos, char := range rawParameterLine {
		switch (char) {
		case '<':
			bracketCount++
		case '>':
			bracketCount--
		case ' ':
			if bracketCount == 0 {
				p.Type = constString + strings.TrimSpace(rawParameterLine[:pos])
				p.Name = strings.TrimSpace(rawParameterLine[pos:])
			}
		}
	}
}