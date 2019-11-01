package cppcomponents

import (
	"time"
	"fmt"
	
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
func NewCopyrightCommentBlock() *CopyrightCommentBlock {
	c := CopyrightCommentBlock{}
	var templateType util.Template = util.CopyrightTemplate
	c.TemplateContents = util.ReadTemplate(templateType)
	c.CompanyName = configurations.Config.UserInfo.CompanyName
	c.Author = configurations.Config.UserInfo.Author
	t := time.Now()
	//TODO: Write documentation on how to format the config.json DateFormat configuration
	c.Date = fmt.Sprintf(t.Format(configurations.Config.Policies.DateFormat))
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
