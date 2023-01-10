package play

import (
	"github.com/aceberg/gans/internal/models"
	"github.com/aceberg/gans/internal/yaml"
)

func writeRepo(conf models.Conf, repo models.Repo, quit chan bool) {
	var newRepos []models.Repo

	close(quit)

	allRepos := yaml.Read(conf.YamlPath)

	for _, oneRepo := range allRepos {
		if oneRepo.Path == repo.Path {
			oneRepo = repo
		}
		newRepos = append(newRepos, oneRepo)
	}

	yaml.Write(conf.YamlPath, newRepos)

	quit = make(chan bool)
	Exec(conf, newRepos, quit)
}
