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
	var lastDate time.Time

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
			nowDate := time.Now()
			plusDate := lastDate.Add(time.Duration(tInt) * time.Second)

			if nowDate.After(plusDate) {

				play(conf, repo)

				lastDate = time.Now()

				repo = yaml.Read(conf.YamlPath)
				conf.GrMap = HostsToMap(repo.Hosts)
			}

			time.Sleep(time.Duration(1) * time.Second) // Cycle to check if quit
		}
	}
}
