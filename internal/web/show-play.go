package web

import (
	"net/http"
	"os"

	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/models"
)

func showPlayHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData
	var text string
	var key models.Key

	guiData.Config = AppConfig
	guiData.Icon = Icon
	guiData.Repo = Repo

	filePath := r.URL.Query().Get("file")

	guiData.File = filePath

	file, err := os.ReadFile(Repo.Path + "/" + filePath)
	check.IfError(err)

	text = string(file)
	guiData.Themes = append(guiData.Themes, text)

	for group := range AppConfig.GrMap {
		key.Name = group
		guiData.Keys = append(guiData.Keys, key)
	}

	execTemplate(w, "show", guiData)
}
