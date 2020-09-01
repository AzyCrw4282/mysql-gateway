package queryHandlers

// input string x=eq.y
type comparators struct {
	Field         string //x, field name
	ComparatorObj string //eq, any operator
	Value         string //y, any operands
}

//possible methods using receiver objects
///get for all fields
