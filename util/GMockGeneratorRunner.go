package util

import (
	"os/exec"
)

// RunGMockGenerator ...
func RunGMockGenerator(interfaceFilePath string) string {
	cmd := exec.Command("python", GMockGeneratorPath, interfaceFilePath)
	output, _ := cmd.Output()
	return string(output)
}