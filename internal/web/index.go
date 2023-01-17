package web

import (
	"net/http"
	"strconv"

	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/db"
	"github.com/aceberg/gans/internal/models"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.Icon = Icon

	Plays = db.Select(AppConfig.DB)

	guiData.Len = len(Plays)

	show, err := strconv.Atoi(AppConfig.Show)
	check.IfError(err)

	if show > guiData.Len {
		show = guiData.Len
	}

	guiData.Plays = Plays[0:show]

	execTemplate(w, "index", guiData)
}
