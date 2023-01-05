package git

import (
	"os/exec"

	"github.com/aceberg/gans/internal/check"
)

// git rev-parse --short HEAD

// Head - get HEAD hash
func Head(path string) string {

	gitDir := "--git-dir=" + path + "/.git"
	gitTree := "--work-tree=" + path

	cmd := exec.Command("git", gitDir, gitTree, "rev-parse", "--short", "HEAD")

	out, err := cmd.CombinedOutput()
	check.IfError(err)

	return string(out)
}
