package web

import (
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

	guiData.Keys = db.SelectKeys(AppConfig.DB)

	execTemplate(w, "keys", guiData)
}

func newKeyHandler(w http.ResponseWriter, r *http.Request) {
	var key models.Key

	key.Name = r.FormValue("name")
	key.Key = r.FormValue("key")

	key.Date = time.Now().Format("2006-01-02 15:04:05")
	key.File = AppConfig.KeyPath + "/" + key.Name

	file, err := os.Create(key.File)
	check.IfError(err)
	defer file.Close()

	_, err = file.WriteString(key.Key)
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
