package web

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/db"
	"github.com/aceberg/gans/internal/models"
)

func keysHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.Icon = Icon

	keys := db.SelectKeys(AppConfig.DB)

	for _, key := range keys {
		_, err := os.Stat(key.File)
		if check.IfError(err) {
			key.State = "Absent"
		} else {
			key.State = "Present"
		}
		guiData.Keys = append(guiData.Keys, key)
	}

	execTemplate(w, "keys", guiData)
}

func newKeyHandler(w http.ResponseWriter, r *http.Request) {
	var key models.Key

	key.Name = r.FormValue("name")

	uploadFile, _, err := r.FormFile("keyfile")
	check.IfError(err)
	defer uploadFile.Close()

	key.Date = time.Now().Format("2006-01-02 15:04:05")
	key.File = AppConfig.KeyPath + "/" + key.Name

	check.Path(AppConfig.KeyPath + "/" + "check.path")
	err = os.Remove(AppConfig.KeyPath + "/" + "check.path")
	check.IfError(err)

	file, err := os.Create(key.File)
	check.IfError(err)
	defer file.Close()

	_, err = io.Copy(file, uploadFile)
	check.IfError(err)

	err = os.Chmod(key.File, 0600)
	check.IfError(err)

	db.InsertKey(AppConfig.DB, key)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func keyDelHandler(w http.ResponseWriter, r *http.Request) {
	var keys []models.Key
	var file string

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	check.IfError(err)

	keys = db.SelectKeys(AppConfig.DB)

	for _, oneKey := range keys {
		if oneKey.ID == id {
			file = oneKey.File
			err = os.Remove(oneKey.File)
			check.IfError(err)
		}
	}

	db.DeleteKey(AppConfig.DB, id)

	log.Println("INFO: deleting key", file)

	http.Redirect(w, r, "/keys/", 302)
}
