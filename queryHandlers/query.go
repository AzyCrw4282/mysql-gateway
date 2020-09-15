package queryHandlers

import (
	"errors"
	"strconv"
	"strings"
)

/* For interface ->
interface{} means you can put value of any type, including your own custom type.
All types in Go satisfy an empty interface (interface{} is an empty interface).
*/

//struct obj for queries
type Query struct {
	Table       string
	column      string
	Select      []string
	id          int
	limit       int
	comparisons []comparators
}

//SQL syntax: SELECT column_name(s) FROM table_name WHERE condition LIMIT number;

/* SELECT stmt formatter - create string, add in table and call to Where() to formatter WHERE
   Limit should also be appended to the string POST-where operation using relevant field
   output: p1: the formatted string. p2: bindedArrayData for (possible?) use!
*/
func (q *Query) formatSelectStmt() (queryStmt string, bindArray []interface{}) {
	//generate select
	queryStmt = generateSelect(q.Select)
	queryStmt += "FROM " + string(q.Table) + "as tbl"

	if q.limit != 0 {
		queryStmt += "LIMIT " + strconv.Itoa(q.limit)
		bindArray = append(bindArray, strconv.Itoa(q.limit))
	}
	return
}

/*
Check for single/multiple selects and update accordingly
*/
func generateSelect(selectVal []string) (selectString string) {
	selectString = "SELECT"
	if len(selectVal) == 0 {
		selectString += ""
	}

	return
}

//Aggregate functions checks and limits, etc.
func (q *Query) ProcessAggregateValues(aggFunc string, value string) (exists bool, err error) {
	if strings.ToLower(aggFunc) == "limit" {
		intLimit, _ := strconv.Atoi(value)
		q.limit = intLimit
		return
	} else if aggFunc == "select" {
		q.Select = strings.Split(value, "")
		return
	}
	err = errors.New("Error processing parameters: AggFunc-> " + aggFunc + " and value-> " + value)
	return
}
