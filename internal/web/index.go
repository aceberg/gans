package web

import (
	"net/http"

	"github.com/aceberg/gans/internal/models"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.Icon = Icon

	guiData.Repo = Repo

	execTemplate(w, "index", guiData)
}
