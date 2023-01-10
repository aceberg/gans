package db

import (
	"fmt"

	"github.com/aceberg/gans/internal/models"
)

// Create - create table if not exists
func Create(path string) {
	sqlStatement := `CREATE TABLE IF NOT EXISTS plays (
		"ID"		INTEGER PRIMARY KEY,
		"HOST"		TEXT,
		"FILE"		TEXT,
		"HEAD"		TEXT,
		"OUT"		TEXT,
		"ERROR"		TEXT
	);`
	exec(path, sqlStatement)
}

// Insert - insert one play into DB
func Insert(path string, play models.Play) {

	sqlStatement := `INSERT INTO plays (HOST, FILE, HEAD, OUT, ERROR) 
	VALUES ('%s','%s','%s','%s','%s');`

	sqlStatement = fmt.Sprintf(sqlStatement, play.Host, play.File, play.Head, play.Out, play.Error)

	exec(path, sqlStatement)
}
