package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //for side-effect use with the sql package
)

//channels and func will handle pool of db connections
var CurrConnection func()

func MakeConneciton() {
	db, err := connectOrFail()
	if err != true {
		if db.Ping() != nil {
			fmt.Print("db connection not able to be pinged")
		}

	}

}

func cosntantPing() {

	go func() {
		//for {
		//		//TODO: ping check if conenction drops
		//
		//}
	}()

}
