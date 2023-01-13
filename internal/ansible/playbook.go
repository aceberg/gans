package ansible

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/db"
	"github.com/aceberg/gans/internal/models"
)

// Playbook - run one playbook for one host from inventory
func Playbook(conf models.Conf, play models.Play, path string) {

	commandString := "cd " + path + " && " + "ansible-playbook -i " + play.Inv + " -l " + play.Host + " " + play.File

	log.Println("EXEC:", commandString)

	cmd := exec.Command("sh", "-c", commandString)

	out, err := cmd.CombinedOutput()

	play.Out = string(out)

	if check.IfError(err) {
		play.Error = fmt.Sprintf("%s", err)
	} else {
		play.Error = ""
	}

	if strings.Contains(play.Out, "skipping: no hosts matched") {
		play.Error = "Skipped"
	}

	play.Date = time.Now().Format("2006-01-02 15:04:05")

	db.Insert(conf.DB, play)

	log.Println(play.Out)
}
