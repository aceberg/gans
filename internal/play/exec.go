package play

import (
	"log"
	"strconv"
	"time"

	"github.com/aceberg/gans/internal/ansible"
	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/git"
	"github.com/aceberg/gans/internal/models"
)

// Exec - execute playbooks
func Exec(conf models.Conf, allRepos []models.Repo, quit chan bool) {
	var newRepos []models.Repo

	// All checks here
	for _, repo := range allRepos {
		if check.IsRepo(repo.Path) {
			newRepos = append(newRepos, repo)
		} else {
			log.Println("ERROR: not a git repo", repo.Path)
		}
	}

	tSec, err := check.TimeToSec(conf.Timeout)
	check.IfError(err)

	tInt, _ := strconv.Atoi(tSec)

	// Endless cycle with timeout
	for {
		for _, repo := range newRepos {
			play(conf, repo, quit)
		}

		time.Sleep(time.Duration(tInt) * time.Second)
	}
}

func play(conf models.Conf, repo models.Repo, quit chan bool) {

	head := git.Head(repo.Path)
	head = head[0:7]

	if head != repo.Head {
		if repo.Head == "" {
			log.Println("INFO: empty repo head. Writing head", head, "for", repo.Path)
			repo.Head = head
			writeRepo(conf, repo, quit)
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
			writeRepo(conf, repo, quit)
		}
	}
}
