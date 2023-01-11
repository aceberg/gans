package ansible

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/db"
	"github.com/aceberg/gans/internal/models"
)

// Playbook - run one playbook for one host from inventory
func Playbook(conf models.Conf, play models.Play) {

	log.Println("EXEC: ansible-playbook -i ", play.Inv, "-l", play.Host, play.File)

	cmd := exec.Command("ansible-playbook", "-i", play.Inv, "-l", play.Host, play.File)

	out, err := cmd.CombinedOutput()

	play.Out = string(out)

	if check.IfError(err) {
		play.Error = fmt.Sprintf("%s", err)
	} else {
		play.Error = ""
	}

	db.Insert(conf.DB, play)

	log.Println(play.Out)
}
