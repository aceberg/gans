package check

// IsPlay - check if file exists and yaml
func IsPlay(path string) bool {

	if IsYaml(path) {

		_, pin := PlayHosts(path)
		if pin {
			return true
		}
	}

	return false
}
