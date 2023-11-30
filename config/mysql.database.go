package database

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connection() (*sql.DB) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/testshop")
	
	if err != nil {
		fmt.Println("connect fail")
	} else {
		fmt.Println("connect success")
	}

	return	db
}