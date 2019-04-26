package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	mysqlHost     = "10.10.11.141"
	mysqlPort     = "3306"
	mysqlUsername = "admin"
	mysqlPassword = "1q2w3e4r"
)

type LandMartInfo struct {
	id       int
	ip       string
	typenum  int
	name     string
	country  string
	province string
	city     string
	asn      int
	jingdu   float32
	weidu    float32
}

func GetLandmarkType() (idToType map[int]string, err error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/landmark", mysqlUsername, mysqlPassword, mysqlHost, mysqlPort))
	if err != nil {
		log.Fatal(err)
		return
	}

	rows, err := db.Query("SELECT * FROM landmark.landmark_type;")
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()
	idToType = make(map[int]string)
	var id int
	var typename string
	for rows.Next() {
		err := rows.Scan(&id, &typename)
		if err != nil {
			log.Fatal(err)
		}
		idToType[id] = typename
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	return
}

func GetLandmarkData() (result map[string]LandMartInfo, typeToIp map[int][]string, err error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/landmark", mysqlUsername, mysqlPassword, mysqlHost, mysqlPort))
	if err != nil {
		log.Fatal(err)
		return
	}

	rows, err := db.Query("SELECT * FROM landmark.landmark_info;")
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()
	result = make(map[string]LandMartInfo)
	typeToIp = make(map[int][]string)
	for rows.Next() {
		var t LandMartInfo
		err := rows.Scan(&t.id, &t.ip, &t.typenum, &t.name, &t.country, &t.province, &t.city, &t.asn, &t.jingdu, &t.weidu)
		if err != nil {
			log.Fatal(err)
		}
		if _, ok := typeToIp[t.typenum]; ok {
			typeToIp[t.typenum] = append(typeToIp[t.typenum], t.ip)
		} else {
			typeToIp[t.typenum] = []string{t.ip}
		}
		result[t.ip] = t
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	return
}

//insertSql := "INSERT INTO `landmark`.`landmark_info` (`ip`, `type`, `name`, `country`, `province`, `city`, `asn`) " +
//				"VALUES (?, ?, ?, ?, ?, ?, ?);"
//stmt,err := db.Prepare(insertSql)
///*
//*    这个stmt的主要方法:Exec、Query、QueryRow、Close
//*/
//if err != nil {
//	log.Println(err)
//	return
//}
//
//
//
//res,err := stmt.Exec()
//if err != nil {
//	log.Println(err)
//	return
//}
//
//rowNum, err = res.RowsAffected()    //插入的是后RowsAffected 返回的是插入的条数
//if err != nil {
//	log.Println(err)
//	return
//}
