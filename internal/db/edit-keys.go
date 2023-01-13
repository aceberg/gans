package db

import (
	"fmt"

	"github.com/aceberg/gans/internal/models"
)

// InsertKey - insert one key into DB
func InsertKey(path string, key models.Key) {

	sqlStatement := `INSERT INTO keys (DATE, NAME, FILE, STATE) 
	VALUES ('%s','%s','%s','%s');`

	key.Name = quoteStr(key.Name)

	sqlStatement = fmt.Sprintf(sqlStatement, key.Date, key.Name, key.File, key.State)

	exec(path, sqlStatement)
}

// DeleteKey - delete one key
func DeleteKey(path string, id int) {

	sqlStatement := `DELETE FROM keys WHERE ID='%d';`

	sqlStatement = fmt.Sprintf(sqlStatement, id)

	exec(path, sqlStatement)
}
