package database

import (
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
	"mysql-gateway/queryHandlers"
)

var (
	ctx context.Context
)

/* Return a channel of bytes of the results, then fed into writerResponse module
* Input: A comparator object: Field, comparatorObj( operator) and value
* Output: returns the final results, i.e. whatever returned by the mysql-driver-API
 */

//ToDO: Add SQL implementation using sqldriver for relevant functions.
//ToDO: Pass the output data rowsToChannels.go to output TUI

func GetData(q queryHandlers.Query) (results chan []byte, err error) {
	results = make(chan []byte)
	newConn := CurrConnection()
	defer newConn.Close()
	sqlQuery, compObj := q.FormatSelectStmt()
	var rows *sql.Rows

	if len(compObj) == 0 {
		rows, err = newConn.QueryContext(ctx, sqlQuery)
	} else {
		rows, err = newConn.QueryContext(ctx, sqlQuery, compObj...)
	}

	if err != nil {
		logrus.Errorln(sqlQuery)
		logrus.Errorln(err)
		logrus.Errorln(compObj, "len of compobj ", len(compObj))

		return
	}
	go RowsToChan(rows, results)
	close(results)
	return
}

func InsertData() {

}

func DeleteData() {

}

func UpdateData() {

}

func GetDataFields(table string, field string, id int) {

}
