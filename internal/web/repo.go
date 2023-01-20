package web

import (
	"net/http"

	"github.com/aceberg/gans/internal/models"
	"github.com/aceberg/gans/internal/play"
	"github.com/aceberg/gans/internal/yaml"
)

func repoHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.Icon = Icon

	Repo = yaml.Read(AppConfig.YamlPath)
	AppConfig.GrMap = play.HostsToMap(Repo.Hosts)
	guiData.Repo = Repo

	execTemplate(w, "repo", guiData)
}

func saveRepoHandler(w http.ResponseWriter, r *http.Request) {

	Repo.Path = r.FormValue("path")
	Repo.Head = r.FormValue("head")
	Repo.Inv = r.FormValue("inv")

	close(AppConfig.Quit)
	yaml.Write(AppConfig.YamlPath, Repo)
	AppConfig.GrMap = play.HostsToMap(Repo.Hosts)

	AppConfig.Quit = make(chan bool)
	go play.Exec(AppConfig, Repo)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
