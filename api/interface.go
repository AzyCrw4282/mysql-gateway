package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"mysql-gateway/client"
	"mysql-gateway/database"
	"mysql-gateway/queryHandlers"
	"net/http"
	"strconv"
)

/* TODO: create login:
get the vars
check for errors efficiently, if so throw an http error
then call to db.method() in package, and in there, a call to process query
finally write the header using responseWriter and give a stuaus code

*/
// :param-> expectedCode : the expected parms in the string, where the sum of all represents 3
func processString(w http.ResponseWriter, vars map[string]string, expectedCode int) (err error) {

	//entity,table check
	var error error
	for k, v := range vars {
		fmt.Println("k", k, "v", v)

		switch expectedCode {
		case 3:
			if vars[k] == "" {
				logrus.Println("Issue parsing your table/entity. Please make sure your cmds are correct")
				break
			}
		case 2:
			if vars[k] == "" {
				logrus.Println("Issue parsing your row Please make sure your cmds are correct")
				break
			}
		case 1:
			if vars[k] == "" {
				logrus.Println("Issue parsing your id. Please make sure your cmds are correct")
				break
			}
		default:
			fmt.Println("Error in case statements!!!")
		}

	}
	http.Error(w, "Error with your request /{entity}", http.StatusBadRequest)
	return error
}

func parseString(vars map[string]string) (table string, field string, id int) {

	switch len(vars) {
	case 3:
		table = vars["entity"]
	case 2:
		field = vars["field"]
	case 1:
		id, _ = strconv.Atoi(vars["id"]) //can use atoi, or itoa for the reversal ops
	}
	return

}

//writer sets the header for what is to be retutrned
func HandleAllHeaderOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "GET, PATCH, DELETE, POST, OPTIONS")
}

//includes one and many operation
func HandleGetOneOrMany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	//check for all possibities
	err := processString(w, vars, len(vars))

	if err != nil {
		fmt.Println("You are missing an empty field.")
		http.Error(w, "Error with your request /{entity}", http.StatusBadRequest)
	}

	parseString(vars) //TODO:check feasibility of this

	query, err := queryHandlers.GetQueryFromUrl(r.URL.String()) //forms the query from the url

	if err != nil {
		fmt.Println("You are missing an empty field.")
		CallHTTPBadRequest(w)
	}

	result, err := database.GetData(query)

	client.WriteToClient(result) //passes the chan to write it to client

	w.WriteHeader(http.StatusOK)

}

func HandleDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	//check for all possibities
	err := processString(w, vars, len(vars))

	if err != nil {
		fmt.Println("You are missing an empty field.")
		http.Error(w, "Error with your request /{entity}", http.StatusBadRequest)
	}

	database.DeleteData()

	w.WriteHeader(http.StatusOK)

}

func HandleInsert(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	//check for all possibities
	err := processString(w, vars, len(vars))

	if err != nil {
		fmt.Println("You are missing an empty field.")

	}

	database.InsertData()

	w.WriteHeader(http.StatusOK)

}

func HandleUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	//check for all possibities
	err := processString(w, vars, len(vars))

	if err != nil {
		fmt.Println("You are missing an empty field.")
		http.Error(w, "Error with your request /{entity}", http.StatusBadRequest)
	}

	database.UpdateData()

	w.WriteHeader(http.StatusOK)

}

//TODO: To handle relational table operation at later stage.
func HandleRelationalOps(w http.ResponseWriter, r *http.Request) {

}

func CallHTTPBadRequest(w http.ResponseWriter) {
	http.Error(w, "Error with your request /{entity}", http.StatusBadRequest)
}
