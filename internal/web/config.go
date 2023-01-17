package web

import (
	"log"
	"net/http"

	"github.com/aceberg/gans/internal/conf"
	"github.com/aceberg/gans/internal/db"
	"github.com/aceberg/gans/internal/models"
	"github.com/aceberg/gans/internal/play"
	"github.com/aceberg/gans/internal/yaml"
)

func configHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.Icon = Icon

	guiData.Themes = []string{"cerulean", "cosmo", "cyborg", "darkly", "flatly", "journal", "litera", "lumen", "lux", "materia", "minty", "morph", "pulse", "quartz", "sandstone", "simplex", "sketchy", "slate", "solar", "spacelab", "superhero", "united", "vapor", "yeti", "zephyr"}

	execTemplate(w, "config", guiData)
}

func saveConfigHandler(w http.ResponseWriter, r *http.Request) {

	AppConfig.DB = r.FormValue("db")
	AppConfig.Host = r.FormValue("host")
	AppConfig.Port = r.FormValue("port")
	AppConfig.Theme = r.FormValue("theme")
	AppConfig.Show = r.FormValue("show")
	AppConfig.YamlPath = r.FormValue("yamlpath")
	AppConfig.KeyPath = r.FormValue("keypath")
	AppConfig.Interval = r.FormValue("interval")

	close(AppConfig.Quit)

	conf.Write(AppConfig)
	log.Println("INFO: new config", AppConfig)

	Repo = yaml.Read(AppConfig.YamlPath)

	AppConfig.Quit = make(chan bool)
	go play.Exec(AppConfig, Repo)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func clearHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("INFO: deleting all plays from DB")

	db.Clear(AppConfig.DB)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
