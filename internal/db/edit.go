package db

import (
	"fmt"

	"github.com/aceberg/gans/internal/models"
)

// Create - create table if not exists
func Create(path string) {
	sqlStatement := `CREATE TABLE IF NOT EXISTS plays (
		"ID"		INTEGER PRIMARY KEY,
		"DATE"		TEXT,
		"HOST"		TEXT,
		"FILE"		TEXT,
		"HEAD"		TEXT,
		"INV"		TEXT,
		"OUT"		TEXT,
		"ERROR"		TEXT
	);`
	exec(path, sqlStatement)
}

// Insert - insert one play into DB
func Insert(path string, play models.Play) {

	sqlStatement := `INSERT INTO plays (DATE, HOST, FILE, HEAD, INV, OUT, ERROR) 
	VALUES ('%s','%s','%s','%s','%s','%s','%s');`

	sqlStatement = fmt.Sprintf(sqlStatement, play.Date, play.Host, play.File, play.Head, play.Inv, play.Out, play.Error)

	exec(path, sqlStatement)
}

// Delete - delete one play
func Delete(path string, id int) {

	sqlStatement := `DELETE FROM plays WHERE ID='%d';`

	sqlStatement = fmt.Sprintf(sqlStatement, id)

	exec(path, sqlStatement)
}
