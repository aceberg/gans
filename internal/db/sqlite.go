package db

import (
	"github.com/jmoiron/sqlx"

	// Import module for SQLite DB
	_ "modernc.org/sqlite"

	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/models"
)

func exec(path string, sqlStatement string) {
	dbx, err := sqlx.Connect("sqlite", path)
	check.IfError(err)

	_, err = dbx.Exec(sqlStatement)
	check.IfError(err)
}

// Select - select all from DB
func Select(path string) []models.Play {
	var playList []models.Play

	dbx, err := sqlx.Connect("sqlite", path)
	check.IfError(err)

	err = dbx.Select(&playList, "SELECT * FROM plays ORDER BY ID DESC")
	check.IfError(err)

	return playList
}
