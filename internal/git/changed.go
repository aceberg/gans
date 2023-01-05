package git

import (
	"os/exec"
	"strings"

	"github.com/aceberg/gans/internal/check"
)

// git diff --name-only 9fc2add HEAD

// ChangedFiles - get changed files
func ChangedFiles(path, prevCommit string) []string {

	gitDir := "--git-dir=" + path + "/.git"
	gitTree := "--work-tree=" + path

	cmd := exec.Command("git", gitDir, gitTree, "diff", "--name-only", prevCommit, "HEAD")

	out, err := cmd.CombinedOutput()
	check.IfError(err)

	files := strings.Split(string(out), "\n")

	return files
}
