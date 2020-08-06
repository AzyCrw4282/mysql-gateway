package queryHandlers

/*
Consists of:
	equalities cmparison HM
	a struct for query
	And methods for:
					---select, ---upadte, --where,
*/
//keyMap
var symbolsValues = map[string]string{
	"gt":    ">",
	"lt":    "<",
	"lte":   "<=",
	"gte":   ">=",
	"eq":    "=",
	"NotEq": "!=",
}

//struct obj for queries
type Query struct {
	Table       string
	column      string
	id          int
	limit       int
	comparisons []comparators
}
