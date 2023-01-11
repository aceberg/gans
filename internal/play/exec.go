package play

import (
	"log"
	"strconv"
	"time"

	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/models"
	"github.com/aceberg/gans/internal/yaml"
)

// Exec - execute playbooks
func Exec(conf models.Conf, repo models.Repo) {

	if !check.IsRepo(repo.Path) {
		log.Println("ERROR: not a git repo", repo.Path)
		return
	}

	tSec, err := check.TimeToSec(conf.Interval)
	check.IfError(err)

	tInt, _ := strconv.Atoi(tSec)

	// Endless cycle with timeout
	for {
		select {
		case <-conf.Quit:
			return
		default:
			play(conf, repo)

			time.Sleep(time.Duration(tInt) * time.Second)

			repo = yaml.Read(conf.YamlPath)
		}
	}
}
