package check

import (
	"os"
)

// IsRepo checks if the path is git repo
func IsRepo(path string) bool {

	_, err := os.Stat(path + "/.git")

	return err == nil
}
