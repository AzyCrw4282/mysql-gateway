package database

import (
	"database/sql"
	"github.com/sirupsen/logrus"
)

/*
   converts the inputs of rowData to channelData
	Input: RowsData
	Output: adds
*/
func RowsToChan(rows *sql.Rows, results chan []byte) {
	for rows.Next() {
		dest := []byte{}

		if err := rows.Scan(dest); err != nil {
			logrus.Errorln(err)
			continue
		}
		results <- dest
	}
	rows.Close()
}

/*
Responsible for feeding the []bytes into channels
Input: []bytes of data
Output: Nothing -  channel modified by byRef
*/
func BytesDataToChan(rows *sql.Rows, results chan []byte) (err error) {

	return err
}
