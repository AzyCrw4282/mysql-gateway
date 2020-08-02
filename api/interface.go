package api

import "net/http"

/* TODO: create login:
get the vars
check for errors efficiently, if so throw an http error
then call to db.method() in package, and in there, a call to process query
finally write the header using responseWriter and give a stuaus code

*/

func handleError()

//writer sets the header for what is to be retutrned
func HandleAllHeaderOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "GET, PATCH, DELETE, POST, OPTIONS")
}

//includes one and many operation
func HandleGet(w http.ResponseWriter, r *http.Request) {

}

func HandleDelete(w http.ResponseWriter, r *http.Request) {

}

func HandleInsert(w http.ResponseWriter, r *http.Request) {

}

//TODO: To handle relational table operation
func HandleRelationalOps(w http.ResponseWriter, r *http.Request) {

}
