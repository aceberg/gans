package web

import (
	"net/http"
	"strconv"

	"github.com/aceberg/gans/internal/ansible"
	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/db"
)

func runHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	check.IfError(err)

	for _, play := range Plays {
		if play.ID == id {
			ansible.Playbook(AppConfig, play)
			break
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
