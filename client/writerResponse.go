package client

import "net/http"

/*
Writes the given channel to the client
Input: []byte data from process queries
Output: writes it to CLI
*/

var LB = []byte("{")
var RB = []byte("}")
var sep = []byte(" | ")

func WriteToClient(w http.ResponseWriter, results chan []byte) {
	w.Write(LB)
	sepCount := 0
	for row := range results {
		if sepCount > 0 {
			w.Write(sep)
		}
		w.Write(row)
		sepCount++
	}
	w.Write(RB)
}
