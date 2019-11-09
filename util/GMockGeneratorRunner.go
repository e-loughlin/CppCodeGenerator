package util

import (
	"os/exec"
	"fmt"
)

// RunGMockGenerator ...
func RunGMockGenerator(interfaceFilePath string) string {
	runCommand := fmt.Sprintf("python %v %v", GmockGeneratorPath, interfaceFilePath)
	cmd := exec.Command(runCommand)
	output, _ := cmd.Output()
	return string(output)
}