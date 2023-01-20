package yaml

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/models"
)

// Read - read .yaml file to struct
func Read(path string) models.Repo {

	file, err := os.ReadFile(path)
	check.IfError(err)

	var repo models.Repo
	err = yaml.Unmarshal(file, &repo)
	check.IfError(err)

	return repo
}

// Write - write struct to  .yaml file
func Write(path string, repo models.Repo) {

	yamlData, err := yaml.Marshal(&repo)
	check.IfError(err)

	err = os.WriteFile(path, yamlData, 0644)
	check.IfError(err)

	log.Println("INFO: writing new repo file to", path, "\n", string(yamlData))
}

// GetPlayHosts - get valuse of "hosts:"
func GetPlayHosts(path string) string {

	file, err := os.ReadFile(path)
	check.IfError(err)

	var data []interface{}

	err = yaml.Unmarshal(file, &data)
	check.IfError(err)

	myMap := data[0].(map[string]interface{})

	str := fmt.Sprintf("%v", myMap["hosts"])

	return str
}
