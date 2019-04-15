package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


var (
	mysqlHost     = "10.10.11.141"
	mysqlPort     = "3306"
	mysqlUsername = "admin"
	mysqlPassword = "1q2w3e4r"
)

func GetMysqlConnect() (err error) {
	db, err := sql.Open("mysql", "root:123@tcp(192.168.10.15:3306)/gamedata_tian?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

}