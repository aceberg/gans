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
	var play models.Play

	head := git.Head(repo.Path)
	head = head[0:7]

	if head != repo.Head {
		if repo.Head == "" {
			log.Println("INFO: empty repo head. Writing head", head)
			repo.Head = head
			yaml.Write(conf.YamlPath, repo)
		} else {
			files := git.ChangedFiles(repo.Path, repo.Head)
			log.Println("INFO: changed files", files)

			play.Head = head
			play.Inv = repo.Inv

			for _, file := range files {
				play.File = file

				if check.IsPlay(repo.Path+"/"+play.File) && (play.File != play.Inv) {

					group := check.PlayHosts(repo.Path + "/" + play.File)

					hosts := conf.GrMap[group]

					log.Println("INFO: playbook group", group, "hosts:", hosts)

					for _, host := range hosts {
						play.Host = host
						ansible.Playbook(conf, play, repo.Path)
					}
				}
			}
			repo.Head = head
			yaml.Write(conf.YamlPath, repo)
		}
	}
}
