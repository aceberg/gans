package play

import (
	"strconv"
	"time"

	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/models"
	"github.com/aceberg/gans/internal/yaml"
)

// Exec - execute playbooks
func Exec(conf models.Conf, repo models.Repo) {

	tSec, err := check.TimeToSec(conf.Timeout)
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
