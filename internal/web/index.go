package web

import (
	"net/http"

	"github.com/aceberg/gans/internal/db"
	"github.com/aceberg/gans/internal/models"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.Icon = Icon

	Plays = db.Select(AppConfig.DB)
	guiData.Plays = Plays

	execTemplate(w, "index", guiData)
}
