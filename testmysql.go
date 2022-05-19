package main

//user:password@tcp(localhost:5555)/dbname?tls=skip-verify&autocommit=true
//https://zetcode.com/golang/mysql/

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func testmysql() {

	db, err := sql.Open("mysql", "root:cst@tcp(127.0.0.1:3306)/testdb")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var version string

	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("mysql db")
	fmt.Println(version)
}
