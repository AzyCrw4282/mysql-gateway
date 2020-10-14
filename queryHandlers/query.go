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
	Select      []string // for multi case, in which SELECT stores the extra fields
	Limit       int      // need to be expanded to allow for multiple LIMIT queries
	Comparisons []Comparators
}

//SQL syntax: SELECT column_name(s) FROM table_name WHERE (condition) LIMIT number;

/* SELECT stmt formatter - create string, add in table and call to Where() to formatter WHERE
   Limit should also be appended to the string POST-where operation using relevant field
   output: p1: the formatted string. p2: bindedArrayData for (possible?) use!
*/
func (q *Query) formatSelectStmt() (queryStmt string, bindArray []interface{}) {
	queryStmt = generateSelect(q.Select, q)
	queryStmt += "FROM " + string(q.Table) + "as tbl"
	queryStmt, bindArray = generateWhere(q, queryStmt)

	if q.Limit != 0 {
		queryStmt += "LIMIT " + strconv.Itoa(q.Limit)
		bindArray = append(bindArray, strconv.Itoa(q.Limit))
	}
	return
}

/*
Check for single/multiple selects and update accordingly
** For more than 1 select, requires nested selects which needs to be formatted...
** Ref here  -> https://stackoverflow.com/questions/1775168/multiple-select-statements-in-single-query
*/
func generateSelect(selectVal []string, q *Query) (selectString string) {
	selectString = "SELECT "
	if len(selectVal) == 0 {
		selectString += "AS inst"
	} else {
		//code to handle nested select in SQL without row_to_json Format
		selectString += "(SELECT (tbl."
		for id, selectStmt := range q.Select {
			if id != 0 {
				selectString += ", tbl."
			}
			selectString += selectStmt
		}
		selectString += ") AS mainQuery"
	}
	return
}

//generate WHERE stmt using comparisons objects
func generateWhere(q *Query, queryString string) (whereString string, bindArr []interface{}) {
	whereString = " " + queryString
	if len(q.Comparisons) == 0 { //if no comparators no Where clause needed
		return
	}
	whereString += "WHERE "
	for ind, compObj := range q.Comparisons {
		if ind != 0 {
			whereString += "AND "
		}
		//forms the where query for each check
		whereString += compObj.Field + " " + compObj.ComparatorObj + "$" + strconv.Itoa(ind) // Poss?-> compObj.Value
		bindArr = append(bindArr, compObj.Value)
	}
	return
}

//Aggregate functions checks and limits, etc.
func (q *Query) ProcessAggregateValues(aggFunc string, value string) (exists bool, err error) {
	if strings.ToLower(aggFunc) == "limit" {
		intLimit, _ := strconv.Atoi(value)
		q.Limit = intLimit
		return
	} else if strings.ToLower(aggFunc) == "select" {
		q.Select = strings.Split(value, ",")
		return
	}
	err = errors.New("Error processing parameters: AggFunc-> " + aggFunc + " and value-> " + value)
	return
}
