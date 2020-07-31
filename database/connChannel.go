package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

//channels and func will handle pool of db connections
var CurrConnection func()

func MakeConneciton() {
	connectOrFail()

}

func cosntantPing() {

	go func() {
		for {
			sql.Conn{}

		}
	}()

}
