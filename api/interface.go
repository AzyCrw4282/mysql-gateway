package api

import "net/http"

//writer sets the header for what is to be retutrned
func handleAllHeaderOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "GET, PATCH, DELETE, POST, OPTIONS")
}

func handleGet(w http.ResponseWriter, r *http.Request) {

}

func handleDelete(w http.ResponseWriter, r *http.Request) {

}

func handleInsert(w http.ResponseWriter, r *http.Request) {

}
