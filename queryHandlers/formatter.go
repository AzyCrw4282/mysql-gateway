package queryHandlers

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

//Allows to 1) BuildTheQuery, by breaking it into components of the `r` module
func GetQueryFromUrl(url string) (resultQuery Query, err error) {
	splitURL := SplitUrlWithExclusion(url)
	resultQuery = Query{
		Table:       splitURL[0],
		Comparisons: []Comparators{}, // comparator struct to hold the value of the
	}
	//pass pointer of the results?
	resultQuery, err = SplitsToCohesiveForm(&resultQuery, splitURL)
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
	splitArrayResult := make([]string, 1)
	splitArrayResult[0] = splitResult[0]
	secondSplit := strings.Split(splitResult[1], "&")
	//splitArrayResult = append(splitArrayResult, splitResult[0]) //adds entity/table name
	splitArrayResult = append(splitArrayResult, secondSplit...) //passes in a variadic function

	return splitArrayResult
}

/* Formats and updates the string using comparator struct and saving values to the relevant fields
   Input -> [users,a=eq.b,c=gt.d] ( `,` separates an `&`)
   Output -> results query of all splits, consisting of field, and all fields using type comparators (struct)
*/
func SplitsToCohesiveForm(resultQuery *Query, splitdata []string) (resultObj Query, err error) {
	for index, elem := range splitdata {
		if index == 0 { // skip table name as it is already added
			continue
		}
		data := strings.Split(elem, "=")
		comp := Comparators{
			Field: data[0],
		}
		secSplit := strings.Split(data[1], ".") //if no `.` means an aggregate function and it wont split and resolves existing element to the array
		//check for Limit or aggregate function
		if len(secSplit) == 1 {
			_, Innererr := resultQuery.ProcessAggregateValues(data[0], secSplit[0])
			if Innererr != nil {
				logrus.Error(err)
				return
			}
			fmt.Println("Boolean eval of aggregate or limit functions SET") //-- using print() interfers with line interrupt so add \n break or use println
			continue
		}
		comp.ComparatorObj = secSplit[0]
		comp.Value = secSplit[1]
		resultQuery.Comparisons = append(resultQuery.Comparisons, comp)
	}
	resultObj = *resultQuery //dereference the val
	return
}
