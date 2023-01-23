package main

import (
	"flag"

	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/models"
	"github.com/aceberg/gans/internal/web"
)

const confPath = "/data/gans/config.yaml"
const logPath = ""
const dbPath = ""
const yamlPath = ""

func main() {
	var conf models.Conf

	dbPtr := flag.String("d", dbPath, "Path to sqlite DB file")
	confPtr := flag.String("c", confPath, "Path to config yaml file")
	logPtr := flag.String("l", logPath, "Path to log file")
	yamlPtr := flag.String("r", yamlPath, "Path to repo yaml file")
	flag.Parse()

	conf.ConfPath = *confPtr
	conf.DB = *dbPtr
	conf.YamlPath = *yamlPtr
	conf.LogPath = *logPtr

	check.Path(conf.ConfPath)
	check.Path(conf.DB)
	check.Path(conf.YamlPath)
	check.Path(conf.LogPath)

	web.Gui(conf)
}
