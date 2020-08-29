package queryHandlers

// input string x=eq.y
type comparators struct {
	Field         string //x
	ComparatorObj string //eq
	Value         string //y
}

//possible methods using receiver objects
///get for all fields
