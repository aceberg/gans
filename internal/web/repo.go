package web

import (
	"net/http"
	"strings"

	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/git"
	"github.com/aceberg/gans/internal/models"
	"github.com/aceberg/gans/internal/play"
	"github.com/aceberg/gans/internal/yaml"
)

func repoHandler(w http.ResponseWriter, r *http.Request) {
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

	execTemplate(w, "repo", guiData)
}

func saveRepoHandler(w http.ResponseWriter, r *http.Request) {

	Repo.Path = r.FormValue("path")
	Repo.Head = r.FormValue("head")
	Repo.Inv = r.FormValue("inv")
	hostsStr := r.FormValue("hosts")

	hosts := strings.Split(hostsStr, " ")

	Repo.Hosts = []string{}
	for _, host := range hosts {
		if host != "" {
			Repo.Hosts = append(Repo.Hosts, host)
		}
	}

	close(AppConfig.Quit)
	yaml.Write(AppConfig.YamlPath, Repo)

	AppConfig.Quit = make(chan bool)
	go play.Exec(AppConfig, Repo)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
