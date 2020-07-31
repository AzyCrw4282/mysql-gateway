package api

import "net/http"

//writer sets the header for what is to be retutrned
func HandleAllHeaderOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "GET, PATCH, DELETE, POST, OPTIONS")
}

//includees one and many operation
//includees one and many operation
func HandleGet(w http.ResponseWriter, r *http.Request) {

}

func HandleDelete(w http.ResponseWriter, r *http.Request) {

}

func HandleInsert(w http.ResponseWriter, r *http.Request) {

}
