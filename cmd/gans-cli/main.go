package main

import (
	// "log"
	"flag"

	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/cli"
	"github.com/aceberg/gans/internal/db"
	"github.com/aceberg/gans/internal/models"
)

const dbPath = "/data/gans/sqlite.db"
const yamlPath = "/data/gans/repo.yaml"
const interval = "5s"

func main() {
	var conf models.Conf

	dbPtr := flag.String("d", dbPath, "Path to sqlite DB file")
	yamlPtr := flag.String("r", yamlPath, "Path to repo yaml file")
	timePtr := flag.String("t", interval, "Interval between repo scans (s, m, h)")
	flag.Parse()

	conf.DB = *dbPtr
	conf.YamlPath = *yamlPtr
	conf.Interval = *timePtr

	check.Path(conf.DB)
	db.Create(conf.DB)

	cli.Start(conf)
}
