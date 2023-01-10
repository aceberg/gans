package cli

import (
	"log"

	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/models"
	"github.com/aceberg/gans/internal/play"
	"github.com/aceberg/gans/internal/yaml"
)

// Start cli app
func Start(conf models.Conf) {
	log.Println("INFO: starting gans-cli with", conf.YamlPath)

	repo := yaml.Read(conf.YamlPath)
	log.Println("INFO: repo", repo)

	if check.IsRepo(repo.Path) {
		conf.Quit = make(chan bool)
		play.Exec(conf, repo)
	} else {
		log.Println("ERROR: not a git repo", repo.Path)
	}
}
