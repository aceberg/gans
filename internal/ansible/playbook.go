package ansible

import (
	"log"
	"os/exec"

	"github.com/aceberg/gans/internal/check"
)

// Playbook - run one playbook for one host from inventory
func Playbook(host, inv, play string) {

	log.Println("PLAY: ansible-playbook -i ", inv, "-l", host, play)

	cmd := exec.Command("ansible-playbook", "-i", inv, "-l", host, play)

	out, err := cmd.CombinedOutput()
	check.IfError(err)

	log.Println(string(out))
}
