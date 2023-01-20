package play

import (
	"github.com/aceberg/gans/internal/models"
)

// HostsToMap - create a map of groups and hosts
func HostsToMap(repoHosts []models.Host) map[string][]string {
	var try bool

	m := make(map[string][]string)

	for _, host := range repoHosts {
		m["all"] = append(m["all"], host.Name)
		_, try = m[host.Name]
		if !try {
			m[host.Name] = append(m[host.Name], host.Name)
		}

		for _, gr := range host.Groups {
			if !contains(m[gr], host.Name) {
				m[gr] = append(m[gr], host.Name)
			}
		}
	}

	return m
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
