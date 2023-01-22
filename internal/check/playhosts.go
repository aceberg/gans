package check

import (
	"os"
	"strings"
)

// PlayHosts - get valuse of "hosts:"
func PlayHosts(path string) ([]string, bool) {
	var out string
	var pin bool

	file, err := os.ReadFile(path)
	IfError(err)

	str := string(file)
	lines := strings.Split(str, "\n")

	for _, line := range lines {
		if strings.Contains(line, "hosts:") {
			out = line
			break
		}
	}

	if out == "" {
		lines = []string{}
		pin = false
	} else {
		hosts1 := strings.Split(out, "hosts: ")
		hosts2 := strings.Split(hosts1[1], "#")
		lines = strings.Split(hosts2[0], " ")
		pin = true
	}

	return lines, pin
}
