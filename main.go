package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"mysql-gateway/api"
	"mysql-gateway/database"
	"net/http"
	"os"
	//"github.com/AzyCrw4282/mysql-gateway/api"
)

/*
main file in which the router initiation is made and then calls the db
*/

func main() {
	fmt.Println("In main")
	err := godotenv.Load("*.env")

	if err != nil {
		fmt.Println("Your .env file didnt get loaded. Make sure it's in the root directory")
		return
	}
	//start db connection cache for required pool size
	database.MakeSingleConnection()

	//database.CurrConnection = database.StartConnectionCache(4)

	r := mux.NewRouter()
	setRouterHandlerAndServe(r)

}

func setRouterHandlerAndServe(router *mux.Router) {
	//handle functions for now the first 3 functions

	router.HandleFunc("/", api.HandleAllHeaderOptions).Methods(http.MethodOptions)

	//router functions for all ops handling. patterns comes with variables separated by '/', as
	//can be seen in the router call

	router.HandleFunc("/{TableEntity}", api.HandleGetOneOrMany).Methods(http.MethodGet)
	router.HandleFunc("/{TableEntity}/{field}", api.HandleInsert).Methods(http.MethodPost)
	router.HandleFunc("/{TableEntity}/{field}/{id}", api.HandleDelete).Methods(http.MethodDelete)

	panic(http.ListenAndServe(os.Getenv("listenAddress"), router))

}
