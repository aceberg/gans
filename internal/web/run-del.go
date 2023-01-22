package web

import (
	"log"
	"net/http"
	"strconv"

	"github.com/aceberg/gans/internal/ansible"
	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/db"
	"github.com/aceberg/gans/internal/models"
)

func runGroupHandler(w http.ResponseWriter, r *http.Request) {
	var play models.Play

	play.Head = Repo.Head
	play.Inv = Repo.Inv
	play.File = r.FormValue("file")

	group := r.FormValue("group")
	hosts := AppConfig.GrMap[group]

	log.Println("INFO: playbook group", group, "hosts:", hosts)

	for _, host := range hosts {
		play.Host = host
		ansible.Playbook(AppConfig, play, Repo.Path)
	}

	http.Redirect(w, r, "/", 302)
}

func runHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")
	file := r.URL.Query().Get("file")

	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		check.IfError(err)

		for _, play := range Plays {
			if play.ID == id {
				ansible.Playbook(AppConfig, play, Repo.Path)
				break
			}
		}
	}

	if file != "" {
		var play models.Play

		play.Head = Repo.Head
		play.Inv = Repo.Inv
		play.File = file

		group := check.PlayHosts(Repo.Path + "/" + play.File)

		hosts := AppConfig.GrMap[group]

		log.Println("INFO: playbook group", group, "hosts:", hosts)

		for _, host := range hosts {
			play.Host = host
			ansible.Playbook(AppConfig, play, Repo.Path)
		}
	}

	http.Redirect(w, r, "/", 302)
}

func delHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	check.IfError(err)

	db.Delete(AppConfig.DB, id)

	http.Redirect(w, r, "/", 302)
}
