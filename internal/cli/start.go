package cli

import (
	"log"

	"github.com/aceberg/gans/internal/play"
	"github.com/aceberg/gans/internal/yaml"
)

// Start cli app
func Start(yamlPath string) {
	log.Println("INFO: starting gans-cli with", yamlPath)

	allRepos := yaml.Read(yamlPath)
	log.Println("INFO: all repos", allRepos)

	if allRepos == nil {
		log.Printf("ERROR: no repos")
		return
	}

	// quit := make(chan bool)
	play.Exec(yamlPath, allRepos)

	// select {}
}
