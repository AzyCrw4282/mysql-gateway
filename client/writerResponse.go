package client

import "net/http"

/*
Writes the given channel to the client
Input: []byte data from process queries
Output: writes it to CLI
*/
func WriteToClient(w http.ResponseWriter, data chan []byte) {

}
