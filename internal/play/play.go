package play

import (
	"log"

	"github.com/aceberg/gans/internal/ansible"
	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/git"
	"github.com/aceberg/gans/internal/models"
	"github.com/aceberg/gans/internal/yaml"
)

func play(conf models.Conf, repo models.Repo) {

	head := git.Head(repo.Path)
	head = head[0:7]

	if head != repo.Head {
		if repo.Head == "" {
			log.Println("INFO: empty repo head. Writing head", head, "for", repo.Path)
			repo.Head = head
			yaml.Write(conf.YamlPath, repo)
		} else {
			files := git.ChangedFiles(repo.Path, repo.Head)
			log.Println("INFO: changed files", files)

			for _, host := range repo.Hosts {
				for _, file := range files {
					if check.IsYaml(repo.Path + "/" + file) {

						ansible.Playbook(host, repo.Path+"/"+repo.Inv, repo.Path+"/"+file)
					}
				}
			}
			repo.Head = head
			yaml.Write(conf.YamlPath, repo)
		}
	}
}
