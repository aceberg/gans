package check

import (
	"os"
	"path/filepath"
)

// IsYaml - check if file exists and yaml
func IsYaml(path string) bool {

	_, err := os.Stat(path)
	if err != nil {
		return false
	}

	ext := filepath.Ext(path)
	if ext != ".yaml" && ext != ".yml" {
		return false
	}

	return true
}
