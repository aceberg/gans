package check

// IsPlay - check if file exists and yaml
func IsPlay(path string) bool {

	if IsYaml(path) {

		if PlayHosts(path) != "" {
			return true
		}
	}

	return false
}
