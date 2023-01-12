package web

import (
	"net/http"
	"reflect"
	"sort"

	"github.com/aceberg/gans/internal/models"
)

func sortHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.Icon = Icon

	sortMethod := r.URL.Query().Get("meth")

	switch sortMethod {
	case "date-up":
		sortByField("asc", "Date")
	case "date-down":
		sortByField("desc", "Date")
	case "host-up":
		sortByField("asc", "Host")
	case "host-down":
		sortByField("desc", "Host")
	case "file-up":
		sortByField("asc", "File")
	case "file-down":
		sortByField("desc", "File")
	case "head-up":
		sortByField("asc", "Head")
	case "head-down":
		sortByField("desc", "Head")
	case "stat-up":
		sortByField("asc", "Error")
	case "stat-down":
		sortByField("desc", "Error")
	}

	guiData.Plays = Plays
	guiData.Len = len(Plays)

	execTemplate(w, "index", guiData)
}

func sortByField(method string, field string) {
	type fieldPlay struct {
		Play models.Play
		F    string
	}
	toSort := []fieldPlay{}
	var oneSort fieldPlay

	for _, play := range Plays {
		oneSort.Play = play
		r := reflect.ValueOf(&play)
		f := reflect.Indirect(r).FieldByName(field)
		oneSort.F = f.String()

		toSort = append(toSort, oneSort)
	}

	switch method {
	case "asc":
		sort.Slice(toSort, func(i, j int) bool {
			return toSort[i].F < toSort[j].F
		})
	case "desc":
		sort.Slice(toSort, func(i, j int) bool {
			return toSort[i].F > toSort[j].F
		})
	}

	Plays = []models.Play{}
	for _, oneSort := range toSort {
		Plays = append(Plays, oneSort.Play)
	}
}
