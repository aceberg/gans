package web

import (
	"net/http"

	"github.com/aceberg/gans/internal/models"
)

func filterHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.Icon = Icon

	filter := r.URL.Query().Get("by")
	tag := r.URL.Query().Get("tag")

	for _, play := range Plays {
		switch filter {
		case "host":
			if play.Host == tag {
				guiData.Plays = append(guiData.Plays, play)
			}
		case "file":
			if play.File == tag {
				guiData.Plays = append(guiData.Plays, play)
			}
		case "head":
			if play.Head == tag {
				guiData.Plays = append(guiData.Plays, play)
			}
		default:
			guiData.Plays = append(guiData.Plays, play)
		}
	}

	Plays = guiData.Plays
	guiData.Len = len(Plays)

	execTemplate(w, "index", guiData)
}
