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
	AppConfig = mergeConfig(config)

	log.Println("INFO: starting web gui with", AppConfig.ConfPath)

	Repo = yaml.Read(AppConfig.YamlPath)
	AppConfig.GrMap = play.HostsToMap(Repo.Hosts)
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
	http.HandleFunc("/host_add/", hostAddHandler)
	http.HandleFunc("/host_del/", hostDelHandler)
	http.HandleFunc("/keys/", keysHandler)
	http.HandleFunc("/key_del/", keyDelHandler)
	http.HandleFunc("/log/", logHandler)
	http.HandleFunc("/new_key/", newKeyHandler)
	http.HandleFunc("/plays/", playsHandler)
	http.HandleFunc("/repo/", repoHandler)
	http.HandleFunc("/run/", runHandler)
	http.HandleFunc("/run_group/", runGroupHandler)
	http.HandleFunc("/save_config/", saveConfigHandler)
	http.HandleFunc("/save_repo/", saveRepoHandler)
	http.HandleFunc("/sort/", sortHandler)
	http.HandleFunc("/status/", statusHandler)
	err := http.ListenAndServe(address, nil)
	check.IfError(err)
}

func mergeConfig(config models.Conf) models.Conf {

	newConfig := conf.Get(config.ConfPath)
	newConfig.ConfPath = config.ConfPath

	if config.DB != "" {
		newConfig.DB = config.DB
	}
	if config.YamlPath != "" {
		newConfig.YamlPath = config.YamlPath
	}
	if config.LogPath != "" {
		newConfig.LogPath = config.LogPath
	}

	return newConfig
}
