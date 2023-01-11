package web

import (
	"log"
	"net/http"

	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/conf"
	"github.com/aceberg/gans/internal/db"
	"github.com/aceberg/gans/internal/models"
	"github.com/aceberg/gans/internal/play"
	"github.com/aceberg/gans/internal/yaml"
)

// Gui - start web server
func Gui(config models.Conf) {
	AppConfig = config

	log.Println("INFO: starting web gui with", AppConfig.ConfPath)

	Repo = yaml.Read(AppConfig.YamlPath)
	log.Println("INFO: repo", Repo)

	tmpConfig := conf.Get(AppConfig.ConfPath)

	AppConfig.Host = tmpConfig.Host
	AppConfig.Port = tmpConfig.Port
	AppConfig.Theme = tmpConfig.Theme

	db.Create(AppConfig.DB)

	AppConfig.Quit = make(chan bool)
	go play.Exec(AppConfig, Repo)

	address := AppConfig.Host + ":" + AppConfig.Port

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/config/", configHandler)
	// http.HandleFunc("/filter/", filterHandler)
	// http.HandleFunc("/repo/", repoHandler)
	http.HandleFunc("/save_config/", saveConfigHandler)
	http.HandleFunc("/status/", statusHandler)
	err := http.ListenAndServe(address, nil)
	check.IfError(err)
}
