package main

import (
	query2 "mysql-gateway/queryHandlers" // Aliased import names. used for main package reference instead of original name
	"testing"
)

/*The splitterURLTest focuses on the parsing and splitting of the given URL string.
 * Input: http/https://localhost:8080/entity?id=eq.10
 * Output: Pass/fail assertion of all test cases
 */
func TestSimpleUnitForSplitters(t *testing.T) {
	urlString := "entity?Number=eq.10"
	query, err := query2.GetQueryFromUrl(urlString)
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
		t.Log("Obtained a wrong field. Got: ", query.Comparisons[0].Field, " Expected: ", "Number")
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

/*The splitterURLTest focuses on the parsing and splitting of the mid-complex URL string.
 * Input: http/https://localhost:8080/entity?Number=eq.10&EvenNumber=is.True
 * Output: Pass/fail assertion of all test cases
 */
func TestMidUnitForSplitters(t *testing.T) {
	urlString := "Nums?Number=eq.10&EvenNumber=is.True"
	query, err := query2.GetQueryFromUrl(urlString)
	if err != nil {
		t.Error("Error!!! Failed to parse and split a perfect string") //eq = log, followed by fail
		return
	}
	if query.Table != "Nums" {
		t.Log("Error with entity value. Got->", query.Table, " Expected->", query.Table)
		t.Fail()
	}
	if len(query.Comparisons) != 2 {
		t.Log("Failed to obtain the required comparators value")
		t.Fail()
	}
	if query.Limit != 0 {
		t.Log("Obtained a wrong LIMIT value for the query")
		t.Fail()
	}
	if query.Comparisons[0].Field != "Number" && query.Comparisons[1].Field != "EvenNumber" {
		t.Log("Obtained a wrong field. Got: ", query.Comparisons[0].Field, " Expected: ", "Number or EvenNumber")
		t.Fail()
	}
	if query.Comparisons[0].ComparatorToSQL() != "=" && query.Comparisons[1].ComparatorToSQL() != "is" {
		t.Log("Error! Wrong comparator sign obtained-> ", query.Comparisons[0].ComparatorToSQL(), "Expected ", "=")
		t.Fail()
	}
	if query.Comparisons[1].Value != "True" {
		t.Log("Wrong passed value obtained. Got ", query.Comparisons[0].Value, " Expected ", "True")
		t.Fail()
	}
}

/*The splitterURLTest focuses on the parsing and splitting of the complex URL string.
 * Input: http/https://localhost:8080/entity?Number=eq.10&EvenNumber=is.True
 * Output: Pass/fail assertion of all test cases
 */
func TestComplexUnitForSplitters(t *testing.T) {
	urlString := "Nums?Number=eq.10&EvenNumber=is.True&NumberType=eq.Even&LIMIT=100"
	query, err := query2.GetQueryFromUrl(urlString)
	if err != nil {
		t.Error("Error!!! Failed to parse and split a perfect string") //eq = log, followed by fail
		return
	}
	if query.Table != "Nums" {
		t.Log("Error with entity value. Got->", query.Table, " Expected->", query.Table)
		t.Fail()
	}
	if len(query.Comparisons) != 3 {
		t.Log("Failed to obtain the required comparators value")
		t.Fail()
	}
	if query.Limit != 100 {
		t.Log("Obtained a wrong LIMIT value for the query, GOT ", query.Limit)
		t.Fail()
	}
	if query.Comparisons[0].Field != "Number" {
		t.Log("Obtained a wrong field. Got: ", query.Comparisons[0].Field, " Expected: ", "Number ")
		t.Fail()
	}
	if query.Comparisons[0].ComparatorToSQL() != "=" {
		t.Log("Error! Wrong comparator sign obtained-> ", query.Comparisons[0].ComparatorToSQL(), "Expected ", "=")
		t.Fail()
	}
	if query.Comparisons[0].Value != "10" {
		t.Log("Wrong passed value obtained. Got ", query.Comparisons[0].Value, " Expected ", "10")
		t.Fail()
	}
	if query.Comparisons[1].Field != "EvenNumber" {
		t.Log("Obtained a wrong field. Got: ", query.Comparisons[1].Field, " Expected: ", "Number ")
		t.Fail()
	}
	if query.Comparisons[1].ComparatorToSQL() != "is" {
		t.Log("Error! Wrong comparator sign obtained-> ", query.Comparisons[1].ComparatorToSQL(), "Expected ", "is")
		t.Fail()
	}
	if query.Comparisons[1].Value != "True" {
		t.Log("Wrong passed value obtained. Got ", query.Comparisons[1].Value, " Expected ", "True")
		t.Fail()
	}
	if query.Comparisons[2].Field != "NumberType" {
		t.Log("Obtained a wrong field. Got: ", query.Comparisons[2].Field, " Expected: ", "NumberType")
		t.Fail()
	}
	if query.Comparisons[2].ComparatorToSQL() != "=" {
		t.Log("Error! Wrong comparator sign obtained-> ", query.Comparisons[2].ComparatorToSQL(), "Expected ", "=")
		t.Fail()
	}
	if query.Comparisons[2].Value != "Even" {
		t.Log("Wrong passed value obtained. Got ", query.Comparisons[2].Value, " Expected ", "Even")
		t.Fail()
	}
	t.Log("Successful")
}

/*
Add tests for:
	Single Limit checks
	Select stmt checks - (TODO: single and multi)
*/

func TestLimitFuncforSplitters(t *testing.T) {
	url := "Numbers?Limit=5"
	query, err := query2.GetQueryFromUrl(url)

	if err != nil {
		t.Error("Can't parse a simple LIMIT query") //eq = log, followed by fail
		return
	}

	if query.Limit != 5 && query.Table != "Numbers" {
		t.Log("Wrong LIMIT or TABLE value obtained", " Got ", query.Limit, " table-> ", query.Table)
		t.Fail()
	}

}
