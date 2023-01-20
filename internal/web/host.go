package web

import (
	"net/http"
	"strings"

	"github.com/aceberg/gans/internal/models"
	"github.com/aceberg/gans/internal/play"
	"github.com/aceberg/gans/internal/yaml"
)

func hostAddHandler(w http.ResponseWriter, r *http.Request) {
	var oneHost models.Host

	oneHost.Name = r.FormValue("host")
	grStr := r.FormValue("groups")

	if oneHost.Name != "" {
		groups := strings.Split(grStr, " ")

		for _, gr := range groups {
			if gr != "" {
				oneHost.Groups = append(oneHost.Groups, gr)
			}
		}

		Repo = yaml.Read(AppConfig.YamlPath)
		Repo.Hosts = append(Repo.Hosts, oneHost)

		AppConfig.GrMap = play.HostsToMap(Repo.Hosts)

		close(AppConfig.Quit)
		yaml.Write(AppConfig.YamlPath, Repo)

		AppConfig.Quit = make(chan bool)
		go play.Exec(AppConfig, Repo)
	}

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func hostDelHandler(w http.ResponseWriter, r *http.Request) {
	var newHosts []models.Host

	host := r.FormValue("name")

	Repo = yaml.Read(AppConfig.YamlPath)

	for _, oneHost := range Repo.Hosts {
		if oneHost.Name != host {
			newHosts = append(newHosts, oneHost)
		}
	}

	Repo.Hosts = newHosts
	AppConfig.GrMap = play.HostsToMap(Repo.Hosts)

	close(AppConfig.Quit)
	yaml.Write(AppConfig.YamlPath, Repo)

	AppConfig.Quit = make(chan bool)
	go play.Exec(AppConfig, Repo)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
