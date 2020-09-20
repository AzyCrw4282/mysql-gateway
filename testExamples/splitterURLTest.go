package main

import (
	"fmt"
	"mysql-gateway/queryHandlers"
	"testing"
)

/*The splitterURLTest focuses on the parsing and splitting of the given URL string.
 * Input: http/https://localhost:8080/entity?id=eq.10
 * Output: Pass/fail assertion of all test cases
 */

func SimpleUnitTestForSplitters(t *testing.T) {
	urlString := "entity?id=eq.10"
	query, err := queryHandlers.GetQueryFromUrl(urlString)

	if err != nil {
		t.Error("Error! Failed to parse and split a perfect string")
		t.Fail()
		return
	}

}
