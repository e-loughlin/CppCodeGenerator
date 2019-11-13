package gmockgenrunner

import (
	"os/exec"
	"bufio"
	"strings"
	
	"github.com/emloughl/CppCodeGenerator/util/paths"
)

// RunGMockGenerator ...
func RunGMockGenerator(interfaceFilePath string) string {
	cmd := exec.Command("python", paths.GMockGeneratorPath, interfaceFilePath)
	output, _ := cmd.Output()
	return string(output)
}

// GetGMockGeneratorFunctionRegistrations ... Runs gmock_gen_wrapper.py, 
// but returns only the GoogleMock registration macros.
func GetGMockGeneratorFunctionRegistrations(interfaceFilePath string) string {
	gmockContents := RunGMockGenerator(interfaceFilePath)

	gmockMacros := ""
	scanner := bufio.NewScanner(strings.NewReader(gmockContents))
	getLineCounter := 0
	for scanner.Scan() {
		if(strings.Contains(scanner.Text(), "MOCK_METHOD")){
			gmockMacros += scanner.Text() + "\n"
			getLineCounter = 1
		} else if getLineCounter > 0 {
			gmockMacros += scanner.Text() + "\n"
			getLineCounter--
		}
	}
	return gmockMacros
}