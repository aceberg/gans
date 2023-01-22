package check

import (
	"os"
	"strings"
)

// PlayHosts - get valuse of "hosts:"
func PlayHosts(path string) string {
	var out string

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
		return out
	}

	hostsLine := strings.Split(out, "hosts: ")
	hosts := strings.Split(hostsLine[1], "#")
	return hosts[0]
}
