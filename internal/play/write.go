package play

import (
	"github.com/aceberg/gans/internal/models"
	"github.com/aceberg/gans/internal/yaml"
)

func writeRepo(path string, repo models.Repo) {
	var newRepos []models.Repo

	allRepos := yaml.Read(path)

	for _, oneRepo := range allRepos {
		if oneRepo.Path == repo.Path {
			oneRepo = repo
		}
		newRepos = append(newRepos, oneRepo)
	}

	yaml.Write(path, newRepos)
}
