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

	guiData.Plays = db.Select(AppConfig.DB)

	execTemplate(w, "index", guiData)
}
