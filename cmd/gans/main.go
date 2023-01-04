package main

import (
	"flag"

	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/web"
)

const confPath = "/data/gans/config.yaml"
const yamlPath = "/data/gans/repos.yaml"

func main() {
	confPtr := flag.String("c", confPath, "Path to config yaml file")
	yamlPtr := flag.String("r", yamlPath, "Path to repos yaml file")
	flag.Parse()

	check.Path(*confPtr)
	check.Path(*yamlPtr)

	web.Gui(*confPtr, *yamlPtr)
}
