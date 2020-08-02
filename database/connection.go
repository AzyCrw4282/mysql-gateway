package database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
)

//credentials to construct the datasourcename

var userName = ""
var password = ""
var database = ""

func connectOrFail() (db *sql.Conn, err error) {
	loadCredentials()
	datab, err := sql.Open("mysql", userName+":"+password+"@/"+database)

	if err != nil {
		fmt.Println("Failed to connect. Likely due to wrong credentails in .env or your sql server is not running")
		panic(err.Error())
	}
	ctx, _ := context.WithCancel(context.Background())
	db, _ = datab.Conn(ctx)

	return
}

func loadCredentials() {
	err := godotenv.Load("*.env")

	if err != nil {
		fmt.Println("error loading creds.") // apply printf to check for the actual values
	}
	userName = os.Getenv("userName")
	password = os.Getenv("password")
	database = os.Getenv("database")

}
