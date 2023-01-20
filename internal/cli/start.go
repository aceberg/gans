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

	repo := yaml.Read(conf.YamlPath)
	conf.GrMap = play.HostsToMap(repo.Hosts)
	log.Println("INFO: repo", repo)

	conf.Quit = make(chan bool)
	play.Exec(conf, repo)
}
