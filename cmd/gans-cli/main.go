package main

import (
	// "log"
	"flag"

	"github.com/aceberg/gans/internal/cli"
	"github.com/aceberg/gans/internal/models"
)

const dbPath = "/data/gans/sqlite.db"
const yamlPath = "/data/gans/repos.yaml"
const timeout = "5s"

func main() {
	var conf models.Conf

	dbPtr := flag.String("d", dbPath, "Path to sqlite DB file")
	yamlPtr := flag.String("r", yamlPath, "Path to repos yaml file")
	timePtr := flag.String("t", timeout, "Scan for updates timeout (s, m, h)")
	flag.Parse()

	conf.DB = *dbPtr
	conf.YamlPath = *yamlPtr
	conf.Timeout = *timePtr

	cli.Start(conf)
}
