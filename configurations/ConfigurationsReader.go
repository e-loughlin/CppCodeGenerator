package configurations

import (
	"encoding/json"

	"github.com/emloughl/CppCodeGenerator/util"
)

func readConfigurations() Configurations {
	var config Configurations
	jsonData := []byte(util.ReadContents(util.ConfigurationsPath))
	err := json.Unmarshal(jsonData, &config)
	util.Check(err)

	return config
}