package git

import (
	"os/exec"
	"strings"

	"github.com/aceberg/gans/internal/check"
)

// git ls-files

// List - list all files
func List(path string) []string {

	gitDir := "--git-dir=" + path + "/.git"
	gitTree := "--work-tree=" + path

	cmd := exec.Command("git", gitDir, gitTree, "ls-files")

	out, err := cmd.CombinedOutput()
	check.IfError(err)

	files := strings.Split(string(out), "\n")

	return files
}
