package queryHandlers

import "strings"

//Allows to 1) BuildTheQuery, by breaking it into components of the `r` module
func GetQueryFromUrl(url string) (resultQuery Query, err error) {
	splitURL := SplitUrlWithExclusion(url)

	resultQuery = Query{
		Table:       splitURL[0],
		comparisons: []comparators{}, // comparator struct to hold the value of the
	}

	//pass pointer of the results?
	resultQuery.comparisons, err = SplitsToCohesiveForm(splitURL, resultQuery)

	return
}

/*splits the url with respect to comparative ops and terms
E.g. Input string: users?a=eq.b&c=gt.d
output: [users,a=eq.b,c=gt.d] ( `,` separates an `&`)
*/
func SplitUrlWithExclusion(url string) []string {
	splitResult := strings.Split(url, "?")

	if len(splitResult) == 1 { //no split char present
		return splitResult
	}

	splitArrayResult := make([]string, 10)
	secondSplit := strings.Split(splitResult[1], "&")
	splitArrayResult = append(splitArrayResult, splitResult[0]) //adds entity/table name
	splitArrayResult = append(splitArrayResult, secondSplit...) //passes in a variadic function

	return splitArrayResult
}

/* Formats and updates the string using comparator struct and saving values to the relevant fields
   Input -> [users,a=eq.b,c=gt.d] ( `,` separates an `&`)
   Output -> resultsquery of all splits consisting of field, and all fields using type comparators (struct)
*/

//limits value, (select?) and others not in the splits
//Iterates through and updates it to resultObj as an array of comparators object
func SplitsToCohesiveForm(splitdata []string, resultQuery Query) (resultObj []comparators, err error) {

	for index, elem := range splitdata {
		if index == 0 { // skip table name as it is already added
			continue
		}
		data := strings.Split(elem, "=")
		comp := comparators{
			Field: data[0],
		}
		comparators := strings.Split(data[1], ".")

		comp.ComparatorObj = comparators[0]
		comp.Value = comparators[1]
		resultObj = append(resultObj, comp)
	}

	return
}
