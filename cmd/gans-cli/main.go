package main

import (
	"log"

	"github.com/aceberg/gans/internal/ansible"
	"github.com/aceberg/gans/internal/check"
)

func main() {
	repo := "/home/data/repo/01-cloned/testrepo"
	play := "playbook.yaml"
	inv := "inventory"
	host := "bf"

	out, err := ansible.Playbook(host, repo+"/"+inv, repo+"/"+play)
	check.IfError(err)

	log.Println("OUTPUT:", out)
}
