package main

import (
	"mysql-gateway/queryHandlers"
	"testing"
)

/*The splitterURLTest focuses on the parsing and splitting of the given URL string.
 * Input: http/https://localhost:8080/entity?id=eq.10
 * Output: Pass/fail assertion of all test cases
 */
func TestSimpleUnitForSplitters(t *testing.T) {
	urlString := "entity?Number=eq.10"
	query, err := queryHandlers.GetQueryFromUrl(urlString)
	if err != nil {
		t.Error("Error!!! Failed to parse and split a perfect string") //eq = log, followed by fail
		return
	}
	if query.Table != "entity" {
		t.Log("Error with entity value. Got->", query.Table, " Expected->", query.Table)
		t.Fail()
	}
	if len(query.Comparisons) != 1 {
		t.Log("Failed to obtain the required comparators value")
		t.Fail()
	}
	if query.Limit != 0 {
		t.Log("Obtained a wrong LIMIT value for the query")
		t.Fail()
	}
	if query.Comparisons[0].Field != "Number" {
		t.Log("Obtained a wrong field. Got: ", query.Column, " Expected: ", "Number")
		t.Fail()
	}
	if query.Comparisons[0].ComparatorToSQL() != "=" {
		t.Log("Error! Wrong comparator sign obtained-> ", query.Comparisons[0].ComparatorToSQL(), "Expected ", "=")
		t.Fail()
	}
	if query.Comparisons[0].Value != "10" {
		t.Log("Wrong passed value obtained. Got ", query.Comparisons[0].Value, " Expected ", 10)
		t.Fail()
	}
}
