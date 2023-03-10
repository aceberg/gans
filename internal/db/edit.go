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
		"ERROR"		TEXT,
		"COLOR"		TEXT
	);`
	exec(path, sqlStatement)

	sqlStatement = `CREATE TABLE IF NOT EXISTS keys (
		"ID"		INTEGER PRIMARY KEY,
		"DATE"		TEXT,
		"NAME"		TEXT,
		"FILE"		TEXT,
		"STATE"		TEXT
	);`
	exec(path, sqlStatement)
}

// Insert - insert one play into DB
func Insert(path string, play models.Play) {

	sqlStatement := `INSERT INTO plays (DATE, HOST, FILE, HEAD, INV, OUT, ERROR, COLOR) 
	VALUES ('%s','%s','%s','%s','%s','%s','%s','%s');`

	play.Out = quoteStr(play.Out)
	play.Error = quoteStr(play.Error)

	sqlStatement = fmt.Sprintf(sqlStatement, play.Date, play.Host, play.File, play.Head, play.Inv, play.Out, play.Error, play.Color)

	exec(path, sqlStatement)
}

// Delete - delete one play
func Delete(path string, id int) {

	sqlStatement := `DELETE FROM plays WHERE ID='%d';`

	sqlStatement = fmt.Sprintf(sqlStatement, id)

	exec(path, sqlStatement)
}

// Clear - delete all plays from table
func Clear(path string) {
	sqlStatement := `DELETE FROM plays;`
	exec(path, sqlStatement)
}
