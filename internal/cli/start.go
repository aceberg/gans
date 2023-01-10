package cli

import (
	"log"

	"github.com/aceberg/gans/internal/models"
	"github.com/aceberg/gans/internal/play"
	"github.com/aceberg/gans/internal/yaml"
)

// Start cli app
func Start(conf models.Conf) {
	log.Println("INFO: starting gans-cli with", conf.YamlPath)

	allRepos := yaml.Read(conf.YamlPath)
	log.Println("INFO: all repos", allRepos)

	if allRepos == nil {
		log.Printf("ERROR: no repos")
		return
	}

	quit := make(chan bool)
	play.Exec(conf, allRepos, quit)

	// select {}
}
