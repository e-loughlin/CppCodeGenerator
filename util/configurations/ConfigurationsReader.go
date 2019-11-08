package configurations

import (
	"encoding/json"
	"fmt"

	"github.com/emloughl/CppCodeGenerator/util"
)

func readConfigurations() Configurations {
	var config Configurations
	jsonData := []byte(util.ReadContents(util.ConfigurationsPath))
	err := json.Unmarshal(jsonData, &config)
	
	if(err != nil) {
		fmt.Println("Unable to successfully read config.json. Check that it is valid.")
		util.Check(err)
	}

	return config
}