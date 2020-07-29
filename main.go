package main

import (
	"fmt"
	//"mysql-gateway/api" local import issue
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
	"github.com/AzyCrw4282/mysql-gateway/api"
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

	r := mux.NewRouter()
	setRouterHandlerAndServe(r)

}

func setRouterHandlerAndServe(r *mux.Router) {
	//handle functions for now the first 3 functions

	r.HandleFunc("/",api.)

}
