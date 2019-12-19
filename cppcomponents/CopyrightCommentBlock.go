// Copyright (c) 2019 Evan Loughlin
//
// MIT License

package cppcomponents

import (
	"time"
	"fmt"
	
	"github.com/emloughl/CppCodeGenerator/util/configurations"
	"github.com/emloughl/CppCodeGenerator/util/templates"
)

// CopyrightCommentBlock ... Implements File
type CopyrightCommentBlock struct {
	TemplateContents string
	CompanyName string
	Author string
	Date string
}

// NewCopyrightCommentBlock ...
func NewCopyrightCommentBlock() *CopyrightCommentBlock {
	c := CopyrightCommentBlock{}
	var templateType templates.Template = templates.Copyright
	c.TemplateContents = templates.ReadTemplate(templateType)
	c.CompanyName = configurations.Config.UserInfo.CompanyName
	c.Author = configurations.Config.UserInfo.Author
	t := time.Now()
	c.Date = fmt.Sprintf(t.Format(configurations.Config.Policies.DateFormat))
	return &c
}

// Fields ... The fields within templates to be replaced.
func (c CopyrightCommentBlock) Fields() map[string]string {
	fields := make(map[string]string)
	fields["{{Copyright}}"] = c.TemplateContents
	fields["{{UserInfo.Company}}"] = c.CompanyName
	fields["{{UserInfo.Author}}"] = c.Author
	fields["{{Date}}"] = c.Date
	return fields
}
