package web

import (
	"log"
	"net/http"

	"github.com/aceberg/gans/internal/models"
)

func keysHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.Icon = Icon

	execTemplate(w, "keys", guiData)
}

func newKeyHandler(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	key := r.FormValue("key")

	log.Println("NAME:", name)
	log.Println("KEY:", key)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
