package ansible

import (
	"bufio"
	"os"

	"github.com/aceberg/gans/internal/check"
)

// ParseInventory - get hosts and groups from inventory
func ParseInventory(path string) ([]string, []string) {
	var (
		text   string
		hosts  []string
		groups []string
	)

	file, err := os.Open(path)
	check.IfError(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text = scanner.Text()
		if text != "" && text[0:1] != "#" {
			if text[0:1] == "[" {
				groups = append(groups, text)
			} else {
				hosts = append(hosts, text)
			}
		}
	}

	file.Close()

	return hosts, groups
}
