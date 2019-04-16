package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	mysqlHost     = "10.10.11.141"
	mysqlPort     = "3306"
	mysqlUsername = "admin"
	mysqlPassword = "1q2w3e4r"
)

func GetMysqlConnect() (err error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s:%s", mysqlUsername, mysqlPassword, mysqlHost, mysqlPort))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()
	return

}
