package main

import (
	"flag"

	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/models"
	"github.com/aceberg/gans/internal/web"
)

const confPath = "/data/gans/config.yaml"
const dbPath = ""
const yamlPath = ""
const interval = ""

func main() {
	var conf models.Conf

	dbPtr := flag.String("d", dbPath, "Path to sqlite DB file")
	confPtr := flag.String("c", confPath, "Path to config yaml file")
	yamlPtr := flag.String("r", yamlPath, "Path to repo yaml file")
	timePtr := flag.String("t", interval, "Interval between repo scans (s, m, h)")
	flag.Parse()

	conf.ConfPath = *confPtr
	conf.DB = *dbPtr
	conf.YamlPath = *yamlPtr
	conf.Interval = *timePtr

	check.Path(conf.ConfPath)
	check.Path(conf.DB)
	check.Path(conf.YamlPath)

	web.Gui(conf)
}
