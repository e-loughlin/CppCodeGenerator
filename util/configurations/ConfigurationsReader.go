package configurations

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/emloughl/CppCodeGenerator/util/paths"
	"github.com/emloughl/CppCodeGenerator/util/io"
	"github.com/emloughl/CppCodeGenerator/util/errorhandler"
)

func ReadConfigurations() Configurations {
	var config Configurations
	jsonData := []byte(io.ReadContents(paths.ConfigurationsPath))
	err := json.Unmarshal(jsonData, &config)
	
	if(err != nil) {
		fmt.Println("Unable to successfully read config.json. Check that it is valid.")
		errorhandler.Check(err)
	}

	return config
}

func SetTemplateFilePathsFromConfiguration() {
	paths.InterfaceTemplatePath = filepath.Join(paths.TemplatesDir, Config.TemplateFileNames.Interface)
	paths.ClassImplementionTemplatePath = filepath.Join(paths.TemplatesDir, Config.TemplateFileNames.ClassImplementation)
	paths.ClassHeaderTemplatePath = filepath.Join(paths.TemplatesDir, Config.TemplateFileNames.ClassHeader)
	paths.CommentBlockTemplatePath = filepath.Join(paths.TemplatesDir, Config.TemplateFileNames.FunctionCommentBlock)
	paths.CopyrightTemplatePath = filepath.Join(paths.TemplatesDir, Config.TemplateFileNames.CopyrightBlock)
	paths.TestTemplatePath = filepath.Join(paths.TemplatesDir, Config.TemplateFileNames.Test)
	paths.MockImplementationTemplatePath = filepath.Join(paths.TemplatesDir, Config.TemplateFileNames.MockImplementation)
	paths.MockHeaderTemplatePath = filepath.Join(paths.TemplatesDir, Config.TemplateFileNames.MockHeader)

	for _, filename := range Config.TemplateFileNames.MockHelperDeclarations {
		paths.MockHelperFunctionDeclarationPaths = append(paths.MockHelperFunctionDeclarationPaths, filepath.Join(paths.TemplatesDir, filename))
	}
	
	for _, filename := range Config.TemplateFileNames.MockHelperDefinitions {
		paths.MockHelperFunctionDefinitionPaths = append(paths.MockHelperFunctionDefinitionPaths, filepath.Join(paths.TemplatesDir, filename))
	}
}