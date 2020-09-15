package database

import (
	"mysql-gateway/queryHandlers"
)

/*
This will exec all prepared statements for all methods operations below.
*/
func InsertData() {

}

func DeleteData() {

}

func UpdateData() {

}

/* Return a channel of bytes of the results, then fed into writerResponse module
* Input: A comparator object: Field, comparatorObj( operator) and value
* Output: returns the final results, i.e. whatever returned by the mysql-driver-API
 */
func GetData(q queryHandlers.Query) (results chan []byte, err error) {

	//call to -> formatSelectStmt()

	return results, err
}

func GetDataFields(table string, field string, id int) {

}
