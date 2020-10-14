package database

import (
	"database/sql"
)

/*
Responsible for feeding the []bytes into channels
Input: []bytes of data
Output: Nothing -  channel modified by byRef
*/
func bytesDataToChan(rows *sql.Rows, results chan []byte) (err error) {

	return err
}
