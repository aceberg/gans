package web

import (
	"net/http"

	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/git"
	"github.com/aceberg/gans/internal/models"
)

func playsHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData
	var key models.Key
	var hosts []string

	guiData.Config = AppConfig
	guiData.Icon = Icon

	guiData.Repo = Repo

	files := git.List(Repo.Path)

	for _, oneFile := range files {
		if oneFile != "" && check.IsPlay(Repo.Path+"/"+oneFile) && oneFile != Repo.Inv {
			key.Name = oneFile

			hosts, _ = check.PlayHosts(Repo.Path + "/" + oneFile)

			key.State = ""
			for _, oneHost := range hosts {
				key.State = key.State + " " + oneHost
			}
			guiData.Keys = append(guiData.Keys, key)
		}
	}

	execTemplate(w, "plays", guiData)
}
