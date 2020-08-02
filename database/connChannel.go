package database

import (
	//"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //for side-effect use with the sql package
	"os"
)

//channels and func will handle pool of db connections
var CurrConnection func() *sql.Conn

//type def for multiple connections
type connctionCache chan *sql.Conn

func MakeSingleConnection() {

	db, err := connectOrFail()
	if err != nil {
		fmt.Printf("failed to conenct, db data %v", db)
		os.Exit(1)
	}

	fmt.Println("connection established")

}

func (con connctionCache) MakeCacheConnection() {
	for {
		db, _ := connectOrFail()
		con <- db
	}
}

//called from main class to create the 1.m connections
func StartConnectionCache(size int) func() *sql.Conn {

	var conCache connctionCache = make(chan *sql.Conn, size)
	go conCache.MakeCacheConnection()
	return func() *sql.Conn { return <-conCache }

}
