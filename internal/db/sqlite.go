package db

import (
	"github.com/jmoiron/sqlx"

	// Import module for SQLite DB
	_ "modernc.org/sqlite"

	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/models"
)

func connect(path string) *sqlx.DB {
	dbx, err := sqlx.Connect("sqlite", path)
	check.IfError(err)

	return dbx
}

func exec(path string, sqlStatement string) {

	dbx := connect(path)

	_, err := dbx.Exec(sqlStatement)
	check.IfError(err)
}

// Select - select all from DB
func Select(path string) []models.Play {
	var playList []models.Play

	dbx := connect(path)

	err := dbx.Select(&playList, "SELECT * FROM plays ORDER BY ID DESC")
	check.IfError(err)

	return playList
}

// SelectKeys - select all keys from DB
func SelectKeys(path string) []models.Key {
	var keyList []models.Key

	dbx := connect(path)

	err := dbx.Select(&keyList, "SELECT * FROM keys ORDER BY ID DESC")
	check.IfError(err)

	return keyList
}
