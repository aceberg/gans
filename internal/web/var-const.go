package web

import (
	"github.com/aceberg/gans/internal/models"
)

var (
	// AppConfig - config for Web Gui
	AppConfig models.Conf
	// Repo - repository
	Repo models.Repo
)

// TemplPath - path to html templates
const TemplPath = "../../internal/web/templates/"

// Icon - favicon
const Icon = "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"
