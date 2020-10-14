package database

import "mysql-gateway/queryHandlers"

/* Return a channel of bytes of the results, then fed into writerResponse module
* Input: A comparator object: Field, comparatorObj( operator) and value
* Output: returns the final results, i.e. whatever returned by the mysql-driver-API
 */

//ToDO: Add SQL implementation using sqldriver for relevant functions.
//ToDO: Pass the output data rowsToChannels.go to output TUI
//ToDO: Add compound tests for all implements

func GetData(q queryHandlers.Query) (results chan []byte, err error) {

	//call to -> formatSelectStmt()

	return results, err
}

func InsertData() {

}

func DeleteData() {

}

func UpdateData() {

}

func GetDataFields(table string, field string, id int) {

}
