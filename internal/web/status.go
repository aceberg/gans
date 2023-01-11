package web

import (
	"net/http"
	"strconv"

	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/models"
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.Icon = Icon

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	check.IfError(err)

	for _, play := range Plays {
		if play.ID == id {
			guiData.Plays = append(guiData.Plays, play)
			break
		}
	}

	execTemplate(w, "status", guiData)
}
