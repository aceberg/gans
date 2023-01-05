package main

import (
	"log"

	// "github.com/aceberg/gans/internal/ansible"
	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/git"
)

func main() {
	repo := "/home/data/repo/01-cloned/testrepo"
	play := "playbook.yaml"
	inv := "inventory"
	// host := "bf"

	// hosts, groups := ansible.ParseInventory(repo + "/" + inv)

	// log.Println("HOSTS:", hosts)
	// log.Println("GROUPS:", groups)

	// out, err := ansible.Playbook(host, repo+"/"+inv, repo+"/"+play)
	// check.IfError(err)

	// log.Println("OUTPUT:", out)

	head := git.Head(repo)

	log.Println("HEAD:", head)

	log.Println("PLAY:", check.IsYaml(repo+"/"+play))
	log.Println("INV:", check.IsYaml(repo+"/"+inv))
	log.Println("REPO:", check.IsYaml(repo))

	files := git.ChangedFiles(repo, "9fc2add")
	log.Println("FILES:", files)
}
