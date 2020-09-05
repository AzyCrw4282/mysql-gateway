package queryHandlers

import (
	"errors"
	"strconv"
	"strings"
)

/*
Consists of:
	equalities cmparison HM
	a struct for query
	And methods for:
					---select, ---upadte, --where,
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

/*
TODO: Add methods for the query based classes
*/

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
