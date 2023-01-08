package play

import (
	"log"

	"github.com/aceberg/gans/internal/ansible"
	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/git"
	"github.com/aceberg/gans/internal/models"
)

// Exec - execute playbooks
func Exec(path string, allRepos []models.Repo) {
	var newRepos []models.Repo

	// All checks here
	for _, repo := range allRepos {
		if check.IsRepo(repo.Path) {
			newRepos = append(newRepos, repo)
		} else {
			log.Println("ERROR: not a git repo", repo.Path)
		}
	}

	// Endless cycle with timeout
	// for {
	for _, repo := range newRepos {
		play(path, repo)
	}
	// }
}

func play(path string, repo models.Repo) {

	head := git.Head(repo.Path)
	head = head[0:7]

	log.Println("HEAD:", head, ".")
	log.Println("REPO HEAD:", repo.Head, ".")

	if head != repo.Head {
		if repo.Head == "" {
			repo.Head = head
			writeRepo(path, repo)
			log.Println("Empty repo head")
		} else {
			files := git.ChangedFiles(repo.Path, repo.Head)
			log.Println("FILES:", files)

			for _, host := range repo.Hosts {
				for _, file := range files {
					if check.IsYaml(repo.Path + "/" + file) {
						log.Println("PLAY:", host, repo.Path+"/"+repo.Inv, repo.Path+"/"+file)
						out, err := ansible.Playbook(host, repo.Path+"/"+repo.Inv, repo.Path+"/"+file)
						check.IfError(err)

						log.Println(out)
					}
				}
			}
			repo.Head = head
			writeRepo(path, repo)
		}
	}
}
