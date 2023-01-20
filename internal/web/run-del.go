package web

import (
	"net/http"
	"strconv"

	"github.com/aceberg/gans/internal/ansible"
	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/db"
	"github.com/aceberg/gans/internal/models"
)

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

		for _, host := range Repo.Hosts {
			play.Host = host.Name
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
