package database

import (
	"mysql-gateway/queryHandlers"
)

/*
This will exec all prepared statements for all methods ops below.

*/
func InsertData() {

}

func DeleteData() {

}

func UpdateData() {

}

//return a channel of byes
func GetData(q queryHandlers.Query) (results chan []byte, err error) {

	return results, err

}

func GetDataFields(table string, field string, id int) {

}
