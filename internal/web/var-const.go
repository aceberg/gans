package web

import (
	"github.com/aceberg/gans/internal/models"
)

var (
	// AppConfig - config for Web Gui
	AppConfig models.Conf
	// AllRepos - all repositories
	AllRepos []models.Repo
	// ConfigPath - path to Gui config file
	ConfigPath string
	// YamlPath - path to repos file
	YamlPath string
	// Quit - channel
	Quit chan bool
)

// TemplPath - path to html templates
const TemplPath = "../../internal/web/templates/"
