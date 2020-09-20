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
	Table string
	//Column      string
	Select      []string // only applies for MULTI select ( of nested types)
	Limit       int
	Comparisons []Comparators
}

//SQL syntax: SELECT column_name(s) FROM table_name WHERE condition LIMIT number;

/* SELECT stmt formatter - create string, add in table and call to Where() to formatter WHERE
   Limit should also be appended to the string POST-where operation using relevant field
   output: p1: the formatted string. p2: bindedArrayData for (possible?) use!
*/
func (q *Query) formatSelectStmt() (queryStmt string, bindArray []interface{}) {
	queryStmt = generateSelect(q.Select)
	queryStmt += "FROM " + string(q.Table) + "as tbl"

	if q.Limit != 0 {
		queryStmt += "LIMIT " + strconv.Itoa(q.Limit)
		bindArray = append(bindArray, strconv.Itoa(q.Limit))
	}
	return
}

/* TODO: start with the unit tests and build it there on....
Check for single/multiple selects and update accordingly
** For more than 1 select, requires nested selects which needs to be formatted accordingly...
*/
func generateSelect(selectVal []string) (selectString string) {
	selectString = "SELECT "
	if len(selectVal) == 0 {
		selectString += ""
	} else {
		//code to handle nested select in SQL without row_tojson Format

	}

	return
}

//Aggregate functions checks and limits, etc.
func (q *Query) ProcessAggregateValues(aggFunc string, value string) (exists bool, err error) {
	if strings.ToLower(aggFunc) == "limit" {
		intLimit, _ := strconv.Atoi(value)
		q.Limit = intLimit
		return
	} else if aggFunc == "select" {
		q.Select = strings.Split(value, "")
		return
	}
	err = errors.New("Error processing parameters: AggFunc-> " + aggFunc + " and value-> " + value)
	return
}
