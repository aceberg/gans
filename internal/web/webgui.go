package web

import (
	"log"
	"net/http"

	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/conf"
	"github.com/aceberg/gans/internal/yaml"
)

// Gui - start web server
func Gui(confPath, yamlPath string) {

	// AllRepos = allRepos
	ConfigPath = confPath
	YamlPath = yamlPath

	AllRepos = yaml.Read(YamlPath)
	log.Println("INFO: all repos", AllRepos)

	log.Println("INFO: starting web gui with config", ConfigPath)
	AppConfig = conf.Get(ConfigPath)

	address := AppConfig.Host + ":" + AppConfig.Port

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(address, nil)
	check.IfError(err)
}