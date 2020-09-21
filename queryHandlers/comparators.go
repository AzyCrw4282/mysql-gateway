package main

// input string x=eq.y
type Comparators struct {
	Field         string //x, field name
	ComparatorObj string //eq, any operator
	Value         string //y, any operands
}

//keyMap
var symbolsValues = map[string]string{
	"gt":    ">",
	"lt":    "<",
	"lte":   "<=",
	"gte":   ">=",
	"eq":    "=",
	"NotEq": "!=",
	"is":    "is",
}

/*
 Returns the actual value in hashmap for SQL use
*/
func (c *Comparators) ComparatorToSQL() string {
	return symbolsValues[c.ComparatorObj]
}

//Methods to be added
