package web

import (
	"bufio"
	"net/http"
	"os"
	"strconv"

	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/models"
)

func logHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData
	var text string
	var logs []string
	var show int

	lines := r.FormValue("lines")

	guiData.Config = AppConfig
	guiData.Icon = Icon

	file, err := os.Open(AppConfig.LogPath)
	check.IfError(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text = scanner.Text()
		logs = append(logs, text)
	}

	file.Close()

	guiData.Len = len(logs)

	if lines == "" {
		show, err = strconv.Atoi(AppConfig.Show)
		check.IfError(err)
	} else {
		if lines == "all" {
			show = guiData.Len
		} else {
			show, err = strconv.Atoi(lines)
			check.IfError(err)
		}
	}

	if show > guiData.Len {
		show = guiData.Len
	}

	start := guiData.Len - show

	guiData.Themes = logs[start:guiData.Len]

	execTemplate(w, "log", guiData)
}
