package web

import (
	"html"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/models"
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.Icon = Icon

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	check.IfError(err)

	for _, play := range Plays {
		if play.ID == id {
			play.Out = colorOutput(play.Out)
			guiData.Plays = append(guiData.Plays, play)
			break
		}
	}

	execTemplate(w, "status", guiData)
}

func colorOutput(str string) string {
	var out string

	out = strings.ReplaceAll(str, "ok", "&lt;b&gt;ok&lt;/b&gt;")
	out = html.UnescapeString(out)

	log.Println("OUT:\n", out)

	return out
}
