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
	AppConfig = conf.Get(config.ConfPath)

	// Add if != "" later
	AppConfig.DB = config.DB
	AppConfig.ConfPath = config.ConfPath
	AppConfig.YamlPath = config.YamlPath
	AppConfig.Interval = config.Interval

	log.Println("INFO: starting web gui with", AppConfig.ConfPath)

	Repo = yaml.Read(AppConfig.YamlPath)
	log.Println("INFO: repo", Repo)

	db.Create(AppConfig.DB)

	AppConfig.Quit = make(chan bool)
	go play.Exec(AppConfig, Repo)

	address := AppConfig.Host + ":" + AppConfig.Port

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/clear/", clearHandler)
	http.HandleFunc("/config/", configHandler)
	http.HandleFunc("/del/", delHandler)
	http.HandleFunc("/filter/", filterHandler)
	http.HandleFunc("/keys/", keysHandler)
	http.HandleFunc("/key_del/", keyDelHandler)
	http.HandleFunc("/new_key/", newKeyHandler)
	http.HandleFunc("/repo/", repoHandler)
	http.HandleFunc("/run/", runHandler)
	http.HandleFunc("/save_config/", saveConfigHandler)
	http.HandleFunc("/save_repo/", saveRepoHandler)
	http.HandleFunc("/sort/", sortHandler)
	http.HandleFunc("/status/", statusHandler)
	err := http.ListenAndServe(address, nil)
	check.IfError(err)
}
