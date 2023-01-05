package ansible

import (
	"os/exec"
)

// Playbook - run one playbook for one host from inventory
func Playbook(host, inv, play string) (string, error) {

	cmd := exec.Command("ansible-playbook", "-i", inv, "-l", host, play)

	out, err := cmd.CombinedOutput()

	return string(out), err
}
