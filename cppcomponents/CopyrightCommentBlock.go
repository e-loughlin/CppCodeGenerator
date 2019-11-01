package cppcomponents

import (
	"time"
	
	"github.com/emloughl/CppCodeGenerator/util"
	"github.com/emloughl/CppCodeGenerator/configurations"
)

// CopyrightCommentBlock ... Implements File
type CopyrightCommentBlock struct {
	TemplateContents string
	CompanyName string
	Author string
	Date string
}

// NewCopyrightCommentBlock ...
func NewCopyrightCommentBlock(templateFilePath string) *CopyrightCommentBlock {
	c := CopyrightCommentBlock{}
	var templateType util.Template = util.CommentBlockTemplate
	c.TemplateContents = util.ReadTemplate(templateType)
	c.CompanyName = configurations.Config.UserInfo.CompanyName
	c.Author = configurations.Config.UserInfo.Author
	t := time.Now()
	t.Format(configurations.Config.Policies.DateFormat)
	c.Date = t.String()
	return &c
}

// Fields ... The fields within templates to be replaced.
func (c CopyrightCommentBlock) Fields() map[string]string {
	fields := make(map[string]string)
	fields["{{COPYRIGHT}}"] = c.TemplateContents
	fields["{{UserInfo.Company}}"] = c.CompanyName
	fields["{{UserInfo.Author}}"] = c.Author
	fields["{{Date}}"] = c.Date
	return fields
}