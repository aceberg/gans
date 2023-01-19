package web

import (
	"net/http"

	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/git"
	"github.com/aceberg/gans/internal/models"
	"github.com/aceberg/gans/internal/yaml"
)

func playsHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.Icon = Icon

	Repo = yaml.Read(AppConfig.YamlPath)
	guiData.Repo = Repo

	files := git.List(Repo.Path)

	for _, oneFile := range files {
		if oneFile != "" && check.IsYaml(Repo.Path+"/"+oneFile) && oneFile != Repo.Inv {
			guiData.Themes = append(guiData.Themes, oneFile)
		}
	}

	execTemplate(w, "plays", guiData)
}